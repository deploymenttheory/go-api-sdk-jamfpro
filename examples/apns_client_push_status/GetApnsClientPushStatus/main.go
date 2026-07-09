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

	// Retrieve the list of clients that currently have push notifications disabled.
	statuses, err := client.GetApnsClientPushStatus(nil)
	if err != nil {
		log.Fatalf("Error retrieving APNS client push status: %v", err)
	}

	out, err := json.MarshalIndent(statuses, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling APNS client push status: %v", err)
	}
	fmt.Println("APNS Client Push Status:\n", string(out))
}
