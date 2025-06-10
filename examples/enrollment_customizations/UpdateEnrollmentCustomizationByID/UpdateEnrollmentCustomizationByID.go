package main

import (
	"encoding/json"
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

	// Define the ID of the enrollment customization to update
	customizationID := "2"

	// Create updated customization object
	updatedCustomization := jamfpro.ResourceEnrollmentCustomization{
		SiteID:      "-1",
		DisplayName: "Updated Test Customization",
		Description: "Updated test description",
		BrandingSettings: jamfpro.EnrollmentCustomizationSubsetBrandingSettings{
			TextColor:       "000000",
			ButtonColor:     "0066CC",
			ButtonTextColor: "FFFFFF",
			BackgroundColor: "F5F5F5",
			IconUrl:         "https://example.com/icon.png",
		},
	}

	// Update the enrollment customization
	response, err := client.UpdateEnrollmentCustomizationByID(customizationID, updatedCustomization)
	if err != nil {
		log.Fatalf("Failed to update enrollment customization: %v", err)
	}

	// Pretty print the updated enrollment customization details in JSON
	JSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated enrollment customization data: %v", err)
	}
	fmt.Println("Updated Enrollment Customization Details:\n", string(JSON))
}
