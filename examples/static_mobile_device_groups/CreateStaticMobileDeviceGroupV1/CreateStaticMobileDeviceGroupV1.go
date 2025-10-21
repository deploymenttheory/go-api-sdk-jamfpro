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

	// Create new static mobile device group
	siteId := "-1"
	newGroup := jamfpro.ResourceStaticMobileDeviceGroupV1{
		GroupName:        "Static Test Group",
		SiteId:           siteId,
		GroupDescription: "Description goes here",
	}

	// Call function
	created, err := client.CreateStaticMobileDeviceGroupV1(newGroup)
	if err != nil {
		log.Fatalf("Error creating static mobile device group: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(created, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created group data: %v", err)
	}
	fmt.Println("Created Static Mobile Device Group:\n", string(response))
}
