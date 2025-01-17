// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	storagea "cloud.google.com/go/storage"
	"github.com/phogolabs/google-cloud/storage"
)

type FakeWriter struct {
	AttrsStub        func() *storagea.ObjectAttrs
	attrsMutex       sync.RWMutex
	attrsArgsForCall []struct {
	}
	attrsReturns struct {
		result1 *storagea.ObjectAttrs
	}
	attrsReturnsOnCall map[int]struct {
		result1 *storagea.ObjectAttrs
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	CloseWithErrorStub        func(error) error
	closeWithErrorMutex       sync.RWMutex
	closeWithErrorArgsForCall []struct {
		arg1 error
	}
	closeWithErrorReturns struct {
		result1 error
	}
	closeWithErrorReturnsOnCall map[int]struct {
		result1 error
	}
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
	SetCRC32CStub        func(uint32)
	setCRC32CMutex       sync.RWMutex
	setCRC32CArgsForCall []struct {
		arg1 uint32
	}
	SetChunkSizeStub        func(int)
	setChunkSizeMutex       sync.RWMutex
	setChunkSizeArgsForCall []struct {
		arg1 int
	}
	SetProgressFuncStub        func(func(int64))
	setProgressFuncMutex       sync.RWMutex
	setProgressFuncArgsForCall []struct {
		arg1 func(int64)
	}
	WriteStub        func([]byte) (int, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	writeReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWriter) Attrs() *storagea.ObjectAttrs {
	fake.attrsMutex.Lock()
	ret, specificReturn := fake.attrsReturnsOnCall[len(fake.attrsArgsForCall)]
	fake.attrsArgsForCall = append(fake.attrsArgsForCall, struct {
	}{})
	stub := fake.AttrsStub
	fakeReturns := fake.attrsReturns
	fake.recordInvocation("Attrs", []interface{}{})
	fake.attrsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWriter) AttrsCallCount() int {
	fake.attrsMutex.RLock()
	defer fake.attrsMutex.RUnlock()
	return len(fake.attrsArgsForCall)
}

func (fake *FakeWriter) AttrsCalls(stub func() *storagea.ObjectAttrs) {
	fake.attrsMutex.Lock()
	defer fake.attrsMutex.Unlock()
	fake.AttrsStub = stub
}

func (fake *FakeWriter) AttrsReturns(result1 *storagea.ObjectAttrs) {
	fake.attrsMutex.Lock()
	defer fake.attrsMutex.Unlock()
	fake.AttrsStub = nil
	fake.attrsReturns = struct {
		result1 *storagea.ObjectAttrs
	}{result1}
}

func (fake *FakeWriter) AttrsReturnsOnCall(i int, result1 *storagea.ObjectAttrs) {
	fake.attrsMutex.Lock()
	defer fake.attrsMutex.Unlock()
	fake.AttrsStub = nil
	if fake.attrsReturnsOnCall == nil {
		fake.attrsReturnsOnCall = make(map[int]struct {
			result1 *storagea.ObjectAttrs
		})
	}
	fake.attrsReturnsOnCall[i] = struct {
		result1 *storagea.ObjectAttrs
	}{result1}
}

func (fake *FakeWriter) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWriter) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeWriter) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeWriter) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWriter) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWriter) CloseWithError(arg1 error) error {
	fake.closeWithErrorMutex.Lock()
	ret, specificReturn := fake.closeWithErrorReturnsOnCall[len(fake.closeWithErrorArgsForCall)]
	fake.closeWithErrorArgsForCall = append(fake.closeWithErrorArgsForCall, struct {
		arg1 error
	}{arg1})
	stub := fake.CloseWithErrorStub
	fakeReturns := fake.closeWithErrorReturns
	fake.recordInvocation("CloseWithError", []interface{}{arg1})
	fake.closeWithErrorMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWriter) CloseWithErrorCallCount() int {
	fake.closeWithErrorMutex.RLock()
	defer fake.closeWithErrorMutex.RUnlock()
	return len(fake.closeWithErrorArgsForCall)
}

func (fake *FakeWriter) CloseWithErrorCalls(stub func(error) error) {
	fake.closeWithErrorMutex.Lock()
	defer fake.closeWithErrorMutex.Unlock()
	fake.CloseWithErrorStub = stub
}

