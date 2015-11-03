package tsa_test

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudfoundry-incubator/garden"
	gfakes "github.com/cloudfoundry-incubator/garden/fakes"
	"github.com/concourse/atc"
	. "github.com/concourse/tsa"
	"github.com/concourse/tsa/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/ghttp"
	"github.com/pivotal-golang/lager"
	"github.com/pivotal-golang/lager/lagertest"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"
	"github.com/tedsuo/rata"
)

var _ = Describe("Heartbeater", func() {
	type registration struct {
		worker       atc.Worker
		ttl          time.Duration
		lastInterval time.Duration
	}

	var (
		logger lager.Logger

		addrToRegister string
		interval       time.Duration
		cprInterval    time.Duration
		resourceTypes  []atc.WorkerResourceType

		expectedWorker     atc.Worker
		fakeTokenGenerator *fakes.FakeTokenGenerator
		fakeGardenClient   *gfakes.FakeClient
		fakeATC            *ghttp.Server

		heartbeater ifrit.Process

		verifyHeartbeat http.HandlerFunc

		registrations <-chan registration
		clientWriter  *gbytes.Buffer
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")

		addrToRegister = "1.2.3.4:7777"
		interval = time.Second
		cprInterval = 100 * time.Millisecond
		resourceTypes = []atc.WorkerResourceType{
			{
				Type:  "git",
				Image: "docker:///concourse/git-resource",
			},
		}

		expectedWorker = atc.Worker{
			GardenAddr:       addrToRegister,
			ActiveContainers: 2,
			ResourceTypes:    resourceTypes,
			Platform:         "some-platform",
			Tags:             []string{"some", "tags"},
		}

		fakeATC = ghttp.NewServer()

		registerRoute, found := atc.Routes.FindRouteByName(atc.RegisterWorker)
		Expect(found).To(BeTrue())

		registered := make(chan registration, 100)
		registrations = registered

		lastRequestTime := time.Now()
		verifyHeartbeat = ghttp.CombineHandlers(
			ghttp.VerifyRequest(registerRoute.Method, registerRoute.Path),
			func(w http.ResponseWriter, r *http.Request) {
				var worker atc.Worker
				Expect(r.Header.Get("Authorization")).To(Equal("Bearer yo"))

				err := json.NewDecoder(r.Body).Decode(&worker)
				Expect(err).NotTo(HaveOccurred())

				ttl, err := time.ParseDuration(r.URL.Query().Get("ttl"))
				Expect(err).NotTo(HaveOccurred())

				requestTime := time.Now()
				lastInterval := requestTime.Sub(lastRequestTime)
				lastRequestTime = requestTime

				registered <- registration{worker, ttl, lastInterval}
			},
		)

		fakeGardenClient = new(gfakes.FakeClient)
		fakeTokenGenerator = new(fakes.FakeTokenGenerator)

		fakeTokenGenerator.GenerateTokenReturns("yo", nil)
		clientWriter = gbytes.NewBuffer()
	})

	JustBeforeEach(func() {
		atcEndpoint := rata.NewRequestGenerator(fakeATC.URL(), atc.Routes)
		heartbeater = ifrit.Invoke(
			NewHeartbeater(
				logger,
				interval,
				cprInterval,
				fakeGardenClient,
				atcEndpoint,
				fakeTokenGenerator,
				atc.Worker{
					GardenAddr:    addrToRegister,
					ResourceTypes: resourceTypes,
					Platform:      "some-platform",
					Tags:          []string{"some", "tags"},
				},
				clientWriter,
			),
		)
	})

	AfterEach(func() {
		ginkgomon.Interrupt(heartbeater)
	})

	Context("when Garden returns containers", func() {
		BeforeEach(func() {
			containers := make(chan []garden.Container, 4)

			containers <- []garden.Container{
				new(gfakes.FakeContainer),
				new(gfakes.FakeContainer),
			}

			containers <- []garden.Container{
				new(gfakes.FakeContainer),
				new(gfakes.FakeContainer),
				new(gfakes.FakeContainer),
				new(gfakes.FakeContainer),
				new(gfakes.FakeContainer),
			}

			containers <- []garden.Container{
				new(gfakes.FakeContainer),
				new(gfakes.FakeContainer),
				new(gfakes.FakeContainer),
				new(gfakes.FakeContainer),
			}

			containers <- []garden.Container{
				new(gfakes.FakeContainer),
				new(gfakes.FakeContainer),
				new(gfakes.FakeContainer),
			}

			close(containers)

			fakeGardenClient.ContainersStub = func(garden.Properties) ([]garden.Container, error) {
				return <-containers, nil
			}
		})

		Context("when the ATC responds to heartbeat requests", func() {
			BeforeEach(func() {
				fakeATC.AppendHandlers(verifyHeartbeat, verifyHeartbeat)
			})

			It("immediately registers", func() {
				actualRegistration := <-registrations
				Expect(actualRegistration.worker).To(Equal(expectedWorker))
				Expect(actualRegistration.ttl).To(Equal(2 * interval))
				Expect(actualRegistration.lastInterval).To(BeNumerically("~", 0, 100*time.Millisecond))
			})

			It("heartbeats", func() {
				<-registrations
				actualRegistration := <-registrations
				expectedWorker.ActiveContainers = 5
				Expect(actualRegistration.worker).To(Equal(expectedWorker))
				Expect(actualRegistration.ttl).To(Equal(2 * interval))
				Expect(actualRegistration.lastInterval).To(BeNumerically("~", interval, 20*time.Millisecond))
			})
		})

		Context("when the ATC doesn't respond to the first heartbeat", func() {
			BeforeEach(func() {
				fakeATC.AppendHandlers(
					verifyHeartbeat,
					ghttp.CombineHandlers(
						verifyHeartbeat,
						func(w http.ResponseWriter, r *http.Request) { fakeATC.CloseClientConnections() },
					),
					verifyHeartbeat,
					verifyHeartbeat,
				)
			})

			It("heartbeats faster according to cprInterval", func() {
				<-registrations
				<-registrations
				actualRegistration := <-registrations
				expectedWorker.ActiveContainers = 4
				Expect(actualRegistration.worker).To(Equal(expectedWorker))
				Expect(actualRegistration.ttl).To(Equal(2 * interval))
				Expect(actualRegistration.lastInterval).To(BeNumerically("~", cprInterval, 20*time.Millisecond))
			})

			It("goes back to normal after the heartbeat succeeds", func() {
				<-registrations
				<-registrations
				<-registrations
				actualRegistration := <-registrations
				expectedWorker.ActiveContainers = 3
				Expect(actualRegistration.worker).To(Equal(expectedWorker))
				Expect(actualRegistration.ttl).To(Equal(2 * interval))
				Expect(actualRegistration.lastInterval).To(BeNumerically("~", interval, 20*time.Millisecond))
			})
		})
	})
})
