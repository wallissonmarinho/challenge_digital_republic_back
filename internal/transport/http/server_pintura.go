package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log/level"
	"github.com/sirupsen/logrus"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/domain"
)

func (s *server) ObterQuantidadeDeLatasdeTinta(c *gin.Context) {
	var paredes []domain.Parede

	err := c.Bind(&paredes)
	if err != nil {
		_ = level.Error(*s.logger).Log("message", "invalid request")
	}

	resp, err := s.endpoint.ObterQuantidadeDeLatasdeTintaEndpoint(c, paredes)
	if err != nil {
		logrus.Error(err)
	}

	c.JSON(int(resp.(domain.CustomerResponse).Code.Int64), resp.(domain.CustomerResponse).Response)

}
