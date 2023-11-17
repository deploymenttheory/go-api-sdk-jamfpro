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
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the licensed software details
	licensedSoftware := &jamfpro.ResponseLicensedSoftware{
		General: jamfpro.LicensedSoftwareGeneral{
			Name:                               "Adobe Creative Suite",
			Publisher:                          "Adobe Systems Incorporated",
			Platform:                           "Mac",
			SendEmailOnViolation:               true,
			RemoveTitlesFromInventoryReports:   false,
			ExcludeTitlesPurchasedFromAppStore: false,
			Notes:                              "string",
			Site: jamfpro.LicensedSoftwareSite{
				ID:   -1,
				Name: "None",
			},
		},
		SoftwareDefinitions: []jamfpro.SoftwareDefinition{
			{
				CompareType: "like",
				Name:        "string",
				Version:     14,
			},
		},
		FontDefinitions: []jamfpro.LicensedSoftwareFontDefinition{
			{
				CompareType: "like",
				Name:        "string",
				Version:     14,
			},
		},
		PluginDefinitions: []jamfpro.LicensedSoftwarePluginDefinition{
			{
				CompareType: "like",
				Name:        "string",
				Version:     14,
			},
		},
		Licenses: []jamfpro.LicensedSoftwareLicense{
			{
				Size: 1,
				License: jamfpro.LicenseDetail{
					SerialNumber1:    "string",
					SerialNumber2:    "string",
					OrganizationName: "string",
					RegisteredTo:     "string",
					LicenseType:      "Standard",
					LicenseCount:     500,
					Notes:            "string",
					Purchasing: jamfpro.PurchasingDetail{
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
					Attachments: []jamfpro.LicensedSoftwareAttachment{
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

	// Create the licensed software in Jamf Pro
	createdSoftware, err := client.CreateLicensedSoftware(licensedSoftware)
	if err != nil {
		log.Fatalf("Error creating licensed software: %v", err)
	}

	// Pretty print the created software details
	createdSoftwareXML, err := xml.MarshalIndent(createdSoftware, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created software data: %v", err)
	}
	fmt.Println("Created Licensed Software:", string(createdSoftwareXML))
}
