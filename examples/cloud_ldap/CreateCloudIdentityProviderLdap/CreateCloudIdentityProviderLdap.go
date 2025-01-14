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

	// Create the request body
	request := &jamfpro.ResourceCloudLdap{
		CloudIdPCommon: &jamfpro.CloudIdPCommon{
			ProviderName: "GOOGLE",
			DisplayName:  "test",
		},
		Server: &jamfpro.CloudLdapServer{
			Enabled: true,
			Keystore: &jamfpro.CloudLdapKeystore{
				Password:  "thing",
				FileBytes: "WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09",
				FileName:  "keystore.jks", // Added example filename
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
				AdditionalSearchBase:  "thing",
				UserID:                "mail",
				Username:              "uid",
				RealName:              "displayName",
				EmailAddress:          "mail",
				Department:            "departmentNumber",
				Building:              "building",
				Room:                  "room",
				Phone:                 "phone",
				Position:              "position",
				UserUuid:              "uid",
			},
			GroupMappings: jamfpro.CloudIdentityProviderDefaultMappingsSubsetGroupMappings{
				ObjectClassLimitation: "ANY_OBJECT_CLASSES",
				SearchScope:           "ALL_SUBTREES",
				ObjectClasses:         "groupOfNames",
				SearchBase:            "ou=Groups",
				GroupID:               "cn=",
				GroupName:             "cn=",
				GroupUuid:             "gidNumber",
			},
			MembershipMappings: jamfpro.CloudIdentityProviderDefaultMappingsSubsetMembershipMappings{
				GroupMembershipMapping: "memberOf",
			},
		},
	}

	// Create the Cloud LDAP configuration
	resp, err := client.CreateCloudIdentityProviderLdap(request)
	if err != nil {
		log.Fatalf("Error creating cloud LDAP configuration: %v", err)
	}

	// Pretty print the response in JSON
	responseJSON, err := json.MarshalIndent(resp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling response data: %v", err)
	}
	fmt.Println("Created Cloud LDAP Configuration:\n", string(responseJSON))
}
