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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Building details to be updated
	buildingUpdate := &jamfpro.ResourceBuilding{
		Name:           "Updated Building Name",
		StreetAddress1: "Updated Address 1",
		StreetAddress2: "Updated Address 2",
		City:           "Updated City",
		StateProvince:  "Updated State",
		ZipPostalCode:  "Updated Zip Code",
		Country:        "Updated Country",
	}

	// Update the building with a specific ID
	buildingID := "1326" // Replace with the actual ID of the building you want to update
	updatedBuilding, err := client.UpdateBuildingByID(buildingID, buildingUpdate)
	if err != nil {
		log.Fatalf("Error updating building: %v", err)
	}

	// Print the details of the updated building
	fmt.Printf("Updated Building: %+v\n", updatedBuilding)
}
