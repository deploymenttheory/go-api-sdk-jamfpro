package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	settingsUpdate := jamfpro.ResourcePolicyProperties{
		PoliciesRequireNetworkStateChange: true,
		AllowNetworkStateChangeTriggers:   true,
	}

	clientChecinInfo, err := client.UpdatePolicyProperties(settingsUpdate)
	if err != nil {
		log.Fatalf("Failed to get client checkin info, %v", err)
	}

	checkinInfoJson, err := json.MarshalIndent(clientChecinInfo, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal json, %v", err)
	}

	fmt.Println(string(checkinInfoJson))
}
