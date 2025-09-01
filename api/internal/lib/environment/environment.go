package environment

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

type Config struct {
	Port    string `json:"port"`
	Env     string `json:"env"`
	Version string `json:"version"`
}

var config *Config

func setupConfig() {
	// load .env into process env (only once)
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on system env vars")
	}

	config = &Config{
		Port:    os.Getenv("PORT"),
		Env:     os.Getenv("ENV"),
		Version: os.Getenv("VERSION"),
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
