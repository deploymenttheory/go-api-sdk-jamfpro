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

	// Modern Jamf Pro API (/api/v1/users) users resource.
	users, err := client.GetUsersV1(nil)
	if err != nil {
		log.Fatalf("Error retrieving users: %v", err)
	}

	out, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling users: %v", err)
	}
	fmt.Println("Users (Jamf Pro API v1):\n", string(out))
}
