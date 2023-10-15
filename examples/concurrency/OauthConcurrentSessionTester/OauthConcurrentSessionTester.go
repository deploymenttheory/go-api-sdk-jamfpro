package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	concurrentRequests           = 10 // Number of simultaneous requests.
	maxConcurrentRequestsAllowed = 5  // Maximum allowed concurrent requests.
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:             authConfig.InstanceName,
		DebugMode:                true,
		Logger:                   jamfpro.NewDefaultLogger(),
		MaxConcurrentRequests:    maxConcurrentRequestsAllowed,
		TokenLifespan:            defaultTokenLifespan,
		TokenRefreshBufferPeriod: defaultBufferPeriod,
		ClientID:                 authConfig.ClientID,
		ClientSecret:             authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Print the obtained token
	//fmt.Println("Successfully obtained token:", client.HTTP.Token)

	// Use a wait group to wait for all goroutines to finish.
	var wg sync.WaitGroup

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Fetch SSO failover settings
			failoverSettings, err := client.GetSSOFailoverSettings()
			if err != nil {
				log.Printf("[Goroutine %d] Error fetching SSO failover settings: %v", id, err)
				return
			}

			fmt.Printf("[Goroutine %d] Failover URL: %s\n", id, failoverSettings.FailoverURL)
			fmt.Printf("[Goroutine %d] Generation Time: %d\n", id, failoverSettings.GenerationTime)
		}(i)
	}

	// Wait for all requests to complete.
	wg.Wait()
	fmt.Println("All requests completed.")
}
