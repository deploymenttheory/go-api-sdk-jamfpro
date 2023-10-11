// ldapServers.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
)

const uriLdapServers = "/JSSResource/ldapservers"

type ResponseLdapServer struct {
	Connection       LdapServerDataSubsetLdapConnection       `json:"connection" xml:"connection"`
	MappingsForUsers LdapServerDataSubsetLdapMappingsForUsers `json:"mappings_for_users" xml:"mappings_for_users"`
}

type LdapServerDataSubsetLdapConnection struct {
	ID                  int                             `json:"id,omitempty" xml:"id,omitempty"`
	Name                string                          `json:"name" xml:"name"`
	Hostname            string                          `json:"hostname" xml:"hostname"`
	CertificateUsed     string                          `json:"certificate_used" xml:"certificate_used"`
	ConnectionIsUsedFor string                          `json:"connection_is_used_for" xml:"connection_is_used_for"`
	ServerType          string                          `json:"server_type" xml:"server_type"`
	Port                int                             `json:"port" xml:"port"`
	UseSSL              bool                            `json:"use_ssl" xml:"use_ssl"`
	AuthenticationType  string                          `json:"authentication_type" xml:"authentication_type"`
	Account             LdapServerDataSubsetLdapAccount `json:"account" xml:"account"`
	OpenCloseTimeout    int                             `json:"open_close_timeout" xml:"open_close_timeout"`
	SearchTimeout       int                             `json:"search_timeout" xml:"search_timeout"`
	ReferralResponse    string                          `json:"referral_response" xml:"referral_response"`
	UseWildcards        bool                            `json:"use_wildcards" xml:"use_wildcards"`
}

type LdapServerDataSubsetLdapAccount struct {
	DistinguishedUsername string `json:"distinguished_username" xml:"distinguished_username"`
	Password              string `json:"password" xml:"password"`
}

type LdapServerDataSubsetLdapMappingsForUsers struct {
	UserMappings                LdapServerDataSubsetLdapUserMappings                `json:"user_mappings" xml:"user_mappings"`
	UserGroupMappings           LdapServerDataSubsetLdapUserGroupMappings           `json:"user_group_mappings" xml:"user_group_mappings"`
	UserGroupMembershipMappings LdapServerDataSubsetLdapUserGroupMembershipMappings `json:"user_group_membership_mappings" xml:"user_group_membership_mappings"`
}

type LdapServerDataSubsetLdapUserMappings struct {
	MapObjectClassToAnyOrAll string `json:"map_object_class_to_any_or_all" xml:"map_object_class_to_any_or_all"`
	ObjectClasses            string `json:"object_classes" xml:"object_classes"`
	SearchBase               string `json:"search_base" xml:"search_base"`
	SearchScope              string `json:"search_scope" xml:"search_scope"`
	MapUserId                string `json:"map_user_id" xml:"map_user_id"`
	MapUsername              string `json:"map_username" xml:"map_username"`
	MapRealname              string `json:"map_realname" xml:"map_realname"`
	MapEmailAddress          string `json:"map_email_address" xml:"map_email_address"`
	AppendToEmailResults     string `json:"append_to_email_results" xml:"append_to_email_results"`
	MapDepartment            string `json:"map_department" xml:"map_department"`
	MapBuilding              string `json:"map_building" xml:"map_building"`
	MapRoom                  string `json:"map_room" xml:"map_room"`
	MapTelephone             string `json:"map_telephone" xml:"map_telephone"`
	MapPosition              string `json:"map_position" xml:"map_position"`
	MapUserUuid              string `json:"map_user_uuid" xml:"map_user_uuid"`
}

type LdapServerDataSubsetLdapUserGroupMappings struct {
	MapObjectClassToAnyOrAll          string `json:"map_object_class_to_any_or_all" xml:"map_object_class_to_any_or_all"`
	ObjectClasses                     string `json:"object_classes" xml:"object_classes"`
	MapUserMembershipToGroupField     bool   `json:"map_user_membership_to_group_field" xml:"map_user_membership_to_group_field"`
	MapUserMembershipUseDn            bool   `json:"map_user_membership_use_dn" xml:"map_user_membership_use_dn"`
	UserGroupMembershipUseLdapCompare bool   `json:"user_group_membership_use_ldap_compare" xml:"user_group_membership_use_ldap_compare"`
	SearchBase                        string `json:"search_base" xml:"search_base"`
	SearchScope                       string `json:"search_scope" xml:"search_scope"`
	MapGroupId                        string `json:"map_group_id" xml:"map_group_id"`
	MapGroupName                      string `json:"map_group_name" xml:"map_group_name"`
	MapGroupUuid                      string `json:"map_group_uuid" xml:"map_group_uuid"`
}

