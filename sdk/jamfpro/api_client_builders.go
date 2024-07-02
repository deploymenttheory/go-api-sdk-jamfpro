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
	"go.uber.org/zap"
)

const jamfLoadBalancerCookieName = "jpro-ingress"

type Client struct {
	HTTP *httpclient.Client
}

type ConfigContainer struct {
	// Logger
	LogLevel          string `json:"log_level"`
	LogExportPath     string `json:"log_export_path"`
	HideSensitiveData bool   `json:"hide_sensitive_data"`

	// API Integration
	InstanceDomain       string `json:"instance_domain"`
	AuthMethod           string `json:"auth_method"`
	ClientID             string `json:"client_id"`
	ClientSecret         string `json:"client_secret"`
	Username             string `json:"basic_auth_username"`
	Password             string `json:"basic_auth_password"`
	JamfLoadBalancerLock bool   `json:"jamf_load_balancer_lock"`

	// Client
	CustomCookies               []CustomCookie `json:"custom_cookies"`
	MaxRetryAttempts            int            `json:"max_retry_attempts"`
	MaxConcurrentRequests       int            `json:"max_concurrent_requests"`
	EnableDynamicRateLimiting   bool           `json:"enable_dynamic_rate_limiting"`
	CustomTimeout               int            `json:"custom_timeout_seconds"`
	TokenRefreshBufferPeriod    int            `json:"token_refresh_buffer_period_seconds"`
	TotalRetryDuration          int            `json:"total_retry_duration_seconds"`
	FollowRedirects             bool           `json:"follow_redirects"`
	MaxRedirects                int            `json:"max_redirects"`
	EnableConcurrencyManagement bool           `json:"enable_concurrency_management"`
	MandatoryRequestDelay       int            `json:"mandatory_request_delay_milliseconds"`
	RetryEligiableRequests      bool           `json:"retry_eligiable_requests"`
}

type CustomCookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func BuildClient(config *ConfigContainer) (*Client, error) {
	var err error

	DefaultLoggerConfig := zap.NewProductionConfig()
	DefaultLoggerConfig.Level, err = ConvertLogLevel(config.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to set log level: %v", err)
	}

	if config.LogExportPath != "" {
		DefaultLoggerConfig.OutputPaths = append(DefaultLoggerConfig.OutputPaths, config.LogExportPath)
	}

	Logger, err := DefaultLoggerConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %v", err)
	}

	Sugar := Logger.Sugar()

	// Initialize API integration
	integration, err := initializeAPIIntegration(config, Sugar)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize integration: %w", err)
	}

	// Handle jamf pro load balancer lock and custom cookies
	customCookies, err := handleLoadBalancerLock(config, integration, convertCustomCookies(config.CustomCookies), Sugar)
	if err != nil {
		return nil, err
	}

	// HttpClient
	httpClientConfig := &httpclient.ClientConfig{
		Sugar:                       Sugar,
		Integration:                 integration,
		HideSensitiveData:           config.HideSensitiveData,
		CustomCookies:               customCookies,
		MaxRetryAttempts:            config.MaxRetryAttempts,
		MaxConcurrentRequests:       config.MaxConcurrentRequests,
		EnableDynamicRateLimiting:   config.EnableDynamicRateLimiting,
		CustomTimeout:               time.Duration(config.CustomTimeout) * time.Second,
		TokenRefreshBufferPeriod:    time.Duration(config.TokenRefreshBufferPeriod) * time.Second,
		TotalRetryDuration:          time.Duration(config.TotalRetryDuration) * time.Second,
		FollowRedirects:             config.FollowRedirects,
		MaxRedirects:                config.MaxRedirects,
		EnableConcurrencyManagement: config.EnableConcurrencyManagement,
		MandatoryRequestDelay:       time.Duration(config.MandatoryRequestDelay) * time.Millisecond,
		RetryEligiableRequests:      config.RetryEligiableRequests,
	}

	httpClient, err := httpClientConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build HTTP client: %w", err)
	}

	// Wrap into SDK & return
	return &Client{HTTP: httpClient}, nil
}

// BuildClientWithConfigFile initializes a new Jamf Pro client using a configuration file for the HTTP client, logger, and integration.
func BuildClientWithConfigFile(configFilePath string) (*Client, error) {
	config, err := loadConfigFromJSONFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration from file: %w", err)
	}

	return BuildClient(config)
}

// BuildClientWithEnv initializes a new Jamf Pro client using environment variables for the HTTP client, logger, and integration.
func BuildClientWithEnv() (*Client, error) {
	config, err := loadConfigFromEnv()
	if err != nil {
		return nil, err
	}

	return BuildClient(config)

}

// initializeAPIIntegration initializes the API integration based on the configuration
func initializeAPIIntegration(config *ConfigContainer, Sugar *zap.SugaredLogger) (httpclient.APIIntegration, error) {
	var integration *jamfprointegration.Integration
	var err error

	switch config.AuthMethod {
	case "oauth2":
		integration, err = jamfprointegration.BuildWithOAuth(
			config.InstanceDomain,
			Sugar,
			time.Duration(config.TokenRefreshBufferPeriod)*time.Second,
			config.ClientID,
			config.ClientSecret,
			config.HideSensitiveData,
		)
	case "basic":
		integration, err = jamfprointegration.BuildWithBasicAuth(
			config.InstanceDomain,
			Sugar,
			time.Duration(config.TokenRefreshBufferPeriod)*time.Second,
			config.Username,
			config.Password,
			config.HideSensitiveData,
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
		LogLevel: getEnv("LOG_LEVEL", "warning"),
		// ExportLogs:                  getEnvAsBool("EXPORT_LOGS", false),
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
		MandatoryRequestDelay:       getEnvAsInt("MANDATORY_REQUEST_DELAY_MILLISECONDS", 0),
		RetryEligiableRequests:      getEnvAsBool("RETRY_ELIGIABLE_REQUESTS", true),
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
func handleLoadBalancerLock(config *ConfigContainer, integration httpclient.APIIntegration, customCookies []*http.Cookie, Sugar *zap.SugaredLogger) ([]*http.Cookie, error) {
	if config.JamfLoadBalancerLock {
		jamfIntegration, ok := integration.(*jamfprointegration.Integration)
		if !ok {
			return nil, fmt.Errorf("integration is not of type *jamfprointegration.Integration")
		}
		cookies, err := jamfIntegration.GetSessionCookies()
		if err != nil {
			Sugar.Error("Failed to get session cookies for load balancer lock", zap.Error(err))
			return customCookies, nil
		}

		for _, cookie := range cookies {
			if cookie.Name == jamfLoadBalancerCookieName {
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

func ConvertLogLevel(inLevel string) (zap.AtomicLevel, error) {
	levelMap := map[string]zap.AtomicLevel{
		"debug":  zap.NewAtomicLevelAt(zap.DebugLevel),
		"info":   zap.NewAtomicLevelAt(zap.InfoLevel),
		"warn":   zap.NewAtomicLevelAt(zap.WarnLevel),
		"dpanic": zap.NewAtomicLevelAt(zap.DPanicLevel),
		"error":  zap.NewAtomicLevelAt(zap.ErrorLevel),
		"fatal":  zap.NewAtomicLevelAt(zap.FatalLevel),
	}

	outLevel, ok := levelMap[inLevel]
	if !ok {
		return zap.AtomicLevel{}, nil
	}

	return outLevel, nil

}
