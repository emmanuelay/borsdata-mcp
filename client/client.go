package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultBaseURL = "https://apiservice.borsdata.se"
	userAgent      = "borsdata-mcp/"
)

// Client is the Börsdata API client.
type Client struct {
	apiKey     string
	version    string
	baseURL    string
	httpClient *http.Client
}

// New creates a new Börsdata API client with the given API key.
func New(apiKey, version string) *Client {
	return &Client{
		apiKey:  apiKey,
		version: version,
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// NewWithBaseURL creates a new client with a custom base URL (useful for testing).
func NewWithBaseURL(apiKey, baseURL, version string) *Client {
	c := New(apiKey, version)
	c.baseURL = baseURL
	return c
}

// get performs a GET request and returns the response body.
func (c *Client) get(path string, params url.Values) ([]byte, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("authKey", c.apiKey)

	reqURL := fmt.Sprintf("%s%s?%s", c.baseURL, path, params.Encode())

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	agent := fmt.Sprintf("%s%s", userAgent, c.version)
	req.Header.Set("User-Agent", agent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// getJSON performs a GET request and unmarshals the JSON response into dest.
func (c *Client) getJSON(path string, params url.Values, dest interface{}) error {
	body, err := c.get(path, params)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, dest); err != nil {
		return fmt.Errorf("unmarshaling response: %w", err)
	}
	return nil
}
