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
			Biometric:         jamfpro.FalsePtr(),
			TermsOfAddress:    jamfpro.FalsePtr(),
			FileVault:         jamfpro.FalsePtr(),
			ICloudDiagnostics: jamfpro.FalsePtr(),
			Diagnostics:       jamfpro.FalsePtr(),
			Accessibility:     jamfpro.FalsePtr(),
			AppleID:           jamfpro.FalsePtr(),
			ScreenTime:        jamfpro.FalsePtr(),
			Siri:              jamfpro.FalsePtr(),
			DisplayTone:       jamfpro.FalsePtr(),
			Restore:           jamfpro.FalsePtr(),
			Appearance:        jamfpro.FalsePtr(),
			Privacy:           jamfpro.FalsePtr(),
			Payment:           jamfpro.FalsePtr(),
			Registration:      jamfpro.FalsePtr(),
			TOS:               jamfpro.FalsePtr(),
			ICloudStorage:     jamfpro.FalsePtr(),
			Location:          jamfpro.FalsePtr(),
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
			VersionLock:       0, // Not required for creates
		},
		EnrollmentCustomizationId:        "0",
		AutoAdvanceSetup:                 jamfpro.FalsePtr(),
		InstallProfilesDuringSetup:       jamfpro.TruePtr(),
		PrestageInstalledProfileIds:      []string{},
		CustomPackageIds:                 []string{},
		CustomPackageDistributionPointId: "-1",
		EnableRecoveryLock:               jamfpro.FalsePtr(),
		RecoveryLockPasswordType:         "MANUAL",
		RotateRecoveryLockPassword:       jamfpro.FalsePtr(),
		SiteId:                           "-1",
		VersionLock:                      0, // Not required for creates
		AccountSettings: jamfpro.ComputerPrestageSubsetAccountSettings{
			PayloadConfigured:                       jamfpro.TruePtr(),
			LocalAdminAccountEnabled:                jamfpro.FalsePtr(),
			AdminUsername:                           "",
			AdminPassword:                           "",
			HiddenAdminAccount:                      jamfpro.FalsePtr(),
			LocalUserManaged:                        jamfpro.FalsePtr(),
			UserAccountType:                         "ADMINISTRATOR",
			VersionLock:                             0, // Not required for creates
			PrefillPrimaryAccountInfoFeatureEnabled: jamfpro.FalsePtr(),
			PrefillType:                             "UNKNOWN",
			PrefillAccountFullName:                  "",
			PrefillAccountUserName:                  "",
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
