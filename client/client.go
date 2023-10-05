package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL   string
	AuthToken string
}

type ClientConfig struct {
	BaseURL   string
	AuthToken string
}

func NewClient(config *ClientConfig) (*Client, error) {
	if config.AuthToken == "" {
		return nil, errors.New("AuthToken is required")
	}

	if config.BaseURL == "" {
		return nil, errors.New("BaseUrl is required")
	}

	return &Client{
		BaseURL:   config.BaseURL,
		AuthToken: config.AuthToken,
	}, nil
}

func (c *Client) NewRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	// Convert body to JSON if it's not nil
	var buf *bytes.Buffer
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
		buf = bytes.NewBuffer(bodyBytes)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, c.BaseURL+endpoint, buf)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set common headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.AuthToken)

	return req, nil
}
