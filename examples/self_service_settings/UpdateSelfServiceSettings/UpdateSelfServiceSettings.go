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

	// Set up new settings for self service
	newSettings := jamfpro.ResourceSelfServiceSettings{
		InstallSettings: jamfpro.InstallSettings{
			InstallAutomatically: true,
			InstallLocation:      "/Applications",
		},
		LoginSettings: jamfpro.LoginSettings{
			UserLoginLevel:  "Anonymous",
			AllowRememberMe: false,
			AuthType:        "Saml", // Basic / Saml
		},
		ConfigurationSettings: jamfpro.ConfigurationSettings{
			NotificationsEnabled:  true,
			AlertUserApprovedMdm:  true,
			DefaultLandingPage:    "HOME",
			DefaultHomeCategoryId: -1,
			BookmarksName:         "Bookmarks",
		},
	}

	// Update self service settings
	updatedSettings, err := client.UpdateSelfServiceSettings(&newSettings)
	if err != nil {
		fmt.Printf("Error updating self service settings: %s\n", err)
		return
	}

	fmt.Printf("self service settings updated successfully: %+v\n", updatedSettings)
}
