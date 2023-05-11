package services

import (
	"context"
	"docker-go-project/pkg/repository"
)

type IAirService interface {
	GetAirActual(ctx context.Context) (string, error)
}

type airService struct {
	repository repository.IRepository
}

func NewAirService(repository repository.IRepository) IAirService {
	return &airService{
		repository: repository,
	}
}

func (a *airService) GetAirActual(ctx context.Context) (string, error) {
	rs, err := a.repository.Get(ctx)
	if err != nil {
		return "", err
	}
	return rs, nil
}
