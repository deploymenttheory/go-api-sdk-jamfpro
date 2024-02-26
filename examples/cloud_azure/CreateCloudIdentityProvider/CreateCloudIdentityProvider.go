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
