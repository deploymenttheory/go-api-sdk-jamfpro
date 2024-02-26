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

	scriptIDToDelete := "3855" // Replace with the ID of the script you want to delete

	err = client.DeleteScriptByID(scriptIDToDelete)
	if err != nil {
		log.Fatalf("Error deleting script by ID: %v", err)
	}

	fmt.Println("Script deleted successfully!")
}
