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
		DisplayName:                       "Example Mobile Prestage Name",
		Mandatory:                         jamfpro.FalsePtr(),
		MDMRemovable:                      jamfpro.TruePtr(),
		SupportPhoneNumber:                "5555555555",
		SupportEmailAddress:               "example@example.com",
		Department:                        "Oxbow",
		DefaultPrestage:                   jamfpro.FalsePtr(),
		EnrollmentSiteId:                  "-1",
		KeepExistingSiteMembership:        jamfpro.TruePtr(),
		KeepExistingLocationInformation:   jamfpro.TruePtr(),
		RequireAuthentication:             jamfpro.TruePtr(),
		AuthenticationPrompt:              "LDAP authentication prompt",
		PreventActivationLock:             jamfpro.TruePtr(),
		EnableDeviceBasedActivationLock:   jamfpro.TruePtr(),
		DeviceEnrollmentProgramInstanceId: "1",
		SkipSetupItems: jamfpro.ComputerPrestageSubsetSkipSetupItems{
			Location: jamfpro.TruePtr(),
			Privacy:  jamfpro.FalsePtr(),
		},
		LocationInformation: jamfpro.ComputerPrestageSubsetLocationInformation{
			Username:     "name",
			Realname:     "realName",
			Phone:        "123-456-7890",
			Email:        "test@jamf.com",
			Room:         "room",
			Position:     "postion",
			DepartmentId: "1",
			BuildingId:   "1",
			//ID:           "-1",
			VersionLock: 1,
		},
		PurchasingInformation: jamfpro.ComputerPrestageSubsetPurchasingInformation{
			//ID:                "-1",
			Leased:            jamfpro.TruePtr(),
			Purchased:         jamfpro.TruePtr(),
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
		EnrollmentCustomizationId:  "0",
		Language:                   "en",
		Region:                     "US",
		AutoAdvanceSetup:           jamfpro.TruePtr(),
		InstallProfilesDuringSetup: jamfpro.TruePtr(),
		PrestageInstalledProfileIds: []string{
			"3847",
			"3864",
			"3806",
		},
		CustomPackageIds:                 []string{},
		CustomPackageDistributionPointId: "-1",
		EnableRecoveryLock:               jamfpro.TruePtr(),
		RecoveryLockPasswordType:         "MANUAL",
		RotateRecoveryLockPassword:       jamfpro.TruePtr(),
		ProfileUuid:                      "0386E7C8D455A040106850A8A2033968",
		SiteId:                           "-1",
		VersionLock:                      1,
		AccountSettings: jamfpro.ComputerPrestageSubsetAccountSettings{
			//ID:                                      "-1",
			PayloadConfigured:                       jamfpro.TruePtr(),
			LocalAdminAccountEnabled:                jamfpro.TruePtr(),
			AdminUsername:                           "admin",
			AdminPassword:                           "password",
			HiddenAdminAccount:                      jamfpro.FalsePtr(),
			LocalUserManaged:                        jamfpro.TruePtr(),
			UserAccountType:                         "STANDARD",
			VersionLock:                             1,
			PrefillPrimaryAccountInfoFeatureEnabled: jamfpro.TruePtr(),
			PrefillType:                             "DEVICE_OWNER",
			PrefillAccountFullName:                  "TestUser FullName",
			PrefillAccountUserName:                  "UserName",
			PreventPrefillInfoFromModification:      jamfpro.FalsePtr(),
		},
		RecoveryLockPassword: "password123",
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
