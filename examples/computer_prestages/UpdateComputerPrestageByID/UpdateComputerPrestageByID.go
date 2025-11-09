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

	// The ID of the computer prestage you want to update
	prestageID := "95"

	// First, get the current prestage to obtain the current version lock
	currentPrestage, err := client.GetComputerPrestageByID(prestageID)
	if err != nil {
		log.Fatalf("Error fetching current computer prestage: %v", err)
	}

	update := &jamfpro.ResourceComputerPrestage{
		DisplayName:                        "jamfpro-sdk-example-computerPrestageFull-config",
		Mandatory:                          true,
		MDMRemovable:                       true,
		SupportPhoneNumber:                 "111-222-3333",
		SupportEmailAddress:                "email@company.com",
		Department:                         "department name",
		DefaultPrestage:                    false,
		EnrollmentSiteId:                   "-1",
		KeepExistingSiteMembership:         false,
		KeepExistingLocationInformation:    false,
		RequireAuthentication:              false,
		AuthenticationPrompt:               "hello welcome to your enterprise managed macOS device",
		PreventActivationLock:              false,
		EnableDeviceBasedActivationLock:    false,
		DeviceEnrollmentProgramInstanceId:  "1",
		EnrollmentCustomizationId:          "0",
		AutoAdvanceSetup:                   false,
		InstallProfilesDuringSetup:         true,
		PrestageInstalledProfileIds:        []string{},
		CustomPackageIds:                   []string{},
		CustomPackageDistributionPointId:   "-1",
		EnableRecoveryLock:                 false,
		RecoveryLockPasswordType:           "MANUAL",
		RotateRecoveryLockPassword:         false,
		PrestageMinimumOsTargetVersionType: "MINIMUM_OS_SPECIFIC_VERSION",
		MinimumOsSpecificVersion:           "14.6",
		SiteId:                             "-1",
		VersionLock:                        currentPrestage.VersionLock,
	}

	// Update SkipSetupItems
	update.SkipSetupItems = jamfpro.ComputerPrestageSubsetSkipSetupItems{
		Biometric:          true,
		TermsOfAddress:     true,
		FileVault:          true,
		ICloudDiagnostics:  true,
		Diagnostics:        true,
		Accessibility:      true,
		AppleID:            true,
		ScreenTime:         true,
		Siri:               true,
		DisplayTone:        false,
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
	}

	// Update LocationInformation
	update.LocationInformation = jamfpro.ComputerPrestageSubsetLocationInformation{
		ID:           "-1", // As we are updating the prestage, we need to set the ID to -1. Jamf will increment the ID for us.
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
		ID:                "-1", // As we are updating the prestage, we need to set the ID to -1. Jamf will increment the ID for us.
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
		VersionLock:       currentPrestage.PurchasingInformation.VersionLock + 1,
	}

	// Update AccountSettings
	update.AccountSettings = jamfpro.ComputerPrestageSubsetAccountSettings{
		ID:                                      "-1", // As we are updating the prestage, we need to set the ID to -1. Jamf will increment the ID for us.
		PayloadConfigured:                       true,
		LocalAdminAccountEnabled:                true,
		AdminUsername:                           "testadmin",
		AdminPassword:                           "testpassword",
		HiddenAdminAccount:                      true,
		LocalUserManaged:                        false,
		UserAccountType:                         "ADMINISTRATOR",
		VersionLock:                             0,
		PrefillPrimaryAccountInfoFeatureEnabled: true,
		PrefillType:                             "UNKNOWN",
		PrefillAccountFullName:                  "",
		PrefillAccountUserName:                  "",
		PreventPrefillInfoFromModification:      false,
	}

	// Call UpdateComputerPrestageByID to update the prestage
	updatedPrestage, err := client.UpdateComputerPrestageByID(prestageID, update)
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
