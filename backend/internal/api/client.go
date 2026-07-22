package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const tokenName = "X-Auth-Token"
const DateLayout = "2006-01-02"

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	config     Config
	httpClient HttpClient
}

/*
func IsCountry(parentArea *string) bool {
	return parentArea != nil && *parentArea != "World"
}
*/

func NewClient(config Config, httpClient HttpClient) *Client {
	return &Client{
		config:     config,
		httpClient: httpClient,
	}
}

func NewClientFromEnv() *Client {
	config := NewConfigFromEnv()
	return &Client{
		config:     config,
		httpClient: &http.Client{Timeout: config.Timeout},
	}
}

func (c *Client) fetch(path string, out any) error {
	return c.fetchWithQuery(path, out, make(map[string]string))
}

func (c *Client) fetchWithQuery(path string, out any, queryParamsToAdd map[string]string) error {
	req, err := http.NewRequest(http.MethodGet, c.config.BaseURL+path, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	queryParams := req.URL.Query()
	for key, value := range queryParamsToAdd {
		queryParams.Set(key, value)
	}
	req.URL.RawQuery = queryParams.Encode()

	req.Header.Set(tokenName, c.config.APIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status %s", resp.Status)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("failed to decode response from %s: %w", path, err)
	}

	return nil
}
