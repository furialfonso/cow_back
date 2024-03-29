// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	request "cow_back/api/dto/request"

	mock "github.com/stretchr/testify/mock"

	response "cow_back/api/dto/response"
)

// ITeamService is an autogenerated mock type for the ITeamService type
type ITeamService struct {
	mock.Mock
}

// ComposeTeam provides a mock function with given fields: ctx, code, teamRequest
func (_m *ITeamService) ComposeTeam(ctx context.Context, code string, teamRequest request.TeamRequest) error {
	ret := _m.Called(ctx, code, teamRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, request.TeamRequest) error); ok {
		r0 = rf(ctx, code, teamRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DecomposeTeam provides a mock function with given fields: ctx, code, teamRequest
func (_m *ITeamService) DecomposeTeam(ctx context.Context, code string, teamRequest request.TeamRequest) error {
	ret := _m.Called(ctx, code, teamRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, request.TeamRequest) error); ok {
		r0 = rf(ctx, code, teamRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTeamByGroup provides a mock function with given fields: ctx, code
func (_m *ITeamService) GetTeamByGroup(ctx context.Context, code string) (response.UsersByTeamResponse, error) {
	ret := _m.Called(ctx, code)

	var r0 response.UsersByTeamResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (response.UsersByTeamResponse, error)); ok {
		return rf(ctx, code)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) response.UsersByTeamResponse); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(response.UsersByTeamResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamsByUser provides a mock function with given fields: ctx, userID
func (_m *ITeamService) GetTeamsByUser(ctx context.Context, userID string) (response.TeamsByUserResponse, error) {
	ret := _m.Called(ctx, userID)

	var r0 response.TeamsByUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (response.TeamsByUserResponse, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) response.TeamsByUserResponse); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(response.TeamsByUserResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewITeamService creates a new instance of ITeamService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITeamService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITeamService {
	mock := &ITeamService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
