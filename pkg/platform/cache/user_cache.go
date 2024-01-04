package cache

import (
	"sync"
)

var cacheMap = sync.Map{}

type ICache interface {
	Set(key string, user User)
	Get(key string) (User, bool)
	GetByNickName(nickName string) (User, bool)
}

type cache struct{}

func NewCache() ICache {
	return &cache{}
}

func (u *cache) Set(key string, user User) {
	cacheMap.Store(key, user)
}

func (u *cache) Get(key string) (User, bool) {
	value, exists := cacheMap.Load(key)
	if !exists {
		return User{}, false
	}
	return value.(User), true
}

func (u *cache) GetByNickName(nickName string) (User, bool) {
	var user User
	var exists bool
	cacheMap.Range(func(key, value any) bool {
		if value.(User).NickName == nickName {
			user = value.(User)
			exists = true
			return false
		}
		return true
	})
	return user, exists
}
