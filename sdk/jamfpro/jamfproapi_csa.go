// jamfproapi_csa.go
// Jamf Pro Api - CSA Tokens
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

// TODO

const uriCSAToken = "/api/v1/csa/token"

// Structs

type ResourceCSATokenExchange struct {
	RefreshExpiration int      `json:"refreshExpiration"`
	Scopes            []string `json:"scopes"`
}

// CRUD

func (c *Client) GetCSATokenExchangeInfo() (*ResourceCSATokenExchange, error) {
	endpoint := uriCSAToken
	var out ResourceCSATokenExchange

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "csa token exchange info", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

func (c *Client) RefreshCSATokenExchange(username, password string) (*ResourceCSATokenExchange, error) {
	endpoint := uriCSAToken
	var out ResourceCSATokenExchange

	payload := struct {
		Username string
		Password string
	}{
		Username: username,
		Password: password,
	}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, payload, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "csa token exchange", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

func (c *Client) InitializeCSATokenExchange(username, password string) (*ResourceCSATokenExchange, error) {
	endpoint := uriCSAToken
	var out ResourceCSATokenExchange

	payload := struct {
		Username string
		Password string
	}{
		Username: username,
		Password: password,
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, payload, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "csa token exchange", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

func (c *Client) DeleteCSATokenExchange() error {
	endpoint := uriCSAToken

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDelete, "csa token exchange", err)
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDelete, "csa token exchange", err)
	}

	return nil
}
