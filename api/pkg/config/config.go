package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string        `mapstructure:"PORT"`
	DbConn             string        `mapstructure:"DB_CONN"`
	Domain             string        `mapstructure:"DOMAIN"`
	Env                string        `mapstructure:"ENV"`
	MaxHeaderBytes     int           `json:"max_header_bytes"`
	ReadTimeout        time.Duration `json:"read_timeout"`
	WriteTimeout       time.Duration `json:"write_timeout"`
	IdleTimeout        time.Duration `json:"idle_timeout"`
	MaxBodySizeAllowed int64         `json:"max_body_size_allowed"`
}

func NewConfig(path string) (*Config, error) {

	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	return &Config{
		DbConn:             os.Getenv("DB_CONN"),
		Env:                os.Getenv("ENV"),
		Port:               os.Getenv("PORT"),
		MaxHeaderBytes:     1 << 20,
		ReadTimeout:        10 * time.Second,
		WriteTimeout:       10 * time.Second,
		IdleTimeout:        10 * time.Second,
		MaxBodySizeAllowed: 1 * 1024 * 1024,
	}, nil
}
