package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the new prestage with the provided example
	newPrestage := jamfpro.ComputerPrestagesItem{
		DisplayName:                     "computer-prestage-name",
		Mandatory:                       true,
		MDMRemovable:                    false,
		DefaultPrestage:                 false,
		KeepExistingSiteMembership:      true,
		KeepExistingLocationInformation: true,
		RequireAuthentication:           true,
		PreventActivationLock:           true,
		EnableDeviceBasedActivationLock: true,
		SkipSetupItems: map[string]bool{
			"newKey": false,
		},
		LocationInformation: jamfpro.ComputerPrestagesLocationInformation{
			Phone:        "123-456-7890",
			Realname:     "realName",
			Room:         "room",
			Position:     "position",
			DepartmentId: "-1",
			BuildingId:   "-1",
			ID:           "-1",
			VersionLock:  1,
			Username:     "name",
			Email:        "test@jamf.com",
		},
		PurchasingInformation: jamfpro.ComputerPrestagesPurchasingInformation{
			Leased:            true,
			Purchased:         true,
			ID:                "-1",
			AppleCareId:       "abcd",
			PONumber:          "53-1",
			Vendor:            "Example Vendor",
			PurchasePrice:     "$500",
			LifeExpectancy:    5,
			PurchasingAccount: "admin",
			PurchasingContact: "true",
			LeaseDate:         "2019-01-01",
			PODate:            "2019-01-01",
			WarrantyDate:      "2019-01-01",
			VersionLock:       1,
		},
		AutoAdvanceSetup:                  true,
		InstallProfilesDuringSetup:        true,
		SupportPhoneNumber:                "5555555555",
		SupportEmailAddress:               "example@example.com",
		Department:                        "Oxbow",
		EnrollmentSiteId:                  "-1",
		AuthenticationPrompt:              "LDAP authentication prompt",
		DeviceEnrollmentProgramInstanceId: "5",
		AnchorCertificates: []string{
			"xNE5HRgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
		},
		EnrollmentCustomizationId:        "2",
		Language:                         "en",
		Region:                           "US",
		PrestageInstalledProfileIds:      []string{"1"},
		CustomPackageIds:                 []string{"1"},
		CustomPackageDistributionPointId: "1",
		RecoveryLockPasswordType:         "RANDOM",
		EnableRecoveryLock:               true,
		RotateRecoveryLockPassword:       true,
		RecoveryLockPassword:             "password123",
		AccountSettings: jamfpro.ComputerPrestagesAccountSettings{
			PayloadConfigured:                       false,
			LocalAdminAccountEnabled:                false,
			HiddenAdminAccount:                      false,
			LocalUserManaged:                        false,
			UserAccountType:                         "STANDARD",
			VersionLock:                             0,
			PrefillPrimaryAccountInfoFeatureEnabled: false,
			PrefillType:                             "CUSTOM",
			PreventPrefillInfoFromModification:      false,
			ID:                                      "1",
			AdminUsername:                           "admin",
			AdminPassword:                           "password", // Assuming this field exists in your struct
			PrefillAccountFullName:                  "TestUser FullName",
			PrefillAccountUserName:                  "username",
		},
	}

	// Use the address-of operator (&) to pass a pointer to CreateComputerPrestage
	createdPrestage, err := client.CreateComputerPrestage(&newPrestage)
	if err != nil {
		log.Fatalf("Error creating computer prestage: %v", err)
	}

	fmt.Printf("Created computer prestage: %+v\n", createdPrestage)
}
