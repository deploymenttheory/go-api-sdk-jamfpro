package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the updated licensed software details
	licensedSoftware := &jamfpro.ResourceLicensedSoftware{
		General: jamfpro.LicensedSoftwareSubsetGeneral{
			Name:                               "Adobe Creative Suite",
			Publisher:                          "Adobe Systems Incorporated",
			Platform:                           "Mac",
			SendEmailOnViolation:               true,
			RemoveTitlesFromInventoryReports:   false,
			ExcludeTitlesPurchasedFromAppStore: false,
			Notes:                              "string",
			Site: jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		SoftwareDefinitions: []jamfpro.LicensedSoftwareSubsetSoftwareDefinitions{
			{
				CompareType: "like",
				Name:        "string",
				Version:     14,
			},
		},
		FontDefinitions: []jamfpro.LicensedSoftwareSubsetFontDefinitions{
			{
				CompareType: "like",
				Name:        "string",
				Version:     14,
			},
		},
		PluginDefinitions: []jamfpro.LicensedSoftwareSubsetPluginDefinitions{
			{
				CompareType: "like",
				Name:        "string",
				Version:     14,
			},
		},
		Licenses: []jamfpro.LicensedSoftwareSubsetLicenses{
			{
				Size: 1,
				License: jamfpro.LicensedSoftwareSubsetLicense{
					SerialNumber1:    "string",
					SerialNumber2:    "string",
					OrganizationName: "string",
					RegisteredTo:     "string",
					LicenseType:      "Standard",
					LicenseCount:     500,
					Notes:            "string",
					Purchasing: jamfpro.LicensedSoftwareSubsetLicensePurchasing{
						IsPerpetual:       true,
						IsAnnual:          false,
						PONumber:          "string",
						Vendor:            "string",
						PurchasePrice:     "string",
						PurchasingAccount: "string",
						PODate:            "2017-07-07 18:37:04",
						LicenseExpires:    "2017-07-07 18:37:04",
						LifeExpectancy:    0,
						PurchasingContact: "string",
					},
					Attachments: []jamfpro.LicensedSoftwareSubsetLicenseAttachments{
						{
							ID:       1,
							Filename: "icon.png",
							URI:      "https://example.jamfcloud/attachment.html?id=1&amp;o=r",
						},
					},
				},
			},
		},
	}

	// Update the licensed software by ID
	updatedLicensedSoftware, err := client.UpdateLicensedSoftwareByID(1, licensedSoftware)
	if err != nil {
		log.Fatalf("Error updating licensed software by ID: %v", err)
	}

	// Pretty print the created software details
	createdSoftwareXML, err := xml.MarshalIndent(updatedLicensedSoftware, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created software data: %v", err)
	}
	fmt.Println("Created Licensed Software:", string(createdSoftwareXML))
}
