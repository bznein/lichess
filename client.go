package lichess

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const lichessURL = "https://lichess.org"

type Client struct {
	BaseURL    *url.URL
	HttpClient HTTPClient
	Token      string
}

func (c Client) newRequest(method string, path string, buf *bytes.Buffer) (*http.Request, error) {
	var err error
	if c.Token == "" {
		return nil, errors.New("No token specified")
	}
	if c.BaseURL == nil {
		c.BaseURL, err = url.Parse(lichessURL)
	}

	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	if buf == nil {
		buf = &bytes.Buffer{}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	// Default request is json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Golang Lichess Client")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return resp, fmt.Errorf("expected status code 200, got %d - Requested URL was: %s", resp.StatusCode, req.URL)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
