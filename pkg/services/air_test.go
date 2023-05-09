package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockService struct{}

type airMocks struct {
	service func(f *mockService)
}

func Test_GetAirActual(t *testing.T) {
	tests := []struct {
		name   string
		mocks  airMocks
		airExp string
		expErr error
	}{
		{
			name: "full flow",
			mocks: airMocks{
				service: func(f *mockService) {},
			},
			airExp: "LOCAL",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockService{}
			tc.mocks.service(m)
			service := NewAirService()
			air, err := service.GetAirActual()
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.airExp, air)
		})
	}
}
