package user

import (
	"context"
	"docker-go-project/api/dto/request"
	"docker-go-project/api/dto/response"
	"docker-go-project/mocks"
	"docker-go-project/pkg/repository/user"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserService struct {
	userRepository *mocks.IUserRepository
}

type userMocks struct {
	userService func(f *mockUserService)
}

func Test_GetAll(t *testing.T) {
	tests := []struct {
		name   string
		mocks  userMocks
		outPut []response.UserResponse
		expErr error
	}{
		{
			name: "error get users",
			mocks: userMocks{
				userService: func(f *mockUserService) {
					f.userRepository.Mock.On("GetAll", mock.Anything).Return([]user.User{}, errors.New("error x"))
				},
			},
			expErr: errors.New("error x"),
		},
		{
			name: "full flow",
			mocks: userMocks{
				userService: func(f *mockUserService) {
					f.userRepository.Mock.On("GetAll", mock.Anything).Return([]user.User{
						{
							ID:             1,
							Name:           "diego",
							SecondName:     "manuel",
							LastName:       "fernandez",
							SecondLastName: "marin",
							Email:          "diego@gmail.com",
							NickName:       "diegof",
							CreatedAt:      "2023-05-01T08:00:00",
						},
						{
							ID:             2,
							Name:           "petunia",
							SecondName:     "maria",
							LastName:       "ortiz",
							SecondLastName: "marin",
							Email:          "petunia@gmail.com",
							NickName:       "petuniaf",
							CreatedAt:      "2023-05-01T08:00:00",
						},
					}, nil)
				},
			},
			outPut: []response.UserResponse{
				{
					Name:           "diego",
					SecondName:     "manuel",
					LastName:       "fernandez",
					SecondLastName: "marin",
					Email:          "diego@gmail.com",
					NickName:       "diegof",
				},
				{
					Name:           "petunia",
					SecondName:     "maria",
					LastName:       "ortiz",
					SecondLastName: "marin",
					Email:          "petunia@gmail.com",
					NickName:       "petuniaf",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockUserService{
				userRepository: &mocks.IUserRepository{},
			}
			tc.mocks.userService(m)
			service := NewUserService(m.userRepository)
			users, err := service.GetAll(context.Background())
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.outPut, users)
		})
	}
}

func Test_GetByNickName(t *testing.T) {
	tests := []struct {
		name     string
		nickName string
		mocks    userMocks
		outPut   response.UserResponse
		expErr   error
	}{
		{
			name:     "error get users",
			nickName: "petuniaydiego",
			mocks: userMocks{
				userService: func(f *mockUserService) {
					f.userRepository.Mock.On("GetByNickName", mock.Anything, "petuniaydiego").Return(user.User{}, errors.New("user not found"))
				},
			},
			expErr: errors.New("user not found"),
		},
		{
			name:     "full flow",
			nickName: "petuniaf",
			mocks: userMocks{
				userService: func(f *mockUserService) {
					f.userRepository.Mock.On("GetByNickName", mock.Anything, "petuniaf").Return(user.User{
						ID:             2,
						Name:           "petunia",
						SecondName:     "maria",
						LastName:       "ortiz",
						SecondLastName: "marin",
						Email:          "petunia@gmail.com",
						NickName:       "petuniaf",
						CreatedAt:      "2023-05-01T08:00:00",
					}, nil)
				},
			},
			outPut: response.UserResponse{
				Name:           "petunia",
				SecondName:     "maria",
				LastName:       "ortiz",
				SecondLastName: "marin",
				Email:          "petunia@gmail.com",
				NickName:       "petuniaf",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockUserService{
				userRepository: &mocks.IUserRepository{},
			}
			tc.mocks.userService(m)
			service := NewUserService(m.userRepository)
			users, err := service.GetByNickName(context.Background(), tc.nickName)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.outPut, users)
		})
	}
}

func Test_Create(t *testing.T) {
	tests := []struct {
		name   string
		input  request.UserRequest
		mocks  userMocks
		expErr error
	}{
		{
			name: "error",
			input: request.UserRequest{
				Name:           "petunia",
				SecondName:     "maria",
				LastName:       "ortiz",
				SecondLastName: "marin",
				Email:          "petunia@gmail.com",
				NickName:       "petuniaf",
			},
			mocks: userMocks{
				userService: func(f *mockUserService) {
					f.userRepository.Mock.On("Create", mock.Anything, user.User{
						Name:           "petunia",
						SecondName:     "maria",
						LastName:       "ortiz",
						SecondLastName: "marin",
						Email:          "petunia@gmail.com",
						NickName:       "petuniaf",
					}).Return(int64(0), errors.New("error creating user"))
				},
			},
			expErr: errors.New("error creating user"),
		},
		{
			name: "full flow",
			input: request.UserRequest{
				Name:           "petunia",
				SecondName:     "maria",
				LastName:       "ortiz",
				SecondLastName: "marin",
				Email:          "petunia@gmail.com",
				NickName:       "petuniaf",
			},
			mocks: userMocks{
				userService: func(f *mockUserService) {
					f.userRepository.Mock.On("Create", mock.Anything, user.User{
						Name:           "petunia",
						SecondName:     "maria",
						LastName:       "ortiz",
						SecondLastName: "marin",
						Email:          "petunia@gmail.com",
						NickName:       "petuniaf",
					}).Return(int64(1), nil)
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockUserService{
				userRepository: &mocks.IUserRepository{},
			}
			tc.mocks.userService(m)
			service := NewUserService(m.userRepository)
			err := service.Create(context.Background(), tc.input)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
		})
	}
}
func Test_Delete(t *testing.T) {
	tests := []struct {
		name     string
		nickName string
		mocks    userMocks
		expErr   error
	}{
		{
			name:     "error",
			nickName: "petunia",
			mocks: userMocks{
				userService: func(f *mockUserService) {
					f.userRepository.Mock.On("Delete", mock.Anything, "petunia").Return(errors.New("user not found"))
				},
			},
			expErr: errors.New("user not found"),
		},
		{
			name:     "full flow",
			nickName: "petunia",
			mocks: userMocks{
				userService: func(f *mockUserService) {
					f.userRepository.Mock.On("Delete", mock.Anything, "petunia").Return(nil)
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockUserService{
				userRepository: &mocks.IUserRepository{},
			}
			tc.mocks.userService(m)
			service := NewUserService(m.userRepository)
			err := service.Delete(context.Background(), tc.nickName)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
		})
	}
}
