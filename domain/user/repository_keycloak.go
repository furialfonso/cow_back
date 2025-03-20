package user

import (
	"context"
	"shared-wallet-service/infrastructure/cache/dto"
)

type IKeycloakRepository interface {
	GetUsers(ctx context.Context) ([]dto.User, error)
}
