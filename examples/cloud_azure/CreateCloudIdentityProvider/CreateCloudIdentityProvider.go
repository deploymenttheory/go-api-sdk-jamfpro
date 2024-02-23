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
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
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
			LogLevel:            loadedConfig.ClientOptions.LogLevel,
			LogOutputFormat:     loadedConfig.ClientOptions.LogOutputFormat,
			LogConsoleSeparator: loadedConfig.ClientOptions.LogConsoleSeparator,
			HideSensitiveData:   loadedConfig.ClientOptions.HideSensitiveData,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Define the Cloud Identity Provider data
	cloudIdPData := jamfpro.ResourceCloudIdp{
		CloudIdPCommon: jamfpro.CloudIdpListItem{
			DisplayName:  "Cloud Identity Provider",
			ProviderName: "AZURE",
		},
		Server: jamfpro.ResourceCloudIdpServer{
			TenantId: "db65d325-0350-4a17-9af9-b302d0fc386b",
			Enabled:  true,
			Migrated: true, // Note: Ensure this field is required as per your API's specification
			Mappings: jamfpro.CloudIdpServerSubsetCloudIdpServerMappings{
				UserId:     "id",
				UserName:   "userPrincipalName",
				RealName:   "displayName",
				Email:      "mail",
				Department: "department",
				Building:   "companyName",
				Room:       "officeLocation",
				Phone:      "mobilePhone",
				Position:   "jobTitle",
				GroupId:    "id",
				GroupName:  "displayName",
			},
			SearchTimeout:                            30,
			TransitiveMembershipEnabled:              false,
			TransitiveMembershipUserField:            "userPrincipalName",
			TransitiveDirectoryMembershipEnabled:     false,
			MembershipCalculationOptimizationEnabled: true,
			Code:                                     "auth",
		},
	}

	// Create the Cloud Identity Provider
	response, err := client.CreateCloudIdentityProvider(&cloudIdPData)
	if err != nil {
		fmt.Printf("Error creating Cloud Identity Provider: %s\n", err)
		return
	}

	// Marshal the response into pretty JSON
	prettyJSON, err := json.MarshalIndent(response, "", "    ") // Indents with 4 spaces
	if err != nil {
		log.Fatalf("Failed to marshal response into JSON: %v", err)
	}

	// Output the pretty-printed JSON
	fmt.Printf("Cloud Identity Provider Created:\n%s\n", string(prettyJSON))
}
