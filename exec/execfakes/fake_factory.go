// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"sync"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/exec"
	"github.com/concourse/atc/worker"
)

type FakeFactory struct {
	GetStub        func(lager.Logger, int, int, atc.Plan, exec.StepMetadata, db.ContainerMetadata, exec.GetDelegate) exec.StepFactory
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
		arg3 int
		arg4 atc.Plan
		arg5 exec.StepMetadata
		arg6 db.ContainerMetadata
		arg7 exec.GetDelegate
	}
	getReturns struct {
		result1 exec.StepFactory
	}
	getReturnsOnCall map[int]struct {
		result1 exec.StepFactory
	}
	PutStub        func(lager.Logger, int, int, atc.PlanID, exec.StepMetadata, db.ContainerMetadata, exec.PutDelegate, atc.ResourceConfig, atc.Tags, atc.Params, atc.VersionedResourceTypes) exec.StepFactory
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1  lager.Logger
		arg2  int
		arg3  int
		arg4  atc.PlanID
		arg5  exec.StepMetadata
		arg6  db.ContainerMetadata
		arg7  exec.PutDelegate
		arg8  atc.ResourceConfig
		arg9  atc.Tags
		arg10 atc.Params
		arg11 atc.VersionedResourceTypes
	}
	putReturns struct {
		result1 exec.StepFactory
	}
	putReturnsOnCall map[int]struct {
		result1 exec.StepFactory
	}
	DependentGetStub        func(lager.Logger, int, int, atc.PlanID, exec.StepMetadata, worker.ArtifactName, db.ContainerMetadata, exec.GetDelegate, atc.ResourceConfig, atc.Tags, atc.Params, atc.VersionedResourceTypes) exec.StepFactory
	dependentGetMutex       sync.RWMutex
	dependentGetArgsForCall []struct {
		arg1  lager.Logger
		arg2  int
		arg3  int
		arg4  atc.PlanID
		arg5  exec.StepMetadata
		arg6  worker.ArtifactName
		arg7  db.ContainerMetadata
		arg8  exec.GetDelegate
		arg9  atc.ResourceConfig
		arg10 atc.Tags
		arg11 atc.Params
		arg12 atc.VersionedResourceTypes
	}
	dependentGetReturns struct {
		result1 exec.StepFactory
	}
	dependentGetReturnsOnCall map[int]struct {
		result1 exec.StepFactory
	}
	TaskStub        func(lager.Logger, int, int, atc.PlanID, worker.ArtifactName, db.ContainerMetadata, exec.TaskDelegate, exec.Privileged, atc.Tags, exec.TaskConfigSource, atc.VersionedResourceTypes, map[string]string, map[string]string, string, clock.Clock) exec.StepFactory
	taskMutex       sync.RWMutex
	taskArgsForCall []struct {
		arg1  lager.Logger
		arg2  int
		arg3  int
		arg4  atc.PlanID
		arg5  worker.ArtifactName
		arg6  db.ContainerMetadata
		arg7  exec.TaskDelegate
		arg8  exec.Privileged
		arg9  atc.Tags
		arg10 exec.TaskConfigSource
		arg11 atc.VersionedResourceTypes
		arg12 map[string]string
		arg13 map[string]string
		arg14 string
		arg15 clock.Clock
	}
	taskReturns struct {
		result1 exec.StepFactory
	}
	taskReturnsOnCall map[int]struct {
		result1 exec.StepFactory
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFactory) Get(arg1 lager.Logger, arg2 int, arg3 int, arg4 atc.Plan, arg5 exec.StepMetadata, arg6 db.ContainerMetadata, arg7 exec.GetDelegate) exec.StepFactory {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
		arg3 int
		arg4 atc.Plan
		arg5 exec.StepMetadata
		arg6 db.ContainerMetadata
		arg7 exec.GetDelegate
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getReturns.result1
}

func (fake *FakeFactory) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeFactory) GetArgsForCall(i int) (lager.Logger, int, int, atc.Plan, exec.StepMetadata, db.ContainerMetadata, exec.GetDelegate) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2, fake.getArgsForCall[i].arg3, fake.getArgsForCall[i].arg4, fake.getArgsForCall[i].arg5, fake.getArgsForCall[i].arg6, fake.getArgsForCall[i].arg7
}

