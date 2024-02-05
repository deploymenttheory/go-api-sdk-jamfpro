package jamfpro

import (
	http_client "github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
)

type Client struct {
	HTTP *http_client.Client
}

func NewClient(config http_client.Config) (*Client, error) {
	client, err := http_client.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func LoadAuthConfig(configFilePath string) (*http_client.AuthConfig, error) {
	return http_client.LoadAuthConfig(configFilePath)
}
