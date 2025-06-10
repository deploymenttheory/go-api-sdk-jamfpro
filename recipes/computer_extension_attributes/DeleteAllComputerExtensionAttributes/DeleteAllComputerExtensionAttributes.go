package main

import (
	"fmt"
	"log"
	"net/url"

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

	// Fetch all computer extension attributes
	// For more information on how to add parameters to this request, see docs/url_queries.md
	extAtts, err := client.GetComputerExtensionAttributes(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching computer extension attributes: %v", err)
	}

	fmt.Printf("Found %d computer extension attributes. Starting deletion process:\n", extAtts.TotalCount)

	successCount := 0
	failureCount := 0

	// Iterate over each computer extension attribute and delete
	for _, extAtt := range extAtts.Results {
		fmt.Printf("Attempting to delete computer extension attribute - ID: %s, Name: %s\n", extAtt.ID, extAtt.Name)

		err = client.DeleteComputerExtensionAttributeByID(extAtt.ID)
		if err != nil {
			log.Printf("Error deleting computer extension attribute ID %s (%s): %v\n", extAtt.ID, extAtt.Name, err)
			failureCount++
			continue // Move to the next computer extension attribute if there's an error
		}

		fmt.Printf("Successfully deleted computer extension attribute ID %s (%s)\n", extAtt.ID, extAtt.Name)
		successCount++
	}

	// Print summary
	fmt.Printf("\nComputer extension attribute deletion process completed.\n")
	fmt.Printf("Successfully deleted: %d attributes\n", successCount)
	fmt.Printf("Failed to delete: %d attributes\n", failureCount)
	fmt.Printf("Total processed: %d attributes\n", extAtts.TotalCount)
}
