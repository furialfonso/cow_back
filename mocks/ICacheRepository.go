// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	dto "shared-wallet-service/infrastructure/cache/dto"

	mock "github.com/stretchr/testify/mock"
)

// ICacheRepository is an autogenerated mock type for the ICacheRepository type
type ICacheRepository struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: key
func (_m *ICacheRepository) GetUser(key string) (dto.User, bool) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 dto.User
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (dto.User, bool)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) dto.User); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(dto.User)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetUserByNickName provides a mock function with given fields: nickName
func (_m *ICacheRepository) GetUserByNickName(nickName string) (dto.User, bool) {
	ret := _m.Called(nickName)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByNickName")
	}

	var r0 dto.User
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (dto.User, bool)); ok {
		return rf(nickName)
	}
	if rf, ok := ret.Get(0).(func(string) dto.User); ok {
		r0 = rf(nickName)
	} else {
		r0 = ret.Get(0).(dto.User)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(nickName)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: _a0
func (_m *ICacheRepository) SaveUser(_a0 dto.User) {
	_m.Called(_a0)
}

// NewICacheRepository creates a new instance of ICacheRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICacheRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICacheRepository {
	mock := &ICacheRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
