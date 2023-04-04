// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	storagea "cloud.google.com/go/storage"
	"github.com/phogolabs/google-cloud/storage"
)

type FakeCopier struct {
	ObjectAttrsStub        func() *storagea.ObjectAttrs
	objectAttrsMutex       sync.RWMutex
	objectAttrsArgsForCall []struct {
	}
	objectAttrsReturns struct {
		result1 *storagea.ObjectAttrs
	}
	objectAttrsReturnsOnCall map[int]struct {
		result1 *storagea.ObjectAttrs
	}
	RunStub        func(context.Context) (*storagea.ObjectAttrs, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 context.Context
	}
	runReturns struct {
		result1 *storagea.ObjectAttrs
		result2 error
	}
	runReturnsOnCall map[int]struct {
		result1 *storagea.ObjectAttrs
		result2 error
	}
	SetDestinationKMSKeyNameStub        func(string)
	setDestinationKMSKeyNameMutex       sync.RWMutex
	setDestinationKMSKeyNameArgsForCall []struct {
		arg1 string
	}
	SetProgressFuncStub        func(func(uint64, uint64))
	setProgressFuncMutex       sync.RWMutex
	setProgressFuncArgsForCall []struct {
		arg1 func(uint64, uint64)
	}
	SetRewriteTokenStub        func(string)
	setRewriteTokenMutex       sync.RWMutex
	setRewriteTokenArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCopier) ObjectAttrs() *storagea.ObjectAttrs {
	fake.objectAttrsMutex.Lock()
	ret, specificReturn := fake.objectAttrsReturnsOnCall[len(fake.objectAttrsArgsForCall)]
	fake.objectAttrsArgsForCall = append(fake.objectAttrsArgsForCall, struct {
	}{})
	stub := fake.ObjectAttrsStub
	fakeReturns := fake.objectAttrsReturns
	fake.recordInvocation("ObjectAttrs", []interface{}{})
	fake.objectAttrsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCopier) ObjectAttrsCallCount() int {
	fake.objectAttrsMutex.RLock()
	defer fake.objectAttrsMutex.RUnlock()
	return len(fake.objectAttrsArgsForCall)
}

func (fake *FakeCopier) ObjectAttrsCalls(stub func() *storagea.ObjectAttrs) {
	fake.objectAttrsMutex.Lock()
	defer fake.objectAttrsMutex.Unlock()
	fake.ObjectAttrsStub = stub
}

func (fake *FakeCopier) ObjectAttrsReturns(result1 *storagea.ObjectAttrs) {
	fake.objectAttrsMutex.Lock()
	defer fake.objectAttrsMutex.Unlock()
	fake.ObjectAttrsStub = nil
	fake.objectAttrsReturns = struct {
		result1 *storagea.ObjectAttrs
	}{result1}
}

func (fake *FakeCopier) ObjectAttrsReturnsOnCall(i int, result1 *storagea.ObjectAttrs) {
	fake.objectAttrsMutex.Lock()
	defer fake.objectAttrsMutex.Unlock()
	fake.ObjectAttrsStub = nil
	if fake.objectAttrsReturnsOnCall == nil {
		fake.objectAttrsReturnsOnCall = make(map[int]struct {
			result1 *storagea.ObjectAttrs
		})
	}
	fake.objectAttrsReturnsOnCall[i] = struct {
		result1 *storagea.ObjectAttrs
	}{result1}
}

func (fake *FakeCopier) Run(arg1 context.Context) (*storagea.ObjectAttrs, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.RunStub
	fakeReturns := fake.runReturns
	fake.recordInvocation("Run", []interface{}{arg1})
	fake.runMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCopier) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeCopier) RunCalls(stub func(context.Context) (*storagea.ObjectAttrs, error)) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeCopier) RunArgsForCall(i int) context.Context {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCopier) RunReturns(result1 *storagea.ObjectAttrs, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 *storagea.ObjectAttrs
		result2 error
	}{result1, result2}
}

func (fake *FakeCopier) RunReturnsOnCall(i int, result1 *storagea.ObjectAttrs, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 *storagea.ObjectAttrs
			result2 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 *storagea.ObjectAttrs
		result2 error
	}{result1, result2}
}

