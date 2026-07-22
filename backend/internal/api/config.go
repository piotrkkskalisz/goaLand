package api

import (
	"os"
	"time"
)

const defaultTimeout = 10 * time.Second

type Config struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}

func NewConfig(baseURL string, apiKey string, timeout time.Duration) Config {
	return Config{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Timeout: timeout,
	}
}

func NewConfigFromEnv() Config {
	return Config{
		BaseURL: os.Getenv("FOOTBALL_DATA_BASE_URL"),
		APIKey:  os.Getenv("FOOTBALL_DATA_API_TOKEN"),
		Timeout: defaultTimeout,
	}
}
