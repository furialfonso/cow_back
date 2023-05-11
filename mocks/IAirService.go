// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IAirService is an autogenerated mock type for the IAirService type
type IAirService struct {
	mock.Mock
}

// GetAirActual provides a mock function with given fields: ctx
func (_m *IAirService) GetAirActual(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIAirService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIAirService creates a new instance of IAirService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIAirService(t mockConstructorTestingTNewIAirService) *IAirService {
	mock := &IAirService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
