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
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Get all enrollment customizations
	customizations, err := client.GetEnrollmentCustomizations(url.Values{})
	if err != nil {
		log.Fatalf("Failed to get enrollment customizations: %v", err)
	}

	// Pretty print the enrollment customizations details in JSON
	JSON, err := json.MarshalIndent(customizations, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling enrollment customizations data: %v", err)
	}
	fmt.Println("Enrollment Customizations List:\n", string(JSON))

	// Print total count
	fmt.Printf("\nTotal number of enrollment customizations: %d\n", customizations.TotalCount)

	// Optionally, you can iterate through the results
	fmt.Println("\nCustomization Names:")
	for _, customization := range customizations.Results {
		fmt.Printf("- %s (ID: %s)\n", customization.DisplayName, customization.ID)
	}
}
