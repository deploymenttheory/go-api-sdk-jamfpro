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
	newClass := &jamfpro.ResponseClasses{
		Source:      "N/A",
		Name:        "Math 101",
		Description: "Introduction to basic mathematics",
		Site: jamfpro.ClassSite{
			ID:   -1,
			Name: "None",
		},
		MeetingTimes: jamfpro.ClassMeetingTimes{
			MeetingTime: jamfpro.ClassMeetingTime{
				Days:      "M W F",
				StartTime: 1300,
				EndTime:   1345,
			},
		},
		MobileDeviceGroup: jamfpro.ClassDeviceGroup{
			ID:   3,
			Name: "All Managed iPod touches",
		},
		MobileDeviceGroupID: []jamfpro.ClassGroupID{
			{
				ID: 3,
			},
		},
		TeacherIDs: []jamfpro.ClassTeacherID{
			{ID: 1},
			{ID: 2},
		},
		// Do not include TeacherIDs if Teachers is set
	}

	// Create class
	createdClass, err := client.CreateClassesByID(newClass)
	if err != nil {
		log.Fatalf("Error creating class: %s\n", err)
	}

	// Print the XML structure of the created class for verification
	classXML, _ := xml.MarshalIndent(createdClass, "", "  ")
	fmt.Printf("Created Class:\n%s\n", classXML)
}
