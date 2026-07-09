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

	// Modern Jamf Pro API (/api/v1/account-groups) account groups resource.
	groups, err := client.GetAccountGroupsV1(nil)
	if err != nil {
		log.Fatalf("Error retrieving account groups: %v", err)
	}

	out, err := json.MarshalIndent(groups, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling account groups: %v", err)
	}
	fmt.Println("Account Groups (Jamf Pro API v1):\n", string(out))
}
