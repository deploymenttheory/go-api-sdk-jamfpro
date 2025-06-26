package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/ecanault/.go/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	reenrollmentSettingsInfo, err := client.GetReenrollmentSettings()
	if err != nil {
		log.Fatalf("Failed to get re-enrollment settings info, %v", err)
	}

	reenrollmentSettingsInfoJson, err := json.MarshalIndent(reenrollmentSettingsInfo, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal json, %v", err)
	}

	fmt.Println(string(reenrollmentSettingsInfoJson))
}
