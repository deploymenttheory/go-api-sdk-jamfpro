package jamfpro

import (
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) invalidateToken() error {
	req, err := http.NewRequest("POST", c.url+"/api/v1/auth/invalidate-token", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+*c.token)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		c.token = nil
		c.tokenExpiration = nil
		return nil
	} else if resp.StatusCode == 401 {
		return errors.New("Token already invalid")
	}

	return errors.New("An unknown error occurred invalidating the token")
}
