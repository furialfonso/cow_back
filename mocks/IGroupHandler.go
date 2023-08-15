// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// IGroupHandler is an autogenerated mock type for the IGroupHandler type
type IGroupHandler struct {
	mock.Mock
}

// Create provides a mock function with given fields: c
func (_m *IGroupHandler) Create(c *gin.Context) {
	_m.Called(c)
}

// Delete provides a mock function with given fields: c
func (_m *IGroupHandler) Delete(c *gin.Context) {
	_m.Called(c)
}

// GetGroupByCode provides a mock function with given fields: c
func (_m *IGroupHandler) GetGroupByCode(c *gin.Context) {
	_m.Called(c)
}

// GetGroups provides a mock function with given fields: c
func (_m *IGroupHandler) GetGroups(c *gin.Context) {
	_m.Called(c)
}

// NewIGroupHandler creates a new instance of IGroupHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIGroupHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IGroupHandler {
	mock := &IGroupHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
