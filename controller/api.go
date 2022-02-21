package controller

import (
	"github.com/gofiber/fiber/v2"
)

func (o *controller) ApiHome(c *fiber.Ctx) error {
	return c.JSON("Welcome to EyeSuite")
}
