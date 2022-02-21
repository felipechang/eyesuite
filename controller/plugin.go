package controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/hardcake/eyesuite/service/storage"
)

// ReadPlugins read plugin list
func (o *controller) ReadPlugins(c *fiber.Ctx) error {
	rp, err := o.Storage.ReadPlugins()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(rp)
}

// UpsertPlugins update or create plugin list
func (o *controller) UpsertPlugins(c *fiber.Ctx) error {
	var plugins []storage.Plugin
	if err := c.BodyParser(&plugins); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	up, err := o.Storage.UpsertPlugins(plugins)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(up)
}
