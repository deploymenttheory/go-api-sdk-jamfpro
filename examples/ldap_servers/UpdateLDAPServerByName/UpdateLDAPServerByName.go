package main

import (
	"encoding/xml"
	"fmt"
	"log"

	// Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file for OAuth credentials
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

	// Define updated LDAP server details (adjust as needed)
	updatedLDAPServer := &jamfpro.ResponseLDAPServers{
		Connection: jamfpro.LDAPConnection{
			Name:               "Company Active Directory",
			Hostname:           "company.ad.com",
			ServerType:         "Active Directory",
			Port:               389,
			UseSSL:             true,
			AuthenticationType: "simple",
			Account: jamfpro.LDAPAccount{
				DistinguishedUsername: "CN=Administrator,CN=Users,DC=Company,DC=com",
				Password:              "password",
			},
			OpenCloseTimeout: 15,
			SearchTimeout:    60,
			ReferralResponse: "ignore",
			UseWildcards:     true,
			// Additional fields if necessary...
		},
		MappingsForUsers: jamfpro.LDAPMappingsForUsers{
			UserMappings: jamfpro.LDAPUserMappings{
				MapObjectClassToAnyOrAll: "all",
				ObjectClasses:            "organizationalPerson, user",
				SearchBase:               "DC=Company,DC=com",
				SearchScope:              "All Subtrees",
				MapUserID:                "uSNCreated",
				MapUsername:              "sAMAccountName",
				MapRealName:              "displayName",
				MapEmailAddress:          "mail",
				AppendToEmailResults:     "company.com",
				MapDepartment:            "department",
				MapBuilding:              "streetAddress",
				MapRoom:                  "room",
				MapTelephone:             "telephoneNumber",
				MapPosition:              "title",
				MapUserUUID:              "objectGUID",
				// Additional fields if necessary...
			},
			UserGroupMappings: jamfpro.LDAPUserGroupMappings{
				MapObjectClassToAnyOrAll: "all",
				ObjectClasses:            "top, group",
				SearchBase:               "DC=Company,DC=com",
				SearchScope:              "All Subtrees",
				MapGroupID:               "uSNCreated",
				MapGroupName:             "name",
				MapGroupUUID:             "objectGUID",
				// Additional fields if necessary...
			},
			UserGroupMembershipMappings: jamfpro.LDAPGroupMembershipMappings{
				UserGroupMembershipStoredIn:       "user object",
				MapGroupMembershipToUserField:     "memberOf",
				AppendToUsername:                  "company.com",
				UseDN:                             true,
				RecursiveLookups:                  true,
				MapUserMembershipToGroupField:     true,
				MapUserMembershipUseDN:            true,
				MapObjectClassToAnyOrAll:          "all",
				ObjectClasses:                     "group",
				SearchBase:                        "DC=Company,DC=com",
				SearchScope:                       "All Subtrees",
				Username:                          "sAMAccountName",
				GroupID:                           "uSNCreated",
				UserGroupMembershipUseLDAPCompare: true,
				// Additional fields if necessary...
			},
			// Additional fields if necessary...
		},
	}

	// Update LDAP server by Name
	LDAPServerName := "Company Active Directory" // Replace with actual LDAP server Name
	updatedServer, err := client.UpdateLDAPServerByName(LDAPServerName, updatedLDAPServer)
	if err != nil {
		log.Fatalf("Error updating LDAP server by ID: %v", err)
	}

	// Print updated LDAP server details
	serverXML, err := xml.MarshalIndent(updatedServer, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated LDAP server data: %v", err)
	}
	fmt.Println("Updated LDAP Server:", string(serverXML))
}
