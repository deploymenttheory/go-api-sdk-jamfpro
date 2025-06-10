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

	// Define the computer ID and file path
	computerID := "21"
	filePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/icon/UploadIcon/cat.png"

	// Create a slice with the single file path as required by the updated function
	filePaths := []string{filePath}

	// Upload the attachment
	response, err := client.UploadAttachmentAndAssignToComputerByID(computerID, filePaths)
	if err != nil {
		log.Fatalf("Error uploading attachment: %v", err)
	}

	// Print the response details
	fmt.Printf("Attachment uploaded successfully:\n")
	fmt.Printf("  ID: %s\n", response.ID)
	fmt.Printf("  Href: %s\n", response.Href)
}
