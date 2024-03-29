package main

import (
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

	packageName := "firefox-91.0.2.dmg"

	err = client.DeletePackageByName(packageName)
	if err != nil {
		log.Fatalf("Error deleting network package by name: %v", err)
	} else {
		fmt.Printf("Network package '%s' successfully deleted.\n", packageName)
	}
}
