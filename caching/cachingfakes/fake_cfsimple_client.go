// Code generated by counterfeiter. DO NOT EDIT.
package cachingfakes

import (
	"io"
	"sync"

	"github.com/cloudfoundry-community/firehose-to-syslog/caching"
)

type FakeCFSimpleClient struct {
	DoGetStub        func(url string) (io.ReadCloser, error)
	doGetMutex       sync.RWMutex
	doGetArgsForCall []struct {
		url string
	}
	doGetReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	doGetReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFSimpleClient) DoGet(url string) (io.ReadCloser, error) {
	fake.doGetMutex.Lock()
	ret, specificReturn := fake.doGetReturnsOnCall[len(fake.doGetArgsForCall)]
	fake.doGetArgsForCall = append(fake.doGetArgsForCall, struct {
		url string
	}{url})
	fake.recordInvocation("DoGet", []interface{}{url})
	fake.doGetMutex.Unlock()
	if fake.DoGetStub != nil {
		return fake.DoGetStub(url)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.doGetReturns.result1, fake.doGetReturns.result2
}

func (fake *FakeCFSimpleClient) DoGetCallCount() int {
	fake.doGetMutex.RLock()
	defer fake.doGetMutex.RUnlock()
	return len(fake.doGetArgsForCall)
}

func (fake *FakeCFSimpleClient) DoGetArgsForCall(i int) string {
	fake.doGetMutex.RLock()
	defer fake.doGetMutex.RUnlock()
	return fake.doGetArgsForCall[i].url
}

func (fake *FakeCFSimpleClient) DoGetReturns(result1 io.ReadCloser, result2 error) {
	fake.DoGetStub = nil
	fake.doGetReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCFSimpleClient) DoGetReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.DoGetStub = nil
	if fake.doGetReturnsOnCall == nil {
		fake.doGetReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.doGetReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCFSimpleClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doGetMutex.RLock()
	defer fake.doGetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFSimpleClient) recordInvocation(key string, args []interface{}) {
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

var _ caching.CFSimpleClient = new(FakeCFSimpleClient)
