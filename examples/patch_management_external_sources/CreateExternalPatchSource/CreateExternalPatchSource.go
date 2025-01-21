package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create new patch source
	newSource := &jamfpro.ResourcePatchExternalSource{
		Name:       "Test External Source",
		HostName:   "patch.example.com",
		Port:       443,
		SSLEnabled: true,
	}

	// Create the patch external source
	createdSource, err := client.CreateExternalPatchSource(newSource)
	if err != nil {
		log.Fatalf("Error creating patch external source: %v", err)
	}

	// Pretty print the created source
	resp, err := json.MarshalIndent(createdSource, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created source data: %v", err)
	}
	fmt.Printf("Created Patch External Source:\n%s\n", string(resp))
}
