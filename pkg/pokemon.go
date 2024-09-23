package pokemon

import (
	"encoding/json"
	"fmt"
)

func (c Client) Pokemon(id int) (pokemon Pokemon, err error) {
	c.Endpoint = fmt.Sprintf("/pokemon/%d", id)
	resp, err := c.sendRequest()
	if err != nil {
		return
	}
	body, err := readBody(resp)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &pokemon)

	return
}
