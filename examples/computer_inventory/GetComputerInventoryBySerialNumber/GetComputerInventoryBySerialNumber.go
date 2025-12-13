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

	// Define the serial number of the computer inventory you want to retrieve
	serialNumber := "C02ABC123DEF"

	// Call the GetComputerInventoryBySerialNumber function
	computerInventory, err := client.GetComputerInventoryBySerialNumber(serialNumber)
	if err != nil {
		log.Fatalf("Error fetching computer inventory by serial number: %v", err)
	}

	// Pretty print the response
	prettyJSON, err := json.MarshalIndent(computerInventory, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %v", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}
