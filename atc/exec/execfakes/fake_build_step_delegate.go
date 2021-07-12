// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"context"
	"io"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/atc/worker"
	"github.com/concourse/concourse/tracing"
	"go.opentelemetry.io/otel/trace"
)

type FakeBuildStepDelegate struct {
	ConstructAcrossSubstepsStub        func([]byte, []atc.AcrossVar, [][]interface{}) ([]atc.VarScopedPlan, error)
	constructAcrossSubstepsMutex       sync.RWMutex
	constructAcrossSubstepsArgsForCall []struct {
		arg1 []byte
		arg2 []atc.AcrossVar
		arg3 [][]interface{}
	}
	constructAcrossSubstepsReturns struct {
		result1 []atc.VarScopedPlan
		result2 error
	}
	constructAcrossSubstepsReturnsOnCall map[int]struct {
		result1 []atc.VarScopedPlan
		result2 error
	}
	ContainerOwnerStub        func(atc.PlanID) db.ContainerOwner
	containerOwnerMutex       sync.RWMutex
	containerOwnerArgsForCall []struct {
		arg1 atc.PlanID
	}
	containerOwnerReturns struct {
		result1 db.ContainerOwner
	}
	containerOwnerReturnsOnCall map[int]struct {
		result1 db.ContainerOwner
	}
	ErroredStub        func(lager.Logger, string)
	erroredMutex       sync.RWMutex
	erroredArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	FetchImageStub        func(context.Context, atc.ImageResource, atc.VersionedResourceTypes, bool) (worker.ImageSpec, error)
	fetchImageMutex       sync.RWMutex
	fetchImageArgsForCall []struct {
		arg1 context.Context
		arg2 atc.ImageResource
		arg3 atc.VersionedResourceTypes
		arg4 bool
	}
	fetchImageReturns struct {
		result1 worker.ImageSpec
		result2 error
	}
	fetchImageReturnsOnCall map[int]struct {
		result1 worker.ImageSpec
		result2 error
	}
	FinishedStub        func(lager.Logger, bool)
	finishedMutex       sync.RWMutex
	finishedArgsForCall []struct {
		arg1 lager.Logger
		arg2 bool
	}
	InitializingStub        func(lager.Logger)
	initializingMutex       sync.RWMutex
	initializingArgsForCall []struct {
		arg1 lager.Logger
	}
	SelectedWorkerStub        func(lager.Logger, string)
	selectedWorkerMutex       sync.RWMutex
	selectedWorkerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	StartSpanStub        func(context.Context, string, tracing.Attrs) (context.Context, trace.Span)
	startSpanMutex       sync.RWMutex
	startSpanArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 tracing.Attrs
	}
	startSpanReturns struct {
		result1 context.Context
		result2 trace.Span
	}
	startSpanReturnsOnCall map[int]struct {
		result1 context.Context
		result2 trace.Span
	}
	StartingStub        func(lager.Logger)
	startingMutex       sync.RWMutex
	startingArgsForCall []struct {
		arg1 lager.Logger
	}
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct {
	}
	stderrReturns struct {
		result1 io.Writer
	}
	stderrReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct {
	}
	stdoutReturns struct {
		result1 io.Writer
	}
	stdoutReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	WaitingForWorkerStub        func(lager.Logger)
	waitingForWorkerMutex       sync.RWMutex
	waitingForWorkerArgsForCall []struct {
		arg1 lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildStepDelegate) ConstructAcrossSubsteps(arg1 []byte, arg2 []atc.AcrossVar, arg3 [][]interface{}) ([]atc.VarScopedPlan, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []atc.AcrossVar
	if arg2 != nil {
		arg2Copy = make([]atc.AcrossVar, len(arg2))
		copy(arg2Copy, arg2)
	}
	var arg3Copy [][]interface{}
	if arg3 != nil {
		arg3Copy = make([][]interface{}, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.constructAcrossSubstepsMutex.Lock()
	ret, specificReturn := fake.constructAcrossSubstepsReturnsOnCall[len(fake.constructAcrossSubstepsArgsForCall)]
	fake.constructAcrossSubstepsArgsForCall = append(fake.constructAcrossSubstepsArgsForCall, struct {
		arg1 []byte
		arg2 []atc.AcrossVar
		arg3 [][]interface{}
	}{arg1Copy, arg2Copy, arg3Copy})
	stub := fake.ConstructAcrossSubstepsStub
	fakeReturns := fake.constructAcrossSubstepsReturns
	fake.recordInvocation("ConstructAcrossSubsteps", []interface{}{arg1Copy, arg2Copy, arg3Copy})
	fake.constructAcrossSubstepsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBuildStepDelegate) ConstructAcrossSubstepsCallCount() int {
	fake.constructAcrossSubstepsMutex.RLock()
	defer fake.constructAcrossSubstepsMutex.RUnlock()
	return len(fake.constructAcrossSubstepsArgsForCall)
}

func (fake *FakeBuildStepDelegate) ConstructAcrossSubstepsCalls(stub func([]byte, []atc.AcrossVar, [][]interface{}) ([]atc.VarScopedPlan, error)) {
	fake.constructAcrossSubstepsMutex.Lock()
	defer fake.constructAcrossSubstepsMutex.Unlock()
	fake.ConstructAcrossSubstepsStub = stub
}

func (fake *FakeBuildStepDelegate) ConstructAcrossSubstepsArgsForCall(i int) ([]byte, []atc.AcrossVar, [][]interface{}) {
	fake.constructAcrossSubstepsMutex.RLock()
	defer fake.constructAcrossSubstepsMutex.RUnlock()
	argsForCall := fake.constructAcrossSubstepsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBuildStepDelegate) ConstructAcrossSubstepsReturns(result1 []atc.VarScopedPlan, result2 error) {
	fake.constructAcrossSubstepsMutex.Lock()
	defer fake.constructAcrossSubstepsMutex.Unlock()
	fake.ConstructAcrossSubstepsStub = nil
	fake.constructAcrossSubstepsReturns = struct {
		result1 []atc.VarScopedPlan
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildStepDelegate) ConstructAcrossSubstepsReturnsOnCall(i int, result1 []atc.VarScopedPlan, result2 error) {
	fake.constructAcrossSubstepsMutex.Lock()
	defer fake.constructAcrossSubstepsMutex.Unlock()
	fake.ConstructAcrossSubstepsStub = nil
	if fake.constructAcrossSubstepsReturnsOnCall == nil {
		fake.constructAcrossSubstepsReturnsOnCall = make(map[int]struct {
			result1 []atc.VarScopedPlan
			result2 error
		})
	}
	fake.constructAcrossSubstepsReturnsOnCall[i] = struct {
		result1 []atc.VarScopedPlan
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildStepDelegate) ContainerOwner(arg1 atc.PlanID) db.ContainerOwner {
	fake.containerOwnerMutex.Lock()
	ret, specificReturn := fake.containerOwnerReturnsOnCall[len(fake.containerOwnerArgsForCall)]
	fake.containerOwnerArgsForCall = append(fake.containerOwnerArgsForCall, struct {
		arg1 atc.PlanID
	}{arg1})
	stub := fake.ContainerOwnerStub
	fakeReturns := fake.containerOwnerReturns
	fake.recordInvocation("ContainerOwner", []interface{}{arg1})
	fake.containerOwnerMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildStepDelegate) ContainerOwnerCallCount() int {
	fake.containerOwnerMutex.RLock()
	defer fake.containerOwnerMutex.RUnlock()
	return len(fake.containerOwnerArgsForCall)
}

func (fake *FakeBuildStepDelegate) ContainerOwnerCalls(stub func(atc.PlanID) db.ContainerOwner) {
	fake.containerOwnerMutex.Lock()
	defer fake.containerOwnerMutex.Unlock()
	fake.ContainerOwnerStub = stub
}

func (fake *FakeBuildStepDelegate) ContainerOwnerArgsForCall(i int) atc.PlanID {
	fake.containerOwnerMutex.RLock()
	defer fake.containerOwnerMutex.RUnlock()
	argsForCall := fake.containerOwnerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildStepDelegate) ContainerOwnerReturns(result1 db.ContainerOwner) {
	fake.containerOwnerMutex.Lock()
	defer fake.containerOwnerMutex.Unlock()
	fake.ContainerOwnerStub = nil
	fake.containerOwnerReturns = struct {
		result1 db.ContainerOwner
	}{result1}
}

func (fake *FakeBuildStepDelegate) ContainerOwnerReturnsOnCall(i int, result1 db.ContainerOwner) {
	fake.containerOwnerMutex.Lock()
	defer fake.containerOwnerMutex.Unlock()
	fake.ContainerOwnerStub = nil
	if fake.containerOwnerReturnsOnCall == nil {
		fake.containerOwnerReturnsOnCall = make(map[int]struct {
			result1 db.ContainerOwner
		})
	}
	fake.containerOwnerReturnsOnCall[i] = struct {
		result1 db.ContainerOwner
	}{result1}
}

func (fake *FakeBuildStepDelegate) Errored(arg1 lager.Logger, arg2 string) {
	fake.erroredMutex.Lock()
	fake.erroredArgsForCall = append(fake.erroredArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.ErroredStub
	fake.recordInvocation("Errored", []interface{}{arg1, arg2})
	fake.erroredMutex.Unlock()
	if stub != nil {
		fake.ErroredStub(arg1, arg2)
	}
}

func (fake *FakeBuildStepDelegate) ErroredCallCount() int {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	return len(fake.erroredArgsForCall)
}

func (fake *FakeBuildStepDelegate) ErroredCalls(stub func(lager.Logger, string)) {
	fake.erroredMutex.Lock()
	defer fake.erroredMutex.Unlock()
	fake.ErroredStub = stub
}

func (fake *FakeBuildStepDelegate) ErroredArgsForCall(i int) (lager.Logger, string) {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	argsForCall := fake.erroredArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildStepDelegate) FetchImage(arg1 context.Context, arg2 atc.ImageResource, arg3 atc.VersionedResourceTypes, arg4 bool) (worker.ImageSpec, error) {
	fake.fetchImageMutex.Lock()
	ret, specificReturn := fake.fetchImageReturnsOnCall[len(fake.fetchImageArgsForCall)]
	fake.fetchImageArgsForCall = append(fake.fetchImageArgsForCall, struct {
		arg1 context.Context
		arg2 atc.ImageResource
		arg3 atc.VersionedResourceTypes
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.FetchImageStub
	fakeReturns := fake.fetchImageReturns
	fake.recordInvocation("FetchImage", []interface{}{arg1, arg2, arg3, arg4})
	fake.fetchImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBuildStepDelegate) FetchImageCallCount() int {
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	return len(fake.fetchImageArgsForCall)
}

func (fake *FakeBuildStepDelegate) FetchImageCalls(stub func(context.Context, atc.ImageResource, atc.VersionedResourceTypes, bool) (worker.ImageSpec, error)) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = stub
}

func (fake *FakeBuildStepDelegate) FetchImageArgsForCall(i int) (context.Context, atc.ImageResource, atc.VersionedResourceTypes, bool) {
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	argsForCall := fake.fetchImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeBuildStepDelegate) FetchImageReturns(result1 worker.ImageSpec, result2 error) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = nil
	fake.fetchImageReturns = struct {
		result1 worker.ImageSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildStepDelegate) FetchImageReturnsOnCall(i int, result1 worker.ImageSpec, result2 error) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = nil
	if fake.fetchImageReturnsOnCall == nil {
		fake.fetchImageReturnsOnCall = make(map[int]struct {
			result1 worker.ImageSpec
			result2 error
		})
	}
	fake.fetchImageReturnsOnCall[i] = struct {
		result1 worker.ImageSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildStepDelegate) Finished(arg1 lager.Logger, arg2 bool) {
	fake.finishedMutex.Lock()
	fake.finishedArgsForCall = append(fake.finishedArgsForCall, struct {
		arg1 lager.Logger
		arg2 bool
	}{arg1, arg2})
	stub := fake.FinishedStub
	fake.recordInvocation("Finished", []interface{}{arg1, arg2})
	fake.finishedMutex.Unlock()
	if stub != nil {
		fake.FinishedStub(arg1, arg2)
	}
}

func (fake *FakeBuildStepDelegate) FinishedCallCount() int {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return len(fake.finishedArgsForCall)
}

func (fake *FakeBuildStepDelegate) FinishedCalls(stub func(lager.Logger, bool)) {
	fake.finishedMutex.Lock()
	defer fake.finishedMutex.Unlock()
	fake.FinishedStub = stub
}

func (fake *FakeBuildStepDelegate) FinishedArgsForCall(i int) (lager.Logger, bool) {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	argsForCall := fake.finishedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildStepDelegate) Initializing(arg1 lager.Logger) {
	fake.initializingMutex.Lock()
	fake.initializingArgsForCall = append(fake.initializingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.InitializingStub
	fake.recordInvocation("Initializing", []interface{}{arg1})
	fake.initializingMutex.Unlock()
	if stub != nil {
		fake.InitializingStub(arg1)
	}
}

func (fake *FakeBuildStepDelegate) InitializingCallCount() int {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	return len(fake.initializingArgsForCall)
}

func (fake *FakeBuildStepDelegate) InitializingCalls(stub func(lager.Logger)) {
	fake.initializingMutex.Lock()
	defer fake.initializingMutex.Unlock()
	fake.InitializingStub = stub
}

func (fake *FakeBuildStepDelegate) InitializingArgsForCall(i int) lager.Logger {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	argsForCall := fake.initializingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildStepDelegate) SelectedWorker(arg1 lager.Logger, arg2 string) {
	fake.selectedWorkerMutex.Lock()
	fake.selectedWorkerArgsForCall = append(fake.selectedWorkerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.SelectedWorkerStub
	fake.recordInvocation("SelectedWorker", []interface{}{arg1, arg2})
	fake.selectedWorkerMutex.Unlock()
	if stub != nil {
		fake.SelectedWorkerStub(arg1, arg2)
	}
}

func (fake *FakeBuildStepDelegate) SelectedWorkerCallCount() int {
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	return len(fake.selectedWorkerArgsForCall)
}

func (fake *FakeBuildStepDelegate) SelectedWorkerCalls(stub func(lager.Logger, string)) {
	fake.selectedWorkerMutex.Lock()
	defer fake.selectedWorkerMutex.Unlock()
	fake.SelectedWorkerStub = stub
}

func (fake *FakeBuildStepDelegate) SelectedWorkerArgsForCall(i int) (lager.Logger, string) {
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	argsForCall := fake.selectedWorkerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBuildStepDelegate) StartSpan(arg1 context.Context, arg2 string, arg3 tracing.Attrs) (context.Context, trace.Span) {
	fake.startSpanMutex.Lock()
	ret, specificReturn := fake.startSpanReturnsOnCall[len(fake.startSpanArgsForCall)]
	fake.startSpanArgsForCall = append(fake.startSpanArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 tracing.Attrs
	}{arg1, arg2, arg3})
	stub := fake.StartSpanStub
	fakeReturns := fake.startSpanReturns
	fake.recordInvocation("StartSpan", []interface{}{arg1, arg2, arg3})
	fake.startSpanMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBuildStepDelegate) StartSpanCallCount() int {
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	return len(fake.startSpanArgsForCall)
}

func (fake *FakeBuildStepDelegate) StartSpanCalls(stub func(context.Context, string, tracing.Attrs) (context.Context, trace.Span)) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = stub
}

func (fake *FakeBuildStepDelegate) StartSpanArgsForCall(i int) (context.Context, string, tracing.Attrs) {
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	argsForCall := fake.startSpanArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBuildStepDelegate) StartSpanReturns(result1 context.Context, result2 trace.Span) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = nil
	fake.startSpanReturns = struct {
		result1 context.Context
		result2 trace.Span
	}{result1, result2}
}

func (fake *FakeBuildStepDelegate) StartSpanReturnsOnCall(i int, result1 context.Context, result2 trace.Span) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = nil
	if fake.startSpanReturnsOnCall == nil {
		fake.startSpanReturnsOnCall = make(map[int]struct {
			result1 context.Context
			result2 trace.Span
		})
	}
	fake.startSpanReturnsOnCall[i] = struct {
		result1 context.Context
		result2 trace.Span
	}{result1, result2}
}

func (fake *FakeBuildStepDelegate) Starting(arg1 lager.Logger) {
	fake.startingMutex.Lock()
	fake.startingArgsForCall = append(fake.startingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.StartingStub
	fake.recordInvocation("Starting", []interface{}{arg1})
	fake.startingMutex.Unlock()
	if stub != nil {
		fake.StartingStub(arg1)
	}
}

func (fake *FakeBuildStepDelegate) StartingCallCount() int {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	return len(fake.startingArgsForCall)
}

func (fake *FakeBuildStepDelegate) StartingCalls(stub func(lager.Logger)) {
	fake.startingMutex.Lock()
	defer fake.startingMutex.Unlock()
	fake.StartingStub = stub
}

func (fake *FakeBuildStepDelegate) StartingArgsForCall(i int) lager.Logger {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	argsForCall := fake.startingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildStepDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	ret, specificReturn := fake.stderrReturnsOnCall[len(fake.stderrArgsForCall)]
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct {
	}{})
	stub := fake.StderrStub
	fakeReturns := fake.stderrReturns
	fake.recordInvocation("Stderr", []interface{}{})
	fake.stderrMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildStepDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeBuildStepDelegate) StderrCalls(stub func() io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = stub
}

func (fake *FakeBuildStepDelegate) StderrReturns(result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeBuildStepDelegate) StderrReturnsOnCall(i int, result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	if fake.stderrReturnsOnCall == nil {
		fake.stderrReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stderrReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeBuildStepDelegate) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	ret, specificReturn := fake.stdoutReturnsOnCall[len(fake.stdoutArgsForCall)]
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct {
	}{})
	stub := fake.StdoutStub
	fakeReturns := fake.stdoutReturns
	fake.recordInvocation("Stdout", []interface{}{})
	fake.stdoutMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBuildStepDelegate) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakeBuildStepDelegate) StdoutCalls(stub func() io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = stub
}

func (fake *FakeBuildStepDelegate) StdoutReturns(result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeBuildStepDelegate) StdoutReturnsOnCall(i int, result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	if fake.stdoutReturnsOnCall == nil {
		fake.stdoutReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stdoutReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeBuildStepDelegate) WaitingForWorker(arg1 lager.Logger) {
	fake.waitingForWorkerMutex.Lock()
	fake.waitingForWorkerArgsForCall = append(fake.waitingForWorkerArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.WaitingForWorkerStub
	fake.recordInvocation("WaitingForWorker", []interface{}{arg1})
	fake.waitingForWorkerMutex.Unlock()
	if stub != nil {
		fake.WaitingForWorkerStub(arg1)
	}
}

func (fake *FakeBuildStepDelegate) WaitingForWorkerCallCount() int {
	fake.waitingForWorkerMutex.RLock()
	defer fake.waitingForWorkerMutex.RUnlock()
	return len(fake.waitingForWorkerArgsForCall)
}

func (fake *FakeBuildStepDelegate) WaitingForWorkerCalls(stub func(lager.Logger)) {
	fake.waitingForWorkerMutex.Lock()
	defer fake.waitingForWorkerMutex.Unlock()
	fake.WaitingForWorkerStub = stub
}

func (fake *FakeBuildStepDelegate) WaitingForWorkerArgsForCall(i int) lager.Logger {
	fake.waitingForWorkerMutex.RLock()
	defer fake.waitingForWorkerMutex.RUnlock()
	argsForCall := fake.waitingForWorkerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBuildStepDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.constructAcrossSubstepsMutex.RLock()
	defer fake.constructAcrossSubstepsMutex.RUnlock()
	fake.containerOwnerMutex.RLock()
	defer fake.containerOwnerMutex.RUnlock()
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	fake.waitingForWorkerMutex.RLock()
	defer fake.waitingForWorkerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildStepDelegate) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ exec.BuildStepDelegate = new(FakeBuildStepDelegate)
