// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	iread "shared-wallet-service/infrastructure/database/interfaces/read"

	mock "github.com/stretchr/testify/mock"
)

// IReadDataBase is an autogenerated mock type for the IReadDataBase type
type IReadDataBase struct {
	mock.Mock
}

// Read provides a mock function with no fields
func (_m *IReadDataBase) Read() iread.IRead {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 iread.IRead
	if rf, ok := ret.Get(0).(func() iread.IRead); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iread.IRead)
		}
	}

	return r0
}

// NewIReadDataBase creates a new instance of IReadDataBase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIReadDataBase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IReadDataBase {
	mock := &IReadDataBase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
