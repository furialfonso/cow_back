package user

import (
	"shared-wallet-service/infrastructure/cache/dto"
)

type ICacheRepository interface {
	SaveUser(user dto.User)
	GetUser(key string) (dto.User, bool)
	GetUserByNickName(nickName string) (dto.User, bool)
}
