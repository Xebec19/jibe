package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DbConn string `mapstructure:"DB_CONN"`
	Env    string `mapstructure:"ENV"`
}

func NewConfig(path string) (*Config, error) {

	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	return &Config{
		DbConn: os.Getenv("DB_CONN"),
		Env:    os.Getenv("ENV"),
		Port:   os.Getenv("PORT"),
	}, nil
}
