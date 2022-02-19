package storage

import (
	"context"
	"encoding/json"
	"strings"
)

type Config struct {
	Version        string `json:"version"`
	WsUrl          string `json:"ws_url"`
	Realm          string `json:"realm"`
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	TokenID        string `json:"token_id"`
	TokenSecret    string `json:"token_secret"`
}

const CONFIG = "config"

func (s *storage) ReadConfig() (*Config, error) {
	var config Config
	r, err := s.Get(context.Background(), CONFIG).Result()
	if err != nil && err.Error() == NotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(r), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (s *storage) UpsertConfig(config *Config) (*Config, error) {
	config.Realm = strings.ToUpper(config.Realm)
	m, err := json.Marshal(config)
	if err != nil {
		return config, err
	}
	_, err = s.Set(context.Background(), CONFIG, m, 0).Result()
	if err != nil {
		return config, err
	}
	return config, nil
}
