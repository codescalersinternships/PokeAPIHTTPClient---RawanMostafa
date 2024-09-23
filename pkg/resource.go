package pokemon

import "encoding/json"

func (c Client) Resource(options ...int) (resource Resource, err error) {

	resp, err := c.sendRequest(options...)
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
