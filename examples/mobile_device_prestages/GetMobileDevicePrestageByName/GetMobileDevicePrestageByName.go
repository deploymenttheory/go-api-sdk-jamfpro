package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/neilmartin/GitHub/go-api-sdk-jamfpro/client_auth.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// ID of the mobile device prestage you want to retrieve
	prestageName := "jamfpro-sdk-example-mobiledevicePrestage-config" // Replace with the actual Name

	// Call the GetMobileDevicePrestageByID function
	prestage, err := client.GetMobileDevicePrestageByName(prestageName)
	if err != nil {
		log.Fatalf("Error fetching mobile device prestage by Name: %v", err)
	}

	// Pretty print the mobile device prestage in JSON
	prestageJSON, err := json.MarshalIndent(prestage, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling mobile device prestage data: %v", err)
	}
	fmt.Println("Fetched mobile device prestage:\n", string(prestageJSON))
}
