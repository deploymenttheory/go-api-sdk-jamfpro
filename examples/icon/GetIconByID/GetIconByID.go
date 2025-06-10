package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	iconID := 28

	response, err := client.GetIconByID(iconID)
	if err != nil {
		log.Fatalf("Failed to get icon: %v", err)
	}

	// Pretty print the icon details
	iconJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling accounts data: %v", err)
	}
	fmt.Println("Fetched Icon:", string(iconJSON))
}
