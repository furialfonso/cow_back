package user

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"cow_back/mocks"
	"cow_back/pkg/platform/cache"

	"github.com/go-playground/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserService struct {
	restfulService *mocks.IRestfulService
	userCache      *mocks.ICache
}

type userMocks struct {
	userService func(f *mockUserService)
}

func Test_UserLoad(t *testing.T) {
	tests := []struct {
		expErr error
		mocks  userMocks
		name   string
	}{
		{
			name: "error get users",
			mocks: userMocks{
				func(f *mockUserService) {
					users := []cache.User{}
					b, _ := json.Marshal(users)
					f.restfulService.Mock.On("Get", mock.Anything, "http://localhost:9002/users", "5s").Return(b, errors.New("some error"))
				},
			},
			expErr: errors.New("some error"),
		},
		{
			name: "full flow",
			mocks: userMocks{
				func(f *mockUserService) {
					users := []cache.User{
						{
							ID:       "1234",
							Name:     "test",
							LastName: "test",
							Email:    "test@gmail.com",
							NickName: "testy",
						},
					}
					b, _ := json.Marshal(users)
					f.restfulService.Mock.On("Get", mock.Anything, "http://localhost:9002/users", "5s").Return(b, nil)
					f.userCache.Mock.On("Set", "1234", cache.User{
						ID:       "1234",
						Name:     "test",
						LastName: "test",
						Email:    "test@gmail.com",
						NickName: "testy",
					})
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockUserService{
				restfulService: &mocks.IRestfulService{},
				userCache:      &mocks.ICache{},
			}
			tc.mocks.userService(m)
			service := NewUserService(m.restfulService, m.userCache)
			err := service.UserLoad(context.Background())
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
		})
	}
}
