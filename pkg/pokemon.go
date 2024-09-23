// This package implements an http client for resource and pokemon APIs from pokeAPI
package pokemon

import (
	"encoding/json"
	"fmt"

	"golang.org/x/exp/slog"
)
// Pokemon takes the pokemon id as an integer, send the request with the /pokemon endpoint 
// parses the response and returns the Pokemon object and an error if exists
func (c Client) Pokemon(id int) (pokemon Pokemon, err error) {
	c.Endpoint = fmt.Sprintf("/pokemon/%d", id)
	resp, err := c.sendRequest()
	if err != nil {
		slog.Error("error in sending the request: %w", err)
		return
	}
	body, err := readBody(resp)
	if err != nil {
		slog.Error("error in reading the response body: %w", err)
		return
	}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		slog.Error("error in pasing the response body: %w", err)
		return
	}
	slog.Info("your pokemon is ready!")
	return
}
