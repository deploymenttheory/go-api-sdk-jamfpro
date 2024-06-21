package jamfpro

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/deploymenttheory/go-api-http-client-integrations/jamf/jamfprointegration"
	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"go.uber.org/zap"
)

const jamfLoadBalancerCookieName = "jpro-ingress"

type Client struct {
	HTTP *httpclient.Client
}

type ConfigContainer struct {
	// Logger
	LogLevel            string `json:"log_level"`
	LogOutputFormat     string `json:"log_output_format"`
	LogConsoleSeparator string `json:"log_console_separator"`
	LogExportPath       string `json:"log_export_path"`
	ExportLogs          bool   `json:"export_logs"`
	HideSensitiveData   bool   `json:"hide_sensitive_data"`

	// API Integration
	InstanceDomain string `json:"instance_domain"`
	AuthMethod     string `json:"auth_method"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	Username       string `json:"basic_auth_username"`
	Password       string `json:"basic_auth_password"`

	// Client
	CustomCookies               []CustomCookie `json:"custom_cookies"`
	JamfLoadBalancerLock        bool           `json:"jamf_load_balancer_lock"`
	MaxRetryAttempts            int            `json:"max_retry_attempts"`
	EnableDynamicRateLimiting   bool           `json:"enable_dynamic_rate_limiting"`
	MaxConcurrentRequests       int            `json:"max_concurrent_requests"`
	TokenRefreshBufferPeriod    int            `json:"token_refresh_buffer_period_seconds"`
	TotalRetryDuration          int            `json:"total_retry_duration_seconds"`
	CustomTimeout               int            `json:"custom_timeout_seconds"`
	FollowRedirects             bool           `json:"follow_redirects"`
	MaxRedirects                int            `json:"max_redirects"`
	EnableConcurrencyManagement bool           `json:"enable_concurrency_management"`
}

type CustomCookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// BuildClientWithConfigFile initializes a new Jamf Pro client using a configuration file for the HTTP client, logger, and integration.
func BuildClientWithConfigFile(configFilePath string) (*Client, error) {
	config, err := loadConfigFromJSONFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration from file: %w", err)
	}

	// Initialize logger
	logLevel := logger.ParseLogLevelFromString(config.LogLevel)
	log := logger.BuildLogger(
		logLevel,
		config.LogOutputFormat,
		config.LogConsoleSeparator,
		config.LogExportPath,
		config.ExportLogs,
	)

	// Initialize API integration
	integration, err := initializeAPIIntegration(config, log)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize integration: %w", err)
	}

	// Handle jamf pro load balancer lock and custom cookies
	customCookies, err := handleLoadBalancerLock(config, integration, convertCustomCookies(config.CustomCookies), log)
	if err != nil {
		return nil, err
	}

	// HttpClient
	httpClientConfig := &httpclient.ClientConfig{
		Integration:                 integration,
		HideSensitiveData:           config.HideSensitiveData,
		MaxRetryAttempts:            config.MaxRetryAttempts,
		MaxConcurrentRequests:       config.MaxConcurrentRequests,
		EnableDynamicRateLimiting:   config.EnableDynamicRateLimiting,
		CustomTimeout:               time.Duration(config.CustomTimeout) * time.Second,
		TokenRefreshBufferPeriod:    time.Duration(config.TokenRefreshBufferPeriod) * time.Second,
		TotalRetryDuration:          time.Duration(config.TotalRetryDuration) * time.Second,
		FollowRedirects:             config.FollowRedirects,
		MaxRedirects:                config.MaxRedirects,
		EnableConcurrencyManagement: config.EnableConcurrencyManagement,
		CustomCookies:               customCookies,
	}

	httpClient, err := httpclient.BuildClient(*httpClientConfig, true, log)
	if err != nil {
		return nil, fmt.Errorf("failed to build HTTP client: %w", err)
	}

	// Wrap into SDK & return
	return &Client{HTTP: httpClient}, nil
}

// BuildClientWithEnv initializes a new Jamf Pro client using environment variables for the HTTP client, logger, and integration.
func BuildClientWithEnv() (*Client, error) {
	config, err := loadConfigFromEnv()
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration from environment: %w", err)
	}

	// Initialize logger
	logLevel := logger.ParseLogLevelFromString(config.LogLevel)
	log := logger.BuildLogger(
		logLevel,
		config.LogOutputFormat,
		config.LogConsoleSeparator,
		config.LogExportPath,
		config.ExportLogs,
	)

	// Initialize API integration
	integration, err := initializeAPIIntegration(config, log)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize integration: %w", err)
	}

	// Handle jamf pro load balancer lock and custom cookies
	customCookies, err := handleLoadBalancerLock(config, integration, convertCustomCookies(config.CustomCookies), log)
	if err != nil {
		return nil, err
	}

	// HttpClient
	httpClientConfig := &httpclient.ClientConfig{
		Integration:                 integration,
		HideSensitiveData:           config.HideSensitiveData,
		MaxRetryAttempts:            config.MaxRetryAttempts,
		MaxConcurrentRequests:       config.MaxConcurrentRequests,
		EnableDynamicRateLimiting:   config.EnableDynamicRateLimiting,
		CustomTimeout:               time.Duration(config.CustomTimeout) * time.Second,
		TokenRefreshBufferPeriod:    time.Duration(config.TokenRefreshBufferPeriod) * time.Second,
		TotalRetryDuration:          time.Duration(config.TotalRetryDuration) * time.Second,
		FollowRedirects:             config.FollowRedirects,
		MaxRedirects:                config.MaxRedirects,
		EnableConcurrencyManagement: config.EnableConcurrencyManagement,
		CustomCookies:               customCookies,
	}

	httpClient, err := httpclient.BuildClient(*httpClientConfig, true, log)
	if err != nil {
		return nil, fmt.Errorf("failed to build HTTP client: %w", err)
	}

	// Wrap into SDK & return
	return &Client{HTTP: httpClient}, nil
}

// initializeAPIIntegration initializes the API integration based on the configuration
func initializeAPIIntegration(config *ConfigContainer, log logger.Logger) (httpclient.APIIntegration, error) {
	var integration *jamfprointegration.Integration
	var err error

	switch config.AuthMethod {
	case "oauth2":
		integration, err = jamfprointegration.BuildIntegrationWithOAuth(
			config.InstanceDomain,
			log,
			time.Duration(config.TokenRefreshBufferPeriod)*time.Second,
			config.ClientID,
			config.ClientSecret,
		)
	case "basic":
		integration, err = jamfprointegration.BuildIntegrationWithBasicAuth(
			config.InstanceDomain,
			log,
			time.Duration(config.TokenRefreshBufferPeriod)*time.Second,
			config.Username,
			config.Password,
		)
	default:
		return nil, fmt.Errorf("invalid auth method supplied")
	}

	if err != nil {
		return nil, err
	}

	return integration, nil
}

// loadConfigFromJSONFile loads the configuration from a JSON file
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

// loadConfigFromEnv loads the configuration from environment variables
func loadConfigFromEnv() (*ConfigContainer, error) {
	config := &ConfigContainer{
		LogLevel:                    getEnv("LOG_LEVEL", "warning"),
		LogOutputFormat:             getEnv("LOG_OUTPUT_FORMAT", "pretty"),
		LogConsoleSeparator:         getEnv("LOG_CONSOLE_SEPARATOR", " "),
		LogExportPath:               getEnv("LOG_EXPORT_PATH", ""),
		ExportLogs:                  getEnvAsBool("EXPORT_LOGS", false),
		HideSensitiveData:           getEnvAsBool("HIDE_SENSITIVE_DATA", true),
		InstanceDomain:              getEnv("INSTANCE_DOMAIN", ""),
		AuthMethod:                  getEnv("AUTH_METHOD", ""),
		ClientID:                    getEnv("CLIENT_ID", ""),
		ClientSecret:                getEnv("CLIENT_SECRET", ""),
		Username:                    getEnv("BASIC_AUTH_USERNAME", ""),
		Password:                    getEnv("BASIC_AUTH_PASSWORD", ""),
		JamfLoadBalancerLock:        getEnvAsBool("JAMF_LOAD_BALANCER_LOCK", false),
		MaxRetryAttempts:            getEnvAsInt("MAX_RETRY_ATTEMPTS", 3),
		EnableDynamicRateLimiting:   getEnvAsBool("ENABLE_DYNAMIC_RATE_LIMITING", false),
		MaxConcurrentRequests:       getEnvAsInt("MAX_CONCURRENT_REQUESTS", 1),
		TokenRefreshBufferPeriod:    getEnvAsInt("TOKEN_REFRESH_BUFFER_PERIOD_SECONDS", 300),
		TotalRetryDuration:          getEnvAsInt("TOTAL_RETRY_DURATION_SECONDS", 60),
		CustomTimeout:               getEnvAsInt("CUSTOM_TIMEOUT_SECONDS", 60),
		FollowRedirects:             getEnvAsBool("FOLLOW_REDIRECTS", true),
		MaxRedirects:                getEnvAsInt("MAX_REDIRECTS", 5),
		EnableConcurrencyManagement: getEnvAsBool("ENABLE_CONCURRENCY_MANAGEMENT", true),
		CustomCookies:               convertCustomCookiesFromEnv(getEnv("CUSTOM_COOKIES", "")),
	}
	return config, nil
}

// convertCustomCookiesFromEnv converts environment variable string to custom cookie configuration
func convertCustomCookiesFromEnv(customCookiesStr string) []CustomCookie {
	var customCookies []CustomCookie
	if customCookiesStr == "" {
		return customCookies
	}
	err := json.Unmarshal([]byte(customCookiesStr), &customCookies)
	if err != nil {
		fmt.Printf("Error parsing custom cookies from environment: %v\n", err)
		return nil
	}
	return customCookies
}

// getEnv gets the environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// getEnvAsBool gets the environment variable as a boolean or returns a default value
func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// getEnvAsInt gets the environment variable as an integer or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// convertCustomCookies converts custom cookie configuration into http.Cookie objects for client build
func convertCustomCookies(customCookies []CustomCookie) []*http.Cookie {
	var cookies []*http.Cookie
	for _, c := range customCookies {
		cookies = append(cookies, &http.Cookie{
			Name:  c.Name,
			Value: c.Value,
		})
	}
	return cookies
}

// handleLoadBalancerLock handles the load balancer lock by adding appropriate cookies if enabled
func handleLoadBalancerLock(config *ConfigContainer, integration httpclient.APIIntegration, customCookies []*http.Cookie, log logger.Logger) ([]*http.Cookie, error) {
	if config.JamfLoadBalancerLock {
		jamfIntegration, ok := integration.(*jamfprointegration.Integration)
		if !ok {
			return nil, fmt.Errorf("integration is not of type *jamfprointegration.Integration")
		}
		cookies, err := jamfIntegration.GetSessionCookies()
		if err != nil {
			log.Error("Failed to get session cookies for load balancer lock", zap.Error(err))
			return customCookies, nil
		}

		for _, cookie := range cookies {
			if cookie.Name == jamfLoadBalancerCookieName {
				// Ensure no custom cookie conflicts with the load balancer lock cookie
				for i, customCookie := range customCookies {
					if customCookie.Name == jamfLoadBalancerCookieName {
						customCookies = append(customCookies[:i], customCookies[i+1:]...)
						break
					}
				}
				customCookies = append(customCookies, cookie)
			}
		}
	}
	return customCookies, nil
}
