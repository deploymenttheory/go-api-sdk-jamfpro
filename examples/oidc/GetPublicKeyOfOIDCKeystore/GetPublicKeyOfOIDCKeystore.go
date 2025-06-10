package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Get the public key of the OIDC keystore
	publicKeyResponse, err := client.GetPublicKeyOfOIDCKeystore()
	if err != nil {
		log.Fatalf("Error retrieving OIDC public key: %v", err)
	}

	// Pretty print the response in JSON
	publicKeyJSON, err := json.MarshalIndent(publicKeyResponse, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling OIDC public key data: %v", err)
	}
	fmt.Println("OIDC Public Key Details:\n", string(publicKeyJSON))
}
