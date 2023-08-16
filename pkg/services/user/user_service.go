package user

import (
	"context"
	"docker-go-project/api/dto/request"
	"docker-go-project/api/dto/response"
	"docker-go-project/pkg/repository/user"
)

type IUserService interface {
	GetAll(ctx context.Context) ([]response.UserResponse, error)
	GetByNickName(ctx context.Context, nickName string) (response.UserResponse, error)
	Create(ctx context.Context, userRequest request.UserRequest) error
	Delete(ctx context.Context, nickName string) error
}

type userService struct {
	userRepository user.IUserRepository
}

func NewUserService(userRepository user.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (us *userService) GetAll(ctx context.Context) ([]response.UserResponse, error) {
	var userResponse []response.UserResponse
	users, err := us.userRepository.GetAll(ctx)
	if err != nil {
		return userResponse, err
	}
	for _, user := range users {
		userResponse = append(userResponse, response.UserResponse{
			Name:           user.Name,
			SecondName:     user.SecondName,
			LastName:       user.LastName,
			SecondLastName: user.SecondLastName,
			Email:          user.Email,
			NickName:       user.NickName,
		})
	}

	return userResponse, nil
}

func (us *userService) GetByNickName(ctx context.Context, nickName string) (response.UserResponse, error) {
	user, err := us.userRepository.GetByNickName(ctx, nickName)
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.UserResponse{
		Name:           user.Name,
		SecondName:     user.SecondName,
		LastName:       user.LastName,
		SecondLastName: user.SecondLastName,
		Email:          user.Email,
		NickName:       user.NickName,
	}, nil
}

func (us *userService) Create(ctx context.Context, userRequest request.UserRequest) error {
	_, err := us.userRepository.Create(ctx, user.User{
		Name:           userRequest.Name,
		SecondName:     userRequest.SecondName,
		LastName:       userRequest.LastName,
		SecondLastName: userRequest.SecondLastName,
		Email:          userRequest.Email,
		NickName:       userRequest.NickName,
	})
	if err != nil {
		return err
	}
	return nil
}

func (us *userService) Delete(ctx context.Context, nickName string) error {
	err := us.userRepository.Delete(ctx, nickName)
	if err != nil {
		return err
	}
	return nil
}
