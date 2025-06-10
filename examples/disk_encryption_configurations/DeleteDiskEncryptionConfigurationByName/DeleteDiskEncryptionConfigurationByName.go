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

	err = client.DeleteDiskEncryptionConfigurationByName("Corporate Encryption")
	if err != nil {
		log.Fatalf("Error deleting Disk Encryption Configuration by Name: %v", err)
	}

	log.Println("Deleted Disk Encryption Configuration by Name successfully.")
}
