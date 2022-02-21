package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/hardcake/eyesuite/service"
)

type Middleware interface {
	Auth(c *fiber.Ctx) error
	Admin(c *fiber.Ctx) error
}

type middleware struct {
	service *service.Service
	cts     *service.Constants
}

func NewMiddleware(s *service.Service) Middleware {
	cts := service.GetEnv()
	return &middleware{s, cts}
}
