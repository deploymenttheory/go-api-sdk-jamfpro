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

	languageCodes, err := client.GetEnrollmentLanguageCodes()
	if err != nil {
		log.Fatalf("Error getting enrollment language codes: %v", err)
	}

	JSON, err := json.MarshalIndent(languageCodes, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling language codes data: %v", err)
	}
	fmt.Println("Available Language Codes:\n", string(JSON))
}
