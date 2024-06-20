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

// LoggerConfig holds the configuration for the logger.
type LoggerConfig struct {
	LogLevel            logger.LogLevel `json:"log_level"`
	Encoding            string          `json:"encoding"`
	LogConsoleSeparator string          `json:"log_console_separator"`
	LogFilepath         string          `json:"log_filepath"`
	ExportLogs          bool            `json:"export_logs"`
}

// IntegrationConfig holds the configuration for the API integration.
type IntegrationConfig struct {
	BaseDomain           string `json:"base_domain"`
	AuthMethodDescriptor string `json:"auth_method_descriptor"`
}

// CombinedConfig combines HTTP client, logger, and integration configurations.
type CombinedConfig struct {
	HTTPClientConfig  *httpclient.ClientConfig `json:"http_client_config"`
	LoggerConfig      LoggerConfig             `json:"logger_config"`
	IntegrationConfig IntegrationConfig        `json:"api_integration_config"`
}

// BuildClientWithConfigFile initializes a new Jamf Pro client using a
// configuration file for the HTTP client, logger, and integration.
func BuildClientWithConfigFile(configFilePath string) (*Client, error) {
	// Load the combined configuration from the specified file
	combinedConfig, err := loadCombinedConfig(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load combined configuration from file: %w", err)
	}

	// Initialize a logger using the loaded logger configuration
	log := logger.BuildLogger(
		combinedConfig.LoggerConfig.LogLevel,
		combinedConfig.LoggerConfig.Encoding,
		combinedConfig.LoggerConfig.LogConsoleSeparator,
		combinedConfig.LoggerConfig.LogFilepath,
		combinedConfig.LoggerConfig.ExportLogs,
	)

	// Create the API integration using the loaded integration configuration
	integration := &jamfprointegration.Integration{
		BaseDomain:           combinedConfig.IntegrationConfig.BaseDomain,
		AuthMethodDescriptor: combinedConfig.IntegrationConfig.AuthMethodDescriptor,
		Logger:               log,
	}

	// Assign the integration to the HTTP client configuration
	combinedConfig.HTTPClientConfig.Integration = integration

	// Build the HTTP client with the loaded configuration
	httpClient, err := httpclient.BuildClient(*combinedConfig.HTTPClientConfig, true, log)
	if err != nil {
		return nil, fmt.Errorf("failed to build HTTP client: %w", err)
	}

	// Create and return the Jamf Pro client with the HTTP client
	return &Client{HTTP: httpClient}, nil
}

// loadCombinedConfig loads the combined configuration from a JSON file
func loadCombinedConfig(configFilePath string) (*CombinedConfig, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	var config CombinedConfig
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	return &config, nil
}
