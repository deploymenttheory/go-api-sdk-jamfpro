package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
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
	logLevel := http_client.LogLevelDebug

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

	// The ID of the computer prestage you want to update
	prestageName := "YOUR_PRESTAGE_ID_HERE"

	// Define the updated information for the computer prestage
	// Here we are just creating a new instance for demonstration purposes.
	// You would typically populate this struct with your updated data.
	var update jamfpro.ResourceComputerPrestage
	updateData := `{
		"mandatory": true,
		"mdmRemovable": true,
		"defaultPrestage": true,
		"keepExistingSiteMembership": true,
		"keepExistingLocationInformation": true,
		"requireAuthentication": true,
		"preventActivationLock": true,
		"enableDeviceBasedActivationLock": true,
		"skipSetupItems": {
			"newKey": true
		},
		"locationInformation": {
			"email": "test@jamf.com",
			"room": "room",
			"position": "postion",
			"departmentId": "1",
			"id": "-1",
			"versionLock": 1,
			"buildingId": "1",
			"username": "name",
			"realname": "realName",
			"phone": "123-456-7890"
		},
		"purchasingInformation": {
			"leased": true,
			"purchased": true,
			"id": "-1",
			"appleCareId": "abcd",
			"poNumber": "53-1",
			"vendor": "Example Vendor",
			"purchasePrice": "$500",
			"lifeExpectancy": 5,
			"purchasingAccount": "admin",
			"purchasingContact": "true",
			"leaseDate": "2019-01-01",
			"poDate": "2019-01-01",
			"warrantyDate": "2019-01-01",
			"versionLock": 1
		},
		"autoAdvanceSetup": true,
		"installProfilesDuringSetup": true,
		"accountSettings": {
			"payloadConfigured": false,
			"localAdminAccountEnabled": false,
			"hiddenAdminAccount": false,
			"localUserManaged": false,
			"userAccountType": "STANDARD",
			"versionLock": 0,
			"prefillPrimaryAccountInfoFeatureEnabled": false,
			"prefillType": "CUSTOM",
			"preventPrefillInfoFromModification": false,
			"id": "1",
			"adminUsername": "admin",
			"adminPassword": "password"
		},
		"displayName": "test",
		"supportPhoneNumber": "5555555555",
		"supportEmailAddress": "someemail@domain.com",
		"department": "Oxbow",
		"enrollmentSiteId": "-1",
		"authenticationPrompt": "LDAP authentication prompt",
		"deviceEnrollmentProgramInstanceId": "5",
		"anchorCertificates": [
			"xNE5HRgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
		],
		"enrollmentCustomizationId": "2",
		"language": "en",
		"region": "US",
		"prestageInstalledProfileIds": [
			"-1"
		],
		"customPackageDistributionPointId": "1",
		"enableRecoveryLock": true,
		"recoveryLockPasswordType": "MANUAL",
		"rotateRecoveryLockPassword": true,
		"recoveryLockPassword": "password123",
		"customPackageIds": [
			"-1"
		]
	}`
	err = json.Unmarshal([]byte(updateData), &update)
	if err != nil {
		log.Fatalf("Error unmarshaling update data: %v", err)
	}

	// Call UpdateComputerPrestageByName to update the prestage
	updatedPrestage, err := client.UpdateComputerPrestageByName(prestageName, &update)
	if err != nil {
		log.Fatalf("Error updating computer prestage: %v", err)
	}

	// Print the updated prestage
	fmt.Printf("Updated Computer Prestage: %+v\n", updatedPrestage)
}
