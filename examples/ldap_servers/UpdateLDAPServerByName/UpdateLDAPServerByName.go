package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define updated LDAP server details (adjust as needed)
	updatedLDAPServer := &jamfpro.ResourceLDAPServers{
		Connection: jamfpro.LDAPServerSubsetConnection{
			Name:               "Company Active Directory",
			Hostname:           "company.ad.com",
			ServerType:         "Active Directory",
			Port:               389,
			UseSSL:             true,
			AuthenticationType: "simple",
			Account: jamfpro.LDAPServerSubsetConnectionAccount{
				DistinguishedUsername: "CN=Administrator,CN=Users,DC=Company,DC=com",
				Password:              "password",
			},
			OpenCloseTimeout: 15,
			SearchTimeout:    60,
			ReferralResponse: "ignore",
			UseWildcards:     true,
			// Additional fields if necessary...
		},
		MappingsForUsers: jamfpro.LDAPServerContainerMapping{
			UserMappings: jamfpro.LDAPServerSubsetMappingUsers{
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
			UserGroupMappings: jamfpro.LDAPServerSubsetMappingUserGroups{
				MapObjectClassToAnyOrAll: "all",
				ObjectClasses:            "top, group",
				SearchBase:               "DC=Company,DC=com",
				SearchScope:              "All Subtrees",
				MapGroupID:               "uSNCreated",
				MapGroupName:             "name",
				MapGroupUUID:             "objectGUID",
				// Additional fields if necessary...
			},
			UserGroupMembershipMappings: jamfpro.LDAPServerSubsetMappingUserGroupMemberships{
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
