package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

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

	// Define the ID of the self-service branding to fetch
	brandingID := "1" // Replace with the actual ID

	// Call the GetSelfServiceBrandingMacOSByID function and handle any errors
	branding, err := client.GetSelfServiceBrandingMacOSByID(brandingID)
	if err != nil {
		// If there's an error, log it to stderr and exit with a non-zero status code
		fmt.Fprintf(os.Stderr, "Error fetching self-service branding for macOS with ID %s: %v\n", brandingID, err)
		os.Exit(1)
	}

	createdScriptJSON, err := json.MarshalIndent(branding, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created script data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(createdScriptJSON))
}
