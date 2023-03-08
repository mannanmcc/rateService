package config

import (
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	CurrencyAPIUrl         string        `yaml:"currency_api_url"`
	CurrencyAPIConnTimeout time.Duration `yaml:"currency_api_connection_timeout"`
	MongoDbPort            string        `yaml:"mongodb_port"`
}

const (
	configFilePathEnv  = "CONFIG_FILE"
	defaultFilePathEnv = "config.yml"
)

func Load() (Config, error) {
	var cfg Config

	configFilePath, exists := os.LookupEnv(configFilePathEnv)
	if !exists {
		configFilePath = defaultFilePathEnv
	}

	filename, _ := filepath.Abs(configFilePath)
	yamlFile, err := os.ReadFile(filename)

	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
