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

	// Fetch all scripts
	// For more information on how to add parameters to this request, see docs/url_queries.md
	scripts, err := client.GetScripts(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching scripts: %v", err)
	}

	fmt.Println("Scripts fetched. Starting deletion process:")

	// Iterate over each script and delete
	for _, script := range scripts.Results {
		fmt.Printf("Deleting script ID: %s, Name: %s\n", script.ID, script.Name)

		err = client.DeleteScriptByID(script.ID)
		if err != nil {
			log.Printf("Error deleting script ID %s: %v\n", script.ID, err)
			continue // Move to the next script if there's an error
		}

		fmt.Printf("Script ID %s deleted successfully.\n", script.ID)
	}

	fmt.Println("Script deletion process completed.")
}
