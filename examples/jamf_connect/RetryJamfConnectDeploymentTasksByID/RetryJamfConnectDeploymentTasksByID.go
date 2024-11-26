package main

import (
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

	// Example config profile UUID and computer IDs to retry
	configProfileUUID := "24a7bb2a-9871-4895-9009-d1be07ed31b1"
	computerIDs := []string{"1", "2"}

	err = client.RetryJamfConnectDeploymentTasksByID(configProfileUUID, computerIDs)
	if err != nil {
		log.Fatalf("Error retrying tasks: %v", err)
	}

	log.Println("Successfully requested task retry")
}
