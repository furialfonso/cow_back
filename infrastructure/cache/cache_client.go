package cache

import (
	"shared-wallet-service/infrastructure/cache/dto"
	"sync"
)

var cacheMap = sync.Map{}

type ICacheClient interface {
	Set(key string, user dto.User)
	Get(key string) (dto.User, bool)
	GetByNickName(nickName string) (dto.User, bool)
}

type cacheClient struct{}

func NewCacheClient() ICacheClient {
	return &cacheClient{}
}

func (u *cacheClient) Set(key string, user dto.User) {
	cacheMap.Store(key, user)
}

func (u *cacheClient) Get(key string) (dto.User, bool) {
	value, exists := cacheMap.Load(key)
	if !exists {
		return dto.User{}, false
	}
	return value.(dto.User), true
}

func (u *cacheClient) GetByNickName(nickName string) (dto.User, bool) {
	var user dto.User
	var exists bool
	cacheMap.Range(func(key, value any) bool {
		if value.(dto.User).NickName == nickName {
			user = value.(dto.User)
			exists = true
			return false
		}
		return true
	})
	return user, exists
}
