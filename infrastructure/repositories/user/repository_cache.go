package user

import (
	"shared-wallet-service/domain/user"
	"shared-wallet-service/domain/user/dto"
	"shared-wallet-service/infrastructure/cache"
	"shared-wallet-service/infrastructure/cache/model"
)

type cacheRepository struct {
	cacheClient cache.ICacheClient
}

func NewCacheRepository(cacheClient cache.ICacheClient) user.ICacheRepository {
	return &cacheRepository{cacheClient: cacheClient}
}

func (r *cacheRepository) SaveUser(user dto.User) {
	r.cacheClient.Set(user.ID, model.User{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		NickName: user.NickName,
	})
}

func (r *cacheRepository) GetUser(key string) (dto.User, bool) {
	user, exist := r.cacheClient.Get(key)
	return user.ModelToDto(), exist
}

func (r *cacheRepository) GetUserByNickName(nickName string) (dto.User, bool) {
	user, exist := r.cacheClient.GetByNickName(nickName)
	return user.ModelToDto(), exist
}
