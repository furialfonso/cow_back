package handlers

import (
	"bytes"
	"docker-go-project/api/dto/request"
	"docker-go-project/api/dto/response"
	"docker-go-project/mocks"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGroupHandler struct {
	groupService *mocks.IGroupService
}

type groupMocks struct {
	groupHandler func(f *mockGroupHandler)
}

func Test_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		mocks   groupMocks
		expCode int
	}{
		{
			name: "error get groups",
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {
					f.groupService.Mock.On("GetAll", mock.Anything).Return([]response.GroupResponse{}, errors.New("error x"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {
					f.groupService.Mock.On("GetAll", mock.Anything).Return([]response.GroupResponse{}, nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockGroupHandler{
				&mocks.IGroupService{},
			}
			tc.mocks.groupHandler(ms)
			handler := NewGroupHandler(ms.groupService)
			url := "/groups"
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			engine.GET(url, func(ctx *gin.Context) {
				handler.GetAll(ctx)
			})
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, url, nil)
			engine.ServeHTTP(res, req)
			assert.Equal(t, tc.expCode, res.Code)
		})
	}
}

func Test_GetByCode(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		mocks   groupMocks
		expCode int
	}{
		{
			name: "code not sending",
			code: "",
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "error getting group by id",
			code: "ABC",
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {
					f.groupService.Mock.On("GetByCode", mock.Anything, "ABC").Return(response.GroupResponse{}, errors.New("x"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			code: "ABC",
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {
					f.groupService.Mock.On("GetByCode", mock.Anything, "ABC").Return(response.GroupResponse{
						Code: "ABC",
					}, nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockGroupHandler{
				&mocks.IGroupService{},
			}
			tc.mocks.groupHandler(ms)
			handler := NewGroupHandler(ms.groupService)
			url := "/groups"
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			engine.GET(url, func(ctx *gin.Context) {
				if tc.code != "" {
					ctx.AddParam("code", tc.code)
				}
				handler.GetByCode(ctx)
			})
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, url, nil)
			engine.ServeHTTP(res, req)
			assert.Equal(t, tc.expCode, res.Code)
		})
	}
}

func Test_Create(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		mocks   groupMocks
		expCode int
	}{
		{
			name:  "error on input",
			input: "ABC",
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "error creating group",
			input: request.GroupRequest{
				Code: "test1",
			},
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {
					f.groupService.Mock.On("Create", mock.Anything, request.GroupRequest{
						Code: "test1",
					}).Return(errors.New("error x"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			input: request.GroupRequest{
				Code: "test1",
			},
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {
					f.groupService.Mock.On("Create", mock.Anything, request.GroupRequest{
						Code: "test1",
					}).Return(nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockGroupHandler{
				&mocks.IGroupService{},
			}
			tc.mocks.groupHandler(ms)
			handler := NewGroupHandler(ms.groupService)
			url := "/groups/"
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			engine.POST(url, func(ctx *gin.Context) {
				handler.Create(ctx)
			})
			res := httptest.NewRecorder()
			b, _ := json.Marshal(tc.input)
			req := httptest.NewRequest(http.MethodPost, url, io.NopCloser(bytes.NewBuffer(b)))
			engine.ServeHTTP(res, req)
			assert.Equal(t, tc.expCode, res.Code)
		})
	}
}

func Test_Delete(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		mocks   groupMocks
		expCode int
	}{
		{
			name: "code isnt present",
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "code not found",
			code: "test1",
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {
					f.groupService.Mock.On("Delete", mock.Anything, "test1").Return(errors.New("code not found"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			code: "test1",
			mocks: groupMocks{
				groupHandler: func(f *mockGroupHandler) {
					f.groupService.Mock.On("Delete", mock.Anything, "test1").Return(nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockGroupHandler{
				&mocks.IGroupService{},
			}
			tc.mocks.groupHandler(ms)
			handler := NewGroupHandler(ms.groupService)
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			url := "/groups"
			engine.DELETE(url, func(ctx *gin.Context) {
				if tc.code != "" {
					ctx.AddParam("code", tc.code)
				}
				handler.Delete(ctx)
			})
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, url, nil)
			engine.ServeHTTP(res, req)
			assert.Equal(t, tc.expCode, res.Code)
		})
	}
}