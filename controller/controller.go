package controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/hardcake/eyesuite/service"
)

type Controller interface {
	ApiHome(c *fiber.Ctx) error

	UpsertConfig(c *fiber.Ctx) error
	ReadConfig(c *fiber.Ctx) error

	Login(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error

	ReadProfiles(c *fiber.Ctx) error
	UpsertProfiles(c *fiber.Ctx) error

	ReadImageText(c *fiber.Ctx) error
	PostImage(c *fiber.Ctx) error

	ReadUsers(c *fiber.Ctx) error
	UpsertUsers(c *fiber.Ctx) error

	ReadPlugins(c *fiber.Ctx) error
	UpsertPlugins(c *fiber.Ctx) error
}

type controller struct {
	*service.Service
}

func NewController(s *service.Service) Controller {
	return &controller{s}
}
