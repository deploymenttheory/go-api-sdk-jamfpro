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

	// Define the IDs of the computer extension attributes to delete
	attributeIDs := []string{"1", "2"} // Replace with actual IDs you want to delete

	// Call DeleteMultipleComputerExtensionAttributeByID function
	err = client.DeleteMultipleComputerExtensionAttributeByID(attributeIDs)
	if err != nil {
		log.Fatalf("Error deleting multiple Computer Extension Attributes: %v", err)
	}

	fmt.Printf("Successfully deleted Computer Extension Attributes with IDs: %v\n", attributeIDs)
}
