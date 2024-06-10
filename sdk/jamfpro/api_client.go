package jamfpro

import (
	"github.com/deploymenttheory/go-api-http-client/httpclient"
)

type Client struct {
	HTTP *httpclient.Client
}

func BuildClient(config httpclient.ClientConfig) (*Client, error) {
	httpClient, err := httpclient.BuildClient(config, false)
	if err != nil {
		return nil, err
	}
	return &Client{HTTP: httpClient}, nil
}
