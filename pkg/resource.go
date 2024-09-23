// This package implements an http client for resource and pokemon APIs from pokeAPI
package pokemon

import (
	"encoding/json"

	"golang.org/x/exp/slog"
)

// Resource takes optional options, send the request
// with the specified endpoint in the caller client object endpoint
// parses the response and returns the Resource object and an error if exists
func (c Client) Resource(options ...int) (resource Resource, err error) {

	resp, err := c.sendRequest(options...)
	if err != nil {
		slog.Error("error in sending the request: %w", err)
		return
	}
	body, err := readBody(resp)
	if err != nil {
		slog.Error("error in reading the response body: %w", err)
		return
	}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		slog.Error("error in pasing the response body: %w", err)
		return
	}
	slog.Info("your resource is ready!")
	return
}
