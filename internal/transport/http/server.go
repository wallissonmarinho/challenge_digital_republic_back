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

	r.Use(CorsMiddleware())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", rest.HealthCheckHandler)
	r.POST("/pintura", rest.ObterQuantidadeDeLatasdeTinta)

	err := r.Run(":8080")
	logrus.Error(err)

	return r
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
