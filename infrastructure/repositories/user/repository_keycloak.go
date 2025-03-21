package user

import (
	"context"

	"shared-wallet-service/domain/user"
	"shared-wallet-service/domain/user/dto"

	"shared-wallet-service/infrastructure/external/keycloak"
)

type keycloakRepository struct {
	keycloakClient keycloak.IKeycloakClient
}

func NewKeycloakRepository(keycloakClient keycloak.IKeycloakClient) user.IKeycloakRepository {
	return &keycloakRepository{
		keycloakClient: keycloakClient,
	}
}

func (r *keycloakRepository) GetUsers(ctx context.Context) ([]dto.User, error) {
	token, err := r.keycloakClient.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	users, err := r.keycloakClient.GetUser(ctx, token.AccessToken)
	if err != nil {
		return nil, err
	}

	var domainUsers []dto.User
	for _, u := range users {
		domainUsers = append(domainUsers, dto.User{
			ID:       u.ID,
			Name:     u.Name,
			LastName: u.LastName,
			Email:    u.Email,
			NickName: u.NickName,
		})
	}
	return domainUsers, nil
}
