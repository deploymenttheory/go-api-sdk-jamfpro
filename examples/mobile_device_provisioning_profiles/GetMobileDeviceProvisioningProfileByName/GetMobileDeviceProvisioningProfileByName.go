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

	profileName := "profile-name" // Replace with the actual profile Name

	profile, err := client.GetMobileDeviceProvisioningProfileByName(profileName)
	if err != nil {
		log.Fatalf("Error fetching mobile device provisioning profile by profileName: %v", err)
	}

	profileXML, err := xml.MarshalIndent(profile, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling profile data: %v", err)
	}
	fmt.Printf("Mobile Device Provisioning Profile Details:\n%s\n", string(profileXML))
}
