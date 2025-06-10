package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Let's assume you want to delete a department with the name "HR Department"
	departmentName := "TestNewName"

	// Call DeleteDepartmentByName function
	err = client.DeleteDepartmentByName(departmentName)
	if err != nil {
		log.Fatalf("Error deleting department by name: %v", err)
	}

	// Print success message
	fmt.Println("Deleted Department Successfully!")
}
