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

	profileID := "1" // Replace with the actual profile ID

	profile, err := client.GetMobileDeviceProvisioningProfileByID(profileID)
	if err != nil {
		log.Fatalf("Error fetching mobile device provisioning profile by profileID: %v", err)
	}

	profileXML, err := xml.MarshalIndent(profile, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling profile data: %v", err)
	}
	fmt.Printf("Mobile Device Provisioning Profile Details:\n%s\n", string(profileXML))
}
