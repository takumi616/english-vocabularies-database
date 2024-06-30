package infrastructure

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppPort  string `env:"APP_CONTAINER_PORT"`
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	DbName   string `env:"POSTGRES_DB"`
	User     string `env:"POSTGRES_USER"`
	PassWord string `env:"POSTGRES_PASSWORD"`
	Sslmode  string `env:"POSTGRES_SSLMODE"`
}

// type DbConfig struct {
// 	Host     string `env:"POSTGRES_HOST"`
// 	Port     string `env:"POSTGRES_PORT"`
// 	DbName   string `env:"POSTGRES_DB"`
// 	User     string `env:"POSTGRES_USER"`
// 	PassWord string `env:"POSTGRES_PASSWORD"`
// 	Sslmode  string `env:"POSTGRES_SSLMODE"`
// }

// Read environment variables and set them to golang struct
func NewConfig() (*Config, error) {
	config := &Config{}
	//Set environment variables to config.
	if err := env.Parse(config); err != nil {
		return nil, err
	}
	return config, nil
}
