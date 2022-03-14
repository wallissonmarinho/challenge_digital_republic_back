package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log/level"
	"github.com/sirupsen/logrus"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/domain"
)

func (s *server) ObterQuantidadeDeLatasdeTinta(c *gin.Context) {
	var parede []domain.Parede

	teste := c.Request.Header
	logrus.Info(teste)

	err := c.ShouldBind(&parede)
	if err != nil {
		_ = level.Error(*s.logger).Log("message", "invalid request")
	}

	resp, err := s.endpoint.ObterQuantidadeDeLatasdeTintaEndpoint(c, parede)
	if err != nil {
		logrus.Error(err)
	}

	c.JSON(int(200), resp.(domain.CustomerResponse).Response)

}
