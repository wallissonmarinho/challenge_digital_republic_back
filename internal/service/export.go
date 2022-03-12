package service

import (
	"github.com/go-kit/log"
	Pintura "github.com/wallissonmarinho/challenge_digital_republic/internal/service/pintura"
)

type ServiceFactory interface {
	Pintura() Pintura.PinturaService
}

type serviceFactory struct {
	PinturaService Pintura.PinturaService
}

func NewServiceFactory(logger log.Logger) ServiceFactory {
	return &serviceFactory{
		PinturaService: Pintura.NewPinturaService(logger),
	}
}

func (s *serviceFactory) Pintura() Pintura.PinturaService {
	return s.PinturaService
}
