package user

import (
	"context"

	"shared-wallet-service/domain/user/dto"
)

type IKeycloakRepository interface {
	GetUsers(ctx context.Context) ([]dto.User, error)
}
