// http_client.go
/* The `http_client` package provides a configurable HTTP client tailored for interacting with specific APIs.
It supports different authentication methods, including "bearer" and "oauth". The client is designed with a
focus on concurrency management, structured error handling, and flexible configuration options.
The package offers a default timeout, custom backoff strategies, dynamic rate limiting,
and detailed logging capabilities. The main `Client` structure encapsulates all necessary components,
like the baseURL, authentication details, and an embedded standard HTTP client. */
package http_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

const DefaultTimeout = 10 * time.Second

// Client represents an HTTP client to interact with a specific API.
type Client struct {
	InstanceName               string
	authMethod                 string // Specifies the authentication method: "bearer" or "oauth"
	Token                      string
	oAuthCredentials           OAuthCredentials           // ClientID / Client Secret
	bearerTokenAuthCredentials BearerTokenAuthCredentials // Username and Password for Basic Authentication
	Expiry                     time.Time
	httpClient                 *http.Client
	tokenLock                  sync.Mutex
	config                     Config
	logger                     Logger
	ConcurrencyMgr             *ConcurrencyManager
}

// Config holds configuration options for the HTTP Client.
type Config struct {
	DebugMode                 bool
	MaxRetryAttempts          int
	CustomBackoff             func(attempt int) time.Duration
	EnableDynamicRateLimiting bool
	Logger                    Logger
	MaxConcurrentRequests     int
	TokenLifespan             time.Duration
	BufferPeriod              time.Duration
	TotalRetryDuration        time.Duration
}

// StructuredError represents a structured error response from the API.
type StructuredError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// ClientOption defines a function type for modifying client properties during initialization.
type ClientOption func(*Client)

// LoadClientAuthConfig reads a JSON configuration file and decodes it into a ClientAuthConfig struct.
// It is used to retrieve authentication details like BaseURL, Username, and Password for the client.
func LoadClientAuthConfig(filename string) (*ClientAuthConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &ClientAuthConfig{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// NewClient initializes a new http client instance with the given baseURL, logger, concurrency manager and client configuration
/*
If TokenLifespan and BufferPeriod aren't set in the config, they default to 30 minutes and 5 minutes, respectively.
If TotalRetryDuration isn't set in the config, it defaults to 1 minute.
If no logger is provided, a default logger will be used.
Any additional options provided will be applied to the client during initialization.
Detect authentication method based on supplied credential type
*/
func NewClient(instanceName string, config Config, logger Logger, options ...ClientOption) (*Client, error) {
	// Config Check
	if instanceName == "" {
		return nil, fmt.Errorf("instanceName cannot be empty")
	}
	// Default settings if not supplied
	if config.TokenLifespan == 0 {
		config.TokenLifespan = 30 * time.Minute
	}
	if config.BufferPeriod == 0 {
		config.BufferPeriod = 5 * time.Minute
	}
	if config.TotalRetryDuration == 0 {
		config.TotalRetryDuration = 60 * time.Second
	}

	if logger == nil {
		logger = &defaultLogger{}
	}

	client := &Client{
		InstanceName:   instanceName,
		httpClient:     &http.Client{Timeout: DefaultTimeout},
		config:         config,
		logger:         logger,
		ConcurrencyMgr: NewConcurrencyManager(config.MaxConcurrentRequests, logger, config.DebugMode),
	}

	// Apply any additional client options provided during initialization
	for _, opt := range options {
		opt(client)
	}

	if client.config.DebugMode {
		client.logger.Debug(
			"New client initialized with the following details:",
			"InstanceName", client.InstanceName,
			"Timeout", client.httpClient.Timeout,
			"TokenLifespan", client.config.TokenLifespan,
			"BufferPeriod", client.config.BufferPeriod,
			"TotalRetryDuration", client.config.TotalRetryDuration,
			"MaxRetryAttempts", client.config.MaxRetryAttempts,
			"MaxConcurrentRequests", client.config.MaxConcurrentRequests,
			"EnableDynamicRateLimiting", client.config.EnableDynamicRateLimiting,
		)
	}

	return client, nil
}
