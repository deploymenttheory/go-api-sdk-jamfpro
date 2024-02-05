package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

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

	// Create a sample user to be created
	updatedUser := &jamfpro.ResourceUser{
		ID:           1,
		Name:         "AHarrison",
		FullName:     "Ashley Harrison",
		Email:        "aharrison@company.com",
		EmailAddress: "aharrison@company.com",
		PhoneNumber:  "123-555-6789",
		Position:     "Teacher",
		Sites: []jamfpro.SharedResourceSite{
			{
				ID:   -1,
				Name: "None",
			},
		},
	}

	// Replace with the actual ID of the user you want to update
	userID := 1

	// Call UpdateUserByID to update the user
	updatedUser, err = client.UpdateUserByID(userID, updatedUser)
	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}

	// Print the details of the updated user
	fmt.Printf("Updated User Details:\nID: %d\nName: %s\nEmail: %s\n", updatedUser.ID, updatedUser.Name, updatedUser.Email)
}
