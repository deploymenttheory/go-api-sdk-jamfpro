package jamfpro

import (
	"fmt"
)

const uriCSATokenExchange = "/api/v1/csa/token"

// ResponseCSAToken represents the structure of the response when requesting
// a CSA token from the Jamf Pro API.
type ResponseCSAToken struct {
	RefreshExpiration int      `json:"refreshExpiration"`
	Scopes            []string `json:"scopes"`
}

// CSAToken is used to provide the email and password when requesting
// a CSA token from the Jamf Pro API.
type CSAToken struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

// GetCSAToken retrieves a CSA token from the Jamf Pro API.
func (c *Client) GetCSAToken() (*ResponseCSAToken, error) {
	uri := uriCSATokenExchange

	var out ResponseCSAToken
	err := c.DoRequest("GET", uri, nil, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to get CSA token: %v", err)
	}

	return &out, nil
}

// CreateCSAToken creates a new CSA token in the Jamf Pro API using the provided email and password.
func (c *Client) CreateCSAToken(email string, password string) (*ResponseCSAToken, error) {
	uri := uriCSATokenExchange
	in := &CSAToken{
		EmailAddress: email,
		Password:     password,
	}

	var out ResponseCSAToken
	err := c.DoRequest("POST", uri, in, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSA token: %v", err)
	}

	return &out, nil
}

// UpdateCSAToken updates an existing CSA token in the Jamf Pro API using the provided email and password.
func (c *Client) UpdateCSAToken(email string, password string) (*ResponseCSAToken, error) {
	uri := uriCSATokenExchange

	reqBody := &struct {
		*CSAToken
	}{
		CSAToken: &CSAToken{
			EmailAddress: email,
			Password:     password,
		},
	}

	var out ResponseCSAToken
	err := c.DoRequest("PUT", uri, reqBody, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to update CSA token: %v", err)
	}

	return &out, nil
}

// DeleteCSAToken deletes an existing CSA token in the Jamf Pro API.
func (c *Client) DeleteCSAToken() error {
	uri := uriCSATokenExchange

	err := c.DoRequest("DELETE", uri, nil, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf("failed to delete CSA token: %v", err)
	}
	return nil
}
