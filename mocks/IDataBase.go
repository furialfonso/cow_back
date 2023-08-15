// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	sql "docker-go-project/pkg/platform/sql"

	mock "github.com/stretchr/testify/mock"
)

// IDataBase is an autogenerated mock type for the IDataBase type
type IDataBase struct {
	mock.Mock
}

// GetRead provides a mock function with given fields:
func (_m *IDataBase) GetRead() sql.IRead {
	ret := _m.Called()

	var r0 sql.IRead
	if rf, ok := ret.Get(0).(func() sql.IRead); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.IRead)
		}
	}

	return r0
}

// GetWrite provides a mock function with given fields:
func (_m *IDataBase) GetWrite() sql.IWrite {
	ret := _m.Called()

	var r0 sql.IWrite
	if rf, ok := ret.Get(0).(func() sql.IWrite); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.IWrite)
		}
	}

	return r0
}

// NewIDataBase creates a new instance of IDataBase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDataBase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDataBase {
	mock := &IDataBase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
