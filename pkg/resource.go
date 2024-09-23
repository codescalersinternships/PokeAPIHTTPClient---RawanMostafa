package pokemon

import "encoding/json"

func (c Client) Resource(options ...int) (resource Resource, err error) {
	var offset int
	var limit int
	if len(options) == 2 {
		offset = options[0]
		limit = options[1]
	} else if len(options) == 1 {
		offset = options[0]
	}

	resp, err := c.sendRequest(offset, limit)
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
