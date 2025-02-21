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

	imagePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/enrollment_customizations/UploadEnrollmentCustomizationsImage/self_service.png"

	// Upload the image file
	imageResponse, err := client.UploadEnrollmentCustomizationsImage(imagePath)
	if err != nil {
		log.Fatalf("Error uploading icon: %v", err)
	}

	// Pretty print the uploaded image details
	imageJSON, err := json.MarshalIndent(imageResponse, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling image response data: %v", err)
	}
	fmt.Println("Uploaded Image Details:\n", string(imageJSON))

	// Define the new enrollment customization with the uploaded image URL
	newCustomization := jamfpro.ResourceEnrollmentCustomization{
		SiteID:      "-1", // Default site ID
		DisplayName: "Custom Device Enrollment Experience LDAP",
		Description: "Customized enrollment experience for our organization",
		BrandingSettings: jamfpro.EnrollmentCustomizationSubsetBrandingSettings{
			TextColor:       "000000", // ensure that there's no # at the start of the hex code
			ButtonColor:     "007AFF",
			ButtonTextColor: "FFFFFF",
			BackgroundColor: "FFFFFF",
			IconUrl:         imageResponse.Url, // Use the URL from the uploaded image
		},
	}

	// Create the enrollment customization using the client
	response, err := client.CreateEnrollmentCustomization(newCustomization)
	if err != nil {
		log.Fatalf("Failed to create enrollment customization: %v", err)
	}

	// Pretty print the created enrollment customization details in JSON
	customizationJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created enrollment customization data: %v", err)
	}
	fmt.Println("Created Enrollment Customization Details:\n", string(customizationJSON))

	// Get the ID of the newly created customization
	customizationID := response.Id
	fmt.Printf("Created customization with ID: %s\n", customizationID)

	// Create LDAP group access configurations
	ldapGroups := []jamfpro.EnrollmentCustomizationLDAPGroupAccess{
		{
			GroupName:    "Test",
			LDAPServerID: 1006,
		},
		{
			GroupName:    "Test Team",
			LDAPServerID: 1006,
		},
		{
			GroupName:    "Test Operational & Support",
			LDAPServerID: 1006,
		},
		{
			GroupName:    "Test Team (Dunkirk JML)",
			LDAPServerID: 1006,
		},
	}

	// Define a new LDAP prestage pane
	newLDAPPane := jamfpro.ResourceEnrollmentCustomizationLDAPPane{
		Type:               "ldap",
		DisplayName:        "LDAP Authentication",
		Rank:               0,
		Title:              "Authenticate with your credentials",
		UsernameLabel:      "Username",
		PasswordLabel:      "Password",
		BackButtonText:     "Back",
		ContinueButtonText: "Continue",
		LDAPGroupAccess:    ldapGroups,
	}

	// Create the LDAP prestage pane
	ldapPaneResponse, err := client.CreateLDAPPrestagePane(customizationID, newLDAPPane)
	if err != nil {
		log.Fatalf("Failed to create LDAP prestage pane: %v", err)
	}

	// Pretty print the created LDAP pane details
	ldapPaneJSON, err := json.MarshalIndent(ldapPaneResponse, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created LDAP pane data: %v", err)
	}
	fmt.Println("Created LDAP Prestage Pane Details:\n", string(ldapPaneJSON))

	fmt.Println("Enrollment customization with LDAP pane created successfully!")
}
