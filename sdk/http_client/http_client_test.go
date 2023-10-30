package http_client

import (
	"os"
	"testing"
	"time"
)

// Mock Logger for testing
// Mock Logger for testing
type mockLogger struct{}

func (m *mockLogger) Trace(msg string, keysAndValues ...interface{}) {}

func (m *mockLogger) Debug(msg string, keysAndValues ...interface{}) {}

func (m *mockLogger) Info(msg string, keysAndValues ...interface{}) {}

func (m *mockLogger) Warn(msg string, keysAndValues ...interface{}) {}

func (m *mockLogger) Error(msg string, keysAndValues ...interface{}) {}

func (m *mockLogger) Fatal(msg string, keysAndValues ...interface{}) {}

func TestLoadClientAuthConfig(t *testing.T) {
	// Setup
	testFilename := "test_config.json"
	content := `{
        "instanceName": "testInstance",
        "username": "testUser",
        "password": "testPass",
        "clientID": "testClientID",
        "clientSecret": "testClientSecret"
    }`
	os.WriteFile(testFilename, []byte(content), 0644)
	defer os.Remove(testFilename) // Cleanup

	// Execution
	config, err := LoadClientAuthConfig(testFilename)

	// Assertion
	if err != nil {
		t.Fatalf("expected no error but got: %v", err)
	}
	if config.InstanceName != "testInstance" ||
		config.Username != "testUser" ||
		config.Password != "testPass" ||
		config.ClientID != "testClientID" ||
		config.ClientSecret != "testClientSecret" {
		t.Fatal("unexpected values in the returned config")
	}
}

func TestNewClient(t *testing.T) {
	// Setup
	instanceName := "testInstance"
	config := Config{
		DebugMode:                 true,
		MaxRetryAttempts:          3,
		CustomBackoff:             nil,
		EnableDynamicRateLimiting: true,
		Logger:                    nil,
		MaxConcurrentRequests:     5,
		TokenLifespan:             45 * time.Minute,
		TokenRefreshBufferPeriod:  10 * time.Minute,
		TotalRetryDuration:        90 * time.Second,
	}
	logger := &mockLogger{}

	// Execution
	client, err := NewClient(instanceName, config, logger)

	// Assertion
	if err != nil {
		t.Fatalf("expected no error but got: %v", err)
	}
	if client.InstanceName != "testInstance" ||
		client.config.MaxRetryAttempts != 3 ||
		client.config.MaxConcurrentRequests != 5 {
		t.Fatal("unexpected values in the returned client")
	}
}
