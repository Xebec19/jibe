package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	dbConn string `mapstructure:"DB_CONN"`
	env    string `mapstructure:"ENV"`
}

func NewConfig(path string) (*Config, error) {

	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	return &Config{
		dbConn: os.Getenv("DB_CONN"),
		env:    os.Getenv("ENV"),
	}, nil
}
