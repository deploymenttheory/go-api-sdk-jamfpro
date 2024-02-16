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

	// Specify the path to the file you want to upload
	filePath := "/Users/dafyddwatkins/localtesting/support_files/packages/OpenJDK21U-jdk_aarch64_mac_hotspot_21.0.2_13.pkg"

	// Call CreateJCDS2Package with the file path
	response, err := client.CreateJCDS2PackageV2(filePath)
	if err != nil {
		log.Fatalf("Failed to create JCDS 2.0 package: %v", err)
	}

	// Print the response
	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}
	fmt.Println("Response:", string(responseBytes))
}
