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
