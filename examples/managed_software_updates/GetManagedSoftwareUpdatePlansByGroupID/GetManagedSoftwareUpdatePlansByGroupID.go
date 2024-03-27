package main

import (
	"encoding/json"
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

	// Define groupId and groupType for the request
	groupId := "55"               // Example group ID
	groupType := "COMPUTER_GROUP" // Possible values: COMPUTER_GROUP, MOBILE_DEVICE_GROUP

	// Call GetManagedSoftwareUpdatePlansByGroupId function
	updatePlans, err := client.GetManagedSoftwareUpdatePlansByGroupID(groupId, groupType)
	if err != nil {
		log.Fatalf("Error fetching managed software update plans by group ID: %v", err)
	}

	// Pretty print the managed software update plans for the group in json
	updatePlansJSON, err := json.MarshalIndent(updatePlans, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling managed software update plans data: %v", err)
	}
	fmt.Println("Fetched managed software update plans for group:\n", string(updatePlansJSON))
}
