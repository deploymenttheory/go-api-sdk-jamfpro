// http_helpers.go
package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// ParseISO8601Date attempts to parse a string date in ISO 8601 format.
func ParseISO8601Date(dateStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateStr)
}

// EnsureHTTPScheme prefixes a URL with "http://" it defaults to "https://" doesn't already have an "https://".
func EnsureHTTPScheme(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return fmt.Sprintf("https://%s", url)
	}
	return url
}

// CheckDeprecationHeader checks the response headers for the Deprecation header and logs a warning if present.
func CheckDeprecationHeader(resp *http.Response, logger Logger) {
	deprecationHeader := resp.Header.Get("Deprecation")
	if deprecationHeader != "" {
		logger.Warn("API endpoint is deprecated as of", "Date", deprecationHeader)
	}
}

// SetAuthenticationCredentials interprets and sets the credentials for the Client.
func (c *Client) SetAuthenticationCredentials(creds map[string]string) {
	// Check for OAuth credentials
	if clientID, ok := creds["clientID"]; ok {
		if clientSecret, ok := creds["clientSecret"]; ok {
			c.OAuthCredentials = OAuthCredentials{
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
			c.AuthMethod = "oauth"
			return
		}
	}

	// Check for Bearer Token credentials
	if username, ok := creds["username"]; ok {
		if password, ok := creds["password"]; ok {
			c.BearerTokenAuthCredentials = BearerTokenAuthCredentials{
				Username: username,
				Password: password,
			}
			c.AuthMethod = "bearer"
			return
		}
	}
}

// GetOAuthCredentials retrieves the current OAuth credentials (Client ID and Client Secret)
// set for the client instance. Used for test cases.
func (c *Client) GetOAuthCredentials() OAuthCredentials {
	return c.OAuthCredentials
}

// GetBearerAuthCredentials retrieves the current bearer auth credentials (Username and Password)
// set for the client instance. Used for test cases.
func (c *Client) GetBearerAuthCredentials() BearerTokenAuthCredentials {
	return c.BearerTokenAuthCredentials
}

// LoadAuthConfig reads a JSON configuration file and decodes it into a ClientAuthConfig struct.
// It is used to retrieve authentication details like BaseURL, Username, and Password for the client.
func LoadAuthConfig(filename string) (*AuthConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &AuthConfig{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
