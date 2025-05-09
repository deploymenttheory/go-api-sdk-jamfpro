package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create new SSO certificate
	newCert, err := client.CreateSSOCertificate()
	if err != nil {
		log.Fatalf("Error generating new SSO certificate: %v", err)
	}

	fmt.Println("New SSO certificate generated successfully!")
	fmt.Printf("Certificate Type: %s\n", newCert.Keystore.Type)
	fmt.Printf("Setup Type: %s\n", newCert.Keystore.KeystoreSetupType)

	// Print validation status for each key
	for _, key := range newCert.Keystore.Keys {
		fmt.Printf("Key ID: %s, Valid: %v\n", key.ID, key.Valid)
	}
}