func (fake *FakeWriter) CloseWithErrorArgsForCall(i int) error {
	fake.closeWithErrorMutex.RLock()
	defer fake.closeWithErrorMutex.RUnlock()
	argsForCall := fake.closeWithErrorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWriter) CloseWithErrorReturns(result1 error) {
	fake.closeWithErrorMutex.Lock()
	defer fake.closeWithErrorMutex.Unlock()
	fake.CloseWithErrorStub = nil
	fake.closeWithErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWriter) CloseWithErrorReturnsOnCall(i int, result1 error) {
	fake.closeWithErrorMutex.Lock()
	defer fake.closeWithErrorMutex.Unlock()
	fake.CloseWithErrorStub = nil
	if fake.closeWithErrorReturnsOnCall == nil {
		fake.closeWithErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeWithErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWriter) ObjectAttrs() *storagea.ObjectAttrs {
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

func (fake *FakeWriter) ObjectAttrsCallCount() int {
	fake.objectAttrsMutex.RLock()
	defer fake.objectAttrsMutex.RUnlock()
	return len(fake.objectAttrsArgsForCall)
}

func (fake *FakeWriter) ObjectAttrsCalls(stub func() *storagea.ObjectAttrs) {
	fake.objectAttrsMutex.Lock()
	defer fake.objectAttrsMutex.Unlock()
	fake.ObjectAttrsStub = stub
}

func (fake *FakeWriter) ObjectAttrsReturns(result1 *storagea.ObjectAttrs) {
	fake.objectAttrsMutex.Lock()
	defer fake.objectAttrsMutex.Unlock()
	fake.ObjectAttrsStub = nil
	fake.objectAttrsReturns = struct {
		result1 *storagea.ObjectAttrs
	}{result1}
}

func (fake *FakeWriter) ObjectAttrsReturnsOnCall(i int, result1 *storagea.ObjectAttrs) {
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

func (fake *FakeWriter) SetCRC32C(arg1 uint32) {
	fake.setCRC32CMutex.Lock()
	fake.setCRC32CArgsForCall = append(fake.setCRC32CArgsForCall, struct {
		arg1 uint32
	}{arg1})
	stub := fake.SetCRC32CStub
	fake.recordInvocation("SetCRC32C", []interface{}{arg1})
	fake.setCRC32CMutex.Unlock()
	if stub != nil {
		fake.SetCRC32CStub(arg1)
	}
}

func (fake *FakeWriter) SetCRC32CCallCount() int {
	fake.setCRC32CMutex.RLock()
	defer fake.setCRC32CMutex.RUnlock()
	return len(fake.setCRC32CArgsForCall)
}

func (fake *FakeWriter) SetCRC32CCalls(stub func(uint32)) {
	fake.setCRC32CMutex.Lock()
	defer fake.setCRC32CMutex.Unlock()
	fake.SetCRC32CStub = stub
}

func (fake *FakeWriter) SetCRC32CArgsForCall(i int) uint32 {
	fake.setCRC32CMutex.RLock()
	defer fake.setCRC32CMutex.RUnlock()
	argsForCall := fake.setCRC32CArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWriter) SetChunkSize(arg1 int) {
	fake.setChunkSizeMutex.Lock()
	fake.setChunkSizeArgsForCall = append(fake.setChunkSizeArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.SetChunkSizeStub
	fake.recordInvocation("SetChunkSize", []interface{}{arg1})
	fake.setChunkSizeMutex.Unlock()
	if stub != nil {
		fake.SetChunkSizeStub(arg1)
	}
}

func (fake *FakeWriter) SetChunkSizeCallCount() int {
	fake.setChunkSizeMutex.RLock()
	defer fake.setChunkSizeMutex.RUnlock()
	return len(fake.setChunkSizeArgsForCall)
}

func (fake *FakeWriter) SetChunkSizeCalls(stub func(int)) {
	fake.setChunkSizeMutex.Lock()
	defer fake.setChunkSizeMutex.Unlock()
	fake.SetChunkSizeStub = stub
}

func (fake *FakeWriter) SetChunkSizeArgsForCall(i int) int {
	fake.setChunkSizeMutex.RLock()
	defer fake.setChunkSizeMutex.RUnlock()
	argsForCall := fake.setChunkSizeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWriter) SetProgressFunc(arg1 func(int64)) {
	fake.setProgressFuncMutex.Lock()
	fake.setProgressFuncArgsForCall = append(fake.setProgressFuncArgsForCall, struct {
		arg1 func(int64)
	}{arg1})
	stub := fake.SetProgressFuncStub
	fake.recordInvocation("SetProgressFunc", []interface{}{arg1})
	fake.setProgressFuncMutex.Unlock()
	if stub != nil {
		fake.SetProgressFuncStub(arg1)
	}
}

func (fake *FakeWriter) SetProgressFuncCallCount() int {
	fake.setProgressFuncMutex.RLock()
	defer fake.setProgressFuncMutex.RUnlock()
	return len(fake.setProgressFuncArgsForCall)
}

func (fake *FakeWriter) SetProgressFuncCalls(stub func(func(int64))) {
	fake.setProgressFuncMutex.Lock()
	defer fake.setProgressFuncMutex.Unlock()
	fake.SetProgressFuncStub = stub
}

func (fake *FakeWriter) SetProgressFuncArgsForCall(i int) func(int64) {
	fake.setProgressFuncMutex.RLock()
	defer fake.setProgressFuncMutex.RUnlock()
	argsForCall := fake.setProgressFuncArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWriter) Write(arg1 []byte) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.WriteStub
	fakeReturns := fake.writeReturns
	fake.recordInvocation("Write", []interface{}{arg1Copy})
	fake.writeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWriter) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeWriter) WriteCalls(stub func([]byte) (int, error)) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *FakeWriter) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWriter) WriteReturns(result1 int, result2 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeWriter) WriteReturnsOnCall(i int, result1 int, result2 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attrsMutex.RLock()
	defer fake.attrsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.closeWithErrorMutex.RLock()
	defer fake.closeWithErrorMutex.RUnlock()
	fake.objectAttrsMutex.RLock()
	defer fake.objectAttrsMutex.RUnlock()
	fake.setCRC32CMutex.RLock()
	defer fake.setCRC32CMutex.RUnlock()
	fake.setChunkSizeMutex.RLock()
	defer fake.setChunkSizeMutex.RUnlock()
	fake.setProgressFuncMutex.RLock()
	defer fake.setProgressFuncMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWriter) recordInvocation(key string, args []interface{}) {
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

var _ storage.Writer = new(FakeWriter)
