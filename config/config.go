package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	CurrencyAPIUrl string `yaml:"currency_api_url"`
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
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}