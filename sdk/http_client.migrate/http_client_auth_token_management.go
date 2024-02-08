// httpclient_auth_token_management.go
package httpclient

import (
	"fmt"
	"time"
)

// TokenResponse represents the structure of a token response from the API.
type TokenResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

// ValidAuthTokenCheck checks if the current token is valid and not close to expiry.
// If the token is invalid, it tries to refresh it.
// It returns a boolean indicating the validity of the token and an error if there's a failure.
func (c *Client) ValidAuthTokenCheck() (bool, error) {

	if c.Token == "" {
		if c.AuthMethod == "bearer" {
			err := c.ObtainToken()
			if err != nil {
				return false, fmt.Errorf("failed to obtain bearer token: %w", err)
			}
		} else if c.AuthMethod == "oauth" {
			err := c.ObtainOAuthToken(c.config.Auth)
			if err != nil {
				return false, fmt.Errorf("failed to obtain OAuth token: %w", err)
			}
		} else {
			return false, fmt.Errorf("no valid credentials provided. Unable to obtain a token")
		}
	}

	// If token exists and is close to expiry or already expired
	if time.Until(c.Expiry) < c.config.TokenRefreshBufferPeriod {
		var err error
		if c.BearerTokenAuthCredentials.Username != "" && c.BearerTokenAuthCredentials.Password != "" {
			err = c.RefreshToken()
		} else if c.OAuthCredentials.ClientID != "" && c.OAuthCredentials.ClientSecret != "" {
			err = c.ObtainOAuthToken(c.config.Auth)
		} else {
			return false, fmt.Errorf("unknown auth method: %s", c.AuthMethod)
		}

		if err != nil {
			return false, fmt.Errorf("failed to refresh token: %w", err)
		}
	}

	if time.Until(c.Expiry) < c.config.TokenRefreshBufferPeriod {
		return false, fmt.Errorf("token lifetime setting less than buffer. Buffer setting: %v, Time (seconds) until Exp: %v", c.config.TokenRefreshBufferPeriod, time.Until(c.Expiry))
	}
	return true, nil
}
