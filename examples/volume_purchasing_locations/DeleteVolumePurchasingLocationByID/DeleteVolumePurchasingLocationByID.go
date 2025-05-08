package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// ID of the volume purchasing location you want to delete
	locationID := "13" // Replace with the actual ID you want to delete

	// Call the DeleteVolumePurchasingLocationByID function
	err = client.DeleteVolumePurchasingLocationByID(locationID)
	if err != nil {
		log.Fatalf("Error deleting volume purchasing location: %v", err)
	}

	fmt.Printf("Successfully deleted volume purchasing location with ID: %s\n", locationID)
}
