package main

import (
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Manually define the client configuration
	config := httpclient.ClientConfig{
		Environment: httpclient.EnvironmentConfig{
			InstanceName:       "your-instance-name",
			OverrideBaseDomain: "",        // Only required if you are not on jamfcloud.com
			APIType:            "jamfpro", // Required to specify the API type
		},
		Auth: httpclient.AuthConfig{
			ClientID:     "your-client-id",
			ClientSecret: "your-client-secret",
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:                  "LogLevelDebug",
			LogOutputFormat:           "console",
			HideSensitiveData:         true,
			MaxRetryAttempts:          5,
			EnableDynamicRateLimiting: true,
			MaxConcurrentRequests:     10,
			TokenRefreshBufferPeriod:  5 * time.Minute,
			TotalRetryDuration:        60 * time.Second,
			CustomTimeout:             30 * time.Second,
		},
	}

	// Initialize the Jamf Pro client with the given configuration
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Call the GetComputers method
	computers, err := client.GetComputers()
	if err != nil {
		log.Fatalf("Error fetching computers: %v", err)
	}

	// Print out the fetched computers
	fmt.Println("Fetched Computers:")
	for _, computer := range computers.Results {
		fmt.Printf("ID: %d, Name: %s\n", computer.ID, computer.Name)
	}
}
