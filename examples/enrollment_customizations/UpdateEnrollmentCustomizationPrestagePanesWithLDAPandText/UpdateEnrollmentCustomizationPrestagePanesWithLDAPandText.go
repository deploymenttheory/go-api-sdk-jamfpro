package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

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

	// Specify the enrollment customization ID
	customizationID := "22" // Replace with your actual customization ID

	// Define update templates for each pane type

	// Text pane update template
	updatedTextPane := jamfpro.ResourceEnrollmentCustomizationTextPane{
		DisplayName:        "Updated Text Pane",
		Rank:               0,
		Title:              "Welcome to Enrollment",
		Body:               "This is the updated text content for the enrollment process.",
		Subtext:            "Please follow the instructions to complete enrollment.",
		BackButtonText:     "Back",
		ContinueButtonText: "Continue",
	}

	// LDAP pane update template
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

	// Get all prestage panes for the specified enrollment customization
	prestagePanes, err := client.GetPrestagePanes(customizationID)
	if err != nil {
		log.Fatalf("Failed to get prestage panes: %v", err)
	}

	// Print the list of panes
	fmt.Println("Found the following prestage panes:")
	for _, pane := range prestagePanes.Panels {
		fmt.Printf("- %s (ID: %d, Type: %s, Rank: %d)\n",
			pane.DisplayName,
			pane.ID,
			pane.Type,
			pane.Rank)
	}

	// Update each pane based on its type
	for _, pane := range prestagePanes.Panels {
		paneID := strconv.Itoa(pane.ID)

		switch pane.Type {
		case "text":
			// Clone the template and adjust any panel-specific settings if needed
			textPaneUpdate := updatedTextPane
			textPaneUpdate.Rank = pane.Rank
			textPaneUpdate.DisplayName = fmt.Sprintf("%s - Updated", pane.DisplayName)

			updateTextPane(client, customizationID, paneID, textPaneUpdate)

		case "ldap":
			// Clone the template and adjust any panel-specific settings if needed
			ldapPaneUpdate := updatedLDAPPane
			ldapPaneUpdate.Rank = pane.Rank
			ldapPaneUpdate.DisplayName = fmt.Sprintf("%s - Updated", pane.DisplayName)

			updateLDAPPane(client, customizationID, paneID, ldapPaneUpdate)

		default:
			fmt.Printf("Unknown pane type: %s for pane ID: %d\n", pane.Type, pane.ID)
		}
	}
}

// updateTextPane updates a text pane with new settings
func updateTextPane(client *jamfpro.Client, customizationID, paneID string, updatedPane jamfpro.ResourceEnrollmentCustomizationTextPane) {
	// Make sure the type is set correctly
	updatedPane.Type = "text"

	// Update the text pane
	result, err := client.UpdateTextPrestagePaneByID(customizationID, paneID, updatedPane)
	if err != nil {
		log.Printf("Failed to update text pane %s: %v", paneID, err)
		return
	}

	// Print the result
	prettyJSON, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("\nSuccessfully updated Text pane (ID: %s):\n%s\n", paneID, string(prettyJSON))
}

// updateLDAPPane updates an LDAP pane with new settings
func updateLDAPPane(client *jamfpro.Client, customizationID, paneID string, updatedPane jamfpro.ResourceEnrollmentCustomizationLDAPPane) {
	// Make sure the type is set correctly
	updatedPane.Type = "ldap"

	// Update the LDAP pane
	result, err := client.UpdateLDAPPrestagePaneByID(customizationID, paneID, updatedPane)
	if err != nil {
		log.Printf("Failed to update LDAP pane %s: %v", paneID, err)
		return
	}

	// Print the result
	prettyJSON, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("\nSuccessfully updated LDAP pane (ID: %s):\n%s\n", paneID, string(prettyJSON))
}
