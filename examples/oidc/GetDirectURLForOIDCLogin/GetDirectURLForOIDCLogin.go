package main

import (
	"encoding/json"
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

	// Get the OIDC Direct IdP Login URL
	oIDCDirectIdPLoginURLResponse, err := client.GetDirectURLForOIDCLogin()
	if err != nil {
		log.Fatalf("Error retrieving OIDC direct IdP login URL: %v", err)
	}

	// Pretty print the response in JSON
	oIDCDirectIdPLoginURLJSON, err := json.MarshalIndent(oIDCDirectIdPLoginURLResponse, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling OIDC direct IdP login URL data: %v", err)
	}
	fmt.Println("OIDC Direct IdP Login URL Details:\n", string(oIDCDirectIdPLoginURLJSON))
}
