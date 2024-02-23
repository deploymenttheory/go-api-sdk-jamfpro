package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

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

	// ID of the advanced user search to retrieve
	advancedUserSearchID := 29 // Replace with the actual ID

	// Call GetAdvancedUserSearchByID function
	advancedUserSearch, err := client.GetAdvancedUserSearchByID(advancedUserSearchID)
	if err != nil {
		log.Fatalf("Error fetching advanced user search by ID: %v", err)
	}

	// Pretty print the advanced user search in XML
	advancedUserSearchXML, err := xml.MarshalIndent(advancedUserSearch, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling advanced user search data: %v", err)
	}
	fmt.Println("Fetched Advanced User Search:\n", string(advancedUserSearchXML))
}
