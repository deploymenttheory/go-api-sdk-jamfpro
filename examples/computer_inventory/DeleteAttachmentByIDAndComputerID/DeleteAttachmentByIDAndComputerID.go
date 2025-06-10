package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	computerID := "8"   // Example computer ID
	attachmentID := "2" // ID of the attachment to be deleted

	err = client.DeleteAttachmentByIDAndComputerID(computerID, attachmentID)
	if err != nil {
		log.Fatalf("Error deleting attachment: %v", err)
	}

	fmt.Println("Attachment deleted successfully.")
}
