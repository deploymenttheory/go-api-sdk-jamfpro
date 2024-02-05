package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
)

func main() {
	clientID := "your-client-secret"
	clientSecret := "your-clientid"
	instanceName := "your-jamf-instance"

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: instanceName,
		LogLevel:     logLevel,
		Auth: http_client.AuthConfig{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
	}

	// Create a new client instance
	client, err := http_client.NewClient(config)
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
