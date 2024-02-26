package main

import (
	"encoding/xml"
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
	// Define the package update
	updatedPackage := &jamfpro.ResourcePackage{
		Name:     "Updated Package Name",
		Category: "Productivity",
		Filename: "updated_package.zip",
		Info:     "This is an updated package for office productivity tools.",
		Notes:    "Please ensure compatibility before deploying.",
		Priority: 10,
	}

	packageID := 1 // Replace with actual ID

	// Update the package by ID
	updated, err := client.UpdatePackageByID(packageID, updatedPackage)
	if err != nil {
		log.Fatalf("Error updating package by ID: %v", err)
	}

	// Print the updated package details
	packageXML, _ := xml.MarshalIndent(updated, "", "    ")
	fmt.Println("Updated Package:", string(packageXML))
}
