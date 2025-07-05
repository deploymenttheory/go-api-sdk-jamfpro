package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the mobile group name you want to retrieve
	mobileGroupName := "All Managed iPads" // Replace with the actual mobile group name

	// Call GetMobileGroupByJamfProName function
	mobileGroup, err := client.GetMobileGroupByJamfProName(mobileGroupName)
	if err != nil {
		log.Fatalf("Error fetching mobile group by name '%s': %v", mobileGroupName, err)
	}

	// Pretty print the mobile group in JSON
	groupJSON, err := json.MarshalIndent(mobileGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling mobile group data: %v", err)
	}
	fmt.Printf("Fetched Mobile Group (Name: %s):\n%s\n", mobileGroupName, string(groupJSON))
}
