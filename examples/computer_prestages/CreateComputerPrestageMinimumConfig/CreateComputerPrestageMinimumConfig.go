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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Define the payload for creating a new computer prestage
	// Manually create a ResourceComputerPrestage struct with mapped values
	prestage := jamfpro.ResourceComputerPrestage{
		DisplayName:                       "jamfpro-sdk-example-computerPrestage-config",
		Mandatory:                         true,
		MDMRemovable:                      false,
		SupportPhoneNumber:                "",
		SupportEmailAddress:               "",
		Department:                        "",
		DefaultPrestage:                   false,
		EnrollmentSiteId:                  "-1",
		KeepExistingSiteMembership:        false,
		KeepExistingLocationInformation:   false,
		RequireAuthentication:             false,
		AuthenticationPrompt:              "",
		PreventActivationLock:             true,
		EnableDeviceBasedActivationLock:   false,
		DeviceEnrollmentProgramInstanceId: "1",
		SkipSetupItems: jamfpro.ComputerPrestageSubsetSkipSetupItems{
			Accessibility:     true,
			Appearance:        true,
			AppleID:           true,
			Biometric:         true,
			Diagnostics:       true,
			DisplayTone:       true,
			FileVault:         true,
			Location:          true,
			Payment:           true,
			Privacy:           true,
			Registration:      true,
			Restore:           true,
			ScreenTime:        true,
			Siri:              true,
			TOS:               true,
			TermsOfAddress:    true,
			ICloudDiagnostics: true,
			ICloudStorage:     true,
		},
		LocationInformation: jamfpro.ComputerPrestageSubsetLocationInformation{
			Username:     "",
			Realname:     "",
			Phone:        "",
			Email:        "",
			Room:         "",
			Position:     "",
			DepartmentId: "-1",
			BuildingId:   "-1",
			ID:           "1",
			VersionLock:  0,
		},
		PurchasingInformation: jamfpro.ComputerPrestageSubsetPurchasingInformation{
			ID:                "1",
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
		AnchorCertificates:               []string{},
		EnrollmentCustomizationId:        "0",
		Language:                         "en",
		Region:                           "GB",
		AutoAdvanceSetup:                 true,
		InstallProfilesDuringSetup:       true,
		PrestageInstalledProfileIds:      []string{},
		CustomPackageIds:                 []string{},
		CustomPackageDistributionPointId: "-1",
		EnableRecoveryLock:               false,
		RecoveryLockPasswordType:         "MANUAL",
		RotateRecoveryLockPassword:       false,
		ID:                               "1",
		ProfileUuid:                      "C101330EE870D6082D5D08FA013ADE51",
		SiteId:                           "-1",
		VersionLock:                      3,
		AccountSettings: jamfpro.ComputerPrestageSubsetAccountSettings{
			ID:                                      "1",
			PayloadConfigured:                       true,
			LocalAdminAccountEnabled:                true,
			AdminUsername:                           "localAdmin",
			HiddenAdminAccount:                      true,
			LocalUserManaged:                        false,
			UserAccountType:                         "ADMINISTRATOR",
			VersionLock:                             1,
			PrefillPrimaryAccountInfoFeatureEnabled: false,
			PrefillType:                             "CUSTOM",
			PrefillAccountFullName:                  "",
			PrefillAccountUserName:                  "",
			PreventPrefillInfoFromModification:      false,
		},
	}

	// Call the CreateComputerPrestage function
	createdPrestage, err := client.CreateComputerPrestage(&prestage)
	if err != nil {
		log.Fatalf("Error creating computer prestage: %v", err)
	}

	// Pretty print the computer prestage in JSON
	prestageJSON, err := json.MarshalIndent(createdPrestage, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling computer prestage data: %v", err)
	}
	fmt.Println("Created computer prestage:\n", string(prestageJSON))
}
