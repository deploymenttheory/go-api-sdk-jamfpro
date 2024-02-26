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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Specify the path to the file you want to upload
	filePath := "/Users/dafyddwatkins/localtesting/support_files/packages/microsoft-edge-121-0-2277-106.pkg"

	// Call CreateJCDS2Package with the file path
	response, err := client.CreateJCDS2PackageV2(filePath)
	if err != nil {
		log.Fatalf("Failed to create JCDS 2.0 package: %v", err)
	}

	// Print the response
	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	fmt.Println("Response:", string(responseBytes))
}
