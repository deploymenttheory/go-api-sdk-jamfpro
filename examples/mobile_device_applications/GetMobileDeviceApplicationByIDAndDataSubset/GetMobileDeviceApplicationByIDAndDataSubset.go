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

	// Replace "123" with an actual ID and "subset" with the desired data subset
	id := "1"
	subset := "General" // General / Scope / Selfservice / VPPCodes / VPP / AppConfiguration

	app, err := client.GetMobileDeviceApplicationByIDAndDataSubset(id, subset)
	if err != nil {
		fmt.Println("Error fetching application by ID and Data Subset:", err)
	} else {
		fmt.Println("Fetched Application by ID and Data Subset:", app)
	}
}
