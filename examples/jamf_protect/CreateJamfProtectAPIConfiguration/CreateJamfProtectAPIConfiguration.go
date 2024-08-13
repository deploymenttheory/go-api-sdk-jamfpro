package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create a Jamf Protect API configuration request
	config := jamfpro.ResourceJamfProtectRegisterRequest{
		ProtectURL: "https://examplejamfprotect.jamfcloud.com/graphql",
		ClientID:   "uzPJXlArmzTAmPRQtZEnQ2OFtNw8qQV",
		Password:   "7fyP6BphUUQ5B_zoLrkYhM5j1HTcf-4PxshettZbK0ZcnzV57gyHwF23U3F96F",
	}

	// Call CreateJamfProtectAPIConfiguration function
	response, err := client.CreateJamfProtectAPIConfiguration(config)
	if err != nil {
		log.Fatalf("Error creating Jamf Protect API configuration: %v", err)
	}

	// Convert the response struct to pretty-printed JSON
	responseJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling Jamf Protect API configuration response to JSON: %v", err)
	}

	// Print the pretty-printed JSON
	fmt.Println("Jamf Protect API Configuration Response:")
	fmt.Println(string(responseJSON))
}
