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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	computerID := "8"                                                                                                                                                              // Example computer ID
	filePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/ebooks_pdf/Apple-Developer-Program-License-Agreement-20230828-English.pdf" // File to be uploaded

	response, err := client.UploadAttachmentAndAssignToComputerByID(computerID, filePath)
	if err != nil {
		log.Fatalf("Error uploading attachment: %v", err)
	}

	fmt.Printf("Attachment uploaded successfully. ID: %s, Href: %s\n", response.ID, response.Href)
}
