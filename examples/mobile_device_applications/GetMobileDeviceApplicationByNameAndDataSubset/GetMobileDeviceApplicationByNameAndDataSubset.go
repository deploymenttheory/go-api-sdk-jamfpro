package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Replace these with actual name and subset
	name := "YourAppName"
	subset := "YourSubset"

	app, err := client.GetMobileDeviceApplicationByNameAndDataSubset(name, subset)
	if err != nil {
		fmt.Println("Error fetching application by Name and Data Subset:", err)
	} else {
		fmt.Println("Fetched Application by Name and Data Subset:", app)
	}
}
