package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Create a new VPP account
	newAccount := &jamfpro.ResourceVPPAccount{
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
