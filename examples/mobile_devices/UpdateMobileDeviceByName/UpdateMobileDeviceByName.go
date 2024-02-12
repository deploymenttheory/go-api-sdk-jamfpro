package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

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

	// Construct the mobile device object
	mobileDevice := &jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			DisplayName:  "iPad",
			DeviceName:   "iPad",
			Name:         "iPad",
			AssetTag:     "string",
			SerialNumber: "C02Q7KHTGFWD",
			UDID:         "270aae10800b6e61a2ee2bbc285eb967050b6984",
		},
	}

	mobileDeviceName := "iPad"

	// Create the mobile device in Jamf Pro
	responseDevice, err := client.UpdateMobileDeviceByName(mobileDeviceName, mobileDevice)
	if err != nil {
		log.Fatalf("Failed to create mobile device: %v", err)
	}

	// Print the response
	fmt.Printf("Created Mobile Device: %+v\n", responseDevice)
}
