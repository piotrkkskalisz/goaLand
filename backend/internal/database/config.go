package database

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

func NewConfig(host string, port int, user string, password string, Name string, sslMode string) Config {
	return Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     Name,
		SSLMode:  sslMode,
	}
}

func NewConfigFromEnv() (Config, error) {
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid POSTGRES_PORT: %w", err)
	}

	return Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     port,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Name:     os.Getenv("POSTGRES_NAME"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}, nil
}
