package jamfpro

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
)

type Client struct {
	HTTP *httpclient.Client
}

// ClientConfig combines authentication and environment settings for the client.
type ClientConfig struct {
	Auth        httpclient.AuthConfig
	Environment httpclient.EnvironmentConfig
}

// BuildClient initializes a new Jamf Pro client with the given configuration.
func BuildClient(config httpclient.Config) (*Client, error) {
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

	// Debugging: Log the values loaded from the JSON file to verify correctness
	log.Printf("Loaded Configuration: %+v\n", config)

	// Validate loaded configuration to ensure necessary fields are populated
	if config.Auth.ClientID == "" || config.Auth.ClientSecret == "" {
		return nil, errors.New("authentication configuration incomplete")
	}
	if config.Environment.APIType == "" || config.Environment.InstanceName == "" {
		return nil, errors.New("environment configuration incomplete")
	}

	return &config, nil
}
