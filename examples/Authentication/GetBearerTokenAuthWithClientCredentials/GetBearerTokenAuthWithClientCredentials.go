package main

import (
	"fmt"

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
	client := http_client.NewClient(baseURL, config, nil)

	// Set OAuth credentials for the client
	oAuthCreds := http_client.OAuthCredentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
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
