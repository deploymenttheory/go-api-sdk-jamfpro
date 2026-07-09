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

	sessions, err := client.GetActiveUserSessions()
	if err != nil {
		log.Fatalf("Error retrieving active user sessions: %v", err)
	}

	out, err := json.MarshalIndent(sessions, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling active user sessions: %v", err)
	}
	fmt.Println("Active User Sessions:\n", string(out))
}
