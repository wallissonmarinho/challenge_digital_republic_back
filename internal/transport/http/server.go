package transport

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/sirupsen/logrus"
	"github.com/wallissonmarinho/challenge_digital_republic/internal/endpoint"
)

type server struct {
	endpoint *endpoint.Endpoints
	logger   *log.Logger
}

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(context context.Context, endpoint *endpoint.Endpoints, logger *log.Logger) http.Handler {
	rest := &server{
		endpoint: endpoint,
		logger:   logger,
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", rest.HealthCheckHandler)

	err := r.Run(":8080")
	logrus.Error(err)

	return r
}
