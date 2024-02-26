package main

import (
	"encoding/xml"
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

	// Define the variable for the building name
	buildingName := "Apple Park" // Change this value as needed

	// Call GetBuildingByNameByID function
	building, err := client.GetBuildingByName(buildingName)
	if err != nil {
		log.Fatalf("Error fetching building by Name: %v", err)
	}

	// Pretty print the building details using XML marshaling
	buildingXML, err := xml.MarshalIndent(building, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling building data: %v", err)
	}
	fmt.Println("Fetched Building Details:", string(buildingXML))
}
