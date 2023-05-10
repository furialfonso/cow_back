package services

import (
	"docker-go-project/pkg/repository"
)

type IAirService interface {
	GetAirActual() (string, error)
}

type airService struct {
	repository repository.IRepository
}

func NewAirService(repository repository.IRepository) IAirService {
	return &airService{
		repository: repository,
	}
}

func (a *airService) GetAirActual() (string, error) {
	rs, err := a.repository.Get()
	if err != nil {
		return "", err
	}
	return rs, nil
}
