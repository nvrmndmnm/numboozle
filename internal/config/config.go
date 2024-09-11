package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Datasource string `yaml:"datasource"`
}

func MustLoadConfig() *Config {
	var path string

	flag.StringVar(&path, "config", "./config.yaml", "path to config file")
	flag.Parse()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("failed to find config file at: " + path)
	}

	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config file: " + err.Error())
	}

	return &config
}
