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

	// Fetch all restricted software
	restrictedSoftware, err := client.GetRestrictedSoftwares()
	if err != nil {
		log.Fatalf("Error fetching restricted software: %v", err)
	}

	fmt.Printf("Found %d restricted software entries. Starting deletion process:\n", restrictedSoftware.Size)

	successCount := 0
	failureCount := 0

	// Iterate over each restricted software entry and delete
	for _, software := range restrictedSoftware.RestrictedSoftware {
		fmt.Printf("Attempting to delete restricted software - ID: %d, Name: %s\n", software.ID, software.Name)

		err = client.DeleteRestrictedSoftwareByID(strconv.Itoa(software.ID))
		if err != nil {
			log.Printf("Error deleting restricted software ID %d (%s): %v\n", software.ID, software.Name, err)
			failureCount++
			continue // Move to the next entry if there's an error
		}

		fmt.Printf("Successfully deleted restricted software ID %d (%s)\n", software.ID, software.Name)
		successCount++
	}

	// Print summary
	fmt.Printf("\nRestricted software deletion process completed.\n")
	fmt.Printf("Successfully deleted: %d entries\n", successCount)
	fmt.Printf("Failed to delete: %d entries\n", failureCount)
	fmt.Printf("Total processed: %d entries\n", restrictedSoftware.Size)
}
