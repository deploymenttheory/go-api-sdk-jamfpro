// shared_api_client.go
// The jamfpro package offers a client for interacting with the Jamf Pro API.
// This client extends the foundational capabilities of the http_client package,
// adding methods specifically tailored for Jamf Pro's API endpoints.
// By embedding the http_client's Client, it leverages core HTTP methods,
// authentication mechanisms, and other utilities, while also enabling
// Jamf Pro-specific functionalities.

package jamfpro

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
)

const (
	concurrentRequests           = 10 // Number of simultaneous requests.
	maxConcurrentRequestsAllowed = 5  // Maximum allowed concurrent requests.
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
)

type Client struct {
	HTTP *http_client.Client
}

type Config struct {
	InstanceName             string
	DebugMode                bool
	Logger                   http_client.Logger
	MaxConcurrentRequests    int
	TokenLifespan            time.Duration
	TokenRefreshBufferPeriod time.Duration
	ClientID                 string
	ClientSecret             string
}

func NewClient(config Config) (*Client, error) {

	// If not provided, use the default values from constants
	if config.MaxConcurrentRequests == 0 {
		config.MaxConcurrentRequests = maxConcurrentRequestsAllowed
	}
	if config.TokenLifespan == 0 {
		config.TokenLifespan = defaultTokenLifespan
	}
	if config.TokenRefreshBufferPeriod == 0 {
		config.TokenRefreshBufferPeriod = defaultBufferPeriod
	}

	httpConfig := http_client.Config{
		DebugMode:                config.DebugMode,
		Logger:                   config.Logger,
		MaxConcurrentRequests:    config.MaxConcurrentRequests,
		TokenLifespan:            config.TokenLifespan,
		TokenRefreshBufferPeriod: config.TokenRefreshBufferPeriod,
	}

	httpCli, err := http_client.NewClient(config.InstanceName, httpConfig, nil)
	if err != nil {
		return nil, err // Return the error if HTTP client initialization fails
	}

	client := &Client{
		HTTP: httpCli,
	}

	// Set auth credential configuration
	creds := http_client.OAuthCredentials{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
	}
	client.SetClientOAuthCredentials(creds)

	// validate credentials oauth credentials exist
	if client.HTTP.GetOAuthCredentials().ClientID == "" || client.HTTP.GetOAuthCredentials().ClientSecret == "" {
		return nil, fmt.Errorf("OAuth credentials (ClientID and ClientSecret) must be provided")
	}

	return client, nil
}

func (c *Client) SetClientOAuthCredentials(creds http_client.OAuthCredentials) {
	c.HTTP.SetOAuthCredentials(creds)
}

func (c *Client) SetAuthenticationCredentials(creds map[string]string) {
	c.HTTP.SetAuthenticationCredentials(creds)
}

// LoadClientAuthConfig reads a JSON configuration file and decodes it into a ClientAuthConfig struct.
func LoadClientAuthConfig(filename string) (*http_client.ClientAuthConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &http_client.ClientAuthConfig{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
