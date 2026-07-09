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

	clientManagementID := "550e8400-e29b-41d4-a716-446655440000"

	details, err := client.GetMdmRenewalDeviceCommonDetailsByClientManagementID(clientManagementID)
	if err != nil {
		log.Fatalf("Error retrieving MDM renewal device common details: %v", err)
	}

	out, err := json.MarshalIndent(details, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling MDM renewal device common details: %v", err)
	}
	fmt.Println("MDM Renewal Device Common Details:\n", string(out))
}
