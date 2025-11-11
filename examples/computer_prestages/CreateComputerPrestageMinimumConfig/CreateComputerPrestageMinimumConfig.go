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
		DisplayName:                       "jamfpro-sdk-example-computerPrestageMinimum-config",
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
			Biometric:          false,
			TermsOfAddress:     false,
			FileVault:          false,
			ICloudDiagnostics:  false,
			Diagnostics:        false,
			Accessibility:      false,
			AppleID:            false,
			ScreenTime:         false,
			Siri:               false,
			DisplayTone:        false,
			Restore:            false,
			Appearance:         false,
			Privacy:            false,
			Payment:            false,
			Registration:       false,
			TOS:                false,
			ICloudStorage:      false,
			Location:           false,
			Intelligence:       false,
			EnableLockdownMode: false,
			Welcome:            false,
			Wallpaper:          false,
		},
		LocationInformation: jamfpro.ComputerPrestageSubsetLocationInformation{
			ID:           "-1", // Required
			Username:     "",
			Realname:     "",
			Phone:        "",
			Email:        "",
			Room:         "",
			Position:     "",
			DepartmentId: "-1",
			BuildingId:   "-1",
			VersionLock:  0, // Not required for creates
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
			VersionLock:       0, // Not required for creates
		},
		EnrollmentCustomizationId:          "0",
		AutoAdvanceSetup:                   false,
		InstallProfilesDuringSetup:         true,
		PrestageInstalledProfileIds:        []string{},
		CustomPackageIds:                   []string{},
		CustomPackageDistributionPointId:   "-1", // -2 for all cloud distribution points , -1 for not used, then id for all distribution points
		EnableRecoveryLock:                 false,
		RecoveryLockPasswordType:           "",
		RecoveryLockPassword:               "",
		RotateRecoveryLockPassword:         false,
		PrestageMinimumOsTargetVersionType: "NO_ENFORCEMENT", // NO_ENFORCEMENT / MINIMUM_OS_LATEST_VERSION / MINIMUM_OS_LATEST_MAJOR_VERSION / MINIMUM_OS_LATEST_MINOR_VERSION / MINIMUM_OS_SPECIFIC_VERSION
		MinimumOsSpecificVersion:           "",
		SiteId:                             "-1",
		VersionLock:                        0, // Not required for creates
		AccountSettings: jamfpro.ComputerPrestageSubsetAccountSettings{
			PayloadConfigured:                       true,
			LocalAdminAccountEnabled:                false,
			AdminUsername:                           "",
			AdminPassword:                           "",
			HiddenAdminAccount:                      false,
			LocalUserManaged:                        false,
			UserAccountType:                         "ADMINISTRATOR",
			VersionLock:                             0, // Not required for creates
			PrefillPrimaryAccountInfoFeatureEnabled: false,
			PrefillType:                             "UNKNOWN",
			PrefillAccountFullName:                  "",
			PrefillAccountUserName:                  "",
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
