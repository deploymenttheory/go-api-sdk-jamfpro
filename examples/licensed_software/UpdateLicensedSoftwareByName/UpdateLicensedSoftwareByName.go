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
	updatedLicensedSoftware, err := client.UpdateLicensedSoftwareByName("Adobe Creative Suite", licensedSoftware)
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
