package jamfpro

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/deploymenttheory/go-api-http-client-integrations/jamf/jamfprointegration"
	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
)

type Client struct {
	HTTP *httpclient.Client
}

type ConfigContainer struct {
	// Logger
	LogLevel            logger.LogLevel `json:"log_level"`
	Encoding            string          `json:"encoding"`
	LogConsoleSeparator string          `json:"log_console_separator"`
	LogFilepath         string          `json:"log_filepath"`
	ExportLogs          bool            `json:"export_logs"`

	// Integration
	BaseDomain           string `json:"base_domain"`
	AuthMethodDescriptor string `json:"auth_method_descriptor"`
	ClientId             string `json:"client_id"`
	ClientSecret         string `json:"client_secret"`
	BasicAuthUsername    string `json:"basic_auth_username"`
	BasicAuthPassword    string `json:"basic_auth_password"`

	// Client
	HTTPClientConfig *httpclient.ClientConfig `json:"http_client_config"`
}

// BuildClientWithConfigFile initializes a new Jamf Pro client using a
// configuration file for the HTTP client, logger, and integration.
// func BuildClientWithConfigFile(configFilePath string, httpClientConfig *httpclient.ClientConfig) (*Client, error) {
func BuildClientWithConfigFile(configFilePath string, httpClientConfig *httpclient.ClientConfig) (*Client, error) {
	config, err := loadConfigFromJSONFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load combined configuration from file: %w", err)
	}

	// Logger
	log := logger.BuildLogger(
		config.LogLevel,
		config.Encoding,
		config.LogConsoleSeparator,
		config.LogFilepath,
		config.ExportLogs,
	)

	var integration *jamfprointegration.Integration

	// Integration
	switch config.AuthMethodDescriptor {
	case "oauth2":
		integration, err = jamfprointegration.BuildIntegrationWithOAuth(
			config.BaseDomain,
			log,
			config.HTTPClientConfig.TokenRefreshBufferPeriod,
			config.ClientId,
			config.ClientSecret,
		)

	case "basic":
		integration, err = jamfprointegration.BuildIntegrationWithBasicAuth(
			config.BaseDomain,
			log,
			config.HTTPClientConfig.TokenRefreshBufferPeriod,
			config.BasicAuthUsername,
			config.BasicAuthPassword,
		)

	default:
		return nil, fmt.Errorf("invalid auth method supplied")

	}

	// HttpClient
	config.HTTPClientConfig.Integration = integration
	httpClient, err := httpclient.BuildClient(*config.HTTPClientConfig, true, log)
	if err != nil {
		return nil, fmt.Errorf("failed to build HTTP client: %w", err)
	}

	// Wrap into SDK & return

	return &Client{HTTP: httpClient}, nil
}

// loadCombinedConfig loads the combined configuration from a JSON file
func loadConfigFromJSONFile(configFilePath string) (*ConfigContainer, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	var config ConfigContainer
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	return &config, nil
}
