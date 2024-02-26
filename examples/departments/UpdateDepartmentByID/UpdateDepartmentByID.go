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

	// ID of the department you want to update
	departmentID := "23514" // Placeholder ID, replace with the correct ID you want to update

	// New name for the department you want to update
	updatedDepartment := &jamfpro.ResourceDepartment{
		Name: "jamf pro go sdk Department",
	}

	// Call UpdateDepartmentByID function
	departmentItem, err := client.UpdateDepartmentByID(departmentID, updatedDepartment)
	if err != nil {
		log.Fatalf("Error updating department: %v", err)
	}

	// Fetch the updated department's details
	fetchedDepartment, err := client.GetDepartmentByID(departmentItem.ID)
	if err != nil {
		log.Fatalf("Error fetching updated department: %v", err)
	}

	// Pretty print the fetched department in JSON
	fetchedDepartmentJSON, err := json.MarshalIndent(fetchedDepartment, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling fetched department data: %v", err)
	}
	fmt.Println("Fetched Updated Department:\n", string(fetchedDepartmentJSON))
}
