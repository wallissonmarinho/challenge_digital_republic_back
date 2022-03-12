package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/service"
)

// Endpoints holds all Go kit endpoints for the Order service.
type Endpoints struct {
	Health                                endpoint.Endpoint
	ObterQuantidadeDeLatasdeTintaEndpoint endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Order service.
func MakeEndpoints(s service.ServiceFactory, logger log.Logger) Endpoints {
	return Endpoints{
		Health:                                makeHealthEndpoint(s, logger),
		ObterQuantidadeDeLatasdeTintaEndpoint: makeObterQuantidadeDeLatasdeTintaEndpoint(s, logger),
	}
}
