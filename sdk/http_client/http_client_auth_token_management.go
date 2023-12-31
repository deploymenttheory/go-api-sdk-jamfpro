// http_client_auth_token_management.go
package http_client

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
	// If token doesn't exist
	if c.Token == "" {
		if c.BearerTokenAuthCredentials.Username != "" && c.BearerTokenAuthCredentials.Password != "" {
			err := c.ObtainToken()
			if err != nil {
				return false, fmt.Errorf("failed to obtain bearer token: %w", err)
			}
		} else if c.OAuthCredentials.ClientID != "" && c.OAuthCredentials.ClientSecret != "" {
			err := c.ObtainOAuthToken(c.OAuthCredentials)
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
			err = c.RefreshOAuthToken()
		} else {
			return false, fmt.Errorf("unknown auth method: %s", c.AuthMethod)
		}

		if err != nil {
			return false, fmt.Errorf("failed to refresh token: %w", err)
		}
	}

	return true, nil
}
