package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	//logLevel := logger.LogLevelInfo // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// The ID of the computer prestage you want to update
	prestageID := "YOUR_PRESTAGE_ID_HERE"

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

	// Call UpdateComputerPrestageByID to update the prestage
	updatedPrestage, err := client.UpdateComputerPrestageByID(prestageID, &update)
	if err != nil {
		log.Fatalf("Error updating computer prestage: %v", err)
	}

	// Pretty print the computer prestage in JSON
	prestageJSON, err := json.MarshalIndent(updatedPrestage, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling computer prestage data: %v", err)
	}
	fmt.Println("Fetched computer prestage:\n", string(prestageJSON))
}
