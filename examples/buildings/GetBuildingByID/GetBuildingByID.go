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

	// Define the variable for the building ID
	buildingID := 1327 // Change this value as needed

	// Call GetBuildingByID function
	building, err := client.GetBuildingByID(fmt.Sprint(buildingID))
	if err != nil {
		log.Fatalf("Error fetching building by ID: %v", err)
	}

	// Print the building details
	fmt.Printf("Fetched Building Details:\n%+v\n", building)
}
