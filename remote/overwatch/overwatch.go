package overwatch

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	baseURI = "https://overwatch-api.net/api/v1"
)

// Client contains any internals utilised for connecting with the Overwatch API.
type Client struct{}

// New instantiates a client for consuming data from the Overwatch API.
func New() (*Client, error) {
	c := &Client{}

	return c, nil
}

// do manipulates all requests to the remote API and applies any defaults, such as user-agent.
func (c *Client) do(req *http.Request, v interface{}) error {
	// Modify request as seen fit
	req.Header.Set("User-Agent", "Go_OverwatchBot/1.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return formatResponseError(resp.StatusCode)
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(&v)
}

func formatResponseError(code int) error {
	switch code {
	case http.StatusNotFound:
		return errors.New("404 not found")
	case http.StatusTooManyRequests:
		return errors.New("429 too many requests, please slow down")
	default:
		return errors.New("received non-200 request from overwatch-api")
	}
}
