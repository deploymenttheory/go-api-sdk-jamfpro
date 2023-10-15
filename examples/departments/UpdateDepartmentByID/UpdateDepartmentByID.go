package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// ID of the department you want to update
	departmentID := 6 // Placeholder ID, replace with the correct ID you want to update
	// New name for the department you want to update
	newDepartmentName := "UpdatedDepartmentNameBySDK10" // Replace with the desired updated department name

	// Call UpdateDepartmentByID function
	updatedDepartment, err := client.UpdateDepartmentByID(departmentID, newDepartmentName)
	if err != nil {
		log.Fatalf("Error updating department: %v", err)
	}

	// Fetch the updated department's details
	fetchedDepartment, err := client.GetDepartmentByID(updatedDepartment.Id)
	if err != nil {
		log.Fatalf("Error fetching updated department: %v", err)
	}

	// Pretty print the fetched department in XML
	fetchedDepartmentXML, err := xml.MarshalIndent(fetchedDepartment, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling fetched department data: %v", err)
	}
	fmt.Println("Fetched Updated Department:\n", string(fetchedDepartmentXML))
}
