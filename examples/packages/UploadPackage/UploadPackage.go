package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example ID and file path to upload. The package manifest must exist in Jamf Pro
	// before uploading the package file using CreatePackage or UpdatePackage functions.
	packageID := "297"
	filePaths := []string{
		"/Users/dafyddwatkins/localtesting/terraform/support_files/packages/powershell-7.4.1-osx-x64.pkg",
		// Add more file paths if needed
	}

	// Upload the package
	response, err := client.UploadPackage(packageID, filePaths)
	if err != nil {
		fmt.Println("Error uploading package:", err)
		return
	}

	// Pretty print the response details in JSON
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling response data: %v", err)
	}
	fmt.Println("Upload Package Response:\n", string(responseJSON))
}
