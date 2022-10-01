// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"sync"

	"github.com/petewall/eink-radiator-image-source-blank/v2/internal"
)

type FakeImageContextMaker struct {
	Stub        func(int, int) internal.ImageContext
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 int
		arg2 int
	}
	returns struct {
		result1 internal.ImageContext
	}
	returnsOnCall map[int]struct {
		result1 internal.ImageContext
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageContextMaker) Spy(arg1 int, arg2 int) internal.ImageContext {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	stub := fake.Stub
	returns := fake.returns
	fake.recordInvocation("ImageContextMaker", []interface{}{arg1, arg2})
	fake.mutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return returns.result1
}

func (fake *FakeImageContextMaker) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeImageContextMaker) Calls(stub func(int, int) internal.ImageContext) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *FakeImageContextMaker) ArgsForCall(i int) (int, int) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1, fake.argsForCall[i].arg2
}

func (fake *FakeImageContextMaker) Returns(result1 internal.ImageContext) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 internal.ImageContext
	}{result1}
}

func (fake *FakeImageContextMaker) ReturnsOnCall(i int, result1 internal.ImageContext) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 internal.ImageContext
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 internal.ImageContext
	}{result1}
}

func (fake *FakeImageContextMaker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageContextMaker) recordInvocation(key string, args []interface{}) {
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

var _ internal.ImageContextMaker = new(FakeImageContextMaker).Spy
