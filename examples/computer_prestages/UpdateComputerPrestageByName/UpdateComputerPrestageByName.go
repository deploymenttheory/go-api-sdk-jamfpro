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

	// The name of the computer prestage you want to update
	prestageName := "YOUR_PRESTAGE_NAME_HERE"

	// First, get the current prestage to obtain the current version lock
	currentPrestage, err := client.GetComputerPrestageByName(prestageName)
	if err != nil {
		log.Fatalf("Error fetching current computer prestage: %v", err)
	}

	// Prepare the update using known good settings
	update := &jamfpro.ResourceComputerPrestage{
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
		EnrollmentCustomizationId:         "0",
		AutoAdvanceSetup:                  jamfpro.FalsePtr(),
		InstallProfilesDuringSetup:        jamfpro.TruePtr(),
		PrestageInstalledProfileIds:       []string{},
		CustomPackageIds:                  []string{},
		CustomPackageDistributionPointId:  "-1",
		EnableRecoveryLock:                jamfpro.FalsePtr(),
		RecoveryLockPasswordType:          "MANUAL",
		RotateRecoveryLockPassword:        jamfpro.FalsePtr(),
		ProfileUuid:                       "0386E7C8D455A040106850A8A2033968",
		SiteId:                            "-1",
		VersionLock:                       currentPrestage.VersionLock,
	}

	// Update SkipSetupItems
	update.SkipSetupItems = jamfpro.ComputerPrestageSubsetSkipSetupItems{
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
	}

	// Update LocationInformation
	update.LocationInformation = jamfpro.ComputerPrestageSubsetLocationInformation{
		ID:           jamfpro.IncrementStringID(currentPrestage.LocationInformation.ID),
		Username:     "dafydd",
		Realname:     "",
		Phone:        "",
		Email:        "",
		Room:         "",
		Position:     "",
		DepartmentId: "-1",
		BuildingId:   "-1",
		VersionLock:  currentPrestage.LocationInformation.VersionLock + 1,
	}

	// Update PurchasingInformation
	update.PurchasingInformation = jamfpro.ComputerPrestageSubsetPurchasingInformation{
		ID:                jamfpro.IncrementStringID(currentPrestage.PurchasingInformation.ID),
		Leased:            jamfpro.FalsePtr(),
		Purchased:         jamfpro.TruePtr(),
		AppleCareId:       "",
		PONumber:          "PO123456",
		Vendor:            "",
		PurchasePrice:     "",
		LifeExpectancy:    0,
		PurchasingAccount: "",
		PurchasingContact: "",
		LeaseDate:         "2024-01-01",
		PODate:            "2024-01-01",
		WarrantyDate:      "2024-01-01",
		VersionLock:       currentPrestage.PurchasingInformation.VersionLock + 1,
	}

	// Update AccountSettings
	update.AccountSettings = jamfpro.ComputerPrestageSubsetAccountSettings{
		ID:                                      jamfpro.IncrementStringID(currentPrestage.AccountSettings.ID),
		PayloadConfigured:                       jamfpro.TruePtr(),
		LocalAdminAccountEnabled:                jamfpro.TruePtr(),
		AdminUsername:                           "testadmin",
		AdminPassword:                           "testpassword",
		HiddenAdminAccount:                      jamfpro.TruePtr(),
		LocalUserManaged:                        jamfpro.FalsePtr(),
		UserAccountType:                         "ADMINISTRATOR",
		VersionLock:                             currentPrestage.AccountSettings.VersionLock + 1,
		PrefillPrimaryAccountInfoFeatureEnabled: jamfpro.TruePtr(),
		PrefillType:                             "CUSTOM",
		PrefillAccountFullName:                  "Firstname.Surname",
		PrefillAccountUserName:                  "Firstname.Surname",
		PreventPrefillInfoFromModification:      jamfpro.FalsePtr(),
	}

	// Call UpdateComputerPrestageByName to update the prestage
	updatedPrestage, err := client.UpdateComputerPrestageByName(prestageName, update)
	if err != nil {
		log.Fatalf("Error updating computer prestage: %v", err)
	}

	// Pretty print the updated computer prestage in JSON
	updatedPrestageJSON, err := json.MarshalIndent(updatedPrestage, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated prestage data: %v", err)
	}
	fmt.Println("Updated computer prestage:")
	fmt.Println(string(updatedPrestageJSON))
}
