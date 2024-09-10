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

	// Prepare the update using values from the create code
	update := &jamfpro.ResourceComputerPrestage{
		DisplayName:                        "jamfpro-sdk-example-computerPrestageFull-config",
		Mandatory:                          jamfpro.TruePtr(),
		MDMRemovable:                       jamfpro.TruePtr(),
		SupportPhoneNumber:                 "111-222-3333",
		SupportEmailAddress:                "email@company.com",
		Department:                         "department name",
		DefaultPrestage:                    jamfpro.FalsePtr(),
		EnrollmentSiteId:                   "-1",
		KeepExistingSiteMembership:         jamfpro.FalsePtr(),
		KeepExistingLocationInformation:    jamfpro.FalsePtr(),
		RequireAuthentication:              jamfpro.FalsePtr(),
		AuthenticationPrompt:               "hello welcome to your enterprise managed macOS device",
		PreventActivationLock:              jamfpro.FalsePtr(),
		EnableDeviceBasedActivationLock:    jamfpro.FalsePtr(),
		DeviceEnrollmentProgramInstanceId:  "1",
		EnrollmentCustomizationId:          "0",
		AutoAdvanceSetup:                   jamfpro.FalsePtr(),
		InstallProfilesDuringSetup:         jamfpro.TruePtr(),
		PrestageInstalledProfileIds:        []string{},
		CustomPackageIds:                   []string{},
		CustomPackageDistributionPointId:   "-1",
		EnableRecoveryLock:                 jamfpro.FalsePtr(),
		RecoveryLockPasswordType:           "MANUAL",
		RotateRecoveryLockPassword:         jamfpro.FalsePtr(),
		PrestageMinimumOsTargetVersionType: "MINIMUM_OS_SPECIFIC_VERSION",
		MinimumOsSpecificVersion:           "14.6",
		SiteId:                             "-1",
		VersionLock:                        currentPrestage.VersionLock,
	}

	// Update SkipSetupItems
	update.SkipSetupItems = jamfpro.ComputerPrestageSubsetSkipSetupItems{
		Biometric:          jamfpro.TruePtr(),
		TermsOfAddress:     jamfpro.TruePtr(),
		FileVault:          jamfpro.TruePtr(),
		ICloudDiagnostics:  jamfpro.TruePtr(),
		Diagnostics:        jamfpro.TruePtr(),
		Accessibility:      jamfpro.TruePtr(),
		AppleID:            jamfpro.TruePtr(),
		ScreenTime:         jamfpro.TruePtr(),
		Siri:               jamfpro.TruePtr(),
		DisplayTone:        jamfpro.FalsePtr(),
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
	}

	// Update LocationInformation
	update.LocationInformation = jamfpro.ComputerPrestageSubsetLocationInformation{
		ID:           jamfpro.IncrementStringID(currentPrestage.LocationInformation.ID),
		Username:     "Deployment",
		Realname:     "Theory",
		Phone:        "+44-1234-567890",
		Email:        "dummy@domain.com",
		Room:         "",
		Position:     "IT",
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
		PONumber:          "",
		Vendor:            "",
		PurchasePrice:     "",
		LifeExpectancy:    0,
		PurchasingAccount: "",
		PurchasingContact: "",
		LeaseDate:         "1970-01-01",
		PODate:            "1970-01-01",
		WarrantyDate:      "1970-01-01",
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
		VersionLock:                             0,
		PrefillPrimaryAccountInfoFeatureEnabled: jamfpro.TruePtr(),
		PrefillType:                             "UNKNOWN",
		PrefillAccountFullName:                  "",
		PrefillAccountUserName:                  "",
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
