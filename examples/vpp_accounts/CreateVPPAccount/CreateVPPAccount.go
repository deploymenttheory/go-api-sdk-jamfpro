package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Create a new VPP account
	newAccount := &jamfpro.ResponseVPPAccount{
		Name:         "Company VPP Account",
		Contact:      "Company Admin",
		ServiceToken: "eyJvcmdOYWadveaz40d2FyZSIsImV4cERhdGUiOiIyMDE3LTA5LTEzVDA5OjQ5OjA5LTA3MDAiLCJ0b2tlbiI6Ik5yVUtPK1RXeityUXQyWFpIeENtd0xxby8ydUFmSFU1NW40V1FTZU8wR1E5eFh4UUZTckVJQjlzbGdYei95WkpaeVZ3SklJbW0rWEhJdGtKM1BEZGRRPT0ifQ==",
		AccountName:  "Company Name",
		AppleID:      "vpp@company.com",
		// Site information is optional, defaults will be set if not provided
		PopulateCatalogFromVPPContent: true,
		NotifyDisassociation:          true,
		AutoRegisterManagedUsers:      false,
	}

	createdAccount, err := client.CreateVPPAccount(newAccount)
	if err != nil {
		log.Fatalf("Error creating VPP account: %v", err)
	}

	fmt.Printf("Created VPP Account: %+v\n", createdAccount)
}
