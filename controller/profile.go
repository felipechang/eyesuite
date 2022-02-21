package controller

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/hardcake/eyesuite/service/storage"
)

// ReadProfiles read profile list
func (o *controller) ReadProfiles(c *fiber.Ctx) error {
	rp, err := o.Storage.ReadProfiles()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(rp)
}

// UpsertProfiles update or create profile list
func (o *controller) UpsertProfiles(c *fiber.Ctx) error {
	var profiles []storage.Profile
	if err := c.BodyParser(&profiles); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	up, err := o.Storage.UpsertProfiles(profiles)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(up)
}
