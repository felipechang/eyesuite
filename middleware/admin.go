package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func (m *middleware) Admin(c *fiber.Ctx) error {

	accessToken := m.service.Token.ExtractHeaderToken(c)
	if accessToken == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "authorization token not found in header")
	}
	key, err := m.service.Token.ExtractUserKey(accessToken, m.cts.AccessSecret)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	user, err := m.service.Storage.ReadUser(key)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	if user.Admin != true {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized access")
	}
	return c.Next()
}
