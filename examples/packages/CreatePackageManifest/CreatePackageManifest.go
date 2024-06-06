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

	// Set boolean pointers for the FillUserTemplate field
	falsePointer := false

	// Define the package manifest payload. the settings below is the minimum required
	// to create a package with the api
	pkg := jamfpro.ResourcePackage{
		PackageName:          "microsoft-edge-121-0-2277-106.pkg",
		FileName:             "microsoft-edge-121-0-2277-106.pkg",
		CategoryID:           "-1",
		Priority:             3,
		FillUserTemplate:     &falsePointer,
		SWU:                  &falsePointer,
		RebootRequired:       &falsePointer,
		OSInstall:            &falsePointer,
		SuppressUpdates:      &falsePointer,
		SuppressFromDock:     &falsePointer,
		SuppressEula:         &falsePointer,
		SuppressRegistration: &falsePointer,
	}

	// Use the CreatePackageManifest function with the package payload
	response, err := client.CreatePackageManifest(pkg)
	if err != nil {
		fmt.Println("Error creating package manifest:", err)
		return
	}

	// Pretty print the created package details in XML
	packageJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created package data: %v", err)
	}
	fmt.Println("Created Package Manifest Details:\n", string(packageJSON))
}
