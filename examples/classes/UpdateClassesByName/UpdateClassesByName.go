package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the new class details
	updatedClass := &jamfpro.ResourceClass{
		Source:      "N/A",
		Name:        "Math 101",
		Description: "Introduction to basic mathematics",
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		MobileDeviceGroup: jamfpro.ClassSubsetMobileDeviceGroup{
			ID:   3,
			Name: "All Managed iPod touches",
		},
		MobileDeviceGroupID: []jamfpro.ClassSubsetMobileDeviceGroupID{
			{
				ID: 3,
			},
		},
		TeacherIDs: []jamfpro.ClassSubsetTeacherIDs{
			{ID: 1},
			{ID: 2},
		},
		MeetingTimes: jamfpro.ClassContainerMeetingTimes{
			MeetingTime: jamfpro.ClassSubsetMeetingTime{
				Days:      "M W F",
				StartTime: 1300,
				EndTime:   1345,
			},
		},
		// Ensure other fields are aligned with the ResourceClass struct definition
	}

	className := "name-of-class"

	// Call the update function with the class ID and the new details.
	err = client.UpdateClassByName(className, updatedClass)
	if err != nil {
		log.Fatalf("Error updating class: %s\n", err)
	}

	// Print the XML structure of the created class for verification
	classXML, _ := xml.MarshalIndent(updatedClass, "", "  ")
	fmt.Printf("Created Class:\n%s\n", classXML)
}
