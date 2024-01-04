package user

import (
	"context"
	"encoding/json"
	"fmt"

	"cow_back/pkg/config"
	"cow_back/pkg/platform/cache"
	"cow_back/pkg/platform/restful"
)

const (
	_getUsers = "/users"
)

type IUserService interface {
	UserLoad(ctx context.Context) error
}

type userService struct {
	restfulService restful.IRestfulService
	userCache      cache.ICache
}

func NewUserService(restfulService restful.IRestfulService,
	userCache cache.ICache,
) IUserService {
	return &userService{
		restfulService: restfulService,
		userCache:      userCache,
	}
}

func (us *userService) UserLoad(ctx context.Context) error {
	var users []cache.User
	url := fmt.Sprintf("%s%s", config.Get().UString("users-api.url"), _getUsers)
	timeOut := config.Get().UString("users-api.timeout")
	resp, err := us.restfulService.Get(ctx, url, timeOut)
	if err != nil {
		fmt.Println("error consuming users api")
		return err
	}

	err = json.Unmarshal(resp, &users)
	if err != nil {
		fmt.Println("error with entity")
		return err
	}

	totalUsers := 0
	for _, user := range users {
		us.userCache.Set(user.ID, user)
		totalUsers++
	}
	fmt.Println("cache loaded successfully")
	fmt.Println("total users:", totalUsers)
	return nil
}
