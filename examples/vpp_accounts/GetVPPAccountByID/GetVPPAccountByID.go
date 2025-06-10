package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example VPP Account ID
	vppAccountID := "1"

	// Call GetVPPAccountByID
	vppAccount, err := client.GetVPPAccountByID(vppAccountID)
	if err != nil {
		log.Fatalf("Error fetching VPP account by ID: %v", err)
	}

	// Print the retrieved VPP account details
	fmt.Printf("ID: %d\nName: %s\nContact: %s\nService Token: %s\nAccount Name: %s\nExpiration Date: %s\nCountry: %s\nApple ID: %s\nSite ID: %d\nSite Name: %s\nPopulate Catalog From VPP Content: %t\nNotify Disassociation: %t\nAuto Register Managed Users: %t\n",
		vppAccount.ID, vppAccount.Name, vppAccount.Contact, vppAccount.ServiceToken, vppAccount.AccountName, vppAccount.ExpirationDate, vppAccount.Country, vppAccount.AppleID, vppAccount.Site.ID, vppAccount.Site.Name, vppAccount.PopulateCatalogFromVPPContent, vppAccount.NotifyDisassociation, vppAccount.AutoRegisterManagedUsers)
}