func (fake *FakeCopier) SetDestinationKMSKeyName(arg1 string) {
	fake.setDestinationKMSKeyNameMutex.Lock()
	fake.setDestinationKMSKeyNameArgsForCall = append(fake.setDestinationKMSKeyNameArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetDestinationKMSKeyNameStub
	fake.recordInvocation("SetDestinationKMSKeyName", []interface{}{arg1})
	fake.setDestinationKMSKeyNameMutex.Unlock()
	if stub != nil {
		fake.SetDestinationKMSKeyNameStub(arg1)
	}
}

func (fake *FakeCopier) SetDestinationKMSKeyNameCallCount() int {
	fake.setDestinationKMSKeyNameMutex.RLock()
	defer fake.setDestinationKMSKeyNameMutex.RUnlock()
	return len(fake.setDestinationKMSKeyNameArgsForCall)
}

func (fake *FakeCopier) SetDestinationKMSKeyNameCalls(stub func(string)) {
	fake.setDestinationKMSKeyNameMutex.Lock()
	defer fake.setDestinationKMSKeyNameMutex.Unlock()
	fake.SetDestinationKMSKeyNameStub = stub
}

func (fake *FakeCopier) SetDestinationKMSKeyNameArgsForCall(i int) string {
	fake.setDestinationKMSKeyNameMutex.RLock()
	defer fake.setDestinationKMSKeyNameMutex.RUnlock()
	argsForCall := fake.setDestinationKMSKeyNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCopier) SetProgressFunc(arg1 func(uint64, uint64)) {
	fake.setProgressFuncMutex.Lock()
	fake.setProgressFuncArgsForCall = append(fake.setProgressFuncArgsForCall, struct {
		arg1 func(uint64, uint64)
	}{arg1})
	stub := fake.SetProgressFuncStub
	fake.recordInvocation("SetProgressFunc", []interface{}{arg1})
	fake.setProgressFuncMutex.Unlock()
	if stub != nil {
		fake.SetProgressFuncStub(arg1)
	}
}

func (fake *FakeCopier) SetProgressFuncCallCount() int {
	fake.setProgressFuncMutex.RLock()
	defer fake.setProgressFuncMutex.RUnlock()
	return len(fake.setProgressFuncArgsForCall)
}

func (fake *FakeCopier) SetProgressFuncCalls(stub func(func(uint64, uint64))) {
	fake.setProgressFuncMutex.Lock()
	defer fake.setProgressFuncMutex.Unlock()
	fake.SetProgressFuncStub = stub
}

func (fake *FakeCopier) SetProgressFuncArgsForCall(i int) func(uint64, uint64) {
	fake.setProgressFuncMutex.RLock()
	defer fake.setProgressFuncMutex.RUnlock()
	argsForCall := fake.setProgressFuncArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCopier) SetRewriteToken(arg1 string) {
	fake.setRewriteTokenMutex.Lock()
	fake.setRewriteTokenArgsForCall = append(fake.setRewriteTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetRewriteTokenStub
	fake.recordInvocation("SetRewriteToken", []interface{}{arg1})
	fake.setRewriteTokenMutex.Unlock()
	if stub != nil {
		fake.SetRewriteTokenStub(arg1)
	}
}

func (fake *FakeCopier) SetRewriteTokenCallCount() int {
	fake.setRewriteTokenMutex.RLock()
	defer fake.setRewriteTokenMutex.RUnlock()
	return len(fake.setRewriteTokenArgsForCall)
}

func (fake *FakeCopier) SetRewriteTokenCalls(stub func(string)) {
	fake.setRewriteTokenMutex.Lock()
	defer fake.setRewriteTokenMutex.Unlock()
	fake.SetRewriteTokenStub = stub
}

func (fake *FakeCopier) SetRewriteTokenArgsForCall(i int) string {
	fake.setRewriteTokenMutex.RLock()
	defer fake.setRewriteTokenMutex.RUnlock()
	argsForCall := fake.setRewriteTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCopier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.objectAttrsMutex.RLock()
	defer fake.objectAttrsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.setDestinationKMSKeyNameMutex.RLock()
	defer fake.setDestinationKMSKeyNameMutex.RUnlock()
	fake.setProgressFuncMutex.RLock()
	defer fake.setProgressFuncMutex.RUnlock()
	fake.setRewriteTokenMutex.RLock()
	defer fake.setRewriteTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCopier) recordInvocation(key string, args []interface{}) {
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

var _ storage.Copier = new(FakeCopier)