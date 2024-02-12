package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	//logLevel := logger.LogLevelInfo // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
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
