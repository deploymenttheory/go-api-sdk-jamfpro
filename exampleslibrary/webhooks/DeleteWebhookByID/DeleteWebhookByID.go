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

	// Example usage of DeleteWebhookByID
	err = client.DeleteWebhookByID(1) // Replace with the desired webhook ID to delete
	if err != nil {
		fmt.Printf("Error deleting webhook by ID: %v\n", err)
		return
	}
	fmt.Println("Webhook by ID deleted successfully")

}
