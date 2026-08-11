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

	params := url.Values{}

	mobileDevicesInventory, err := client.GetMobileDevicesInventory(params)
	if err != nil {
		log.Fatalf("Error fetching mobile devices inventory: %v", err)
	}

	fmt.Printf("Total mobile devices: %d\n\n", mobileDevicesInventory.TotalCount)

	// Pretty print the response
	prettyJSON, err := json.MarshalIndent(mobileDevicesInventory, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %v", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}
