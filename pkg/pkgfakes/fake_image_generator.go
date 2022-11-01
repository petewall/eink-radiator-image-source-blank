// Code generated by counterfeiter. DO NOT EDIT.
package pkgfakes

import (
	"image"
	"sync"

	"github.com/petewall/eink-radiator-image-source-blank/pkg"
)

type FakeImageGenerator struct {
	GenerateImageStub        func(int, int) image.Image
	generateImageMutex       sync.RWMutex
	generateImageArgsForCall []struct {
		arg1 int
		arg2 int
	}
	generateImageReturns struct {
		result1 image.Image
	}
	generateImageReturnsOnCall map[int]struct {
		result1 image.Image
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageGenerator) GenerateImage(arg1 int, arg2 int) image.Image {
	fake.generateImageMutex.Lock()
	ret, specificReturn := fake.generateImageReturnsOnCall[len(fake.generateImageArgsForCall)]
	fake.generateImageArgsForCall = append(fake.generateImageArgsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	stub := fake.GenerateImageStub
	fakeReturns := fake.generateImageReturns
	fake.recordInvocation("GenerateImage", []interface{}{arg1, arg2})
	fake.generateImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImageGenerator) GenerateImageCallCount() int {
	fake.generateImageMutex.RLock()
	defer fake.generateImageMutex.RUnlock()
	return len(fake.generateImageArgsForCall)
}

func (fake *FakeImageGenerator) GenerateImageCalls(stub func(int, int) image.Image) {
	fake.generateImageMutex.Lock()
	defer fake.generateImageMutex.Unlock()
	fake.GenerateImageStub = stub
}

func (fake *FakeImageGenerator) GenerateImageArgsForCall(i int) (int, int) {
	fake.generateImageMutex.RLock()
	defer fake.generateImageMutex.RUnlock()
	argsForCall := fake.generateImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageGenerator) GenerateImageReturns(result1 image.Image) {
	fake.generateImageMutex.Lock()
	defer fake.generateImageMutex.Unlock()
	fake.GenerateImageStub = nil
	fake.generateImageReturns = struct {
		result1 image.Image
	}{result1}
}

func (fake *FakeImageGenerator) GenerateImageReturnsOnCall(i int, result1 image.Image) {
	fake.generateImageMutex.Lock()
	defer fake.generateImageMutex.Unlock()
	fake.GenerateImageStub = nil
	if fake.generateImageReturnsOnCall == nil {
		fake.generateImageReturnsOnCall = make(map[int]struct {
			result1 image.Image
		})
	}
	fake.generateImageReturnsOnCall[i] = struct {
		result1 image.Image
	}{result1}
}

func (fake *FakeImageGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateImageMutex.RLock()
	defer fake.generateImageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageGenerator) recordInvocation(key string, args []interface{}) {
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

var _ pkg.ImageGenerator = new(FakeImageGenerator)
