package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Fetch all departments
	departments, err := client.GetDepartments("")
	if err != nil {
		log.Fatalf("Error fetching departments: %v", err)
	}

	fmt.Println("Departments fetched. Starting deletion process:")

	// Iterate over each department and delete
	for _, department := range departments.Results {
		fmt.Printf("Deleting department ID: %s, Name: %s\n", department.ID, department.Name)

		err = client.DeleteDepartmentByID(department.ID)
		if err != nil {
			log.Printf("Error deleting department ID %s: %v\n", department.ID, err)
			continue // Move to the next department if there's an error
		}

		fmt.Printf("Department ID %s deleted successfully.\n", department.ID)
	}

	fmt.Println("Department deletion process completed.")
}
