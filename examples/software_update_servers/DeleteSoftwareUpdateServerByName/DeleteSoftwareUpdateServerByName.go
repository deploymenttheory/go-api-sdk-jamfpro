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

	// Name of the software update server to delete
	serverName := "New York SUS"

	// Call DeleteSoftwareUpdateServerByName
	err = client.DeleteSoftwareUpdateServerByName(serverName)
	if err != nil {
		log.Fatalf("Error deleting software update server by name: %v", err)
	}

	fmt.Printf("Successfully deleted software update server with name %s\n", serverName)
}
