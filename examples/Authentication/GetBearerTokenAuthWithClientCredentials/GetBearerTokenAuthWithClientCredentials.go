package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
)

func main() {
	clientID := "your-client-secret"
	clientSecret := "your-clientid"
	baseURL := "your-jamf-instance"

	// Configuration for the client
	config := http_client.Config{
		DebugMode: true,
		// You can add other configurations as needed
		Logger: http_client.NewDefaultLogger(),
	}

	// Create a new client instance
	client, err := http_client.NewClient(baseURL, config, nil)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}

	// Set OAuth credentials for the client
	oAuthCreds := http_client.OAuthCredentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	client.SetOAuthCredentials(oAuthCreds)

	// Call the ValidAuthTokenCheck function to ensure that a valid token is set in the client
	isTokenValid, err := client.ValidAuthTokenCheck()
	if err != nil {
		log.Fatalf("Error while validating token: %v", err)
	}
	if !isTokenValid {
		fmt.Println("Error obtaining or refreshing token.")
		return
	}

	// Print the obtained token
	fmt.Println("Successfully obtained token:", client.Token)
}
