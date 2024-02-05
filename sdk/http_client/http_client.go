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
	InstanceName               string                     // Website Instance name without the root domain
	AuthMethod                 string                     // Specifies the authentication method: "bearer" or "oauth"
	Token                      string                     // Authentication Token
	OverrideBaseDomain         string                     // Base domain override used when the default in the api handler isn't suitable
	OAuthCredentials           OAuthCredentials           // ClientID / Client Secret
	BearerTokenAuthCredentials BearerTokenAuthCredentials // Username and Password for Basic Authentication
	Expiry                     time.Time                  // Expiry time set for the auth token
	httpClient                 *http.Client
	tokenLock                  sync.Mutex
	config                     Config
	logger                     Logger
	ConcurrencyMgr             *ConcurrencyManager
	PerfMetrics                ClientPerformanceMetrics
}

// Config holds configuration options for the HTTP Client.
type Config struct {
	LogLevel         LogLevel // Field for defining tiered logging level.
	MaxRetryAttempts int      // Config item defines the max number of retry request attempts for retryable HTTP methods.
	//CustomBackoff             func(attempt int) time.Duration
	EnableDynamicRateLimiting bool
	Logger                    Logger // Field for the packages initailzed logger
	MaxConcurrentRequests     int    // Field for defining the maximum number of concurrent requests allowed in the semaphore
	TokenLifespan             time.Duration
	TokenRefreshBufferPeriod  time.Duration
	TotalRetryDuration        time.Duration
}

// ClientPerformanceMetrics captures various metrics related to the client's
// interactions with the API, providing insights into its performance and behavior.
type ClientPerformanceMetrics struct {
	TotalRequests        int64
	TotalRetries         int64
	TotalRateLimitErrors int64
	TotalResponseTime    time.Duration
	TokenWaitTime        time.Duration
	lock                 sync.Mutex
}

// ClientAuthConfig represents the structure to read authentication details from a JSON configuration file.
type ClientAuthConfig struct {
	InstanceName       string `json:"instanceName,omitempty"`
	OverrideBaseDomain string `json:"overrideBaseDomain,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	ClientID           string `json:"clientID,omitempty"`
	ClientSecret       string `json:"clientSecret,omitempty"`
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

	// Validate MaxRetryAttempts
	if config.MaxRetryAttempts < 0 {
		return nil, fmt.Errorf("MaxRetryAttempts cannot be negative")
	}

	// Validate LogLevel
	if config.LogLevel < LogLevelNone || config.LogLevel > LogLevelDebug {
		return nil, fmt.Errorf("invalid LogLevel")
	}

	// Validate MaxConcurrentRequests
	if config.MaxConcurrentRequests < 0 {
		return nil, fmt.Errorf("MaxConcurrentRequests cannot be negative")
	}

	// Validate TokenLifespan
	if config.TokenLifespan < 0 {
		return nil, fmt.Errorf("TokenLifespan cannot be negative")
	}

	// Validate TokenRefreshBufferPeriod
	if config.TokenRefreshBufferPeriod < 0 {
		return nil, fmt.Errorf("TokenRefreshBufferPeriod cannot be negative")
	}

	// Validate TotalRetryDuration
	if config.TotalRetryDuration < 0 {
		return nil, fmt.Errorf("TotalRetryDuration cannot be negative")
	}

	// Default settings if not supplied
	if config.TokenLifespan == 0 {
		config.TokenLifespan = 30 * time.Minute
	}

	if config.TokenRefreshBufferPeriod == 0 {
		config.TokenRefreshBufferPeriod = 60 * time.Second
	}

	if config.TotalRetryDuration == 0 {
		config.TotalRetryDuration = 60 * time.Second
	}

	if logger == nil {
		logger = NewDefaultLogger()
	}

	// Set the log level of the logger
	logger.SetLevel(config.LogLevel)

	client := &Client{
		InstanceName:   instanceName,
		httpClient:     &http.Client{Timeout: DefaultTimeout},
		config:         config,
		logger:         logger,
		ConcurrencyMgr: NewConcurrencyManager(config.MaxConcurrentRequests, logger, config.LogLevel >= LogLevelDebug),
		PerfMetrics:    ClientPerformanceMetrics{},
	}

	// Apply any additional client options provided during initialization
	for _, opt := range options {
		opt(client)
	}

	// Start the periodic metric evaluation for adjusting concurrency.
	go client.StartMetricEvaluation()

	if client.config.LogLevel >= LogLevelDebug {
		client.logger.Debug(
			"New client initialized with the following details:",
			"InstanceName", client.InstanceName,
			"Timeout", client.httpClient.Timeout,
			"TokenLifespan", client.config.TokenLifespan,
			"TokenRefreshBufferPeriod", client.config.TokenRefreshBufferPeriod,
			"TotalRetryDuration", client.config.TotalRetryDuration,
			"MaxRetryAttempts", client.config.MaxRetryAttempts,
			"MaxConcurrentRequests", client.config.MaxConcurrentRequests,
			"EnableDynamicRateLimiting", client.config.EnableDynamicRateLimiting,
		)
	}

	return client, nil
}
