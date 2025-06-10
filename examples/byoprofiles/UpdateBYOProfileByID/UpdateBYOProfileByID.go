package main

import (
	"encoding/xml"
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

	// Create a BYOProfile structure to send
	updatedProfile := jamfpro.ResourceBYOProfile{
		General: jamfpro.BYOProfileSubsetGeneral{
			Name:        "Personal Device Profile with jamf pro sdk",
			Site:        jamfpro.SharedResourceSite{ID: -1, Name: "None"},
			Enabled:     true,
			Description: "Used for Android or iOS BYO device enrollments",
		},
	}

	profileID := "4" // Use the actual ID of the profile to be updated

	// Convert the profile to XML to see the output (optional, for debug purposes)
	xmlData, err := xml.MarshalIndent(updatedProfile, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling XML: %v", err)
	}
	fmt.Printf("XML Request: %s\n", xmlData)

	// Now call the update function
	updatedProfileResp, err := client.UpdateBYOProfileByID(profileID, &updatedProfile)
	if err != nil {
		log.Fatalf("Error updating BYO Profile by ID: %v", err)
	}
	fmt.Printf("Updated BYO Profile: %+v\n", updatedProfileResp)
}
