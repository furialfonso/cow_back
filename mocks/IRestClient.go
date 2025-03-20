// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IRestClient is an autogenerated mock type for the IRestClient type
type IRestClient struct {
	mock.Mock
}

// EncodeFormData provides a mock function with given fields: data
func (_m *IRestClient) EncodeFormData(data map[string]string) string {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for EncodeFormData")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(map[string]string) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, url, timeOut, headers
func (_m *IRestClient) Get(ctx context.Context, url string, timeOut string, headers map[string]string) ([]byte, error) {
	ret := _m.Called(ctx, url, timeOut, headers)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, map[string]string) ([]byte, error)); ok {
		return rf(ctx, url, timeOut, headers)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, map[string]string) []byte); ok {
		r0 = rf(ctx, url, timeOut, headers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, map[string]string) error); ok {
		r1 = rf(ctx, url, timeOut, headers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: ctx, url, timeOut, headers, data
func (_m *IRestClient) Post(ctx context.Context, url string, timeOut string, headers map[string]string, data string) ([]byte, error) {
	ret := _m.Called(ctx, url, timeOut, headers, data)

	if len(ret) == 0 {
		panic("no return value specified for Post")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, map[string]string, string) ([]byte, error)); ok {
		return rf(ctx, url, timeOut, headers, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, map[string]string, string) []byte); ok {
		r0 = rf(ctx, url, timeOut, headers, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, map[string]string, string) error); ok {
		r1 = rf(ctx, url, timeOut, headers, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIRestClient creates a new instance of IRestClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRestClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRestClient {
	mock := &IRestClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
