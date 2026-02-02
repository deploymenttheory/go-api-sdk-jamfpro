package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	updatePayload := &jamfpro.ResourceCloudDistributionPointV1{
		CdnType: "JAMF_CLOUD",
		Master:  false,
	}

	updated, err := client.UpdateCloudDistributionPointV1(updatePayload)
	if err != nil {
		log.Fatalf("Error updating Cloud Distribution Point: %v", err)
	}

	body, err := json.MarshalIndent(updated, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling response: %v", err)
	}

	fmt.Println("Updated Cloud Distribution Point:\n", string(body))
}
