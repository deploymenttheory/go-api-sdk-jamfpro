package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Directly assign the configuration values
	authConfig := httpclient.AuthConfig{
		ClientID:     "7871012d-98bb-454e-8d17-6cd0f8cbfb7b",
		ClientSecret: "xKwuGxBPBUIXZERLKsbv6d-mVfDvVj6cc2YOiiXVY25mT2aSvmwD61HrzXB_lySy",
		Username:     "apiuser",
		Password:     "password",
	}

	envConfig := httpclient.EnvironmentConfig{
		InstanceName:       "lbgsandbox",
		OverrideBaseDomain: "",
		APIType:            "jamfpro",
	}

	clientOptions := httpclient.ClientOptions{
		LogLevel:                  "LogLevelInfo", // Adjust the log level as needed
		LogOutputFormat:           "console",
		HideSensitiveData:         true,
		EnableDynamicRateLimiting: true,
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth:          authConfig,
		Environment:   envConfig,
		ClientOptions: clientOptions,
	}

	// Create a new client instance using the configuration
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}

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
