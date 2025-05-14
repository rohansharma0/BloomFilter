package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
	}
}

func LoanConfig() *Config {
	cfg := &Config{}
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	cfg.Server.Port = os.Getenv("PORT")
	return cfg
}
