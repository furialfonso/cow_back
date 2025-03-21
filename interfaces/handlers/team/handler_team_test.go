package team

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
	"github.com/go-playground/assert"
	"github.com/stretchr/testify/mock"
)

type mockTeamHandler struct {
	teamService *mocks.ITeamUseCase
}

type teamMocks struct {
	teamHandler func(f *mockTeamHandler)
}

func Test_GetTeamByBudget(t *testing.T) {
	tests := []struct {
		mocks   teamMocks
		name    string
		code    string
		expCode int
	}{
		{
			name: "budget not sending",
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "error get by budget",
			code: "test",
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {
					f.teamService.Mock.On("GetTeamByBudget", mock.Anything, "test").Return(response.UsersByTeamResponse{
						BudgetName: "test",
						Users:      []response.UserResponse{},
					}, errors.New("error searching team"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			code: "test",
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {
					f.teamService.Mock.On("GetTeamByBudget", mock.Anything, "test").Return(response.UsersByTeamResponse{
						BudgetName: "test",
						Users: []response.UserResponse{
							{
								ID:       "1",
								Name:     "Diego",
								LastName: "Malagon",
								Email:    "diego@gmail.com",
								NickName: "diegom",
							},
						},
					}, nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockTeamHandler{
				&mocks.ITeamUseCase{},
			}
			tc.mocks.teamHandler(ms)
			handler := NewTeamHandler(ms.teamService)
			url := "/team"
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			engine.GET(url, func(ctx *gin.Context) {
				if tc.code != "" {
					ctx.AddParam("code", tc.code)
				}
				handler.GetTeamByBudget(ctx)
			})
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, url, nil)
			engine.ServeHTTP(res, req)
			assert.Equal(t, tc.expCode, res.Code)
		})
	}
}

func Test_ComposeTeam(t *testing.T) {
	tests := []struct {
		teamRequest interface{}
		mocks       teamMocks
		name        string
		code        string
		expCode     int
	}{
		{
			name: "budget not sending",
			teamRequest: request.TeamRequest{
				Users: []string{"diego", "petunia"},
			},
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name:        "users empty",
			code:        "test",
			teamRequest: "",
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "error saving users",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diego", "petunia"},
			},
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {
					f.teamService.Mock.On("ComposeTeam", mock.Anything, "test", request.TeamRequest{
						Users: []string{"diego", "petunia"},
					}).Return(errors.New("error saving users"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diego", "petunia"},
			},
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {
					f.teamService.Mock.On("ComposeTeam", mock.Anything, "test", request.TeamRequest{
						Users: []string{"diego", "petunia"},
					}).Return(nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockTeamHandler{
				&mocks.ITeamUseCase{},
			}
			tc.mocks.teamHandler(ms)
			handler := NewTeamHandler(ms.teamService)
			url := "/team"
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			engine.POST(url, func(ctx *gin.Context) {
				if tc.code != "" {
					ctx.AddParam("code", tc.code)
				}
				handler.ComposeTeam(ctx)
			})
			res := httptest.NewRecorder()
			b, _ := json.Marshal(tc.teamRequest)
			req := httptest.NewRequest(http.MethodPost, url, io.NopCloser(bytes.NewBuffer(b)))
			engine.ServeHTTP(res, req)
			assert.Equal(t, tc.expCode, res.Code)
		})
	}
}

func Test_Decompose(t *testing.T) {
	tests := []struct {
		teamRequest interface{}
		mocks       teamMocks
		name        string
		code        string
		expCode     int
	}{
		{
			name: "budget not sending",
			teamRequest: request.TeamRequest{
				Users: []string{"diego", "petunia"},
			},
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name:        "users empty",
			code:        "test",
			teamRequest: "",
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "error saving users",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diego", "petunia"},
			},
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {
					f.teamService.Mock.On("DecomposeTeam", mock.Anything, "test", request.TeamRequest{
						Users: []string{"diego", "petunia"},
					}).Return(errors.New("error saving users"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			code: "test",
			teamRequest: request.TeamRequest{
				Users: []string{"diego", "petunia"},
			},
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {
					f.teamService.Mock.On("DecomposeTeam", mock.Anything, "test", request.TeamRequest{
						Users: []string{"diego", "petunia"},
					}).Return(nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockTeamHandler{
				&mocks.ITeamUseCase{},
			}
			tc.mocks.teamHandler(ms)
			handler := NewTeamHandler(ms.teamService)
			url := "/team"
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			engine.DELETE(url, func(ctx *gin.Context) {
				if tc.code != "" {
					ctx.AddParam("code", tc.code)
				}
				handler.DecomposeTeam(ctx)
			})
			res := httptest.NewRecorder()
			b, _ := json.Marshal(tc.teamRequest)
			req := httptest.NewRequest(http.MethodDelete, url, io.NopCloser(bytes.NewBuffer(b)))
			engine.ServeHTTP(res, req)
			assert.Equal(t, tc.expCode, res.Code)
		})
	}
}
