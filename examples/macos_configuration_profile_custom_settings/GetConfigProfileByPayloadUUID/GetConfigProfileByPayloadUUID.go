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

	PayloadUUID := "f288ac07-014c-4b5e-a95c-d3bd0892d7fa"

	// Call the function to get the profile by name
	profile, err := client.GetConfigProfileByPayloadUUID(PayloadUUID)
	if err != nil {
		fmt.Printf("Error fetching profile: %v\n", err)
		return
	}

	configurationProfileXML, err := json.MarshalIndent(profile, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling profile data: %v", err)
	}
	fmt.Println("Fetched Profile Details:\n", string(configurationProfileXML))
}
