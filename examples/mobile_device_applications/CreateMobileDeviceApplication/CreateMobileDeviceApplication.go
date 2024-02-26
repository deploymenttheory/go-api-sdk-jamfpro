package main

import (
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

	// Define a new mobile device application
	newApp := &jamfpro.ResourceMobileDeviceApplication{
		General: jamfpro.MobileDeviceApplicationSubsetGeneral{
			Name:        "Jamf Self Service",
			DisplayName: "Jamf Self Service",
			Description: "Jamf Self Service empowers you to be more productive...",
			BundleID:    "com.jamfsoftware.selfservice",
			Version:     "10.10.6",
			InternalApp: true,
			OsType:      "iOS", // iOS or tvOS
			Category: jamfpro.SharedResourceCategory{
				ID:   -1,
				Name: "No category assigned",
			},
			IPA: jamfpro.MobileDeviceApplicationSubsetGeneralIPA{
				Name: "IPAName",
				URI:  "http://example.com/ipa",
				Data: "Base64EncodedString",
			},
			Icon: jamfpro.MobileDeviceApplicationSubsetIcon{
				ID:   27,
				Name: "1024x1024bb.png",
				URI:  "http://example.com/icon.png",
				Data: "Base64EncodedString",
			},
			// Populate other fields as necessary...
		},
	}

	createdApp, err := client.CreateMobileDeviceApplication(newApp)
	if err != nil {
		fmt.Println("Error creating mobile device application:", err)
	} else {
		fmt.Println("Created Mobile Device Application:", createdApp)
	}
}
