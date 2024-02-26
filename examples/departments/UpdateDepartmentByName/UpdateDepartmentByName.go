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

	// Define the department's current name and the new name you want to update to
	currentDepartmentName := "JLtestDept"
	newDepartmentName := "jamf pro go sdk Department"

	// New name for the department you want to update
	updatedDepartment := &jamfpro.ResourceDepartment{
		Name: newDepartmentName,
	}

	// Update the department's name using the UpdateDepartmentByName function
	_, err = client.UpdateDepartmentByName(currentDepartmentName, updatedDepartment)
	if err != nil {
		log.Fatalf("Error updating department by name: %v", err)
	}

	// Fetch the updated department's details by its new name
	fetchedDepartment, err := client.GetDepartmentByName(newDepartmentName)
	if err != nil {
		log.Fatalf("Error fetching updated department by name: %v", err)
	}

	// Pretty print the fetched department in JSON
	fetchedDepartmentJSON, err := json.MarshalIndent(fetchedDepartment, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling fetched department data: %v", err)
	}
	fmt.Println("Fetched Updated Department:\n", string(fetchedDepartmentJSON))
}
