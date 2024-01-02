// Refactor Complete

// classicapi_ldap_servers.go
// Jamf Pro Classic Api - LDAP Servers
// api reference: https://developer.jamf.com/jamf-pro/reference/ldapservers
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriLDAPServers = "/JSSResource/ldapservers"

// ResponseLDAPServersList represents the response structure for a list of LDAP servers.

// List

type ResponseLDAPServersList struct {
	Size        int                   `xml:"size"`
	LDAPServers []LDAPServersListItem `xml:"ldap_server"`
}

type LDAPServersListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// ResourceLDAPServers represents the structure of an individual LDAP server.
type ResourceLDAPServers struct {
	Connection       LDAPServerSubsetConnection `xml:"connection"`
	MappingsForUsers LDAPServerContainerMapping `xml:"mappings_for_users"`
}

// Subsets & Containers

// Connection

type LDAPServerSubsetConnection struct {
	ID                 int                               `xml:"id"`
	Name               string                            `xml:"name"`
	Hostname           string                            `xml:"hostname"`
	ServerType         string                            `xml:"server_type"`
	Port               int                               `xml:"port"`
	UseSSL             bool                              `xml:"use_ssl"`
	AuthenticationType string                            `xml:"authentication_type"`
	Account            LDAPServerSubsetConnectionAccount `xml:"account"`
	OpenCloseTimeout   int                               `xml:"open_close_timeout"`
	SearchTimeout      int                               `xml:"search_timeout"`
	ReferralResponse   string                            `xml:"referral_response"`
	UseWildcards       bool                              `xml:"use_wildcards"`
}

type LDAPServerSubsetConnectionAccount struct {
	DistinguishedUsername string `xml:"distinguished_username"`
	Password              string `xml:"password"`
}

// MappingsForUsers

type LDAPServerContainerMapping struct {
	UserMappings                LDAPServerSubsetMappingUsers                `xml:"user_mappings"`
	UserGroupMappings           LDAPServerSubsetMappingUserGroups           `xml:"user_group_mappings"`
	UserGroupMembershipMappings LDAPServerSubsetMappingUserGroupMemberships `xml:"user_group_membership_mappings"`
}

type LDAPServerSubsetMappingUsers struct {
	MapObjectClassToAnyOrAll string `xml:"map_object_class_to_any_or_all"`
	ObjectClasses            string `xml:"object_classes"`
	SearchBase               string `xml:"search_base"`
	SearchScope              string `xml:"search_scope"`
	MapUserID                string `xml:"map_user_id"`
	MapUsername              string `xml:"map_username"`
	MapRealName              string `xml:"map_realname"`
	MapEmailAddress          string `xml:"map_email_address"`
	AppendToEmailResults     string `xml:"append_to_email_results"`
	MapDepartment            string `xml:"map_department"`
	MapBuilding              string `xml:"map_building"`
	MapRoom                  string `xml:"map_room"`
	MapTelephone             string `xml:"map_telephone"`
	MapPosition              string `xml:"map_position"`
	MapUserUUID              string `xml:"map_user_uuid"`
}

type LDAPServerSubsetMappingUserGroups struct {
	MapObjectClassToAnyOrAll string `xml:"map_object_class_to_any_or_all"`
	ObjectClasses            string `xml:"object_classes"`
	SearchBase               string `xml:"search_base"`
	SearchScope              string `xml:"search_scope"`
	MapGroupID               string `xml:"map_group_id"`
	MapGroupName             string `xml:"map_group_name"`
	MapGroupUUID             string `xml:"map_group_uuid"`
}

type LDAPServerSubsetMappingUserGroupMemberships struct {
	UserGroupMembershipStoredIn       string `xml:"user_group_membership_stored_in"`
	MapGroupMembershipToUserField     string `xml:"map_group_membership_to_user_field"`
	AppendToUsername                  string `xml:"append_to_username"`
	UseDN                             bool   `xml:"use_dn"`
	RecursiveLookups                  bool   `xml:"recursive_lookups"`
	MapUserMembershipToGroupField     bool   `xml:"map_user_membership_to_group_field"`
	MapUserMembershipUseDN            bool   `xml:"map_user_membership_use_dn"`
	MapObjectClassToAnyOrAll          string `xml:"map_object_class_to_any_or_all"`
	ObjectClasses                     string `xml:"object_classes"`
	SearchBase                        string `xml:"search_base"`
	SearchScope                       string `xml:"search_scope"`
	Username                          string `xml:"username"`
	GroupID                           string `xml:"group_id"`
	UserGroupMembershipUseLDAPCompare bool   `xml:"user_group_membership_use_ldap_compare"`
}

// CRUD

// GetLDAPServers retrieves a serialized list of LDAP servers.
func (c *Client) GetLDAPServers() (*ResponseLDAPServersList, error) {
	endpoint := uriLDAPServers

	var ldapServers ResponseLDAPServersList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ldapServers)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch LDAP servers: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ldapServers, nil
}

