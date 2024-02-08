package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
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
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
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
