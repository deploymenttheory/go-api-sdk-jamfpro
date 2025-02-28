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

	// Fetch all macOS configuration profiles
	profiles, err := client.GetMacOSConfigurationProfiles()
	if err != nil {
		log.Fatalf("Error fetching macOS configuration profiles: %v", err)
	}

	fmt.Printf("Found %d macOS configuration profiles. Starting deletion process:\n", len(profiles.Results))

	successCount := 0
	failureCount := 0

	// Iterate over each macOS configuration profile and delete
	for _, profile := range profiles.Results {
		profileID := fmt.Sprintf("%d", profile.ID)
		fmt.Printf("Attempting to delete macOS configuration profile - ID: %s, Name: %s\n", profileID, profile.Name)

		err = client.DeleteMacOSConfigurationProfileByID(profileID)
		if err != nil {
			log.Printf("Error deleting macOS configuration profile ID %s (%s): %v\n", profileID, profile.Name, err)
			failureCount++
			continue // Move to the next profile if there's an error
		}

		fmt.Printf("Successfully deleted macOS configuration profile ID %s (%s)\n", profileID, profile.Name)
		successCount++
	}

	// Print summary
	fmt.Printf("\nmacOS configuration profile deletion process completed.\n")
	fmt.Printf("Successfully deleted: %d profiles\n", successCount)
	fmt.Printf("Failed to delete: %d profiles\n", failureCount)
	fmt.Printf("Total processed: %d profiles\n", len(profiles.Results))
}
