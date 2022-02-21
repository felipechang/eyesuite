package controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/hardcake/eyesuite/service/storage"
)

// ReadConfig read configuration
func (o *controller) ReadConfig(c *fiber.Ctx) error {
	rc, err := o.Storage.ReadConfig()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(rc)
}

// UpsertConfig update or create a configuration
func (o *controller) UpsertConfig(c *fiber.Ctx) error {
	config := &storage.Config{}
	if err := c.BodyParser(&config); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	uc, err := o.Storage.UpsertConfig(config)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(uc)
}
