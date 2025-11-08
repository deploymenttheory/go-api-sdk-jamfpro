package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/neilmartin/GitHub/go-api-sdk-jamfpro/client_auth.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// The ID of the mobile device prestage you want to update
	prestageID := "12"

	// First, get the current prestage to obtain the current version lock
	currentPrestage, err := client.GetMobileDevicePrestageByID(prestageID)
	if err != nil {
		log.Fatalf("Error fetching current mobile device prestage: %v", err)
	}

	update := &jamfpro.ResourceMobileDevicePrestage{
		DisplayName:                            "jamfpro-sdk-example-mobiledevicePrestage-config",
		Mandatory:                              true,
		MdmRemovable:                           false,
		SupportPhoneNumber:                     "111-222-3333",
		SupportEmailAddress:                    "email@company.com",
		Department:                             "department name",
		DefaultPrestage:                        false,
		EnrollmentSiteID:                       "-1",
		KeepExistingSiteMembership:             false,
		KeepExistingLocationInformation:        false,
		RequireAuthentication:                  false,
		AuthenticationPrompt:                   "",
		PreventActivationLock:                  true,
		EnableDeviceBasedActivationLock:        false,
		DeviceEnrollmentProgramInstanceID:      "1",
		AnchorCertificates:                     []string{},
		EnrollmentCustomizationID:              "0",
		Language:                               "",
		Region:                                 "",
		AutoAdvanceSetup:                       false,
		AllowPairing:                           true,
		MultiUser:                              false,
		Supervised:                             true,
		MaximumSharedAccounts:                  10,
		ConfigureDeviceBeforeSetupAssistant:    true,
		SendTimezone:                           false,
		Timezone:                               "UTC",
		StorageQuotaSizeMegabytes:              4096,
		UseStorageQuotaSize:                    false,
		TemporarySessionOnly:                   false,
		EnforceTemporarySessionTimeout:         false,
		TemporarySessionTimeout:                nil,
		EnforceUserSessionTimeout:              false,
		UserSessionTimeout:                     nil,
		SiteId:                                 "-1",
		VersionLock:                            currentPrestage.VersionLock,
		PrestageMinimumOsTargetVersionTypeIos:  "MINIMUM_OS_SPECIFIC_VERSION",
		MinimumOsSpecificVersionIos:            "18.4",
		PrestageMinimumOsTargetVersionTypeIpad: "MINIMUM_OS_LATEST_MINOR_VERSION",
		MinimumOsSpecificVersionIpad:           "",
	}

	update.SkipSetupItems = jamfpro.MobileDevicePrestageSubsetSkipSetupItems{
		// Selected items are not displayed in the Setup Assistant during enrollment
		Location:              true,
		Privacy:               true,
		Biometric:             true,
		SoftwareUpdate:        true,
		Diagnostics:           true,
		IMessageAndFaceTime:   true,
		Intelligence:          true,
		TVRoom:                true,
		Passcode:              true,
		SIMSetup:              true,
		ScreenTime:            true,
		RestoreCompleted:      true,
		TVProviderSignIn:      true,
		Siri:                  true,
		Restore:               true,
		ScreenSaver:           true,
		HomeButtonSensitivity: true,
		CloudStorage:          true,
		ActionButton:          true,
		TransferData:          true,
		EnableLockdownMode:    true,
		Zoom:                  true,
		PreferredLanguage:     true,
		VoiceSelection:        true,
		TVHomeScreenSync:      true,
		Safety:                true,
		TermsOfAddress:        true,
		ExpressLanguage:       true,
		CameraButton:          true,
		AppleID:               true,
		DisplayTone:           true,
		WatchMigration:        true,
		UpdateCompleted:       true,
		Appearance:            true,
		Android:               true,
		Payment:               true,
		OnBoarding:            true,
		TOS:                   true,
		Welcome:               true,
		TapToSetup:            true,
	}

	update.LocationInformation = jamfpro.MobileDevicePrestageSubsetLocationInformation{
		ID:           "-1", // Required
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

	update.PurchasingInformation = jamfpro.MobileDevicePrestageSubsetPurchasingInformation{
		ID:                "-1", // Required
		Leased:            false,
		Purchased:         true,
		AppleCareId:       "",
		PoNumber:          "",
		Vendor:            "",
		PurchasePrice:     "",
		LifeExpectancy:    0,
		PurchasingAccount: "",
		PurchasingContact: "",
		LeaseDate:         "1970-01-01",
		PoDate:            "1970-01-01",
		WarrantyDate:      "1970-01-01",
		VersionLock:       currentPrestage.LocationInformation.VersionLock + 1,
	}

	update.Names = jamfpro.MobileDevicePrestageSubsetNames{
		AssignNamesUsing:       "Serial Numbers",
		PrestageDeviceNames:    []jamfpro.MobileDevicePrestageSubsetNamesName{},
		DeviceNamePrefix:       "",
		DeviceNameSuffix:       "",
		SingleDeviceName:       "",
		ManageNames:            true,
		DeviceNamingConfigured: true,
	}

	// Call UpdateMobileDevicePrestageByID to update the prestage
	updatedPrestage, err := client.UpdateMobileDevicePrestageByID(prestageID, update)
	if err != nil {
		log.Fatalf("Error updating mobile device prestage: %v", err)
	}

	// Pretty print the updated mobile device prestage in JSON
	updatedPrestageJSON, err := json.MarshalIndent(updatedPrestage, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated prestage data: %v", err)
	}
	fmt.Println("Updated mobile device prestage:")
	fmt.Println(string(updatedPrestageJSON))
}
