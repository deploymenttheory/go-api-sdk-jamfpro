package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/neilmartin/GitHub/go-api-sdk-jamfpro/localtesting/config.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	engageInfo, err := client.GetEngageSettings()
	if err != nil {
		log.Fatalf("Failed to get engage info, %v", err)
	}

	engageInfoJson, err := json.MarshalIndent(engageInfo, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal json, %v", err)
	}

	fmt.Println(string(engageInfoJson))
}
