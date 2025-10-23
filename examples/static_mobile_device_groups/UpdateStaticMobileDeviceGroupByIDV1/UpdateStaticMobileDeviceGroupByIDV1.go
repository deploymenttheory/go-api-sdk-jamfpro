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

	// Define group ID to update
	groupID := "106"
	siteID := "-1"

	// Create update data
	updateGroup := jamfpro.ResourceStaticMobileDeviceGroupV1{
		GroupName:        "Updated Static Test Group",
		SiteId:           siteID,
		GroupDescription: "Description goes here",
	}
	// Call function
	updated, err := client.UpdateStaticMobileDeviceGroupByIDV1(groupID, updateGroup)
	if err != nil {
		log.Fatalf("Error updating static mobile device group: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(updated, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated group data: %v", err)
	}
	fmt.Println("Updated Static Mobile Device Group:\n", string(response))
}
