package main

import (
	"encoding/xml"
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

	// Call the GetUserExtensionAttributes function
	userExtAttributesList, err := client.GetUserExtensionAttributes()
	if err != nil {
		log.Fatalf("Error fetching user extension attributes: %v", err)
	}

	// Pretty print the user extension attributes details in XML
	userExtAttributesXML, err := xml.MarshalIndent(userExtAttributesList, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling user extension attributes data: %v", err)
	}
	fmt.Println("User Extension Attributes Details:\n", string(userExtAttributesXML))
}
