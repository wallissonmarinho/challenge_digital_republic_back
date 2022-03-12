package endpoint

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/domain"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/service"
	"gopkg.in/guregu/null.v4"
)

func makeObterQuantidadeDeLatasdeTintaEndpoint(s service.ServiceFactory, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		if request == nil {
			return domain.CustomerResponse{
				Code:     null.IntFrom(http.StatusBadRequest),
				Response: errors.New("requisição sem parâmetro"),
			}, nil
		}

		parede := request.([]domain.Parede)
		resp, err := s.Pintura().ObterQuantidadeDeLatasdeTinta(ctx, parede)
		if err != nil {
			_ = level.Error(logger).Log("message", "invalid request")
			return domain.CustomerResponse{
				Code:     null.IntFrom(http.StatusBadRequest),
				Response: err.Error(),
			}, nil
		}

		_ = level.Error(logger).Log("message", "ok")

		return domain.CustomerResponse{
			Code:     null.IntFrom(http.StatusCreated),
			Response: resp,
		}, nil

	}
}
