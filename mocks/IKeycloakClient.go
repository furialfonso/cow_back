// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "shared-wallet-service/infrastructure/external/keycloak/dto"

	mock "github.com/stretchr/testify/mock"
)

// IKeycloakClient is an autogenerated mock type for the IKeycloakClient type
type IKeycloakClient struct {
	mock.Mock
}

// GetToken provides a mock function with given fields: ctx
func (_m *IKeycloakClient) GetToken(ctx context.Context) (dto.TokenResponse, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetToken")
	}

	var r0 dto.TokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (dto.TokenResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) dto.TokenResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(dto.TokenResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, token
func (_m *IKeycloakClient) GetUser(ctx context.Context, token string) ([]dto.UserResponse, error) {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 []dto.UserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]dto.UserResponse, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []dto.UserResponse); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.UserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIKeycloakClient creates a new instance of IKeycloakClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIKeycloakClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *IKeycloakClient {
	mock := &IKeycloakClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
