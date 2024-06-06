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

	// Define the MDM profile renewal request
	renewProfileRequest := &jamfpro.ResourceMDMProfileRenewal{
		UDIDs: []string{
			"6E47EF55-5318-494F-A09E-70F613E0AFD1",
			"6E47EF55-5318-494F-A09E-70F613E0AFD1",
			"6E47EF55-5318-494F-A09E-70F613E0AFD1",
		},
	}

	// Call SendMDMCommandForMDMProfileRenewal function
	response, err := client.SendMDMCommandForMDMProfileRenewal(renewProfileRequest)
	if err != nil {
		log.Fatalf("Error renewing MDM profile: %v", err)
	}

	// Pretty print the response details in JSON
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling response data: %v", err)
	}
	fmt.Println("MDM Profile Renewal Response:\n", string(responseJSON))
}
