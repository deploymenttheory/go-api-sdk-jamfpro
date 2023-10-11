package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := http_client.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Debug print statement to check the loaded configuration
	fmt.Printf("Loaded Config: %+v\n", authConfig)

	// Configuration for the client
	config := http_client.Config{
		DebugMode: true,
		Logger:    http_client.NewDefaultLogger(),
	}

	// Create a new client instance using the loaded InstanceName
	client := http_client.NewClient(authConfig.InstanceName, config, nil)

	// Set OAuth credentials for the client
	oAuthCreds := http_client.OAuthCredentials{
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}
	client.SetOAuthCredentials(oAuthCreds)

	// Call the ValidAuthTokenCheck function to ensure that a valid token is set in the client
	isTokenValid := client.ValidAuthTokenCheck()
	if !isTokenValid {
		fmt.Println("Error obtaining or refreshing token.")
		return
	}

	// Print the obtained token
	fmt.Println("Successfully obtained token:", client.Token)
}
