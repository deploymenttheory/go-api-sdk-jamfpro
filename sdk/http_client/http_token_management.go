// http_token_management.go
package http_client

import (
	"time"
)

// TokenResponse represents the structure of a token response from the API.
type TokenResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

// ValidAuthToken checks if the current token is valid and not close to expiry.
// If the token is invalid, it tries to refresh it.
func (c *Client) ValidAuthTokenCheck() bool {

	// If token doesn't exist
	if c.Token == "" {
		if c.bearerTokenAuthCredentials.Username != "" && c.bearerTokenAuthCredentials.Password != "" {
			err := c.ObtainToken()
			if err != nil {
				return false
			}
		} else if c.oAuthCredentials.ClientID != "" && c.oAuthCredentials.ClientSecret != "" {
			err := c.ObtainOAuthToken(c.oAuthCredentials)
			if err != nil {
				return false
			}
		} else {
			c.logger.Error("No valid credentials provided. Unable to obtain a token.")
			return false
		}
	}

	// If token exists and is close to expiry or already expired
	if time.Until(c.Expiry) < c.config.BufferPeriod {
		if c.config.DebugMode {
			c.logger.Debug("Token is not valid or is close to expiry", "Expiry", c.Expiry)
		}

		var err error
		if c.bearerTokenAuthCredentials.Username != "" && c.bearerTokenAuthCredentials.Password != "" {
			err = c.RefreshToken()
		} else if c.oAuthCredentials.ClientID != "" && c.oAuthCredentials.ClientSecret != "" {
			err = c.RefreshOAuthToken()
		} else {
			c.logger.Error("Unknown auth method", "AuthMethod", c.authMethod)
			return false
		}

		if err != nil {
			return false
		}
	}

	if c.config.DebugMode {
		c.logger.Debug("Token is valid", "Expiry", c.Expiry)
	}
	return true
}
