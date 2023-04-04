// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	storagea "cloud.google.com/go/storage"
	"github.com/phogolabs/google-cloud/storage"
)

type FakeACLHandle struct {
	DeleteStub        func(context.Context, storagea.ACLEntity) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 storagea.ACLEntity
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(context.Context) ([]storagea.ACLRule, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
	}
	listReturns struct {
		result1 []storagea.ACLRule
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []storagea.ACLRule
		result2 error
	}
	SetStub        func(context.Context, storagea.ACLEntity, storagea.ACLRole) error
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 context.Context
		arg2 storagea.ACLEntity
		arg3 storagea.ACLRole
	}
	setReturns struct {
		result1 error
	}
	setReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeACLHandle) Delete(arg1 context.Context, arg2 storagea.ACLEntity) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 storagea.ACLEntity
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeACLHandle) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeACLHandle) DeleteCalls(stub func(context.Context, storagea.ACLEntity) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeACLHandle) DeleteArgsForCall(i int) (context.Context, storagea.ACLEntity) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeACLHandle) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeACLHandle) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeACLHandle) List(arg1 context.Context) ([]storagea.ACLRule, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeACLHandle) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeACLHandle) ListCalls(stub func(context.Context) ([]storagea.ACLRule, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeACLHandle) ListArgsForCall(i int) context.Context {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeACLHandle) ListReturns(result1 []storagea.ACLRule, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []storagea.ACLRule
		result2 error
	}{result1, result2}
}

func (fake *FakeACLHandle) ListReturnsOnCall(i int, result1 []storagea.ACLRule, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []storagea.ACLRule
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []storagea.ACLRule
		result2 error
	}{result1, result2}
}

func (fake *FakeACLHandle) Set(arg1 context.Context, arg2 storagea.ACLEntity, arg3 storagea.ACLRole) error {
	fake.setMutex.Lock()
	ret, specificReturn := fake.setReturnsOnCall[len(fake.setArgsForCall)]
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 context.Context
		arg2 storagea.ACLEntity
		arg3 storagea.ACLRole
	}{arg1, arg2, arg3})
	stub := fake.SetStub
	fakeReturns := fake.setReturns
	fake.recordInvocation("Set", []interface{}{arg1, arg2, arg3})
	fake.setMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeACLHandle) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeACLHandle) SetCalls(stub func(context.Context, storagea.ACLEntity, storagea.ACLRole) error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = stub
}

func (fake *FakeACLHandle) SetArgsForCall(i int) (context.Context, storagea.ACLEntity, storagea.ACLRole) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	argsForCall := fake.setArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeACLHandle) SetReturns(result1 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	fake.setReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeACLHandle) SetReturnsOnCall(i int, result1 error) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = nil
	if fake.setReturnsOnCall == nil {
		fake.setReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeACLHandle) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeACLHandle) recordInvocation(key string, args []interface{}) {
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

var _ storage.ACLHandle = new(FakeACLHandle)