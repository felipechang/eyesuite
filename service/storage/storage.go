package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type Storage interface {
	ReadConfig() (*Config, error)
	UpsertConfig(config *Config) (*Config, error)

	ReadProfiles() ([]Profile, error)
	UpsertProfiles(profile []Profile) ([]Profile, error)

	ReadUser(key string) (*User, error)
	ReadUsers() ([]User, error)
	UpsertUsers(users []User) ([]User, error)

	ReadPlugins() ([]Plugin, error)
	UpsertPlugins(plugins []Plugin) ([]Plugin, error)
}

type storage struct {
	*redis.Client
}

var NotFound = "redis: nil"

func NewStorage(host string, port string, password string) Storage {
	c := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	_, err := c.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	return &storage{c}
}
