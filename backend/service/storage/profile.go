package storage

import (
	"context"
	"encoding/json"
)

type Profile struct {
	Name     string           `json:"name"`
	Mapping  []ProfileMapping `json:"mapping"`
	Plugin   string           `json:"plugin"`
	Template string           `json:"template"`
}

type ProfileMapping struct {
	LabelText   string `json:"labelText"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Placeholder string `json:"placeholder"`
	Readonly    bool   `json:"readonly"`
	Hidden      bool   `json:"hidden"`
}

const PROFILES = "profiles"

func (s *storage) ReadProfiles() ([]Profile, error) {
	var profiles []Profile
	r, err := s.Get(context.Background(), PROFILES).Result()
	if err != nil && err.Error() == NotFound {
		return profiles, nil
	}
	err = json.Unmarshal([]byte(r), &profiles)
	if err != nil {
		return profiles, err
	}
	return profiles, nil
}

func (s *storage) UpsertProfiles(profiles []Profile) ([]Profile, error) {
	m, err := json.Marshal(profiles)
	if err != nil {
		return profiles, err
	}
	_, err = s.Set(context.Background(), PROFILES, m, 0).Result()
	if err != nil {
		return profiles, err
	}
	return profiles, nil
}
