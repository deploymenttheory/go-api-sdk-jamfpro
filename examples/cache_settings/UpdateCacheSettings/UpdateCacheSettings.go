package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
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
