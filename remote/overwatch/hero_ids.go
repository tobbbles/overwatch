package overwatch

import (
	"fmt"
	"net/http"
)

// HerCount provides with a total count of the amount of heros in the database.
// HACK: Do to the remote API's IDs being sequential, we use this as a cheeky hack to just iterate from 1 to 'total'.
func (c *Client) HeroCount() (int, error) {
	var (
		uri = fmt.Sprintf("%s/hero/", baseURI)
	)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return 0, err
	}

	var resp = &Response{}
	if err := c.do(req, resp); err != nil {
		return 0, err
	}

	return resp.Total, nil
}
