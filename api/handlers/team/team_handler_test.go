package team

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

type mockTeamHandler struct {
	teamService *mocks.ITeamService
}

type teamMocks struct {
	teamHandler func(f *mockTeamHandler)
}

func Test_GetUsersByGroup(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		mocks   teamMocks
		expCode int
	}{
		{
			name: "group not sending",
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {},
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "error get by group",
			code: "test",
			mocks: teamMocks{
				teamHandler: func(f *mockTeamHandler) {
					f.teamService.Mock.On("GetUsersByGroup", mock.Anything, "test").Return(response.TeamUsersResponse{
						GroupName: "test",
						Users:     []response.UserResponse{},
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
					f.teamService.Mock.On("GetUsersByGroup", mock.Anything, "test").Return(response.TeamUsersResponse{
						GroupName: "test",
						Users: []response.UserResponse{
							{
								Name:           "Diego",
								SecondName:     "Alejandro",
								LastName:       "Malagon",
								SecondLastName: "Martinez",
								Email:          "diego@gmail.com",
								NickName:       "diegom",
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
				&mocks.ITeamService{},
			}
			tc.mocks.teamHandler(ms)
			handler := NewTeamHandler(ms.teamService)
			url := "/team"
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			engine.GET(url, func(ctx *gin.Context) {
				if tc.code != "" {
					ctx.AddParam("code", tc.code)
				}
				handler.GetUsersByGroup(ctx)
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
		name        string
		code        string
		teamRequest interface{}
		mocks       teamMocks
		expCode     int
	}{
		{
			name: "group not sending",
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
				&mocks.ITeamService{},
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
		name        string
		code        string
		teamRequest interface{}
		mocks       teamMocks
		expCode     int
	}{
		{
			name: "group not sending",
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
				&mocks.ITeamService{},
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
