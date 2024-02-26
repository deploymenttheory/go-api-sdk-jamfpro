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

/*
// BuildClient initializes a new Jamf Pro client with the given configuration.
func BuildClient(config httpclient.ClientConfig) (*Client, error) {
	httpClient, err := httpclient.BuildClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{HTTP: httpClient}, nil
}
// LoadClientConfig loads the full configuration, including both AuthConfig and EnvironmentConfig, from a JSON file.
func LoadClientConfig(configFilePath string) (*ClientConfig, error) {
	bytes, err := os.ReadFile(configFilePath) // Use os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return nil, err
	}

	var config ClientConfig
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	// Validate loaded configuration to ensure necessary fields are populated
	if config.Auth.ClientID == "" || config.Auth.ClientSecret == "" {
		return nil, errors.New("authentication configuration incomplete")
	}
	if config.Environment.APIType == "" || config.Environment.InstanceName == "" {
		return nil, errors.New("environment configuration incomplete")
	}

	return &config, nil
}
// BuildClient initializes a new Jamf Pro client using the configuration file path for the HTTP client.
func BuildClient(httpClientConfigPath string) (*Client, error) {
	// Initialize the Jamf Pro client with optional HTTP client configuration
	httpClientConfig, err := httpclient.SetClientConfiguration(httpClientConfigPath)
	if err != nil {
		return nil, err
	}

	// Build the HTTP client with the loaded or set configuration
	httpClient, err := httpclient.BuildClient(*httpClientConfig)
	if err != nil {
		return nil, err
	}

	// Create the Jamf Pro client with the HTTP client
	jamfProClient := &Client{HTTP: httpClient}

	// Additional Jamf Pro specific settings can be applied here if necessary

	return jamfProClient, nil
}
*/

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
