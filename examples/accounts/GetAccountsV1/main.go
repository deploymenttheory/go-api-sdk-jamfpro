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

	// Modern Jamf Pro API (/api/v1/accounts) user accounts resource.
	accounts, err := client.GetAccountsV1(nil)
	if err != nil {
		log.Fatalf("Error retrieving accounts: %v", err)
	}

	out, err := json.MarshalIndent(accounts, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling accounts: %v", err)
	}
	fmt.Println("Accounts (Jamf Pro API v1):\n", string(out))
}
