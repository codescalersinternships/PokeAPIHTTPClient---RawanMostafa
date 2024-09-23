// This package implements an http client for resource and pokemon APIs from pokeAPI
package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
	"golang.org/x/exp/slog"
)

// Pokemon takes the pokemon id as an integer, send the request with the /pokemon endpoint
// parses the response and returns the Pokemon object and an error if exists
func (c Client) Pokemon(id int) (pokemon Pokemon, err error) {
	c.Endpoint = fmt.Sprintf("/pokemon/%d", id)

	var resp *http.Response

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.MaxElapsedTime = 10 * time.Second

	retryError := backoff.RetryNotify(func() error {
		resp, err = c.sendRequest()
		return err
	}, expBackoff, func(err error, d time.Duration) {
		slog.Warn("Request failed, Retrying ...")
	})

	if retryError != nil {
		slog.Error("failed to make the request after retries: %v", err)
		return pokemon, retryError
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
