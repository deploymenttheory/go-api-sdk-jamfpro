package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	resp, err := client.GetEnrollmentSettings()
	if err != nil {
		log.Fatalf("Error fetching Enrollment Settings: %v", err)
	}

	jsonData, err := json.MarshalIndent(resp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)
	}
	fmt.Println("Fetched data:\n", string(jsonData))

}
