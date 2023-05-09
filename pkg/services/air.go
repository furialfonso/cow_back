package services

import (
	"docker-go-project/pkg/config"
	"errors"
)

type IAirService interface {
	GetAirActual() (string, error)
}

type airService struct{}

func NewAirService() IAirService {
	return &airService{}
}

func (a *airService) GetAirActual() (string, error) {
	sc := config.Get().UString("name")
	if sc == "" {
		return "", errors.New("air not found")
	}
	return sc, nil
}
