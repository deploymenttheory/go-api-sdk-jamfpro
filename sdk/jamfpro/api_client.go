package jamfpro

import (
	http_client "github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
)

type Client struct {
	client *http_client.Client
}

func (c *Client) NewClient(config http_client.Config) (*Client, error) {
	client, err := http_client.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}
