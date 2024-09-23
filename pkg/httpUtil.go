// This package implements an http client for resource and pokemon APIs from pokeAPI
package pokemon

import (
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

func (c Client) sendRequest(options ...int) (*http.Response, error) {
	var offset int
	var limit int
	if len(options) == 2 {
		offset = options[0]
		limit = options[1]
	} else if len(options) == 1 {
		offset = options[0]
	}

	url := baseUrl + c.Endpoint + fmt.Sprintf("?offset=%d&limit=%d", offset, limit)

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
