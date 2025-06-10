package main

import (
	"encoding/xml"
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

	apiRoleID := "1"

	// Call GetJamfApiRolesByID function
	apiRole, err := client.GetJamfApiRoleByID(apiRoleID)
	if err != nil {
		log.Fatalf("Error fetching Jamf API role by ID: %v", err)
	}

	// Pretty print the fetched API role in XML
	apiRoleXML, err := xml.MarshalIndent(apiRole, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf API role data: %v", err)
	}
	fmt.Println("Fetched Jamf API Role by ID:\n", string(apiRoleXML))
}
