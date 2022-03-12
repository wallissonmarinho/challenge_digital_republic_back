package service

import (
	"github.com/go-kit/log"
)

type ServiceFactory interface {
}

type serviceFactory struct {
}

func NewServiceFactory(logger log.Logger) ServiceFactory {
	return &serviceFactory{}
}
