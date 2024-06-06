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
		PackageName:          "Microsoft_365_and_Office_16.82.24021116_Installer.pkg",
		FileName:             "Microsoft_365_and_Office_16.82.24021116_Installer.pkg",
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

	// Use the CreatePackage function with the package payload
	response, err := client.CreatePackage(pkg)
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
