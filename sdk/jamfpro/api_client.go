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

// BuildClient initializes a new Jamf Pro client with the given configuration.
// This is typically used when you want to manually specify the configuration.
// e.g by another caller application such as terraform or a custom application.
func BuildClient(config httpclient.ClientConfig) (*Client, error) {
	httpClient, err := httpclient.BuildClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{HTTP: httpClient}, nil
}

// BuildClientWithEnv initializes a new Jamf Pro client using configurations
// loaded from environment variables. This is typically used when by a user to
// use environment variables to configure the client locally or when running
// in a container or a CI/CD pipeline.
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

// BuildClientWithConfigFile initializes a new Jamf Pro client using a
// configuration file for the HTTP client. This is typically used when a user
// wants to use a configuration file to configure the client locally.
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
