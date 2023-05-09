package handlers

import (
	"docker-go-project/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockHandler struct {
	mockAirService *mocks.IAirService
}

type airMocks struct {
	airService func(f *mockHandler)
}

func Test_GetAir(t *testing.T) {
	tests := []struct {
		name    string
		mocks   airMocks
		expCode int
	}{
		{
			name: "errors",
			mocks: airMocks{
				airService: func(f *mockHandler) {
					f.mockAirService.Mock.On("GetAirActual").Return("", errors.New("error searching air"))
				},
			},
			expCode: http.StatusInternalServerError,
		},
		{
			name: "full flow",
			mocks: airMocks{
				airService: func(f *mockHandler) {
					f.mockAirService.Mock.On("GetAirActual").Return("local", nil)
				},
			},
			expCode: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ms := &mockHandler{
				&mocks.IAirService{},
			}
			tc.mocks.airService(ms)
			handler := NewAirHandler(ms.mockAirService)
			url := "/get-air"
			_, engine := gin.CreateTestContext(httptest.NewRecorder())
			engine.GET(url, func(ctx *gin.Context) {
				handler.GetAir(ctx)
			})
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, url, nil)
			engine.ServeHTTP(res, req)
			assert.Equal(t, tc.expCode, res.Code)
		})
	}
}
