package endpoint

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/domain"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/service"
)

// makeHealthEndpoint return if service up
func makeHealthEndpoint(s service.ServiceFactory, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		_ = level.Error(logger).Log("message", "ok")

		return domain.CustomerResponse{
			Code:     http.StatusOK,
			Response: "Service UP",
		}, nil
	}
}
