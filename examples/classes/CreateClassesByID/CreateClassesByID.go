package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	//logLevel := logger.LogLevelInfo // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the new class details
	newClass := &jamfpro.ResourceClass{
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

	// Create class
	createdClass, err := client.CreateClass(newClass)
	if err != nil {
		log.Fatalf("Error creating class: %s\n", err)
	}

	// Print the XML structure of the created class for verification
	classXML, _ := xml.MarshalIndent(createdClass, "", "  ")
	fmt.Printf("Created Class:\n%s\n", classXML)
}