func (fake *FakeFactory) GetReturns(result1 exec.StepFactory) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) GetReturnsOnCall(i int, result1 exec.StepFactory) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 exec.StepFactory
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Put(arg1 lager.Logger, arg2 int, arg3 int, arg4 atc.PlanID, arg5 exec.StepMetadata, arg6 db.ContainerMetadata, arg7 exec.PutDelegate, arg8 atc.ResourceConfig, arg9 atc.Tags, arg10 atc.Params, arg11 atc.VersionedResourceTypes) exec.StepFactory {
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1  lager.Logger
		arg2  int
		arg3  int
		arg4  atc.PlanID
		arg5  exec.StepMetadata
		arg6  db.ContainerMetadata
		arg7  exec.PutDelegate
		arg8  atc.ResourceConfig
		arg9  atc.Tags
		arg10 atc.Params
		arg11 atc.VersionedResourceTypes
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11})
	fake.recordInvocation("Put", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.putReturns.result1
}

func (fake *FakeFactory) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeFactory) PutArgsForCall(i int) (lager.Logger, int, int, atc.PlanID, exec.StepMetadata, db.ContainerMetadata, exec.PutDelegate, atc.ResourceConfig, atc.Tags, atc.Params, atc.VersionedResourceTypes) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].arg1, fake.putArgsForCall[i].arg2, fake.putArgsForCall[i].arg3, fake.putArgsForCall[i].arg4, fake.putArgsForCall[i].arg5, fake.putArgsForCall[i].arg6, fake.putArgsForCall[i].arg7, fake.putArgsForCall[i].arg8, fake.putArgsForCall[i].arg9, fake.putArgsForCall[i].arg10, fake.putArgsForCall[i].arg11
}

func (fake *FakeFactory) PutReturns(result1 exec.StepFactory) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) PutReturnsOnCall(i int, result1 exec.StepFactory) {
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 exec.StepFactory
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) DependentGet(arg1 lager.Logger, arg2 int, arg3 int, arg4 atc.PlanID, arg5 exec.StepMetadata, arg6 worker.ArtifactName, arg7 db.ContainerMetadata, arg8 exec.GetDelegate, arg9 atc.ResourceConfig, arg10 atc.Tags, arg11 atc.Params, arg12 atc.VersionedResourceTypes) exec.StepFactory {
	fake.dependentGetMutex.Lock()
	ret, specificReturn := fake.dependentGetReturnsOnCall[len(fake.dependentGetArgsForCall)]
	fake.dependentGetArgsForCall = append(fake.dependentGetArgsForCall, struct {
		arg1  lager.Logger
		arg2  int
		arg3  int
		arg4  atc.PlanID
		arg5  exec.StepMetadata
		arg6  worker.ArtifactName
		arg7  db.ContainerMetadata
		arg8  exec.GetDelegate
		arg9  atc.ResourceConfig
		arg10 atc.Tags
		arg11 atc.Params
		arg12 atc.VersionedResourceTypes
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12})
	fake.recordInvocation("DependentGet", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12})
	fake.dependentGetMutex.Unlock()
	if fake.DependentGetStub != nil {
		return fake.DependentGetStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.dependentGetReturns.result1
}

func (fake *FakeFactory) DependentGetCallCount() int {
	fake.dependentGetMutex.RLock()
	defer fake.dependentGetMutex.RUnlock()
	return len(fake.dependentGetArgsForCall)
}

func (fake *FakeFactory) DependentGetArgsForCall(i int) (lager.Logger, int, int, atc.PlanID, exec.StepMetadata, worker.ArtifactName, db.ContainerMetadata, exec.GetDelegate, atc.ResourceConfig, atc.Tags, atc.Params, atc.VersionedResourceTypes) {
	fake.dependentGetMutex.RLock()
	defer fake.dependentGetMutex.RUnlock()
	return fake.dependentGetArgsForCall[i].arg1, fake.dependentGetArgsForCall[i].arg2, fake.dependentGetArgsForCall[i].arg3, fake.dependentGetArgsForCall[i].arg4, fake.dependentGetArgsForCall[i].arg5, fake.dependentGetArgsForCall[i].arg6, fake.dependentGetArgsForCall[i].arg7, fake.dependentGetArgsForCall[i].arg8, fake.dependentGetArgsForCall[i].arg9, fake.dependentGetArgsForCall[i].arg10, fake.dependentGetArgsForCall[i].arg11, fake.dependentGetArgsForCall[i].arg12
}

