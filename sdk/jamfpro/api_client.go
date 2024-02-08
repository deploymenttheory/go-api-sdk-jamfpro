package jamfpro

import (
	"github.com/deploymenttheory/go-api-http-client/httpclient"
)

type Client struct {
	HTTP *httpclient.Client
}

func BuildClient(config httpclient.Config) (*Client, error) {
	client, err := httpclient.BuildClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func LoadAuthConfig(configFilePath string) (*httpclient.AuthConfig, error) {
	return httpclient.LoadAuthConfig(configFilePath)
}
