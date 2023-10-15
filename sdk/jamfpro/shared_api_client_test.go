package jamfpro

import (
	"testing"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/stretchr/testify/assert"
)

func TestOAuthCredentialsSetting(t *testing.T) {
	// Mock config for testing
	config := Config{
		InstanceName:             "testInstance",
		DebugMode:                false,
		Logger:                   http_client.NewDefaultLogger(),
		MaxConcurrentRequests:    5,
		TokenLifespan:            30 * time.Minute,
		TokenRefreshBufferPeriod: 5 * time.Minute,
		ClientID:                 "mockClientID",
		ClientSecret:             "mockClientSecret",
	}

	logger.Info("Initializing client with mock configuration...")
	// Create a new jamfpro client instanceclient,
	client, err := NewClient(config)
	if err != nil {
		logger.Fatal("Failed to create Jamf Pro client: %v", err)
	}

	// Mock new OAuth credentials
	newCreds := http_client.OAuthCredentials{
		ClientID:     "newMockClientID",
		ClientSecret: "newMockClientSecret",
	}
	logger.Info("Setting client's OAuth credentials to: %v", newCreds)
	client.SetClientOAuthCredentials(newCreds)

	// In order to validate the setting of OAuth credentials, the http_client package should expose a method or field to get the current OAuth credentials.
	// Given that it's private in the current code, you might need to add a getter method in the http_client package.
	// For now, I'm assuming such a method has been added called GetOAuthCredentials().

	logger.Info("Asserting that the client's OAuth credentials have been set correctly...")
	assert.Equal(t, newCreds, client.HTTP.GetOAuthCredentials(), "Expected client to have updated OAuth credentials")
}
