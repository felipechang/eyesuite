package storage

import (
	"context"
	"encoding/json"
)

type Plugin struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Enabled     bool             `json:"enabled"`
	Mapping     []ProfileMapping `json:"mapping"`
}

const PLUGINS = "plugins"

func (s *storage) ReadPlugins() ([]Plugin, error) {
	var plugins []Plugin
	r, err := s.Get(context.Background(), PLUGINS).Result()
	if err != nil && err.Error() == NotFound {
		return nil, nil
	}
	err = json.Unmarshal([]byte(r), &plugins)
	if err != nil {
		return nil, err
	}
	return plugins, nil
}

func (s *storage) UpsertPlugins(plugins []Plugin) ([]Plugin, error) {
	m, err := json.Marshal(plugins)
	if err != nil {
		return plugins, err
	}
	_, err = s.Set(context.Background(), PLUGINS, m, 0).Result()
	if err != nil {
		return plugins, err
	}
	return plugins, nil
}
