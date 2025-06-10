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

	// Resource history details to be updated
	historyUpdate := &jamfpro.ResourceBuildingResourceHistory{
		Note: "Sso settings update",
	}

	// Add specified Building history object notes with a specific ID
	buildingID := "4" // Replace with the actual ID of the building you want to update
	updatedHistory, err := client.CreateBuildingResourceHistoryByID(buildingID, historyUpdate)
	if err != nil {
		log.Fatalf("Error updating building resource history: %v", err)
	}

	// Print the details of the updated resource history
	fmt.Printf("Updated Building Resource History: %+v\n", updatedHistory)
}
