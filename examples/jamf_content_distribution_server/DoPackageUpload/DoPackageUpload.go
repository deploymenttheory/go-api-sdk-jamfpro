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

	// Specify the path to the file you want to upload
	filePath := "/Users/dafyddwatkins/localtesting/terraform/support_files/packages/microsoft-edge-121-0-2277-106.pkg"

	// Define package metadata
	packageData := jamfpro.ResourcePackage{
		PackageName:          "Microsoft Edge",
		FileName:             "microsoft-edge-121-0-2277-106.pkg", // Ensure this matches the actual file name
		CategoryID:           "1",                                 // Set appropriate Category ID
		Info:                 "Microsoft Edge Browser Package",
		Notes:                "Version 121.0.2277.106. This package installs Microsoft Edge on macOS devices.",
		Priority:             10,                                      // Set priority (lower number means higher priority)
		RebootRequired:       BoolPtr(false),                          // Set to true if a reboot is required after installation
		FillUserTemplate:     BoolPtr(false),                          // Set to true if the package should fill the user template
		FillExistingUsers:    BoolPtr(false),                          // Set to true if the package should fill existing user directories
		OSInstall:            BoolPtr(true),                           // Set to true if this is an OS install package
		SuppressUpdates:      BoolPtr(false),                          // Set to true to suppress updates
		SuppressFromDock:     BoolPtr(false),                          // Set to true to suppress from dock
		SuppressEula:         BoolPtr(false),                          // Set to true to suppress EULA
		SuppressRegistration: BoolPtr(false),                          // Set to true to suppress registration
		OSRequirements:       "macOS 10.15.x, macOS 11.x, macOS 12.x", // Specify OS requirements
	}

	// Call DoPackageUpload with the file path and package metadata
	uploadResponse, err := client.DoPackageUpload(filePath, &packageData)
	if err != nil {
		log.Fatalf("Failed to upload package: %v", err)
	}

	// Marshal the response to JSON for output
	responseBytes, err := json.Marshal(uploadResponse)
	if err != nil {
		log.Fatalf("Failed to marshal upload response: %v", err)
	}

	// Print the response
	fmt.Println("Response:", string(responseBytes))
}

// Helper function to create a pointer to a bool
func BoolPtr(b bool) *bool {
	return &b
}
