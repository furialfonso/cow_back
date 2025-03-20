package budget

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"shared-wallet-service/interfaces/dto/request"
	"shared-wallet-service/interfaces/dto/response"
	"shared-wallet-service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockBudgetHandler struct {
	budgetService *mocks.IBudgetService
}

type budgetMocks struct {
	budgetHandler func(f *mockBudgetHandler)
}

func Test_GetAll(t *testing.T) {
	tests := []struct {
		mocks   budgetMocks
		name    string
		expCode int
	}{
		{
			name: "error get budgets",
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {
					f.budgetService.Mock.On("GetAll", mock.Anything).Return([]response.BudgetResponse{}, errors.New("error x"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {
					f.budgetService.Mock.On("GetAll", mock.Anything).Return([]response.BudgetResponse{}, nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockBudgetHandler{
				&mocks.IBudgetService{},
			}
			tc.mocks.budgetHandler(ms)
			handler := NewBudgetHandler(ms.budgetService)
			url := "/budgets"
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
		mocks   budgetMocks
		name    string
		code    string
		expCode int
	}{
		{
			name: "code not sending",
			code: "",
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "error getting budget by id",
			code: "ABC",
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {
					f.budgetService.Mock.On("GetByCode", mock.Anything, "ABC").Return(response.BudgetResponse{}, errors.New("x"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			code: "ABC",
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {
					f.budgetService.Mock.On("GetByCode", mock.Anything, "ABC").Return(response.BudgetResponse{
						Code: "ABC",
					}, nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockBudgetHandler{
				&mocks.IBudgetService{},
			}
			tc.mocks.budgetHandler(ms)
			handler := NewBudgetHandler(ms.budgetService)
			url := "/budgets"
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
		input   interface{}
		mocks   budgetMocks
		name    string
		expCode int
	}{
		{
			name:  "error on input",
			input: "ABC",
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "error creating budget",
			input: request.BudgetRequest{
				Code: "test1",
			},
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {
					f.budgetService.Mock.On("Create", mock.Anything, request.BudgetRequest{
						Code: "test1",
					}).Return(errors.New("error x"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			input: request.BudgetRequest{
				Code: "test1",
			},
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {
					f.budgetService.Mock.On("Create", mock.Anything, request.BudgetRequest{
						Code: "test1",
					}).Return(nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockBudgetHandler{
				&mocks.IBudgetService{},
			}
			tc.mocks.budgetHandler(ms)
			handler := NewBudgetHandler(ms.budgetService)
			url := "/budgets/"
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
		mocks   budgetMocks
		name    string
		code    string
		expCode int
	}{
		{
			name: "code isnt present",
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "code not found",
			code: "test1",
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {
					f.budgetService.Mock.On("Delete", mock.Anything, "test1").Return(errors.New("code not found"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			code: "test1",
			mocks: budgetMocks{
				budgetHandler: func(f *mockBudgetHandler) {
					f.budgetService.Mock.On("Delete", mock.Anything, "test1").Return(nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockBudgetHandler{
				&mocks.IBudgetService{},
			}
			tc.mocks.budgetHandler(ms)
			handler := NewBudgetHandler(ms.budgetService)
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			url := "/budgets"
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
