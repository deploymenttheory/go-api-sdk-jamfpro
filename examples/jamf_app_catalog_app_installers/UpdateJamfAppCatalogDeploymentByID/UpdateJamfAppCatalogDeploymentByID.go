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

	// Define the deployment configuration
	deployment := &jamfpro.ResourceJamfAppCatalogDeployment{
		Name:                            "010 Editor",
		Enabled:                         BoolPtr(true),
		AppTitleId:                      "518",
		DeploymentType:                  "SELF_SERVICE",
		UpdateBehavior:                  "AUTOMATIC",
		CategoryId:                      "-1",
		SiteId:                          "-1",
		SmartGroupId:                    "1",
		InstallPredefinedConfigProfiles: BoolPtr(false),
		TitleAvailableInAis:             BoolPtr(true),
		TriggerAdminNotifications:       BoolPtr(false),
		NotificationSettings: jamfpro.JamfAppCatalogDeploymentSubsetNotificationSettings{
			NotificationMessage:  "thing",
			NotificationInterval: 1,
			DeadlineMessage:      "thing",
			Deadline:             1,
			QuitDelay:            1,
			CompleteMessage:      "thing",
			Relaunch:             BoolPtr(true),
		},
		SelfServiceSettings: jamfpro.JamfAppCatalogDeploymentSubsetSelfServiceSettings{
			Description:                 "hello I am a description",
			ForceViewDescription:        BoolPtr(true),
			IncludeInFeaturedCategory:   BoolPtr(true),
			IncludeInComplianceCategory: BoolPtr(true),
			Categories: []jamfpro.JamfAppCatalogDeploymentSubsetCategory{
				{
					ID:       "5",
					Featured: BoolPtr(false),
				},
			},
		},
		LatestAvailableVersion: "14.0.1",
		VersionRemoved:         BoolPtr(false),
	}

	resourceID := "6"

	// Call UpdateJamfAppCatalogAppInstallerDeploymentByID function
	updateResource, err := client.UpdateJamfAppCatalogAppInstallerDeploymentByID(resourceID, deployment)
	if err != nil {
		log.Fatalf("Error updating jamf app catalog deployment: %v", err)
	}

	// Pretty print the created deployment details in JSON
	updateResourceJSON, err := json.MarshalIndent(updateResource, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created jamf app catalog deployment data: %v", err)
	}
	fmt.Println("updated jamf app catalog deployment Details:\n", string(updateResourceJSON))
}

// BoolPtr returns a pointer to the bool value passed in.
func BoolPtr(b bool) *bool {
	return &b
}
