package main

import (
	"encoding/json"
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

	// Define the new enrollment customization
	newCustomization := jamfpro.ResourceEnrollmentCustomization{
		SiteID:      "-1", // Default site ID
		DisplayName: "Custom Enrollment Experience",
		Description: "Customized enrollment experience for our organization",
		BrandingSettings: jamfpro.EnrollmentCustomizationSubsetBrandingSettings{
			TextColor:       "000000", // ensure that there's no # at the start of the hex code
			ButtonColor:     "007AFF",
			ButtonTextColor: "FFFFFF",
			BackgroundColor: "FFFFFF",
			IconUrl:         "https://lbgsandbox.jamfcloud.com/api/v2/enrollment-customizations/images/4", // Replace with your icon URL
		},
	}

	// Create the enrollment customization using the client
	response, err := client.CreateEnrollmentCustomization(newCustomization)
	if err != nil {
		log.Fatalf("Failed to create enrollment customization: %v", err)
	}

	// Pretty print the created enrollment customization details in JSON
	JSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created enrollment customization data: %v", err)
	}
	fmt.Println("Created Enrollment Customization Details:\n", string(JSON))
}
