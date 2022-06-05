package services

import "github.com/Heroin-lab/taxi_service.git/pkg/repositories"

type Authorization interface {
}

type Services struct {
	Authorization
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{}
}
