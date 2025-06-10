package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the client management ID for which to force DDM sync
	clientManagementID := "10" // Replace with the actual client management ID

	// Call ForceDDMSync function
	err = client.ForceDDMSync(clientManagementID)
	if err != nil {
		log.Fatalf("Error forcing DDM sync: %v", err)
	}

	fmt.Printf("Successfully forced DDM sync for client management ID: %s\n", clientManagementID)
}
