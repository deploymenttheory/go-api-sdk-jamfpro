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

	// Define the department name you want to create
	departmentName := &jamfpro.ResourceDepartment{
		Name: "jamf-pro-go-sdk Department",
	}

	// Call CreateDepartment function using the department name
	createdDepartment, err := client.CreateDepartment(departmentName)
	if err != nil {
		log.Fatalf("Error creating department: %v", err)
	}

	// Pretty print the created department in JSON
	createdDepartmentJSON, err := json.MarshalIndent(createdDepartment, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created department data: %v", err)
	}
	fmt.Println("Created Department:\n", string(createdDepartmentJSON))
}
