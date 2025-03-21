package cache

import (
	"sync"

	"shared-wallet-service/infrastructure/cache/model"
)

var cacheMap = sync.Map{}

type ICacheClient interface {
	Set(key string, user model.User)
	Get(key string) (model.User, bool)
	GetByNickName(nickName string) (model.User, bool)
}

type cacheClient struct{}

func NewCacheClient() ICacheClient {
	return &cacheClient{}
}

func (u *cacheClient) Set(key string, user model.User) {
	cacheMap.Store(key, user)
}

func (u *cacheClient) Get(key string) (model.User, bool) {
	value, exists := cacheMap.Load(key)
	if !exists {
		return model.User{}, false
	}
	return value.(model.User), true
}

func (u *cacheClient) GetByNickName(nickName string) (model.User, bool) {
	var user model.User
	var exists bool
	cacheMap.Range(func(key, value any) bool {
		if value.(model.User).NickName == nickName {
			user = value.(model.User)
			exists = true
			return false
		}
		return true
	})
	return user, exists
}
