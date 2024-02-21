package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Example VPP Account ID
	vppAccountID := 1

	// Call GetVPPAccountByID
	vppAccount, err := client.GetVPPAccountByID(vppAccountID)
	if err != nil {
		log.Fatalf("Error fetching VPP account by ID: %v", err)
	}

	// Print the retrieved VPP account details
	fmt.Printf("ID: %d\nName: %s\nContact: %s\nService Token: %s\nAccount Name: %s\nExpiration Date: %s\nCountry: %s\nApple ID: %s\nSite ID: %d\nSite Name: %s\nPopulate Catalog From VPP Content: %t\nNotify Disassociation: %t\nAuto Register Managed Users: %t\n",
		vppAccount.ID, vppAccount.Name, vppAccount.Contact, vppAccount.ServiceToken, vppAccount.AccountName, vppAccount.ExpirationDate, vppAccount.Country, vppAccount.AppleID, vppAccount.Site.ID, vppAccount.Site.Name, vppAccount.PopulateCatalogFromVPPContent, vppAccount.NotifyDisassociation, vppAccount.AutoRegisterManagedUsers)
}
