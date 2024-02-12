package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
)

func main() {
	// Manually set your client credentials
	username := "your-jamf-api-account"
	password := "your-jamf-api-account-password"
	instanceName := "your-jamf-instance" // e.g., "yourcompany.jamfcloud.com"
	apitype := "jamf"

	// Instantiate the default logger and set the desired log level
	logLevel := "logger.LogLevelDebug" // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     username,
			ClientSecret: password,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      apitype,
			InstanceName: instanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel: logLevel,
		},
	}

	// Create a new client instance using the provided BaseURL
	client, err := httpclient.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}

	// Set BearerTokenAuthCredentials
	client.SetBearerTokenAuthCredentials(httpclient.BearerTokenAuthCredentials{
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
