package storage

import (
	"context"
	"encoding/json"
)

type User struct {
	Name         string `json:"name"`
	Key          string `json:"key"`
	Username     string `json:"username"`
	Enabled      bool   `json:"enabled"`
	Admin        bool   `json:"admin"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenUuid    string `json:"token_uuid"`
	RefreshUuid  string `json:"refresh_uuid"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

// USERS is a list to keep track of user keys
const USERS = "users"

func (s *storage) ReadUsers() ([]User, error) {
	var users []User
	r, err := s.Get(context.Background(), USERS).Result()
	if err != nil && err.Error() == NotFound {
		return users, nil
	}
	err = json.Unmarshal([]byte(r), &users)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *storage) ReadUser(key string) (*User, error) {
	users, err := s.ReadUsers()
	if err != nil {
		return nil, err
	}
	for u := range users {
		if users[u].Key == key {
			return &users[u], nil
		}
	}
	return nil, nil
}

func (s *storage) UpsertUsers(users []User) ([]User, error) {
	m, err := json.Marshal(users)
	if err != nil {
		return users, err
	}
	_, err = s.Set(context.Background(), USERS, m, 0).Result()
	if err != nil {
		return users, err
	}
	return users, nil
}
