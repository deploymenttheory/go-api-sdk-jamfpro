package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := http_client.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Debug print statement to check the loaded configuration
	fmt.Printf("Loaded Config: %+v\n", authConfig)

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			Username: "fwfw",
			Password: "fwfw",
		},
		LogLevel: logLevel,
	}

	// Create a new client instance using the loaded InstanceName
	client, err := http_client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}

	// Set OAuth credentials for the client
	oAuthCreds := http_client.OAuthCredentials{
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
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
