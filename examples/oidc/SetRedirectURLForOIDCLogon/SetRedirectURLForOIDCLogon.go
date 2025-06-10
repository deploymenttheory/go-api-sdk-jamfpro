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

	// Define the original URL and email address
	originalURL := "aHR0cHM6Ly9qYW1mLXByby11cmwuY29tL2xvZ2dpbmcuaHRtbA=="
	emailAddress := "admin@domain.name"

	// Create the request struct
	request := &jamfpro.ResourceOIDCRedirectURL{
		OriginalURL:  originalURL,
		EmailAddress: emailAddress,
	}

	// Get the redirect URL for OIDC login
	redirectResponse, err := client.SetRedirectURLForOIDCLogon(request)
	if err != nil {
		log.Fatalf("Error setting OIDC redirect URL: %v", err)
	}

	// Pretty print the response in JSON
	redirectJSON, err := json.MarshalIndent(redirectResponse, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling OIDC redirect URL data: %v", err)
	}
	fmt.Println("OIDC Redirect URL Details:\n", string(redirectJSON))
}
