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

	// Create enrollment message configuration with all available fields
	enrollmentMessage := &jamfpro.ResourceEnrollmentLanguage{
		LanguageCode:     "en",
		Name:             "English",
		Title:            "Enroll Your Device",
		LoginDescription: "Log in to enroll your device.",
		Username:         "Username",
		Password:         "Password",
		LoginButton:      "Log in",

		// Device Class Settings
		DeviceClassDescription:           "Specify if this device is institutionally owned or personally owned.",
		DeviceClassPersonal:              "Personally Owned",
		DeviceClassPersonalDescription:   "For personally owned devices, IT administrators **can**:\n\n*   Lock the device\n*   Apply institutional settings\n*   Install and remove institutional data\n*   Install and remove institutional apps\n\n\nFor personally owned devices, IT administrators **cannot**:\n\n*   Wipe all data and settings from your device\n*   Track the location of your device\n*   Remove anything they did not install\n*   Add/remove configuration profiles\n*   Add/remove provisioning profiles",
		DeviceClassEnterprise:            "Institutionally Owned",
		DeviceClassEnterpriseDescription: "For institutionally owned devices, IT administrators **can**:\n\n*   Wipe all data and settings from the device\n*   Lock the device\n*   Remove the passcode\n*   Apply institutional settings\n*   Install and remove institutional data\n*   Install and remove institutional apps\n*   Add/remove configuration profiles\n*   Add/remove provisioning profiles\n\nFor institutionally owned devices, IT administrators **cannot**:\n\n*   Remove anything they did not install\n*   Track the location of the device",
		DeviceClassButton:                "Enroll",

		// EULA Settings
		PersonalEula:   "",
		EnterpriseEula: "",
		EulaButton:     "Accept",

		// Site Settings
		SiteDescription: "Select the site to use for enrolling this computer or mobile device.",

		// Certificate Settings
		CertificateText:               "To continue with enrollment, you need to install the CA certificate for your organization.",
		CertificateButton:             "Continue",
		CertificateProfileName:        "CA Certificate",
		CertificateProfileDescription: "CA Certificate for mobile device management",

		// Personal Device Settings
		PersonalText:               "To continue with enrollment, you need to install the MDM profile.",
		PersonalButton:             "Continue",
		PersonalProfileName:        "MDM Profile",
		PersonalProfileDescription: "MDM Profile for mobile device management",

		// User Enrollment Settings
		UserEnrollmentText:               "Enter your Managed Apple Id to install the MDM Profile.",
		UserEnrollmentButton:             "Continue",
		UserEnrollmentProfileName:        "MDM Profile",
		UserEnrollmentProfileDescription: "MDM Profile for mobile device management",

		// Enterprise Settings
		EnterpriseText:               "To continue with enrollment, you need to install the MDM profile for your organization.",
		EnterpriseButton:             "Continue",
		EnterpriseProfileName:        "MDM Profile",
		EnterpriseProfileDescription: "MDM Profile for mobile device management",
		EnterprisePending:            "To continue with enrollment, install the CA Certificate and MDM Profile that were downloaded to your computer.",

		// QuickAdd Settings
		QuickAddText:    "Download and install this package.",
		QuickAddButton:  "Download",
		QuickAddName:    "QuickAdd.pkg",
		QuickAddPending: "Install the downloaded QuickAdd.pkg.",

		// Status Messages
		CompleteMessage:        "The enrollment process is complete.",
		FailedMessage:          "The enrollment process could not be completed. Contact your IT administrator.",
		TryAgainButton:         "Try Again",
		CheckNowButton:         "Proceed",
		CheckEnrollmentMessage: "Tap \"Proceed\" to view the enrollment status for this device.",
		LogoutButton:           "Log Out",
	}

	// Specify the language ID to update
	languageId := "en"

	// Call UpdateEnrollmentMessageByLanguageID function
	updatedMessage, err := client.UpdateEnrollmentMessageByLanguageID(languageId, enrollmentMessage)
	if err != nil {
		log.Fatalf("Error updating enrollment language messaging: %v", err)
	}

	// Pretty print the updated configuration in JSON
	JSON, err := json.MarshalIndent(updatedMessage, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated language messaging data: %v", err)
	}
	fmt.Println("Updated Enrollment Language Messaging Configuration:\n", string(JSON))
}
