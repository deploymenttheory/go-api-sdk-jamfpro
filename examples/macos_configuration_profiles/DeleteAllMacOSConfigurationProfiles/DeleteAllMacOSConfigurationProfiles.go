package main

import (
	"fmt"
	"log"
	"strconv"

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

	// Fetch all macOS configuration profiles
	extAtts, err := client.GetMacOSConfigurationProfiles()
	if err != nil {
		log.Fatalf("Error fetching macOS configuration profiles: %v", err)
	}

	fmt.Println("macOS configuration profiles fetched. Starting deletion process:")

	// Iterate over each macOS configuration profile and delete
	for _, extAtt := range extAtts.Results {
		fmt.Printf("Deleting macOS configuration profile ID: %d, Name: %s\n", extAtt.ID, extAtt.Name)

		err = client.DeleteMacOSConfigurationProfileByID(strconv.Itoa(extAtt.ID))
		if err != nil {
			log.Printf("Error deleting macOS configuration profile ID %d: %v\n", extAtt.ID, err)
			continue // Move to the next macOS configuration profile if there's an error
		}

		fmt.Printf("macOS configuration profile ID %d deleted successfully.\n", extAtt.ID)
	}

	fmt.Println("macOS configuration profile deletion process completed.")

}
