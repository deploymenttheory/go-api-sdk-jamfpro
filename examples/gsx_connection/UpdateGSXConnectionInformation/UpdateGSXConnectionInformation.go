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

	// Define new GSX Connection settings
	newGSXSettings := &jamfpro.ResourceGSXConnection{
		Enabled:          false,
		Username:         "",  // Empty string to denote no username
		ServiceAccountNo: "0", // Zero to denote no account number
		ShipToNo:         "0", // Zero to denote no ship-to number
		GsxKeystore: jamfpro.GsxKeystore{
			Name:            "certificate.p12",
			ExpirationEpoch: 169195490000,
			ErrorMessage:    "Certificate error",
		},
	}

	// Call the UpdateGSXConnectionInformation function
	updatedGSXSettings, err := client.UpdateGSXConnectionInformation(newGSXSettings)
	if err != nil {
		log.Fatalf("Error updating GSX Connection Information: %v", err)
	}

	// Pretty print the updated group details
	groupXML, err := json.MarshalIndent(updatedGSXSettings, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling group data: %v", err)
	}
	fmt.Println("Updated GSX Connection Information: Details:", string(groupXML))
}
