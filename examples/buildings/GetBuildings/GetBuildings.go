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

	// Fetch the buildings
	buildings, err := client.GetBuildings("") // Pass nil or a specific sorting criteria
	if err != nil {
		log.Fatalf("Error fetching buildings: %v", err)
	}

	// Iterate through the buildings and print them
	for _, building := range buildings.Results {
		fmt.Printf("ID: %s, Name: %s, Address: %s, %s, %s, %s, %s\n",
			building.ID, building.Name, building.StreetAddress1, building.StreetAddress2,
			building.City, building.StateProvince, building.ZipPostalCode)
	}
}
