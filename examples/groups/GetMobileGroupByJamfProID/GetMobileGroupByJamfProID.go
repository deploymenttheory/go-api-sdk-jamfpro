package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Path to the Jamf Pro client configuration file
	configPath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client
	jamfClient, err := jamfpro.BuildClientWithConfigFile(configPath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// The Jamf Pro Mobile Group ID to retrieve
	mobileGroupID := "1" // Replace with the actual Mobile Group Jamf Pro ID

	// Fetch the mobile group by Jamf Pro ID
	mobileGroup, err := jamfClient.GetMobileGroupByJamfProID(mobileGroupID)
	if err != nil {
		log.Fatalf("Error fetching mobile group by Jamf Pro ID '%s': %v", mobileGroupID, err)
	}

	// Pretty print the mobile group in JSON
	groupJSON, err := json.MarshalIndent(mobileGroup, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling mobile group data: %v", err)
	}
	fmt.Printf("Fetched Mobile Group (Jamf Pro ID: %s):\n%s\n", mobileGroupID, string(groupJSON))
}
