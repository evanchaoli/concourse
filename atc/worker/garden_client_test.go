package worker_test

import (
	"net/http"
	"time"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
	"github.com/onsi/gomega/ghttp"
	"github.com/concourse/concourse/atc/worker"
	"github.com/concourse/concourse/atc/worker/transport/transportfakes"
	"github.com/concourse/retryhttp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("retryable garden client", func() {
	var (
		gServer *ghttp.Server
	)

	BeforeEach(func() {
		gServer = ghttp.NewServer()
	})

	AfterEach(func() {
		gServer.Close()
	})

	Context("blocking forever", func() {
		BeforeEach(func() {
			gServer.Reset()
			gServer.AppendHandlers(func(w http.ResponseWriter, r *http.Request) {
				for {
					time.Sleep(5 * time.Second)
				}
			})
		})

		It("fails once context expires", func() {
			fakeTransport := new(transportfakes.FakeTransportDB)
			fakeLogger := lager.NewLogger("test")
			hostname := new(string)
			*hostname = gServer.Addr()

			clientFactory := worker.NewGardenClientFactory(
				fakeTransport,
				fakeLogger,
				"wont-talk-to-you",
				hostname,
				retryhttp.NewExponentialBackOffFactory(1*time.Second),
				3*time.Second,
			)

			client := clientFactory.NewClient()
			_, err := client.Create(garden.ContainerSpec{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Client.Timeout"))
		})
	})
})
