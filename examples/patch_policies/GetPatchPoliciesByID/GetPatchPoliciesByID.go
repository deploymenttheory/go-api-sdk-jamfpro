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
	// The ID of the patch policy you want to retrieve
	patchPolicyID := 1 // Replace with the actual ID you want to retrieve

	// Call the GetPatchPoliciesByID function
	patchPolicy, err := client.GetPatchPoliciesByID(patchPolicyID)
	if err != nil {
		log.Fatalf("Error fetching patch policy by ID: %v", err)
	}

	// Convert the response into pretty XML for printing
	output, err := xml.MarshalIndent(patchPolicy, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling patch policy to XML: %v", err)
	}

	// Print the pretty XML
	fmt.Printf("Patch Policy (ID: %d):\n%s\n", patchPolicyID, string(output))
}
