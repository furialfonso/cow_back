package user

import (
	"context"
	"fmt"

	"shared-wallet-service/domain/user"
)

const (
	_getUsers = "/users"
)

type IUserUseCase interface {
	UserLoad(ctx context.Context) error
}

type userUseCase struct {
	keycloakRepository user.IKeycloakRepository
	cacheRepository    user.ICacheRepository
}

func NewUserUseCase(keycloakRepository user.IKeycloakRepository,
	cacheRepository user.ICacheRepository,
) IUserUseCase {
	return &userUseCase{
		keycloakRepository: keycloakRepository,
		cacheRepository:    cacheRepository,
	}
}

func (uc *userUseCase) UserLoad(ctx context.Context) error {
	users, err := uc.keycloakRepository.GetUsers(ctx)
	if err != nil {
		fmt.Println("error getting users")
		return err
	}

	for _, user := range users {
		uc.cacheRepository.SaveUser(user)
	}
	fmt.Println("cache loaded successfully")
	return nil
}
