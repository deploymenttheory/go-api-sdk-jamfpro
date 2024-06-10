package jamfpro

import (
	"github.com/deploymenttheory/go-api-http-client/httpclient"
)

type Client struct {
	HTTP *httpclient.Client
}

// func BuildClientWithOAuth(config httpclient.ClientConfig, clientId string, clientSecret string) (*Client, error) {
// 	parsedLogLevel := logger.ParseLogLevelFromString(config.LogLevel)
// 	logger := logger.BuildLogger(parsedLogLevel, config.LogOutputFormat, config.LogConsoleSeparator, config.LogExportPath, config.ExportLogs)

// 	jamfIntegration := jamfprointegration.BuildIntegrationWithOAuth(
// 		config.
// 	)

// 	httpClient, err := httpclient.BuildClient(config, false, logger)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Client{HTTP: httpClient}, nil
// }

// func BuildClientWithBasicAuth(config httpclient.ClientConfig) (*Client, error) {
// 	parsedLogLevel := logger.ParseLogLevelFromString(config.LogLevel)
// 	logger := logger.BuildLogger(parsedLogLevel, config.LogOutputFormat, config.LogConsoleSeparator, config.LogExportPath, config.ExportLogs)

// 	httpClient, err := httpclient.BuildClient(config, false, logger)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Client{HTTP: httpClient}, nil
// }
