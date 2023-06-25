package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string `env:"SERVER_PORT"`
	}
	Database struct {
		Host string `env:"DATABASE_HOST"`
		Port string `env:"DATABASE_PORT"`
		Name string `env:"DATABASE_NAME"`
		Username string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASS"`
	}
}

func LoadEnv(cfg *Config) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load .env config")
	}

	err = env.Parse(cfg)

	if err != nil {
		log.Println("Failed to parse env variables")
	}
}

