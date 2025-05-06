package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Call the GetDeviceEnrollmentsPublicKey method to retrieve the public key
	response, err := client.GetDeviceEnrollmentsPublicKey()
	if err != nil {
		log.Fatalf("Failed to retrieve Public Key: %v", err)
	}
	fmt.Print(response)
}
