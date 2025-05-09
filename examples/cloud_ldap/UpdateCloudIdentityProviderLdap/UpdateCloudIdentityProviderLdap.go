package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the ID of the Cloud LDAP configuration you want to update
	ldapID := "1022" // Replace with your actual LDAP configuration ID

	// Create the update request body
	request := &jamfpro.ResourceCloudLdap{
		CloudIdPCommon: &jamfpro.CloudIdPCommon{
			ID:           ldapID,
			ProviderName: "GOOGLE",
			DisplayName:  "Google LDAP Updated", // Example of updated display name
		},
		Server: &jamfpro.CloudLdapServer{
			Enabled: true,
			Keystore: &jamfpro.CloudLdapKeystore{
				Password:  "supersecretpassword",
				FileBytes: "MIIJsQIBAzCCCXcGCSqGS...",
				FileName:  "keystore.p12",
			},
			UseWildcards:                             true,
			ConnectionType:                           "LDAPS",
			ServerUrl:                                "ldap.google.com",
			DomainName:                               "jamf.com",
			Port:                                     636,
			ConnectionTimeout:                        15,
			SearchTimeout:                            60,
			MembershipCalculationOptimizationEnabled: true,
		},
		Mappings: &jamfpro.CloudLdapMappings{
			UserMappings: jamfpro.CloudIdentityProviderDefaultMappingsSubsetUserMappings{
				ObjectClassLimitation: "ANY_OBJECT_CLASSES",
				SearchScope:           "ALL_SUBTREES",
				ObjectClasses:         "inetOrgPerson",
				SearchBase:            "ou=Users",
				AdditionalSearchBase:  "",
				UserID:                "uid",
				Username:              "mail",
				RealName:              "displayName",
				EmailAddress:          "mail",
				Department:            "departmentNumber",
				Building:              "",
				Room:                  "",
				Phone:                 "",
				Position:              "title",
				UserUuid:              "uid",
			},
			GroupMappings: jamfpro.CloudIdentityProviderDefaultMappingsSubsetGroupMappings{
				ObjectClassLimitation: "ANY_OBJECT_CLASSES",
				SearchScope:           "ALL_SUBTREES",
				ObjectClasses:         "groupOfNames",
				SearchBase:            "ou=Groups",
				GroupID:               "cn",
				GroupName:             "cn",
				GroupUuid:             "gidNumber",
			},
			MembershipMappings: jamfpro.CloudIdentityProviderDefaultMappingsSubsetMembershipMappings{
				GroupMembershipMapping: "memberOf",
			},
		},
	}

	// Update the Cloud LDAP configuration
	resp, err := client.UpdateCloudIdentityProviderLdap(ldapID, request)
	if err != nil {
		log.Fatalf("Error updating cloud LDAP configuration: %v", err)
	}

	// Pretty print the response
	responseJSON, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling response data: %v", err)
	}

	fmt.Printf("Successfully updated Cloud LDAP configuration with ID %s:\n%s\n", ldapID, string(responseJSON))

	// Optionally fetch and display the updated configuration
	updatedConfig, err := client.GetCloudIdentityProviderLdapByID(ldapID)
	if err != nil {
		log.Fatalf("Error fetching updated configuration: %v", err)
	}

	// Pretty print the updated configuration
	updatedJSON, err := json.MarshalIndent(updatedConfig, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated configuration data: %v", err)
	}

	fmt.Printf("\nFull Updated Configuration:\n%s\n", string(updatedJSON))
}
