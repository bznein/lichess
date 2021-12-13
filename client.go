package lichess

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	ndjson "github.com/kandros/go-ndjson"
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

func (c Client) newRequest(method string, path string, buf io.Reader) (*http.Request, error) {
	var err error
	if c.Token == "" {
		return nil, errors.New("No token specified")
	}
	if c.BaseURL == nil {
		c.BaseURL, err = url.Parse(lichessURL)
		if err != nil {
			return nil, err
		}
	}

	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

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
	if req.Header.Get("Accept") == "application/x-ndjson" {
		respToJSON := ndjson.ToJSON(resp.Body)
		// ndjson leaves a trailing , before the last ]
		i := strings.LastIndex(respToJSON, ",")
		if i == len(respToJSON)-2 {
			respToJSON = respToJSON[:i] + strings.Replace(respToJSON[i:], ",", "", 1)
		}
		err = json.Unmarshal([]byte(respToJSON), v)
	} else {
		d := json.NewDecoder(resp.Body)
		err = d.Decode(v)
	}
	return resp, err
}
