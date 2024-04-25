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

	// Fetch mobile device provisioning profiles
	profiles, err := client.GetMobileDeviceProvisioningProfiles()
	if err != nil {
		log.Fatalf("Error fetching mobile device provisioning profiles: %s\n", err)
	}

	// Print the profiles in a formatted XML
	profilesXML, err := xml.MarshalIndent(profiles, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling profiles data: %v", err)
	}
	fmt.Println("Mobile Device Provisioning Profiles:\n", string(profilesXML))
}
