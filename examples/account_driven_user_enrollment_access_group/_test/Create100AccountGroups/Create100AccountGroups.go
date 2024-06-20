package main

import (
	"encoding/xml"
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

	for i := 1; i <= 100; i++ {
		// Assemble the request body for creating an account group with a unique name
		accountGroupName := fmt.Sprintf("jamf sdk test account group %d", i) // Unique name for each group
		accountGroup := &jamfpro.ResourceAccountGroup{
			Name:         accountGroupName,
			AccessLevel:  "Full Access", // Full Access / Site Access
			PrivilegeSet: "Custom",      // Administrator / Auditor / Enrollment Only / Custom
			Site: jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
			Privileges: jamfpro.AccountSubsetPrivileges{
				JSSObjects:    []string{"Update Webhooks", "Delete Webhooks"},
				JSSSettings:   []string{"Read SSO Settings", "Update User-Initiated Enrollment"},
				JSSActions:    []string{"Send Computer Bluetooth Command", "Computer Delete User Account Command"},
				Recon:         []string{"string"},
				CasperAdmin:   []string{"Use Casper Admin", "Save With Casper Admin"},
				CasperRemote:  []string{"string"},
				CasperImaging: []string{"string"},
			},
			Members: jamfpro.AccountGroupSubsetMembers{
				{User: jamfpro.MemberUser{ID: 480, Name: "Barry White"}},
				{User: jamfpro.MemberUser{ID: 2, Name: "dafydd.watkins"}},
			},
		}

		// Call CreateAccountGroupByID function
		createdAccountGroup, err := client.CreateAccountGroup(accountGroup)
		if err != nil {
			log.Printf("Error creating account group '%s': %v", accountGroupName, err)
			continue // Continue with the next iteration in case of an error
		}

		// Pretty print the created account group details
		accountGroupXML, err := xml.MarshalIndent(createdAccountGroup, "", "    ") // Indent with 4 spaces
		if err != nil {
			log.Printf("Error marshaling account group data: %v", err)
			continue
		}
		fmt.Printf("Created Account Group Details for '%s':\n%s\n", accountGroupName, string(accountGroupXML))
	}
}
