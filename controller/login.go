package controller

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/hardcake/eyesuite/service/storage"
	"gitlab.com/hardcake/eyesuite/service/token"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserResponse struct {
	Control      string `json:"control"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login start user session
func (o *controller) Login(c *fiber.Ctx) error {
	key, err := getUserLoginKey(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	user, err := o.refreshUserTokens(key)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	var control = user.Name
	if user.Admin {
		control = control + "1"
	} else {
		control = control + "0"
	}
	return c.JSON(&UserResponse{
		Control:      base64.StdEncoding.EncodeToString([]byte(control)),
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	})
}

// Refresh update user access and refresh tokens
func (o *controller) Refresh(c *fiber.Ctx) error {
	key := fmt.Sprintf("%v", c.Locals("user_key"))
	user, err := o.refreshUserTokens(key)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	return c.JSON(token.Pair{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	})
}

// Logout end user session
func (o *controller) Logout(c *fiber.Ctx) error {
	key := fmt.Sprintf("%v", c.Locals("user_key"))
	_, err := o.refreshUserTokens(key)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	return c.JSON("1")
}

// getUserLoginKey get user key from username and password
func getUserLoginKey(c *fiber.Ctx) (string, error) {
	var u = new(UserLogin)
	if err := c.BodyParser(&u); err != nil {
		return "", err
	}
	if u.Username == "" || u.Password == "" {
		return "", errors.New("username or password missing")
	}
	key := base64.StdEncoding.EncodeToString([]byte(u.Username + ":" + u.Password))
	return key, nil
}

// refreshUserTokens refresh the user tokens and return user
func (o *controller) refreshUserTokens(key string) (*storage.User, error) {

	// read users
	users, err := o.Storage.ReadUsers()
	if err != nil {
		return nil, err
	}

	// find user with key
	var k = o.findUserKey(users, key)
	if k == -1 {
		return nil, errors.New("user not found")
	}

	// apply new tokens to user
	err = o.Token.ApplyTokenValues(key, &users[k])
	if err != nil {
		return nil, err
	}

	// save users
	users, err = o.Storage.UpsertUsers(users)
	if err != nil {
		return nil, err
	}

	// return changed user
	return &users[k], nil
}
