// this file is used to demonstrate how to set up the client with a manual configuration
package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the configuration based on the new ConfigContainer struct
	config := &jamfpro.ConfigContainer{
		LogLevel:          "warning",
		LogExportPath:     "", // Set this if you want to export logs to a file
		HideSensitiveData: true,

		InstanceDomain:       "your-instance-domain",
		AuthMethod:           "oauth2", // Use "basic" for basic authentication
		ClientID:             "your-client-id",
		ClientSecret:         "your-client-secret",
		Username:             "", // Set this if using basic auth
		Password:             "", // Set this if using basic auth
		JamfLoadBalancerLock: true,

		CustomCookies:               []jamfpro.CustomCookie{},
		MaxRetryAttempts:            3,
		MaxConcurrentRequests:       1,
		EnableDynamicRateLimiting:   false,
		CustomTimeout:               60,  // in seconds
		TokenRefreshBufferPeriod:    300, // in seconds
		TotalRetryDuration:          60,  // in seconds
		FollowRedirects:             true,
		MaxRedirects:                5,
		EnableConcurrencyManagement: true,
		MandatoryRequestDelay:       0, // in milliseconds
		RetryEligiableRequests:      true,
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
