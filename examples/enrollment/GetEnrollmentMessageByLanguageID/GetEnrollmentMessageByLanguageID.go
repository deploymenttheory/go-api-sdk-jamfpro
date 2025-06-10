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

	// Specify the language ID (ISO 639-1 code)
	languageId := "en"

	// Call GetEnrollmentMessageByLanguageID function
	languageMsg, err := client.GetEnrollmentMessageByLanguageID(languageId)
	if err != nil {
		log.Fatalf("Error getting enrollment language messaging: %v", err)
	}

	// Pretty print the language messaging configuration in JSON
	JSON, err := json.MarshalIndent(languageMsg, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling language messaging data: %v", err)
	}
	fmt.Println("Enrollment Language Messaging Configuration:\n", string(JSON))
}
