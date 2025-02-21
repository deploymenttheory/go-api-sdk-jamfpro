package main

import (
	"encoding/json"
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

	// Specify the enrollment customization ID and panel ID
	customizationID := "22" // Replace with your actual customization ID
	panelID := "19"         // Replace with your actual panel ID

	// Prepare the updated LDAP pane settings
	updatedLDAPPane := jamfpro.ResourceEnrollmentCustomizationLDAPPane{
		DisplayName:        "Updated LDAP Pane",
		Rank:               1,
		Title:              "LDAP Authentication",
		UsernameLabel:      "Username",
		PasswordLabel:      "Password",
		BackButtonText:     "Back",
		ContinueButtonText: "Log In",
		LDAPGroupAccess: []jamfpro.EnrollmentCustomizationLDAPGroupAccess{
			{
				GroupName:    "IT Staff",
				LDAPServerID: 1,
			},
			{
				GroupName:    "Employees",
				LDAPServerID: 1,
			},
		},
	}

	// Update the LDAP prestage pane
	result, err := client.UpdateLDAPPrestagePaneByID(customizationID, panelID, updatedLDAPPane)
	if err != nil {
		log.Fatalf("Failed to update LDAP prestage pane: %v", err)
	}

	// Pretty print the result in JSON
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling result: %v", err)
	}
	fmt.Println("Updated LDAP Prestage Pane:\n", string(prettyJSON))

}
