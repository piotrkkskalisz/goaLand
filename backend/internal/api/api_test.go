package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewExampleClient(t *testing.T) {
	config := NewConfig(
		"https://example.com",
		"secret-token",
		defaultTimeout,
	)

	client := NewClient(config, nil)

	require.NotNil(t, client)
	require.Equal(t, config, client.config)
	require.Nil(t, client.httpClient)
}

func TestNewClientFromEnv(t *testing.T) {
	client := NewClientFromEnv()

	require.NotNil(t, client)

	require.NotEmpty(t, client.config.BaseURL)
	require.NotEmpty(t, client.config.APIKey)
	require.Equal(t, defaultTimeout, client.config.Timeout)

	require.NotNil(t, client.httpClient)
}