type LdapServerDataSubsetLdapUserGroupMembershipMappings struct {
	UserGroupMembershipStoredIn       string `json:"user_group_membership_stored_in" xml:"user_group_membership_stored_in"`
	MapGroupMembershipToUserField     string `json:"map_group_membership_to_user_field" xml:"map_group_membership_to_user_field"`
	AppendToUsername                  string `json:"append_to_username" xml:"append_to_username"`
	UseDn                             bool   `json:"use_dn" xml:"use_dn"`
	RecursiveLookups                  bool   `json:"recursive_lookups" xml:"recursive_lookups"`
	MapUserMembershipToGroupField     bool   `json:"map_user_membership_to_group_field" xml:"map_user_membership_to_group_field"`
	MapUserMembershipUseDn            bool   `json:"map_user_membership_use_dn" xml:"map_user_membership_use_dn"`
	UserGroupMembershipUseLdapCompare bool   `json:"user_group_membership_use_ldap_compare" xml:"user_group_membership_use_ldap_compare"`
	Username                          string `json:"username" xml:"username"`
	GroupId                           string `json:"group_id" xml:"group_id"`
	MapObjectClassToAnyOrAll          string `json:"map_object_class_to_any_or_all" xml:"map_object_class_to_any_or_all"`
	ObjectClasses                     string `json:"object_classes" xml:"object_classes"`
	SearchBase                        string `json:"search_base" xml:"search_base"`
	SearchScope                       string `json:"search_scope" xml:"search_scope"`
}

type ResponseLdapServerList struct {
	Servers []LdapServerListItem `json:"ldap_server" xml:"ldap_server"`
}

type LdapServerListItem struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// GetLdapServerByID retrieves the LDAP server configuration by its ID.
func (c *Client) GetLdapServerByID(id int) (*ResponseLdapServer, error) {
	url := fmt.Sprintf("%s/id/%d", uriLdapServers, id)

	var server ResponseLdapServer
	if err := c.DoRequest("GET", url, nil, nil, &server); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &server, nil
}

// GetLdapServerByName retrieves the LDAP server configuration by its Name.
func (c *Client) GetLdapServerByName(name string) (*ResponseLdapServer, error) {
	url := fmt.Sprintf("%s/name/%s", uriLdapServers, name)

	var server ResponseLdapServer
	if err := c.DoRequest("GET", url, nil, nil, &server); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &server, nil
}

// GetLdapServers retrieves all LDAP server configurations.
func (c *Client) GetLdapServers() (*ResponseLdapServerList, error) {
	url := uriLdapServers

	var servers ResponseLdapServerList
	if err := c.DoRequest("GET", url, nil, nil, &servers); err != nil {
		return nil, fmt.Errorf("failed to fetch all LDAP servers: %v", err)
	}

	return &servers, nil
}

// CreateLdapServer creates a new LDAP server configuration.
func (c *Client) CreateLdapServer(server *ResponseLdapServer) (*ResponseLdapServer, error) {
	url := fmt.Sprintf("%s/id/0", uriLdapServers)

	reqBody := &struct {
		XMLName struct{} `xml:"ldap_server"`
		*ResponseLdapServer
	}{
		ResponseLdapServer: server,
	}

	var responseServer ResponseLdapServer
	if err := c.DoRequest("POST", url, reqBody, nil, &responseServer); err != nil {
		return nil, fmt.Errorf("failed to create LDAP server: %v", err)
	}

	return &responseServer, nil
}

// UpdateLdapServerById updates an existing LDAP server configuration by its ID.
func (c *Client) UpdateLdapServerById(id int, server *ResponseLdapServer) (*ResponseLdapServer, error) {
	url := fmt.Sprintf("%s/id/%d", uriLdapServers, id)

	reqBody := &struct {
		XMLName struct{} `xml:"ldap_server"`
		*ResponseLdapServer
	}{
		ResponseLdapServer: server,
	}

	var responseServer ResponseLdapServer
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseServer); err != nil {
		return nil, fmt.Errorf("failed to update LDAP server by ID: %v", err)
	}

	return &responseServer, nil
}

// UpdateLdapServerByName updates an existing LDAP server configuration by its name.
func (c *Client) UpdateLdapServerByName(name string, server *ResponseLdapServer) (*ResponseLdapServer, error) {
	url := fmt.Sprintf("%s/name/%s", uriLdapServers, name)

	reqBody := &struct {
		XMLName struct{} `xml:"ldap_server"`
		*ResponseLdapServer
	}{
		ResponseLdapServer: server,
	}

	var responseServer ResponseLdapServer
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseServer); err != nil {
		return nil, fmt.Errorf("failed to update LDAP server by name: %v", err)
	}

	return &responseServer, nil
}

// DeleteLdapServerById deletes an existing LDAP server configuration by its ID.
func (c *Client) DeleteLdapServerById(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriLdapServers, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete LDAP server by ID: %v", err)
	}

	return nil
}

// DeleteLdapServerByName deletes an existing LDAP server configuration by its name.
func (c *Client) DeleteLdapServerByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriLdapServers, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete LDAP server by name: %v", err)
	}

	return nil
}
