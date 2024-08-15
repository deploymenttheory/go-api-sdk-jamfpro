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
		//ID:                                "-1",
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
			ID:           "-1",
			Username:     "",
			Realname:     "",
			Phone:        "",
			Email:        "",
			Room:         "",
			Position:     "",
			DepartmentId: "-1",
			BuildingId:   "-1",
			VersionLock:  0,
		},
		PurchasingInformation: jamfpro.ComputerPrestageSubsetPurchasingInformation{
			ID:                "-1",
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
		EnrollmentCustomizationId:  "0",
		AutoAdvanceSetup:           jamfpro.FalsePtr(),
		InstallProfilesDuringSetup: jamfpro.TruePtr(),
		PrestageInstalledProfileIds: []string{
			"3847",
			"3864",
			"3806",
		},
		CustomPackageIds:                 []string{},
		CustomPackageDistributionPointId: "-1",
		EnableRecoveryLock:               jamfpro.FalsePtr(),
		RecoveryLockPasswordType:         "MANUAL",
		RotateRecoveryLockPassword:       jamfpro.FalsePtr(),
		ProfileUuid:                      "0386E7C8D455A040106850A8A2033968",
		SiteId:                           "-1",
		AccountSettings: jamfpro.ComputerPrestageSubsetAccountSettings{
			ID:                                      "-1",
			PayloadConfigured:                       jamfpro.TruePtr(),
			LocalAdminAccountEnabled:                jamfpro.TruePtr(),
			AdminUsername:                           "Administrator",
			AdminPassword:                           "Administrator",
			HiddenAdminAccount:                      jamfpro.TruePtr(),
			LocalUserManaged:                        jamfpro.FalsePtr(),
			UserAccountType:                         "ADMINISTRATOR",
			PrefillPrimaryAccountInfoFeatureEnabled: jamfpro.TruePtr(),
			PrefillType:                             "DEVICE_OWNER",
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
