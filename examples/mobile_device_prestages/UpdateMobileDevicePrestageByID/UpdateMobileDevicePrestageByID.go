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
		Mandatory:                              jamfpro.TruePtr(),
		MdmRemovable:                           jamfpro.FalsePtr(),
		SupportPhoneNumber:                     "111-222-3333",
		SupportEmailAddress:                    "email@company.com",
		Department:                             "department name",
		DefaultPrestage:                        jamfpro.FalsePtr(),
		EnrollmentSiteID:                       "-1",
		KeepExistingSiteMembership:             jamfpro.FalsePtr(),
		KeepExistingLocationInformation:        jamfpro.FalsePtr(),
		RequireAuthentication:                  jamfpro.FalsePtr(),
		AuthenticationPrompt:                   "",
		PreventActivationLock:                  jamfpro.TruePtr(),
		EnableDeviceBasedActivationLock:        jamfpro.FalsePtr(),
		DeviceEnrollmentProgramInstanceID:      "1",
		AnchorCertificates:                     []string{},
		EnrollmentCustomizationID:              "0",
		Language:                               "",
		Region:                                 "",
		AutoAdvanceSetup:                       jamfpro.FalsePtr(),
		AllowPairing:                           jamfpro.TruePtr(),
		MultiUser:                              jamfpro.FalsePtr(),
		Supervised:                             jamfpro.TruePtr(),
		MaximumSharedAccounts:                  10,
		ConfigureDeviceBeforeSetupAssistant:    jamfpro.TruePtr(),
		SendTimezone:                           jamfpro.FalsePtr(),
		Timezone:                               "UTC",
		StorageQuotaSizeMegabytes:              4096,
		UseStorageQuotaSize:                    jamfpro.FalsePtr(),
		TemporarySessionOnly:                   jamfpro.FalsePtr(),
		EnforceTemporarySessionTimeout:         jamfpro.FalsePtr(),
		TemporarySessionTimeout:                nil,
		EnforceUserSessionTimeout:              jamfpro.FalsePtr(),
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
		Location:              jamfpro.TruePtr(),
		Privacy:               jamfpro.TruePtr(),
		Biometric:             jamfpro.TruePtr(),
		SoftwareUpdate:        jamfpro.TruePtr(),
		Diagnostics:           jamfpro.TruePtr(),
		IMessageAndFaceTime:   jamfpro.TruePtr(),
		Intelligence:          jamfpro.TruePtr(),
		TVRoom:                jamfpro.TruePtr(),
		Passcode:              jamfpro.TruePtr(),
		SIMSetup:              jamfpro.TruePtr(),
		ScreenTime:            jamfpro.TruePtr(),
		RestoreCompleted:      jamfpro.TruePtr(),
		TVProviderSignIn:      jamfpro.TruePtr(),
		Siri:                  jamfpro.TruePtr(),
		Restore:               jamfpro.TruePtr(),
		ScreenSaver:           jamfpro.TruePtr(),
		HomeButtonSensitivity: jamfpro.TruePtr(),
		CloudStorage:          jamfpro.TruePtr(),
		ActionButton:          jamfpro.TruePtr(),
		TransferData:          jamfpro.TruePtr(),
		EnableLockdownMode:    jamfpro.TruePtr(),
		Zoom:                  jamfpro.TruePtr(),
		PreferredLanguage:     jamfpro.TruePtr(),
		VoiceSelection:        jamfpro.TruePtr(),
		TVHomeScreenSync:      jamfpro.TruePtr(),
		Safety:                jamfpro.TruePtr(),
		TermsOfAddress:        jamfpro.TruePtr(),
		ExpressLanguage:       jamfpro.TruePtr(),
		CameraButton:          jamfpro.TruePtr(),
		AppleID:               jamfpro.TruePtr(),
		DisplayTone:           jamfpro.TruePtr(),
		WatchMigration:        jamfpro.TruePtr(),
		UpdateCompleted:       jamfpro.TruePtr(),
		Appearance:            jamfpro.TruePtr(),
		Android:               jamfpro.TruePtr(),
		Payment:               jamfpro.TruePtr(),
		OnBoarding:            jamfpro.TruePtr(),
		TOS:                   jamfpro.TruePtr(),
		Welcome:               jamfpro.TruePtr(),
		TapToSetup:            jamfpro.TruePtr(),
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
		Leased:            jamfpro.FalsePtr(),
		Purchased:         jamfpro.TruePtr(),
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
		ManageNames:            jamfpro.TruePtr(),
		DeviceNamingConfigured: jamfpro.TruePtr(),
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
