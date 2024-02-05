package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/userauth.json"

	// Load the client authentication details from the configuration file
	authConfig, err := http_client.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client authentication configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			Username: authConfig.Username,
			Password: authConfig.Password,
		},
		LogLevel: logLevel,
	}

	// Create a new client instance using the loaded InstanceName
	client, err := http_client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}

	// Set BearerTokenAuthCredentials
	client.SetBearerTokenAuthCredentials(http_client.BearerTokenAuthCredentials{
		Username: authConfig.Username,
		Password: authConfig.Password,
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
