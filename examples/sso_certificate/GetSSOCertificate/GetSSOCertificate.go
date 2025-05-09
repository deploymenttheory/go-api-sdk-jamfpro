package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Get SSO certificate information
	certInfo, err := client.GetSSOCertificate()
	if err != nil {
		log.Fatalf("Error fetching SSO certificate information: %v", err)
	}

	// Print certificate details
	fmt.Printf("Certificate Type: %s\n", certInfo.Keystore.Type)
	fmt.Printf("Keystore Filename: %s\n", certInfo.Keystore.KeystoreFileName)
	fmt.Printf("Setup Type: %s\n", certInfo.Keystore.KeystoreSetupType)

	if certInfo.KeystoreDetails != nil {
		fmt.Printf("Issuer: %s\n", certInfo.KeystoreDetails.Issuer)
		fmt.Printf("Subject: %s\n", certInfo.KeystoreDetails.Subject)
		fmt.Printf("Expiration: %s\n", certInfo.KeystoreDetails.Expiration)
	}
}
