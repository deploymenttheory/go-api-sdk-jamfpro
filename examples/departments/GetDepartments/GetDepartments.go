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

	// Call GetDepartments function
	departments, err := client.GetDepartments("")
	if err != nil {
		log.Fatalf("Error fetching departments: %v", err)
	}

	// Pretty print the departments in JSON
	departmentsJSON, err := json.MarshalIndent(departments, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling departments data: %v", err)
	}
	fmt.Println("Fetched Departments:\n", string(departmentsJSON))
}
