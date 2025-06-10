package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

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

	// Call GetEnrollmentHistory function
	// You can add sorting parameters like "sort=date:desc" to sort by date in descending order
	// For more information on how to add parameters to this request, see docs/url_queries.md
	history, err := client.GetEnrollmentHistory(url.Values{})
	if err != nil {
		log.Fatalf("Error getting enrollment history: %v", err)
	}

	// Print the total count of history records
	fmt.Printf("Total enrollment history records: %d\n\n", history.TotalCount)

	// Pretty print the enrollment history records in JSON
	JSON, err := json.MarshalIndent(history, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling enrollment history data: %v", err)
	}
	fmt.Println("Enrollment History:\n", string(JSON))
}
