package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the MDM blank push request
	blankPushRequest := &jamfpro.ResourceBlankPush{
		ClientManagementIDs: []string{
			"fd68c371-5921-436e-b16b-8a3c1bf90ee5",
			"a1b2c3d4-5678-90ab-cdef-1234567890ab",
		},
	}

	// Call SendMDMCommandForBlankPush function
	response, err := client.SendMDMCommandForBlankPush(blankPushRequest)
	if err != nil {
		log.Fatalf("Error sending MDM blank push command: %v", err)
	}

	// Pretty print the response details in JSON
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling response data: %v", err)
	}
	fmt.Println("MDM Blank Push Response:\n", string(responseJSON))
}
