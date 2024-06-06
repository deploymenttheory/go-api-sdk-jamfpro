package main

import (
	"encoding/json"
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

	// Example ID to fetch package history
	packageID := 233

	response, err := client.GetPackageHistoryByPackageID(packageID, "date:desc", "")
	if err != nil {
		fmt.Println("Error fetching package history by package ID:", err)
		return
	}

	// Pretty print the package history details in JSON
	historyJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling package history data: %v", err)
	}
	fmt.Println("Obtained package history details:\n", string(historyJSON))
}