func (fake *FakeFactory) DependentGetReturns(result1 exec.StepFactory) {
	fake.DependentGetStub = nil
	fake.dependentGetReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) DependentGetReturnsOnCall(i int, result1 exec.StepFactory) {
	fake.DependentGetStub = nil
	if fake.dependentGetReturnsOnCall == nil {
		fake.dependentGetReturnsOnCall = make(map[int]struct {
			result1 exec.StepFactory
		})
	}
	fake.dependentGetReturnsOnCall[i] = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Task(arg1 lager.Logger, arg2 int, arg3 int, arg4 atc.PlanID, arg5 worker.ArtifactName, arg6 db.ContainerMetadata, arg7 exec.TaskDelegate, arg8 exec.Privileged, arg9 atc.Tags, arg10 exec.TaskConfigSource, arg11 atc.VersionedResourceTypes, arg12 map[string]string, arg13 map[string]string, arg14 string, arg15 clock.Clock) exec.StepFactory {
	fake.taskMutex.Lock()
	ret, specificReturn := fake.taskReturnsOnCall[len(fake.taskArgsForCall)]
	fake.taskArgsForCall = append(fake.taskArgsForCall, struct {
		arg1  lager.Logger
		arg2  int
		arg3  int
		arg4  atc.PlanID
		arg5  worker.ArtifactName
		arg6  db.ContainerMetadata
		arg7  exec.TaskDelegate
		arg8  exec.Privileged
		arg9  atc.Tags
		arg10 exec.TaskConfigSource
		arg11 atc.VersionedResourceTypes
		arg12 map[string]string
		arg13 map[string]string
		arg14 string
		arg15 clock.Clock
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15})
	fake.recordInvocation("Task", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15})
	fake.taskMutex.Unlock()
	if fake.TaskStub != nil {
		return fake.TaskStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.taskReturns.result1
}

func (fake *FakeFactory) TaskCallCount() int {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return len(fake.taskArgsForCall)
}

func (fake *FakeFactory) TaskArgsForCall(i int) (lager.Logger, int, int, atc.PlanID, worker.ArtifactName, db.ContainerMetadata, exec.TaskDelegate, exec.Privileged, atc.Tags, exec.TaskConfigSource, atc.VersionedResourceTypes, map[string]string, map[string]string, string, clock.Clock) {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return fake.taskArgsForCall[i].arg1, fake.taskArgsForCall[i].arg2, fake.taskArgsForCall[i].arg3, fake.taskArgsForCall[i].arg4, fake.taskArgsForCall[i].arg5, fake.taskArgsForCall[i].arg6, fake.taskArgsForCall[i].arg7, fake.taskArgsForCall[i].arg8, fake.taskArgsForCall[i].arg9, fake.taskArgsForCall[i].arg10, fake.taskArgsForCall[i].arg11, fake.taskArgsForCall[i].arg12, fake.taskArgsForCall[i].arg13, fake.taskArgsForCall[i].arg14, fake.taskArgsForCall[i].arg15
}

func (fake *FakeFactory) TaskReturns(result1 exec.StepFactory) {
	fake.TaskStub = nil
	fake.taskReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) TaskReturnsOnCall(i int, result1 exec.StepFactory) {
	fake.TaskStub = nil
	if fake.taskReturnsOnCall == nil {
		fake.taskReturnsOnCall = make(map[int]struct {
			result1 exec.StepFactory
		})
	}
	fake.taskReturnsOnCall[i] = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	fake.dependentGetMutex.RLock()
	defer fake.dependentGetMutex.RUnlock()
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFactory) recordInvocation(key string, args []interface{}) {
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

var _ exec.Factory = new(FakeFactory)
