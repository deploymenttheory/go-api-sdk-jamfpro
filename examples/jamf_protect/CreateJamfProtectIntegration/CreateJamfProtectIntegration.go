package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create a Jamf Protect registration request using the new struct
	registration := jamfpro.ResourceJamfProtectRegistration{
		ProtectURL: "https://instance.protect.jamfcloud.com/graphql",
		ClientID:   "supersecretclientid",
		Password:   "supersecretpassword",
	}

	// Create the Jamf Protect integration with auto-install enabled
	response, err := client.CreateJamfProtectIntegration(registration, true)
	if err != nil {
		log.Fatalf("Error creating Jamf Protect integration: %v", err)
	}

	// Convert the response struct to pretty-printed JSON
	responseJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling Jamf Protect integration response to JSON: %v", err)
	}

	// Print the pretty-printed JSON
	fmt.Println("Jamf Protect Integration Response:")
	fmt.Println(string(responseJSON))
}
