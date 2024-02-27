package jamfpro

import (
	"fmt"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
)

type Client struct {
	HTTP *httpclient.Client
}

// ClientConfig combines authentication and environment settings for the client.
type ClientConfig struct {
	Auth          httpclient.AuthConfig
	Environment   httpclient.EnvironmentConfig
	ClientOptions httpclient.ClientOptions
}

// BuildClient initializes a new Jamf Pro client using the provided configuration.
func BuildClient(clientConfig *ClientConfig) (*httpclient.Client, error) {
	httpClientConfig := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			Username:     clientConfig.Auth.Username,
			Password:     clientConfig.Auth.Password,
			ClientID:     clientConfig.Auth.ClientID,
			ClientSecret: clientConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			InstanceName:       clientConfig.Environment.InstanceName,
			OverrideBaseDomain: clientConfig.Environment.OverrideBaseDomain,
			APIType:            clientConfig.Environment.APIType,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:                  clientConfig.ClientOptions.LogLevel,
			LogOutputFormat:           clientConfig.ClientOptions.LogOutputFormat,
			LogConsoleSeparator:       clientConfig.ClientOptions.LogConsoleSeparator,
			HideSensitiveData:         clientConfig.ClientOptions.HideSensitiveData,
			MaxRetryAttempts:          clientConfig.ClientOptions.MaxRetryAttempts,
			EnableDynamicRateLimiting: clientConfig.ClientOptions.EnableDynamicRateLimiting,
			MaxConcurrentRequests:     clientConfig.ClientOptions.MaxConcurrentRequests,
			TokenRefreshBufferPeriod:  clientConfig.ClientOptions.TokenRefreshBufferPeriod,
			TotalRetryDuration:        clientConfig.ClientOptions.TotalRetryDuration,
			CustomTimeout:             clientConfig.ClientOptions.CustomTimeout,
		},
	}

	// Call the httpclient.BuildClient function with the adapted configuration
	return httpclient.BuildClient(httpClientConfig)
}

// BuildClientWithEnv initializes a new Jamf Pro client using configurations loaded from environment variables.
func BuildClientWithEnv() (*Client, error) {
	// Create a new empty ClientConfig
	config := &httpclient.ClientConfig{}

	// Load configurations from environment variables
	loadedConfig, err := httpclient.LoadConfigFromEnv(config)
	if err != nil {
		return nil, fmt.Errorf("failed to load HTTP client configuration from environment variables: %w", err)
	}

	// Build the HTTP client with the loaded configuration
	httpClient, err := httpclient.BuildClient(*loadedConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to build HTTP client: %w", err)
	}

	// Create and return the Jamf Pro client with the HTTP client
	return &Client{HTTP: httpClient}, nil
}

// BuildClientWithConfigFile initializes a new Jamf Pro client using a configuration file for the HTTP client.
func BuildClientWithConfigFile(configFilePath string) (*Client, error) {
	// Load the HTTP client configuration from the specified file
	loadedConfig, err := httpclient.LoadConfigFromFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load HTTP client configuration from file: %w", err)
	}

	// Build the HTTP client with the loaded configuration
	httpClient, err := httpclient.BuildClient(*loadedConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to build HTTP client: %w", err)
	}

	// Create and return the Jamf Pro client with the HTTP client
	return &Client{HTTP: httpClient}, nil
}
