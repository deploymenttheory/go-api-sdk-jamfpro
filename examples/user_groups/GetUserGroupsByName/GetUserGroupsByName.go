package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // Adjust log level as needed

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Replace with the actual name of the user group you want to fetch
	userGroupName := "Teachers"

	// Call GetUserGroupsByName to fetch details of a specific user group
	userGroupDetail, err := client.GetUserGroupByName(userGroupName)
	if err != nil {
		fmt.Println("Error fetching user group details:", err)
		return
	}

	// Pretty print the user group details in XML
	userGroupDetailXML, err := xml.MarshalIndent(userGroupDetail, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling user group data: %v", err)
	}
	fmt.Println("User Group Details:\n", string(userGroupDetailXML))
}
