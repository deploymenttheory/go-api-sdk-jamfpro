package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
)

func main() {
	// Manually set your client credentials
	username := "your-jamf-api-account"
	password := "your-jamf-api-account-password"
	instanceName := "your-jamf-instance" // e.g., "yourcompany.jamfcloud.com"

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: instanceName,
		LogLevel:     logLevel,
		Auth: http_client.AuthConfig{
			Username: username,
			Password: password,
		},
	}

	// Create a new client instance using the provided BaseURL
	client, err := http_client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}

	// Set BearerTokenAuthCredentials
	client.SetBearerTokenAuthCredentials(http_client.BearerTokenAuthCredentials{
		Username: username,
		Password: password,
	})

	// Call the ObtainToken function to get a bearer token
	err = client.ObtainToken()
	if err != nil {
		fmt.Println("Error obtaining token:", err)
		return
	}

	// Print the obtained token
	fmt.Println("Successfully obtained token:", client.Token)
}