// GetLDAPServerByID retrieves the details of a specific LDAP server by its ID.
func (c *Client) GetLDAPServerByID(id int) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriLDAPServers, id)

	var ldapServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ldapServer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch LDAP server by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ldapServer, nil
}

// GetLDAPServerByName retrieves the details of a specific LDAP server by its name.
func (c *Client) GetLDAPServerByName(name string) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriLDAPServers, name)

	var ldapServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ldapServer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch LDAP server by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ldapServer, nil
}

// GetLDAPServerByIDAndUserDataSubset retrieves information about matching users for a specific LDAP server by its ID.
func (c *Client) GetLDAPServerByIDAndUserDataSubset(id int, user string) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/id/%d/user/%s", uriLDAPServers, id, user)

	var ldapServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ldapServer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch LDAP server by ID and User Data Subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ldapServer, nil
}

// GetLDAPServerByIDAndGroupDataSubset retrieves information about matching groups for a specific LDAP server by its ID.
func (c *Client) GetLDAPServerByIDAndGroupDataSubset(id int, group string) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/id/%d/group/%s", uriLDAPServers, id, group)

	var ldapServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ldapServer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch LDAP server by ID and Group Data Subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ldapServer, nil
}

// GetLDAPServerByIDAndUserMembershipInGroupDataSubset retrieves information about user membership in a group for an LDAP server specified by its ID.
func (c *Client) GetLDAPServerByIDAndUserMembershipInGroupDataSubset(id int, group, user string) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/id/%d/group/%s/user/%s", uriLDAPServers, id, group, user)

	var ldapServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ldapServer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch LDAP server user membership in group data: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ldapServer, nil
}

// GetLDAPServerByNameAndUserDataSubset retrieves information about matching users for a specific LDAP server specified by its name.
func (c *Client) GetLDAPServerByNameAndUserDataSubset(name, user string) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/name/%s/user/%s", uriLDAPServers, name, user)

	var ldapServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ldapServer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch LDAP server by name and user data subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ldapServer, nil
}

// GetLDAPServerByNameAndGroupDataSubset retrieves information about groups for a specific LDAP server specified by its name.
func (c *Client) GetLDAPServerByNameAndGroupDataSubset(name, group string) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/name/%s/group/%s", uriLDAPServers, name, group)

	var ldapServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ldapServer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch LDAP server by name and group data subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ldapServer, nil
}

// GetLDAPServerByNameAndUserMembershipInGroupDataSubset retrieves information about user membership in a group for a specific LDAP server by its name.
func (c *Client) GetLDAPServerByNameAndUserMembershipInGroupDataSubset(name, group, user string) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/name/%s/group/%s/user/%s", uriLDAPServers, name, group, user)

	var ldapServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ldapServer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch LDAP server by name and user membership in group data subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ldapServer, nil
}

// CreateLDAPServer creates a new LDAP server in Jamf Pro.
func (c *Client) CreateLDAPServer(ldapServer *ResourceLDAPServers) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriLDAPServers) // '0' typically used for creation in APIs

	// Wrap ldapServer in an anonymous struct to match the expected XML structure
	requestBody := struct {
		XMLName xml.Name `xml:"ldap_server"`
		*ResourceLDAPServers
	}{
		ResourceLDAPServers: ldapServer,
	}

	var responseLDAPServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseLDAPServer)
	if err != nil {
		return nil, fmt.Errorf("failed to create LDAP server: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseLDAPServer, nil
}

// UpdateLDAPServerByID updates an existing LDAP server identified by its ID.
func (c *Client) UpdateLDAPServerByID(id int, ldapServer *ResourceLDAPServers) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriLDAPServers, id)

	requestBody := struct {
		XMLName xml.Name `xml:"ldap_server"`
		*ResourceLDAPServers
	}{
		ResourceLDAPServers: ldapServer,
	}

	var responseLDAPServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseLDAPServer)
	if err != nil {
		return nil, fmt.Errorf("failed to update LDAP server by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseLDAPServer, nil
}

// UpdateLDAPServerByName updates an existing LDAP server identified by its name.
func (c *Client) UpdateLDAPServerByName(name string, ldapServer *ResourceLDAPServers) (*ResourceLDAPServers, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriLDAPServers, name)

	requestBody := struct {
		XMLName xml.Name `xml:"ldap_server"`
		*ResourceLDAPServers
	}{
		ResourceLDAPServers: ldapServer,
	}

	var responseLDAPServer ResourceLDAPServers
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseLDAPServer)
	if err != nil {
		return nil, fmt.Errorf("failed to update LDAP server by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseLDAPServer, nil
}

// DeleteLDAPServerByID deletes an LDAP server identified by its ID.
func (c *Client) DeleteLDAPServerByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriLDAPServers, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete LDAP server by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteLDAPServerByName deletes an LDAP server identified by its name.
func (c *Client) DeleteLDAPServerByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriLDAPServers, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete LDAP server by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
