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

	info, err := client.GetJamfProInformation()
	if err != nil {
		log.Fatalf("Error fetching Jamf Pro Information: %s", err)
	}

	// Convert the info struct to pretty-printed JSON
	infoJSON, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling Jamf Pro Information to JSON: %s", err)
	}

	// Print the pretty-printed JSON
	fmt.Println("Jamf Pro Information:")
	fmt.Println(string(infoJSON))

}
