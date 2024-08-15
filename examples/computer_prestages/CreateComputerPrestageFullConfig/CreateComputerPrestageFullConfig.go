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
	// Manually create a ResourceComputerPrestage struct with mapped values
	prestage := jamfpro.ResourceComputerPrestage{
		DisplayName:                       "jamfpro-sdk-example-computerPrestageFull-config",
		Mandatory:                         jamfpro.TruePtr(),
		MDMRemovable:                      jamfpro.TruePtr(),
		SupportPhoneNumber:                "118-118",
		SupportEmailAddress:               "email@company.com",
		Department:                        "department name",
		DefaultPrestage:                   jamfpro.FalsePtr(),
		EnrollmentSiteId:                  "-1",
		KeepExistingSiteMembership:        jamfpro.TruePtr(),
		KeepExistingLocationInformation:   jamfpro.TruePtr(),
		RequireAuthentication:             jamfpro.TruePtr(),
		AuthenticationPrompt:              "hello welcome to your enterprise managed macOS device",
		PreventActivationLock:             jamfpro.FalsePtr(),
		EnableDeviceBasedActivationLock:   jamfpro.FalsePtr(),
		EnableRecoveryLock:                jamfpro.FalsePtr(),
		RecoveryLockPasswordType:          "MANUAL",
		RotateRecoveryLockPassword:        jamfpro.FalsePtr(),
		DeviceEnrollmentProgramInstanceId: "1",
		SkipSetupItems: jamfpro.ComputerPrestageSubsetSkipSetupItems{
			Accessibility:     jamfpro.TruePtr(),
			Appearance:        jamfpro.TruePtr(),
			AppleID:           jamfpro.TruePtr(),
			Biometric:         jamfpro.TruePtr(),
			Diagnostics:       jamfpro.TruePtr(),
			DisplayTone:       jamfpro.TruePtr(),
			FileVault:         jamfpro.TruePtr(),
			Location:          jamfpro.TruePtr(),
			Payment:           jamfpro.TruePtr(),
			Privacy:           jamfpro.TruePtr(),
			Registration:      jamfpro.TruePtr(),
			Restore:           jamfpro.TruePtr(),
			ScreenTime:        jamfpro.TruePtr(),
			Siri:              jamfpro.TruePtr(),
			TOS:               jamfpro.TruePtr(),
			TermsOfAddress:    jamfpro.TruePtr(),
			ICloudDiagnostics: jamfpro.TruePtr(),
			ICloudStorage:     jamfpro.TruePtr(),
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
		AnchorCertificates:         []string{},
		EnrollmentCustomizationId:  "0",
		Language:                   "en",
		Region:                     "GB",
		AutoAdvanceSetup:           jamfpro.TruePtr(),
		InstallProfilesDuringSetup: jamfpro.TruePtr(),
		PrestageInstalledProfileIds: []string{
			"287",
			"288",
			"289"},
		CustomPackageIds: []string{
			"3",
		},
		CustomPackageDistributionPointId: "-1",
		ID:                               "1",
		ProfileUuid:                      "C101330EE870D6082D5D08FA013ADE51",
		SiteId:                           "-1",
		VersionLock:                      3,
		AccountSettings: jamfpro.ComputerPrestageSubsetAccountSettings{
			ID:                                      "1",
			PayloadConfigured:                       jamfpro.TruePtr(),
			LocalAdminAccountEnabled:                jamfpro.TruePtr(),
			AdminUsername:                           "localAdmin",
			AdminPassword:                           "thingthing1010",
			HiddenAdminAccount:                      jamfpro.TruePtr(),
			LocalUserManaged:                        jamfpro.FalsePtr(),
			UserAccountType:                         "ADMINISTRATOR",
			VersionLock:                             1,
			PrefillPrimaryAccountInfoFeatureEnabled: jamfpro.FalsePtr(),
			PrefillType:                             "CUSTOM",
			PrefillAccountFullName:                  "",
			PrefillAccountUserName:                  "",
			PreventPrefillInfoFromModification:      jamfpro.FalsePtr(),
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
