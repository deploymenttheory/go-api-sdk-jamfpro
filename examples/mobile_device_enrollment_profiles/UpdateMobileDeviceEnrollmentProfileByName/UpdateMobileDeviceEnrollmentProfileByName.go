package main

import (
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

	// Example profile to be updated
	profileToUpdate := jamfpro.ResourceMobileDeviceEnrollmentProfile{
		General: jamfpro.MobileDeviceEnrollmentProfileSubsetGeneral{
			Name:        "Configurator Enrollment Profile",
			Description: "string",
		},
		Location: jamfpro.MobileDeviceEnrollmentProfileSubsetLocation{
			// Initialize with empty or specific values if required
			Username:     "",
			Realname:     "",
			RealName:     "",
			EmailAddress: "",
			Position:     "",
			Phone:        "",
			PhoneNumber:  "",
			Department:   "",
			Building:     "",
			Room:         0, // or specific room number
		},
		Purchasing: jamfpro.MobileDeviceEnrollmentProfileSubsetPurchasing{
			IsPurchased:          true,
			IsLeased:             false,
			PONumber:             "",
			Vendor:               "",
			ApplecareID:          "",
			PurchasePrice:        "",
			PurchasingAccount:    "",
			PODate:               "",
			PODateEpoch:          0,
			PODateUTC:            "",
			WarrantyExpires:      "",
			WarrantyExpiresEpoch: 0,
			WarrantyExpiresUTC:   "",
			LeaseExpires:         "",
			LeaseExpiresEpoch:    0,
			LeaseExpiresUTC:      "",
			LifeExpectancy:       0,
			PurchasingContact:    "",
		},
	}

	profileName := "Configurator Enrollment Profile" // Replace name with the actual enrollment profile name

	updatedProfile, err := client.UpdateMobileDeviceEnrollmentProfileByName(profileName, &profileToUpdate)
	if err != nil {
		log.Fatalf("Error updating profile: %v", err)
	}

	fmt.Println("Updated Profile:", updatedProfile)
}
