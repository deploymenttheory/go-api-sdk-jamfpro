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

	// Specify the updated details for the building
	buildingUpdate := &jamfpro.ResourceBuilding{
		Name:           "Updated Building Name",
		StreetAddress1: "Updated Address 1",
		StreetAddress2: "Updated Address 2",
		City:           "Updated City",
		StateProvince:  "Updated State",
		ZipPostalCode:  "Updated Zip Code",
		Country:        "Updated Country",
	}

	// Update building by name
	buildingName := "Apple Park" // Replace with the actual name
	updatedBuilding, err := client.UpdateBuildingByName(buildingName, buildingUpdate)
	if err != nil {
		log.Fatalf("Error updating building: %v", err)
	}

	// Output the result
	fmt.Printf("Updated Building: %+v\n", updatedBuilding)
}
