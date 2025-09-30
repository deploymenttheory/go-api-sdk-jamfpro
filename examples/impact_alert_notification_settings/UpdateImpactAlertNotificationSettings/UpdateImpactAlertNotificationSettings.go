package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create Impact Alert Notification settings to update
	settings := jamfpro.ResourceImpactAlertNotificationSettings{
		ScopeableObjectsAlertEnabled:             true,
		ScopeableObjectsConfirmationCodeEnabled:  false,
		DeployableObjectsAlertEnabled:            true,
		DeployableObjectsConfirmationCodeEnabled: false,
	}

	// Update the settings
	err = client.UpdateImpactAlertNotificationSettings(settings)
	if err != nil {
		log.Fatalf("Error updating Impact Alert Notification settings: %v", err)
	}

	fmt.Println("Successfully updated Impact Alert Notification settings")
}
