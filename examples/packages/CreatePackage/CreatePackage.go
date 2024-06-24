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

	// Define the package manifest payload. the settings below is the minimum required
	// to create a package with the api
	pkgMetadata := jamfpro.ResourcePackage{
		PackageName:          "powershell-7.4.1-osx-x64.pkg",
		FileName:             "powershell-7.4.1-osx-x64.pkg",
		CategoryID:           "-1",
		Priority:             3,
		FillUserTemplate:     BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}

	// Use the CreatePackage function with the package payload
	response, err := client.CreatePackage(pkgMetadata)
	if err != nil {
		fmt.Println("Error creating package manifest:", err)
		return
	}

	// Pretty print the created package details in XML
	packageJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created package data: %v", err)
	}
	fmt.Println("Created Package Details:\n", string(packageJSON))
}

// BoolPtr is a helper function to create a pointer to a bool.
func BoolPtr(b bool) *bool {
	return &b
}
