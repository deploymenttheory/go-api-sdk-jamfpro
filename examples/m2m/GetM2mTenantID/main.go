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

	tenant, err := client.GetM2mTenantID()
	if err != nil {
		log.Fatalf("Error retrieving M2M tenant id: %v", err)
	}

	out, err := json.MarshalIndent(tenant, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling M2M tenant id: %v", err)
	}
	fmt.Println("M2M Tenant ID:\n", string(out))
}
