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

	// Define the payload for creating a new computer prestage
	prestage := jamfpro.ResourceComputerPrestage{
		DisplayName:                       "jamfpro-sdk-example-computerPrestageFull-config",
		Mandatory:                         true,
		MDMRemovable:                      true,
		SupportPhoneNumber:                "111-222-3333",
		SupportEmailAddress:               "email@company.com",
		Department:                        "department name",
		DefaultPrestage:                   false,
		EnrollmentSiteId:                  "-1",
		KeepExistingSiteMembership:        false,
		KeepExistingLocationInformation:   false,
		RequireAuthentication:             false,
		AuthenticationPrompt:              "hello welcome to your enterprise managed macOS device",
		PreventActivationLock:             false,
		EnableDeviceBasedActivationLock:   false,
		DeviceEnrollmentProgramInstanceId: "1",
		SkipSetupItems: jamfpro.ComputerPrestageSubsetSkipSetupItems{
			// Selected items are not displayed in the Setup Assistant during enrollment
			Biometric:          true,
			TermsOfAddress:     true,
			FileVault:          true,
			ICloudDiagnostics:  true,
			Diagnostics:        true,
			Accessibility:      true,
			AppleID:            true,
			ScreenTime:         true,
			Siri:               true,
			DisplayTone:        false, // Deprecated
			Restore:            true,
			Appearance:         true,
			Privacy:            true,
			Payment:            true,
			Registration:       true,
			TOS:                true,
			ICloudStorage:      true,
			Location:           false,
			Intelligence:       true,
			EnableLockdownMode: true,
			Welcome:            true,
			Wallpaper:          true,
		},
		LocationInformation: jamfpro.ComputerPrestageSubsetLocationInformation{
			ID:           "-1", // Required
			Username:     "Deployment",
			Realname:     "Theory",
			Phone:        "+44-1234-567890",
			Email:        "dummy@domain.com",
			Room:         "",
			Position:     "IT",
			DepartmentId: "-1",
			BuildingId:   "-1",
			VersionLock:  0,
		},
		PurchasingInformation: jamfpro.ComputerPrestageSubsetPurchasingInformation{
			ID:                "-1", // Required
			Leased:            false,
			Purchased:         true,
			AppleCareId:       "",
			PONumber:          "",
			Vendor:            "",
			PurchasePrice:     "",
			LifeExpectancy:    0,
			PurchasingAccount: "",
			PurchasingContact: "",
			LeaseDate:         "1970-01-01",
			PODate:            "1970-01-01",
			WarrantyDate:      "1970-01-01",
			VersionLock:       0,
		},
		EnrollmentCustomizationId:          "0",
		AutoAdvanceSetup:                   false,
		InstallProfilesDuringSetup:         true,
		PrestageInstalledProfileIds:        []string{},
		CustomPackageIds:                   []string{"1", "2"},
		CustomPackageDistributionPointId:   "-1", // -2 for all cloud distribution points , -1 for not used, then id for all distribution points
		EnableRecoveryLock:                 false,
		RecoveryLockPasswordType:           "MANUAL",
		RotateRecoveryLockPassword:         false,
		PrestageMinimumOsTargetVersionType: "MINIMUM_OS_SPECIFIC_VERSION", // NO_ENFORCEMENT / MINIMUM_OS_LATEST_VERSION / MINIMUM_OS_LATEST_MAJOR_VERSION / MINIMUM_OS_LATEST_MINOR_VERSION / MINIMUM_OS_SPECIFIC_VERSION
		MinimumOsSpecificVersion:           "14.6",                        // Required if PrestageMinimumOsTargetVersionType is MINIMUM_OS_SPECIFIC_VERSION
		//ProfileUuid:                        "0386E7C8D455A040106850A8A2033968", // Automated Device Enrollment instance to associate with the PreStage enrollment. Devices associated with the selected Automated Device Enrollment instance can be assigned the PreStage enrollment
		SiteId:      "-1",
		VersionLock: 0,
		AccountSettings: jamfpro.ComputerPrestageSubsetAccountSettings{
			PayloadConfigured:                       true,
			LocalAdminAccountEnabled:                true,
			AdminUsername:                           "testadmin",
			AdminPassword:                           "testpassword",
			HiddenAdminAccount:                      true,
			LocalUserManaged:                        false,
			UserAccountType:                         "ADMINISTRATOR",
			VersionLock:                             0,
			PrefillPrimaryAccountInfoFeatureEnabled: true,
			PrefillType:                             "UNKNOWN", // UNKNOWN / DEVICE_OWNER / CUSTOM
			PrefillAccountFullName:                  "",        // Required if PrefillType is CUSTOM
			PrefillAccountUserName:                  "",        // Required if PrefillType is CUSTOM
			PreventPrefillInfoFromModification:      false,
		},
	}

	// Marshal the prestage struct to JSON and print it
	prestageJSON, err := json.MarshalIndent(prestage, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling prestage data: %v", err)
	}
	fmt.Println("Prestage configuration to be sent:")
	fmt.Println(string(prestageJSON))

	// Call the CreateComputerPrestage function
	createdPrestage, err := client.CreateComputerPrestage(&prestage)
	if err != nil {
		log.Fatalf("Error creating computer prestage: %v", err)
	}

	// Pretty print the created computer prestage in JSON
	createdPrestageJSON, err := json.MarshalIndent(createdPrestage, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created prestage data: %v", err)
	}
	fmt.Println("Created computer prestage:")
	fmt.Println(string(createdPrestageJSON))
}
