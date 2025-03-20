package cache

import (
	"shared-wallet-service/domain/user"
	"shared-wallet-service/infrastructure/cache"
	"shared-wallet-service/infrastructure/cache/dto"
)

type cacheRepository struct {
	cacheClient cache.ICacheClient
}

func NewCacheRepository(cacheClient cache.ICacheClient) user.ICacheRepository {
	return &cacheRepository{cacheClient: cacheClient}
}

func (r *cacheRepository) SaveUser(user dto.User) {
	r.cacheClient.Set(user.ID, user)
}

func (r *cacheRepository) GetUser(key string) (dto.User, bool) {
	return r.cacheClient.Get(key)
}

func (r *cacheRepository) GetUserByNickName(nickName string) (dto.User, bool) {
	return r.cacheClient.GetByNickName(nickName)
}
