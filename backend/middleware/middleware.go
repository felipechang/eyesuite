package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hardcake/eyesuite/service"
)

type Middleware interface {
	Auth(c *gin.Context)
	Admin(c *gin.Context)
}

type middleware struct {
	service *service.Service
	cts     *service.Constants
}

func NewMiddleware(s *service.Service) Middleware {
	cts := service.GetEnv()
	return &middleware{s, cts}
}
