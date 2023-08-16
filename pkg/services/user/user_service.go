package user

import (
	"context"
	"docker-go-project/api/dto/request"
)

type IUserService interface {
	GetAll(ctx context.Context) error
	GetByCode(ctx context.Context, code string) error
	Create(ctx context.Context, groupDTO request.GroupDTO) error
	Delete(ctx context.Context, code string) error
}

type userService struct {
	// groupRepository repository.IGroupRepository
}

func NewUserService() IUserService {
	return &userService{}
}

func (us *userService) GetAll(ctx context.Context) error {
	return nil
}

func (us *userService) GetByCode(ctx context.Context, code string) error {
	return nil
}

func (us *userService) Create(ctx context.Context, groupDTO request.GroupDTO) error {
	return nil
}

func (us *userService) Delete(ctx context.Context, code string) error {
	return nil
}
