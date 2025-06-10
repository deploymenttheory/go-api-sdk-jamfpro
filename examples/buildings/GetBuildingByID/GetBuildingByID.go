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

	// Define the variable for the building ID
	buildingID := "1399"

	// Call GetBuildings function
	building, err := client.GetBuildingByID(buildingID)
	if err != nil {
		log.Fatalf("Error fetching accounts: %v", err)
	}

	// Pretty print the accounts details
	buildingsJSON, err := json.MarshalIndent(building, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling accounts data: %v", err)
	}
	fmt.Println("Fetched Accounts List:", string(buildingsJSON))
}
