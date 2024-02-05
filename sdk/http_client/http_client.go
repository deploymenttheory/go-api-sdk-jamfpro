// http_client.go
/* The `http_client` package provides a configurable HTTP client tailored for interacting with specific APIs.
It supports different authentication methods, including "bearer" and "oauth". The client is designed with a
focus on concurrency management, structured error handling, and flexible configuration options.
The package offers a default timeout, custom backoff strategies, dynamic rate limiting,
and detailed logging capabilities. The main `Client` structure encapsulates all necessary components,
like the baseURL, authentication details, and an embedded standard HTTP client. */
package http_client

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Config holds configuration options for the HTTP Client.
type Config struct {
	// Required
	InstanceName string
	Auth         AuthConfig // User can either supply these values manually or pass from LoadAuthConfig/Env vars

	// Optional
	LogLevel                  LogLevel // Field for defining tiered logging level.
	MaxRetryAttempts          int      // Config item defines the max number of retry request attempts for retryable HTTP methods.
	EnableDynamicRateLimiting bool
	Logger                    Logger // Field for the packages initailzed logger
	MaxConcurrentRequests     int    // Field for defining the maximum number of concurrent requests allowed in the semaphore
	TokenRefreshBufferPeriod  time.Duration
	TotalRetryDuration        time.Duration
	CustomTimeout             time.Duration
}

// ClientPerformanceMetrics captures various metrics related to the client's
// interactions with the API, providing insights into its performance and behavior.
type PerformanceMetrics struct {
	TotalRequests        int64
	TotalRetries         int64
	TotalRateLimitErrors int64
	TotalResponseTime    time.Duration
	TokenWaitTime        time.Duration
	lock                 sync.Mutex
}

// ClientAuthConfig represents the structure to read authentication details from a JSON configuration file.
type AuthConfig struct {
	InstanceName       string `json:"instanceName,omitempty"`
	OverrideBaseDomain string `json:"overrideBaseDomain,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	ClientID           string `json:"clientID,omitempty"`
	ClientSecret       string `json:"clientSecret,omitempty"`
}

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
	PerfMetrics                PerformanceMetrics
}

// NewClient creates a new HTTP client with the provided configuration.
func NewClient(config Config) (*Client, error) {

	// Logging to track client setup process
	var logger Logger
	if config.Logger == nil {
		logger = NewDefaultLogger()
	}

	if config.LogLevel < LogLevelNone || config.LogLevel > LogLevelDebug {
		return nil, fmt.Errorf("invalid LogLevel")
	} else if config.LogLevel == 0 {
		logger.Info("LogLevel not set, setting to default value", "LogLevel", DefaultLogLevel)
		config.LogLevel = DefaultLogLevel
	}

	logger.SetLevel(config.LogLevel)

	// Config Validation & Default Value setting
	if config.InstanceName == "" {
		return nil, fmt.Errorf("instanceName cannot be empty")
	}

	if config.MaxRetryAttempts < 0 {
		logger.Info("MaxRetryAttempts cannot be negative, setting to default value", "MaxRetryAttempts", DefaultMaxRetryAttempts)
		config.MaxRetryAttempts = DefaultMaxRetryAttempts
	}

	if config.MaxConcurrentRequests <= 0 {
		logger.Info("MaxConcurrentRequests cannot be negative, setting to default value", "MaxConcurrentRequests", DefaultMaxConcurrentRequests)
		config.MaxConcurrentRequests = DefaultMaxConcurrentRequests
	}

	if config.TokenRefreshBufferPeriod < 0 {
		logger.Info("TokenRefreshBufferPeriod cannot be negative, setting to default value", "TokenRefreshBufferPeriod", DefaultTokenBufferPeriod)
		config.TokenRefreshBufferPeriod = DefaultTokenBufferPeriod
	}

	if config.TotalRetryDuration < 0 {
		logger.Info("TotalRetryDuration cannot be negative, setting to default value", "TotalRetryDuration", DefaultTotalRetryDuration)
		return nil, fmt.Errorf("TotalRetryDuration cannot be negative")
	}

	if config.TokenRefreshBufferPeriod == 0 {
		logger.Info("TokenRefreshBufferPeriod not set, setting to default value", "TokenRefreshBufferPeriod", DefaultTokenBufferPeriod)
		config.TokenRefreshBufferPeriod = 60 * time.Second
	}

	if config.TotalRetryDuration == 0 {
		logger.Info("TotalRetryDuration not set, setting to default value", "TotalRetryDuration", DefaultTotalRetryDuration)
		config.TotalRetryDuration = 60 * time.Second
	}

	if config.CustomTimeout == 0 {
		logger.Info("CustomTimeout not set, setting to default value", "CustomTimeout", DefaultTimeout)
		config.CustomTimeout = DefaultTimeout
	}

	var AuthMethod string
	if config.Auth.Username != "" && config.Auth.Password != "" {
		AuthMethod = "bearer"
	} else if config.Auth.ClientID != "" && config.Auth.ClientSecret != "" {
		AuthMethod = "oauth"
	} else {
		return nil, fmt.Errorf("invalid AuthConfig")
	}

	client := &Client{
		InstanceName:   config.InstanceName,
		httpClient:     &http.Client{Timeout: DefaultTimeout},
		AuthMethod:     AuthMethod,
		config:         config,
		logger:         logger,
		ConcurrencyMgr: NewConcurrencyManager(config.MaxConcurrentRequests, logger, config.LogLevel >= LogLevelDebug),
		PerfMetrics:    PerformanceMetrics{},
	}

	_, err := client.ValidAuthTokenCheck()
	if err != nil {
		return nil, fmt.Errorf("failed to validate auth: %w", err)
	}

	// Start the periodic metric evaluation for adjusting concurrency.
	go client.StartMetricEvaluation()

	if client.config.LogLevel >= LogLevelDebug {
		client.logger.Debug(
			"New client initialized with the following details:",
			"InstanceName", client.InstanceName,
			"Timeout", client.httpClient.Timeout,
			"TokenRefreshBufferPeriod", client.config.TokenRefreshBufferPeriod,
			"TotalRetryDuration", client.config.TotalRetryDuration,
			"MaxRetryAttempts", client.config.MaxRetryAttempts,
			"MaxConcurrentRequests", client.config.MaxConcurrentRequests,
			"EnableDynamicRateLimiting", client.config.EnableDynamicRateLimiting,
		)
	}

	return client, nil
}
