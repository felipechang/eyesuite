package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/hardcake/eyesuite/service/token"
)

func (m *middleware) Auth(c *fiber.Ctx) error {

	// find and evaluate access token
	key, err := m.evaluateAccessKey(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	if key != "" {
		c.Set("user_key", key)
		return c.Next()

	}

	// find and evaluate refresh token
	key, err = m.evaluateRefreshKey(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())

	}
	if key != "" {
		c.Set("user_key", key)
		return c.Next()
	}

	return fiber.NewError(fiber.StatusUnauthorized, "could not authenticate tokens")
}

// evaluateAccessKey evaluate that access token is correct
func (m *middleware) evaluateAccessKey(c *fiber.Ctx) (string, error) {
	headerToken := m.service.Token.ExtractHeaderToken(c)
	if headerToken == "" {
		return "", nil
	}
	key, err := m.service.Token.ExtractUserKey(headerToken, m.cts.AccessSecret)
	if err != nil {
		return "", err
	}
	user, err := m.service.Storage.ReadUser(key)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}
	if user.AccessToken != headerToken {
		return "", errors.New("invalid token for user")
	}
	return key, nil
}

// evaluateRefreshKey evaluate that refresh token is correct
func (m *middleware) evaluateRefreshKey(c *fiber.Ctx) (string, error) {
	req := &token.Pair{}
	err := c.BodyParser(&req)
	if err != nil && err.Error() != "EOF" {
		return "", errors.New("could not parse body token")
	}
	if req.RefreshToken == "" {
		return "", errors.New("refresh token not found")
	}

	key, err := m.service.Token.ExtractUserKey(req.RefreshToken, m.cts.RefreshSecret)
	if err != nil {
		return "", err
	}
	user, err := m.service.Storage.ReadUser(key)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}
	if user.RefreshToken != req.RefreshToken {
		return "", errors.New("invalid token for user")
	}
	return key, nil
}
