package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
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

	// Sort filter
	// For more information on how to add parameters to this request, see docs/url_queries.md
	params := url.Values{}
	params.Add("sort", "id:desc")

	// Call the GetSelfServiceBrandingMacOS function and handle any errors
	branding, err := client.GetSelfServiceBrandingMacOS(params)
	if err != nil {
		// If there's an error, log it to stderr and exit with a non-zero status code
		fmt.Fprintf(os.Stderr, "Error fetching self-service branding for macOS: %v\n", err)
		os.Exit(1)
	}

	createdScriptJSON, err := json.MarshalIndent(branding, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created script data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(createdScriptJSON))

}
