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

	lastLogin, err := client.GetLastLogin()
	if err != nil {
		log.Fatalf("Error retrieving last login: %v", err)
	}

	out, err := json.MarshalIndent(lastLogin, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling last login: %v", err)
	}
	fmt.Println("Last Login:\n", string(out))
}
