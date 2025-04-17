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

	// Sort filter
	// For more information on how to add parameters to this request, see docs/url_queries.md
	params := url.Values{}
	params.Add("sort", "id:desc")

	// Call the GetComputersFileVaultInventory function
	fileVaultInventory, err := client.GetComputersFileVaultInventory(params)
	if err != nil {
		log.Fatalf("Error fetching FileVault inventory: %v", err)
	}

	// Pretty print the response
	prettyJSON, err := json.MarshalIndent(fileVaultInventory, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %v", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}
