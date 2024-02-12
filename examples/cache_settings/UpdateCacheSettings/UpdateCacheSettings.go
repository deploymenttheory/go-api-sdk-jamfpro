package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
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
