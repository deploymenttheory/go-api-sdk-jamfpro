// httpclient_oauth.go
/* The httpclient_auth package focuses on authentication mechanisms for an HTTP client.
It provides structures and methods for handling OAuth-based authentication
*/
package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// OAuthResponse represents the response structure when obtaining an OAuth access token.
type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Error        string `json:"error,omitempty"`
}

// OAuthCredentials contains the client ID and client secret required for OAuth authentication.
type OAuthCredentials struct {
	ClientID     string
	ClientSecret string
}

// SetOAuthCredentials sets the OAuth credentials (Client ID and Client Secret)
// for the client instance. These credentials are used for obtaining and refreshing
// OAuth tokens for authentication.
func (c *Client) SetOAuthCredentials(credentials OAuthCredentials) {
	c.OAuthCredentials = credentials
}

// ObtainOAuthToken fetches an OAuth access token using the provided OAuthCredentials (Client ID and Client Secret).
// It updates the client's Token and Expiry fields with the obtained values.
func (c *Client) ObtainOAuthToken(credentials AuthConfig) error {
	authenticationEndpoint := c.ConstructAPIAuthEndpoint(OAuthTokenEndpoint)
	data := url.Values{}
	data.Set("client_id", credentials.ClientID)
	data.Set("client_secret", credentials.ClientSecret)
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", authenticationEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Debug: Print the entire raw response body for inspection

	bodyBytes, _ := io.ReadAll(resp.Body)

	// Reset the response body to its original state
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	oauthResp := &OAuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(oauthResp)
	if err != nil {
		return err
	}

	if oauthResp.Error != "" { // Check if the "error" field is non-empty
		return fmt.Errorf("error obtaining OAuth token: %s", oauthResp.Error)
	}

	if oauthResp.AccessToken == "" {
		return fmt.Errorf("empty access token received")
	}

	// Calculate and format token expiration time
	expiresIn := time.Duration(oauthResp.ExpiresIn) * time.Second
	expirationTime := time.Now().Add(expiresIn)
	formattedExpirationTime := expirationTime.Format(time.RFC1123) // or any other preferred format

	// Log the token life expiry details in a human-readable format
	c.logger.Debug("The OAuth token obtained is: ",
		"Valid for", expiresIn.String(),
		"Expires at", formattedExpirationTime)

	c.Token = oauthResp.AccessToken
	c.Expiry = time.Now().Add(time.Second * time.Duration(oauthResp.ExpiresIn))

	return nil
}

// RefreshOAuthToken refreshes the current OAuth token.
// func (c *Client) RefreshOAuthToken() error {
// 	c.tokenLock.Lock()
// 	defer c.tokenLock.Unlock()

// 	tokenRefreshEndpoint := c.ConstructAPIAuthEndpoint(OAuthTokenEndpoint)

// 	req, err := http.NewRequest("POST", tokenRefreshEndpoint, nil)
// 	if err != nil {
// 		c.logger.Error("Failed to create new request for OAuth token refresh", "error", err)
// 		return err
// 	}
// 	req.Header.Add("Authorization", "Bearer "+c.Token)

// 	c.logger.Debug("Attempting to refresh OAuth token", "URL", tokenRefreshEndpoint)

// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		c.logger.Error("Failed to make request for OAuth token refresh", "error", err)
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		c.logger.Warn("OAuth token refresh response status is not OK", "StatusCode", resp.StatusCode)
// 		return c.HandleAPIError(resp)
// 	}

// 	tokenResp := &TokenResponse{}
// 	err = json.NewDecoder(resp.Body).Decode(tokenResp)
// 	if err != nil {
// 		c.logger.Error("Failed to decode OAuth token response", "error", err)
// 		return err
// 	}

// 	c.logger.Debug("OAuth token refreshed successfully", "Expiry", tokenResp.Expires)

// 	c.Token = tokenResp.Token
// 	c.Expiry = tokenResp.Expires
// 	return nil
// }

// InvalidateOAuthToken invalidates the current OAuth access token.
// After invalidation, the token cannot be used for further API requests.
func (c *Client) InvalidateOAuthToken() error {
	invalidateTokenEndpoint := c.ConstructAPIAuthEndpoint(TokenInvalidateEndpoint)
	req, err := http.NewRequest("POST", invalidateTokenEndpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to invalidate token, status code: %d", resp.StatusCode)
	}

	return nil
}
