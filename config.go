package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Versions []Version `json:"versions"`
}

func ConfigReadFromFile(file string) (*Config, error) {
	contents, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := Config{}

	err = json.Unmarshal(contents, &config)
	if err != nil {
		return nil, err
	}

	err = config.validate()
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) validate() error {
	length := len(c.Versions)
	if length == 0 {
		return errors.New("config.validate: no versions specified")
	}
	// TODO figure out better mechanism for verifying if a Version is ready
	for i := 0; i < length; i++ {
		c.Versions[i].Downloads = -1
	}
	return nil
}
