package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Fetch all scripts
	scripts, err := client.GetScripts("")
	if err != nil {
		log.Fatalf("Error fetching scripts: %v", err)
	}

	fmt.Printf("Found %d scripts. Starting deletion process:\n", scripts.Size)

	successCount := 0
	failureCount := 0

	// Iterate over each script and delete
	for _, script := range scripts.Results {
		fmt.Printf("Attempting to delete script - ID: %s, Name: %s\n", script.ID, script.Name)

		err = client.DeleteScriptByID(script.ID)
		if err != nil {
			log.Printf("Error deleting script ID %s (%s): %v\n", script.ID, script.Name, err)
			failureCount++
			continue // Move to the next script if there's an error
		}

		fmt.Printf("Successfully deleted script ID %s (%s)\n", script.ID, script.Name)
		successCount++
	}

	fmt.Printf("\nScript deletion process completed.\n")
	fmt.Printf("Successfully deleted: %d scripts\n", successCount)
	fmt.Printf("Failed to delete: %d scripts\n", failureCount)
	fmt.Printf("Total processed: %d scripts\n", scripts.Size)
}
