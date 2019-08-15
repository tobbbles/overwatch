package overwatch

import (
	"fmt"
	"net/http"

	"service/models"
)

// Hero returns information for the given id.
func (c *Client) Hero(id int) (*models.Hero, error) {
	var (
		uri = fmt.Sprintf("%s/hero/%d/", baseURI, id)
	)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var resp = &models.Hero{}
	if err := c.do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
