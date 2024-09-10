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
		Mandatory:                         jamfpro.TruePtr(),
		MDMRemovable:                      jamfpro.TruePtr(),
		SupportPhoneNumber:                "111-222-3333",
		SupportEmailAddress:               "email@company.com",
		Department:                        "department name",
		DefaultPrestage:                   jamfpro.FalsePtr(),
		EnrollmentSiteId:                  "-1",
		KeepExistingSiteMembership:        jamfpro.FalsePtr(),
		KeepExistingLocationInformation:   jamfpro.FalsePtr(),
		RequireAuthentication:             jamfpro.FalsePtr(),
		AuthenticationPrompt:              "hello welcome to your enterprise managed macOS device",
		PreventActivationLock:             jamfpro.FalsePtr(),
		EnableDeviceBasedActivationLock:   jamfpro.FalsePtr(),
		DeviceEnrollmentProgramInstanceId: "1",
		SkipSetupItems: jamfpro.ComputerPrestageSubsetSkipSetupItems{
			// Selected items are not displayed in the Setup Assistant during enrollment
			Biometric:          jamfpro.TruePtr(),
			TermsOfAddress:     jamfpro.TruePtr(),
			FileVault:          jamfpro.TruePtr(),
			ICloudDiagnostics:  jamfpro.TruePtr(),
			Diagnostics:        jamfpro.TruePtr(),
			Accessibility:      jamfpro.TruePtr(),
			AppleID:            jamfpro.TruePtr(),
			ScreenTime:         jamfpro.TruePtr(),
			Siri:               jamfpro.TruePtr(),
			DisplayTone:        jamfpro.FalsePtr(), // Deprecated
			Restore:            jamfpro.TruePtr(),
			Appearance:         jamfpro.TruePtr(),
			Privacy:            jamfpro.TruePtr(),
			Payment:            jamfpro.TruePtr(),
			Registration:       jamfpro.TruePtr(),
			TOS:                jamfpro.TruePtr(),
			ICloudStorage:      jamfpro.TruePtr(),
			Location:           jamfpro.FalsePtr(),
			Intelligence:       jamfpro.TruePtr(),
			EnableLockdownMode: jamfpro.TruePtr(),
			Welcome:            jamfpro.TruePtr(),
			Wallpaper:          jamfpro.TruePtr(),
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
			Leased:            jamfpro.FalsePtr(),
			Purchased:         jamfpro.TruePtr(),
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
		AutoAdvanceSetup:                   jamfpro.FalsePtr(),
		InstallProfilesDuringSetup:         jamfpro.TruePtr(),
		PrestageInstalledProfileIds:        []string{},
		CustomPackageIds:                   []string{},
		CustomPackageDistributionPointId:   "-1",
		EnableRecoveryLock:                 jamfpro.FalsePtr(),
		RecoveryLockPasswordType:           "MANUAL",
		RotateRecoveryLockPassword:         jamfpro.FalsePtr(),
		PrestageMinimumOsTargetVersionType: "MINIMUM_OS_SPECIFIC_VERSION", // NO_ENFORCEMENT / MINIMUM_OS_LATEST_VERSION / MINIMUM_OS_LATEST_MAJOR_VERSION / MINIMUM_OS_LATEST_MINOR_VERSION / MINIMUM_OS_SPECIFIC_VERSION
		MinimumOsSpecificVersion:           "14.6",                        // Required if PrestageMinimumOsTargetVersionType is MINIMUM_OS_SPECIFIC_VERSION
		//ProfileUuid:                        "0386E7C8D455A040106850A8A2033968", // Automated Device Enrollment instance to associate with the PreStage enrollment. Devices associated with the selected Automated Device Enrollment instance can be assigned the PreStage enrollment
		SiteId:      "-1",
		VersionLock: 0,
		AccountSettings: jamfpro.ComputerPrestageSubsetAccountSettings{
			PayloadConfigured:                       jamfpro.TruePtr(),
			LocalAdminAccountEnabled:                jamfpro.TruePtr(),
			AdminUsername:                           "testadmin",
			AdminPassword:                           "testpassword",
			HiddenAdminAccount:                      jamfpro.TruePtr(),
			LocalUserManaged:                        jamfpro.FalsePtr(),
			UserAccountType:                         "ADMINISTRATOR",
			VersionLock:                             0,
			PrefillPrimaryAccountInfoFeatureEnabled: jamfpro.TruePtr(),
			PrefillType:                             "UNKNOWN", // UNKNOWN / DEVICE_OWNER / CUSTOM
			PrefillAccountFullName:                  "",        // Required if PrefillType is CUSTOM
			PrefillAccountUserName:                  "",        // Required if PrefillType is CUSTOM
			PreventPrefillInfoFromModification:      jamfpro.FalsePtr(),
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
