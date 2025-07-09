package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/ecanault/.go/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create an enrollment configuration with all available options
	enrollmentConfig := &jamfpro.ResourceEnrollment{
		// Basic Configuration
		InstallSingleProfile:     false,
		SigningMdmProfileEnabled: false,

		// MDM Signing Certificate Configuration
		MdmSigningCertificate: &jamfpro.ResourceEnrollmentCertificate{
			Filename:         "null",
			KeystorePassword: "thing",
			IdentityKeystore: "WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09",
		},

		// Enrollment Restrictions and Cleanup
		RestrictReenrollment:            false,

		//
		// Sunsetting Re-enrollment options from this resource
		// Use jamfpro_reenrollment resource instead
		//
		// FlushLocationInformation:        false,
		// FlushLocationHistoryInformation: false,
		// FlushPolicyHistory:              false,
		// FlushExtensionAttributes:        false,
		// FlushSoftwareUpdatePlans:        false,
		// FlushMdmCommandsOnReenroll:      "DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED", // Options: DELETE_NOTHING/DELETE_ERRORS/DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED/DELETE_EVERYTHING

		// macOS Management Configuration
		MacOsEnterpriseEnrollmentEnabled: false,
		ManagementUsername:               "radmin",
		CreateManagementAccount:          true,
		HideManagementAccount:            false,
		AllowSshOnlyManagementAccount:    false,
		EnsureSshRunning:                 true,
		LaunchSelfService:                false,
		SignQuickAdd:                     false,

		// Developer Certificate Configuration
		DeveloperCertificateIdentity: &jamfpro.ResourceEnrollmentCertificate{
			Filename:         "null",
			KeystorePassword: "",
			IdentityKeystore: "",
		},

		// Certificate Details
		DeveloperCertificateIdentityDetails: jamfpro.ResourceCertificateDetails{
			Subject:      "",
			SerialNumber: "",
		},
		MdmSigningCertificateDetails: jamfpro.ResourceCertificateDetails{
			Subject:      "",
			SerialNumber: "",
		},

		// iOS/Mobile Device Enrollment Configuration
		IosEnterpriseEnrollmentEnabled: true,
		IosPersonalEnrollmentEnabled:   false,
		PersonalDeviceEnrollmentType:   "PERSONALDEVICEPROFILES", // Options: PERSONALDEVICEPROFILES/USERENROLLMENT

		// Account Driven Enrollment Configuration
		AccountDrivenUserEnrollmentEnabled:           false,
		AccountDrivenDeviceIosEnrollmentEnabled:      false,
		AccountDrivenDeviceMacosEnrollmentEnabled:    false,
		AccountDrivenUserVisionosEnrollmentEnabled:   false,
		AccountDrivenDeviceVisionosEnrollmentEnabled: false,
	}

	updatedEnrollment, err := client.UpdateEnrollment(enrollmentConfig)
	if err != nil {
		log.Fatalf("Error updating enrollment configuration: %v", err)
	}

	// Pretty print the updated enrollment configuration in JSON
	JSON, err := json.MarshalIndent(updatedEnrollment, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling enrollment data: %v", err)
	}
	fmt.Println("Updated Enrollment Configuration:\n", string(JSON))
}
