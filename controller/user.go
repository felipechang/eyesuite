package controller

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/hardcake/eyesuite/service/storage"
)

type UserSummary struct {
	Name     string `json:"name"`
	Enabled  bool   `json:"enabled"`
	Admin    bool   `json:"admin"`
	Username string `json:"username"`
	Key      string `json:"key"`
	Password string `json:"password"`
}

func (o *controller) ReadUsers(c *fiber.Ctx) error {

	// read users
	users, err := o.Storage.ReadUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// read summary values
	var summary []UserSummary
	for _, u := range users {
		summary = append(summary, UserSummary{
			Name:     u.Name,
			Enabled:  u.Enabled,
			Admin:    u.Admin,
			Username: u.Username,
			Key:      u.Key,
			Password: "",
		})
	}
	return c.JSON(summary)
}

func (o *controller) UpsertUsers(c *fiber.Ctx) error {

	// unmarshall request
	var summary []UserSummary
	if err := c.BodyParser(&summary); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// generate key where required
	for s := range summary {
		if summary[s].Password != "" && summary[s].Username != "" {
			summary[s].Key = base64.StdEncoding.EncodeToString([]byte(summary[s].Username + ":" + summary[s].Password))
		}
	}

	// read users
	users, err := o.Storage.ReadUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// users were deleted
	if len(users) > len(summary) {
		for u := range users {
			var k = o.findSummaryKey(summary, users[u].Key)
			if k == -1 {
				users = append(users[:u], users[u+1:]...)
			}
		}
	}

	// apply changes or create new user
	for s := range summary {
		if summary[s].Key == "" {
			continue
		}
		var k = o.findUserKey(users, summary[s].Key)
		if k != -1 {
			users[k].Name = summary[s].Name
			users[k].Enabled = summary[s].Enabled
			users[k].Admin = summary[s].Admin
			users[k].Key = summary[s].Key
		} else {
			var user storage.User
			err = o.Token.ApplyTokenValues(summary[s].Key, &user)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}
			user.Name = summary[s].Name
			user.Username = summary[s].Username
			user.Enabled = summary[s].Enabled
			user.Admin = summary[s].Admin
			user.Key = summary[s].Key
			users = append(users, user)
		}
	}

	// store and return users
	u, err := o.Storage.UpsertUsers(users)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.JSON(u)
}

// findUserKey find the key position of the user
func (o *controller) findUserKey(users []storage.User, key string) int {
	for u := range users {
		if users[u].Key == key {
			return u
		}
	}
	return -1
}

// findSummaryKey find the key position of the summary
func (o *controller) findSummaryKey(summaries []UserSummary, key string) int {
	for s := range summaries {
		if summaries[s].Key == key {
			return s
		}
	}
	return -1
}
