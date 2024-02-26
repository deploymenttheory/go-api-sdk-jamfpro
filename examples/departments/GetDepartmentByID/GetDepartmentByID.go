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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	departmentID := "25542"

	department, err := client.GetDepartmentByID(departmentID)
	if err != nil {
		log.Fatalf("Error fetching department by ID: %v", err)
	}

	departmentJSON, err := json.MarshalIndent(department, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling department data: %v", err)
	}
	fmt.Println("Fetched Department:\n", string(departmentJSON))
}
