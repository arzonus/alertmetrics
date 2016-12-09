package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config Config

func Get() *Config {
	return &config
}

func Init(configPath string) (*Config, error) {

	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (cfg *Config) Validate() error {
	if cfg.Database.Connection == "" {
		return errors.New("Config: Database connection wasn't specified")
	}
	if cfg.Period <= 0 {
		return errors.New("Config: Period time <= 0")
	}

	return nil
}
