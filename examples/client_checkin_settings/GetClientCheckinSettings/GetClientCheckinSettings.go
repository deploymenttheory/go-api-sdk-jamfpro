package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/joseph/github/go-api-sdk-jamfpro/clientauth.json"
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client auth data, %v", err)
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
		log.Fatalf("Failed to init client, %v", err)
	}

	clientChecinInfo, err := client.GetClientCheckinSettings()
	if err != nil {
		log.Fatalf("Failed to get client checkin info, %v", err)
	}

	checkinInfoJson, err := json.MarshalIndent(clientChecinInfo, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal json, %v", err)
	}

	fmt.Println(string(checkinInfoJson))
}
