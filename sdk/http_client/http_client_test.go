package http_client

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClientValidConfig(t *testing.T) {
	assert := assert.New(t)

	// Define a complete config with some fields explicitly set
	config := Config{
		MaxRetryAttempts:          3,
		LogLevel:                  LogLevelInfo, // Example log level
		EnableDynamicRateLimiting: true,         // Explicitly enabling rate limiting
		Logger:                    nil,          // Assuming nil will use the default logger
		MaxConcurrentRequests:     10,           // Setting max concurrent requests
		// TokenLifespan, TokenRefreshBufferPeriod, and TotalRetryDuration will use default values
	}
	logger := NewDefaultLogger() // Using a default logger for the test
	client, err := NewClient("testInstance", config, logger)

	assert.NoError(err, "Expected no error during client initialization")
	assert.NotNil(client, "Expected client to be initialized, got nil")

	// Assert that the explicit and default config values are correctly set
	assert.Equal(config.MaxRetryAttempts, client.config.MaxRetryAttempts, "MaxRetryAttempts should match config")
	assert.Equal(LogLevelInfo, client.config.LogLevel, "LogLevel should match config")
	assert.Equal(true, client.config.EnableDynamicRateLimiting, "EnableDynamicRateLimiting should match config")
	assert.NotNil(client.logger, "Logger should not be nil")
	assert.Equal(10, client.config.MaxConcurrentRequests, "MaxConcurrentRequests should match config")

	// Default values
	defaultTokenLifespan := 30 * time.Minute
	defaultTokenRefreshBufferPeriod := 5 * time.Minute
	defaultTotalRetryDuration := 60 * time.Second

	assert.Equal(defaultTokenLifespan, client.config.TokenLifespan, "TokenLifespan should have a default value")
	assert.Equal(defaultTokenRefreshBufferPeriod, client.config.TokenRefreshBufferPeriod, "TokenRefreshBufferPeriod should have a default value")
	assert.Equal(defaultTotalRetryDuration, client.config.TotalRetryDuration, "TotalRetryDuration should have a default value")
}

func TestNewClientInvalidConfig(t *testing.T) {
	assert := assert.New(t)

	// Test with an empty instance name
	_, err := NewClient("", Config{}, nil)
	assert.Error(err, "Expected error for empty instance name")

	// Test with a negative MaxRetryAttempts
	configWithNegativeMaxRetry := Config{
		MaxRetryAttempts: -1,
	}
	_, err = NewClient("testInstance", configWithNegativeMaxRetry, nil)
	assert.Error(err, "Expected error for negative MaxRetryAttempts")

	// Test with invalid log level (assuming you have defined constraints for valid log levels)
	configWithInvalidLogLevel := Config{
		LogLevel: LogLevel(-1), // Invalid log level
	}
	_, err = NewClient("testInstance", configWithInvalidLogLevel, nil)
	assert.Error(err, "Expected error for invalid LogLevel")

	// Test with a negative MaxConcurrentRequests
	configWithNegativeConcurrentRequests := Config{
		MaxConcurrentRequests: -1,
	}
	_, err = NewClient("testInstance", configWithNegativeConcurrentRequests, nil)
	assert.Error(err, "Expected error for negative MaxConcurrentRequests")

	// Test with a negative TokenLifespan
	configWithNegativeTokenLifespan := Config{
		TokenLifespan: -1 * time.Minute,
	}
	_, err = NewClient("testInstance", configWithNegativeTokenLifespan, nil)
	assert.Error(err, "Expected error for negative TokenLifespan")

	// Test with a negative TokenRefreshBufferPeriod
	configWithNegativeTokenRefreshBufferPeriod := Config{
		TokenRefreshBufferPeriod: -1 * time.Minute,
	}
	_, err = NewClient("testInstance", configWithNegativeTokenRefreshBufferPeriod, nil)
	assert.Error(err, "Expected error for negative TokenRefreshBufferPeriod")

	// Test with a negative TotalRetryDuration
	configWithNegativeTotalRetryDuration := Config{
		TotalRetryDuration: -1 * time.Second,
	}
	_, err = NewClient("testInstance", configWithNegativeTotalRetryDuration, nil)
	assert.Error(err, "Expected error for negative TotalRetryDuration")
}

func TestClientOptionsApplication(t *testing.T) {
	assert := assert.New(t)

	customOption := func(c *Client) {
		c.config.MaxConcurrentRequests = 5
	}
	client, err := NewClient("testInstance", Config{}, nil, customOption)
	assert.NoError(err, "Client initialization should not return an error")

	assert.Equal(5, client.config.MaxConcurrentRequests, "Expected MaxConcurrentRequests to be set to 5")
}

func createTempConfigFile(content []byte) (string, error) {
	tmpfile, err := os.CreateTemp("", "config*.json")
	if err != nil {
		return "", err
	}

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
		return "", err
	}

	if err := tmpfile.Close(); err != nil {
		os.Remove(tmpfile.Name())
		return "", err
	}

	return tmpfile.Name(), nil
}

func TestLoadClientAuthConfig_Success(t *testing.T) {
	assert := assert.New(t)

	// Create a temporary config file
	config := ClientAuthConfig{
		InstanceName: "testInstance",
		ClientID:     "testClientID",
		ClientSecret: "testClientSecret",
	}
	content, _ := json.Marshal(config)
	filename, err := createTempConfigFile(content)
	require.NoError(t, err)
	defer os.Remove(filename)

	// Test loading the config
	loadedConfig, err := LoadClientAuthConfig(filename)
	assert.NoError(err)
	assert.Equal(config, *loadedConfig)
}

func TestLoadClientAuthConfig_FileNotFound(t *testing.T) {
	assert := assert.New(t)

	_, err := LoadClientAuthConfig("nonexistent.json")
	assert.Error(err)
}

func TestLoadClientAuthConfig_InvalidJSON(t *testing.T) {
	assert := assert.New(t)

	// Create a temporary config file with invalid JSON
	filename, err := createTempConfigFile([]byte("{invalid json"))
	require.NoError(t, err)
	defer os.Remove(filename)

	// Test loading the config
	_, err = LoadClientAuthConfig(filename)
	assert.Error(err)
}

func TestNewClientDefaultSettings(t *testing.T) {
	assert := assert.New(t)

	// Test client initialization with minimal configuration
	config := Config{
		// Only set the fields that are absolutely necessary
		MaxRetryAttempts: 1,
	}
	client, err := NewClient("testInstance", config, nil)

	assert.NoError(err, "Client initialization with minimal config should not return an error")
	assert.NotNil(client, "Expected client to be initialized")
	assert.NotNil(client.logger, "Expected default logger to be set")

	// Test default values
	assert.Equal(DefaultTimeout, client.httpClient.Timeout, "Expected default timeout to be set")
	assert.Equal(30*time.Minute, client.config.TokenLifespan, "Expected default TokenLifespan to be set")
	assert.Equal(5*time.Minute, client.config.TokenRefreshBufferPeriod, "Expected default TokenRefreshBufferPeriod to be set")
	assert.Equal(60*time.Second, client.config.TotalRetryDuration, "Expected default TotalRetryDuration to be set")
}
