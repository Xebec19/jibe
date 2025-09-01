package environment

import (
	"os"
	"sync"
)

var once sync.Once

type Config struct {
	Port string `json:"port"`
	Env  string `json:"env"`
}

var config *Config

func setupConfig() {
	config = &Config{
		Port: os.Getenv("PORT"),
	}
}

func GetConfig() *Config {
	once.Do(setupConfig)

	return config
}

func IsDevEnvironment() bool {
	once.Do(setupConfig)

	return config.Env == "development"
}
