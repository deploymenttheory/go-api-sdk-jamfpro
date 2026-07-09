package main

import (
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Perform the initial platform setup of a fresh Jamf Pro server.
	request := &jamfpro.ResourceSystemPlatformInitialize{
		ActivationCode:  "VFAB-YDAB-DFAB-UDAB-DEAB-EFAB-ABAB-DEAB",
		InstitutionName: "Jamf",
		EulaAccepted:    true,
		Username:        "admin",
		Email:           "admin@jamf.com",
		JssURL:          "https://jamf.jamfcloud.com",
	}

	if err := client.InitializeSystemPlatform(request); err != nil {
		log.Fatalf("Error initializing system platform: %v", err)
	}

	log.Println("System platform initialization request accepted.")
}
