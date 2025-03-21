package keycloak

import (
	"context"
	"encoding/json"
	"fmt"

	"shared-wallet-service/infrastructure/config"
	"shared-wallet-service/infrastructure/external/keycloak/dto"
	"shared-wallet-service/infrastructure/external/restful"
)

const (
	_getToken = "/realms/%s/protocol/openid-connect/token"
	_getUsers = "/admin/realms/%s/users"
)

type IKeycloakClient interface {
	GetToken(ctx context.Context) (dto.TokenResponse, error)
	GetUser(ctx context.Context, token string) ([]dto.UserResponse, error)
}

type keycloakClient struct {
	restClient restful.IRestClient
}

func NewKeycloakClient(restClient restful.IRestClient) IKeycloakClient {
	return &keycloakClient{
		restClient: restClient,
	}
}

func (k *keycloakClient) GetToken(ctx context.Context) (dto.TokenResponse, error) {
	var tokenResponse dto.TokenResponse
	url := fmt.Sprintf(fmt.Sprintf("%s%s",
		config.Get().UString("keycloak.url"),
		_getToken), config.Get().UString("keycloak.realm"))

	timeOut := config.Get().UString("keycloak.timeout")

	data := map[string]string{
		"client_id":     config.Get().UString("keycloak.client_id"),
		"client_secret": config.Get().UString("keycloak.client_secret"),
		"grant_type":    "client_credentials",
	}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	formData := k.restClient.EncodeFormData(data)

	resp, err := k.restClient.Post(ctx, url, timeOut, headers, formData)
	if err != nil {
		fmt.Println("error consuming keyclaok api")
		return tokenResponse, err
	}

	err = json.Unmarshal(resp, &tokenResponse)
	if err != nil {
		fmt.Println("error with entity")
		return tokenResponse, nil
	}

	return tokenResponse, nil
}

func (k *keycloakClient) GetUser(ctx context.Context, token string) ([]dto.UserResponse, error) {
	var userResponse []dto.UserResponse
	url := fmt.Sprintf(fmt.Sprintf("%s%s",
		config.Get().UString("keycloak.url"),
		_getUsers), config.Get().UString("keycloak.realm"))
	timeOut := config.Get().UString("keycloak.timeout")

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	resp, err := k.restClient.Get(ctx, url, timeOut, headers)
	if err != nil {
		fmt.Println("error consuming users api")
		return userResponse, err
	}

	err = json.Unmarshal(resp, &userResponse)
	if err != nil {
		fmt.Println("error with entity")
		return userResponse, nil
	}

	return userResponse, nil
}
