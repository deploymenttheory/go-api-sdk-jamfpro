package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
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
