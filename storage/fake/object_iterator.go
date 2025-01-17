// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	storagea "cloud.google.com/go/storage"
	"github.com/phogolabs/google-cloud/storage"
	"google.golang.org/api/iterator"
)

type FakeObjectIterator struct {
	NextStub        func() (*storagea.ObjectAttrs, error)
	nextMutex       sync.RWMutex
	nextArgsForCall []struct {
	}
	nextReturns struct {
		result1 *storagea.ObjectAttrs
		result2 error
	}
	nextReturnsOnCall map[int]struct {
		result1 *storagea.ObjectAttrs
		result2 error
	}
	PageInfoStub        func() *iterator.PageInfo
	pageInfoMutex       sync.RWMutex
	pageInfoArgsForCall []struct {
	}
	pageInfoReturns struct {
		result1 *iterator.PageInfo
	}
	pageInfoReturnsOnCall map[int]struct {
		result1 *iterator.PageInfo
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeObjectIterator) Next() (*storagea.ObjectAttrs, error) {
	fake.nextMutex.Lock()
	ret, specificReturn := fake.nextReturnsOnCall[len(fake.nextArgsForCall)]
	fake.nextArgsForCall = append(fake.nextArgsForCall, struct {
	}{})
	stub := fake.NextStub
	fakeReturns := fake.nextReturns
	fake.recordInvocation("Next", []interface{}{})
	fake.nextMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeObjectIterator) NextCallCount() int {
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	return len(fake.nextArgsForCall)
}

func (fake *FakeObjectIterator) NextCalls(stub func() (*storagea.ObjectAttrs, error)) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = stub
}

func (fake *FakeObjectIterator) NextReturns(result1 *storagea.ObjectAttrs, result2 error) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	fake.nextReturns = struct {
		result1 *storagea.ObjectAttrs
		result2 error
	}{result1, result2}
}

func (fake *FakeObjectIterator) NextReturnsOnCall(i int, result1 *storagea.ObjectAttrs, result2 error) {
	fake.nextMutex.Lock()
	defer fake.nextMutex.Unlock()
	fake.NextStub = nil
	if fake.nextReturnsOnCall == nil {
		fake.nextReturnsOnCall = make(map[int]struct {
			result1 *storagea.ObjectAttrs
			result2 error
		})
	}
	fake.nextReturnsOnCall[i] = struct {
		result1 *storagea.ObjectAttrs
		result2 error
	}{result1, result2}
}

func (fake *FakeObjectIterator) PageInfo() *iterator.PageInfo {
	fake.pageInfoMutex.Lock()
	ret, specificReturn := fake.pageInfoReturnsOnCall[len(fake.pageInfoArgsForCall)]
	fake.pageInfoArgsForCall = append(fake.pageInfoArgsForCall, struct {
	}{})
	stub := fake.PageInfoStub
	fakeReturns := fake.pageInfoReturns
	fake.recordInvocation("PageInfo", []interface{}{})
	fake.pageInfoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeObjectIterator) PageInfoCallCount() int {
	fake.pageInfoMutex.RLock()
	defer fake.pageInfoMutex.RUnlock()
	return len(fake.pageInfoArgsForCall)
}

func (fake *FakeObjectIterator) PageInfoCalls(stub func() *iterator.PageInfo) {
	fake.pageInfoMutex.Lock()
	defer fake.pageInfoMutex.Unlock()
	fake.PageInfoStub = stub
}

func (fake *FakeObjectIterator) PageInfoReturns(result1 *iterator.PageInfo) {
	fake.pageInfoMutex.Lock()
	defer fake.pageInfoMutex.Unlock()
	fake.PageInfoStub = nil
	fake.pageInfoReturns = struct {
		result1 *iterator.PageInfo
	}{result1}
}

func (fake *FakeObjectIterator) PageInfoReturnsOnCall(i int, result1 *iterator.PageInfo) {
	fake.pageInfoMutex.Lock()
	defer fake.pageInfoMutex.Unlock()
	fake.PageInfoStub = nil
	if fake.pageInfoReturnsOnCall == nil {
		fake.pageInfoReturnsOnCall = make(map[int]struct {
			result1 *iterator.PageInfo
		})
	}
	fake.pageInfoReturnsOnCall[i] = struct {
		result1 *iterator.PageInfo
	}{result1}
}

func (fake *FakeObjectIterator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	fake.pageInfoMutex.RLock()
	defer fake.pageInfoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeObjectIterator) recordInvocation(key string, args []interface{}) {
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

var _ storage.ObjectIterator = new(FakeObjectIterator)
