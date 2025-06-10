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
	configProfileUUID := "5ec9527a-40da-4633-9ccc-7fc3fecea1dc"
	computerIDs := []string{"1", "2"}

	err = client.RetryJamfConnectDeploymentTasksByConfigProfileUUID(configProfileUUID, computerIDs)
	if err != nil {
		log.Fatalf("Error retrying tasks: %v", err)
	}

	log.Println("Successfully requested task retry")
}
