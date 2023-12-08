package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/Joseph/github/go-api-sdk-jamfpro/clientauth.json"
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("failed to load config")
	}

	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelInfo
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create jamfpro client")
	}

	scripts, err := client.GetProScripts()
	if err != nil {
		log.Fatalf("Failure getting scripts, %v", err)
	}

	for _, value := range scripts.Results {
		fmt.Printf("%+v\n", value)
		fmt.Println()
	}
}
