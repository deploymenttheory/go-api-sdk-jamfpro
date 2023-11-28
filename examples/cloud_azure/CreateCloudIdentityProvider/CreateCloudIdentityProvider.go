package main

import (
	"encoding/json"
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

	// Define the Cloud Identity Provider data
	cloudIdPData := jamfpro.ResponseCloudIDP{
		CloudIdPCommon: jamfpro.CloudIdPCommon{
			DisplayName:  "Cloud Identity Provider",
			ProviderName: "AZURE",
		},
		Server: jamfpro.CloudIdPServer{
			TenantId: "db65d325-0350-4a17-9af9-b302d0fc386b",
			Enabled:  true,
			Migrated: true, // Note: Ensure this field is required as per your API's specification
			Mappings: jamfpro.CloudIdPServerMappings{
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
