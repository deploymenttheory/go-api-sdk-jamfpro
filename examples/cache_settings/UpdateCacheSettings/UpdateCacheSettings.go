package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Example cache settings to update
	newSettings := &jamfpro.ResourceCacheSettings{
		CacheType:                  "ehcache",
		TimeToLiveSeconds:          180,
		TimeToIdleSeconds:          180,
		DirectoryTimeToLiveSeconds: 180,
		EhcacheMaxBytesLocalHeap:   "1GB",
		CacheUniqueId:              "84a82a63eecf4213b05d4f7023c9083f",
		Elasticache:                false,
		MemcachedEndpoints: []jamfpro.CacheSettingsSubsetMemcachedEndpoints{
			{
				HostName: "localhost",
				Port:     11211,
				Enabled:  true,
			},
		},
	}

	updatedSettings, err := client.UpdateCacheSettings(newSettings)
	if err != nil {
		log.Fatalf("Error updating cache settings: %s", err)
	}

	fmt.Printf("Updated Cache Settings: %+v\n", updatedSettings)
}
