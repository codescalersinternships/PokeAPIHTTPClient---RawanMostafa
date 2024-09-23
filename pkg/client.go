// This package implements an http client
package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

const baseUrl = "https://pokeapi.co/api/v2"

// Client holds the endpoint of the http client as well as the actual http.Client object
type Client struct {
	Endpoint string
	Client   http.Client
}

// NewClient takes the endpoint and timeout and returns a Client object
func NewClient(endpoint string, timeout time.Duration) Client {
	slog.Info("New Client created! \n")
	return Client{
		Endpoint: endpoint,
		Client:   http.Client{Timeout: timeout},
	}
}

func readBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return make([]byte, 0), fmt.Errorf("error in reading request body: %v", err)

	}
	return body, nil
}

// SendRequest creates the request then send it and returns the response and an error if exists
func (c Client) sendRequest() (*http.Response, error) {
	url := baseUrl+c.Endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("couldn't create the request with url: %s", url)
		return nil, err
	}

	resp, err := c.Client.Do(req)

	if err != nil {
		slog.Error("couldn't send the request with url: %s", url)
		return nil, fmt.Errorf("error in sending request: %v", err)
	}
	slog.Info("your request has been sent successfully!")

	return resp, nil
}

func (c Client) Resource() (resource Resource, err error) {
	resp, err := c.sendRequest()
	if err != nil {
		return
	}
	body, err := readBody(resp)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resource)

	return
}
