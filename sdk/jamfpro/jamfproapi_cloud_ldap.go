// jamfproapi_cloud_ldap.go
// Jamf Pro Api - Cloud LDAP
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v2-cloud-ldaps
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

// TODO - Figure out if we need this.

const uriCloudLdaps = "/api/v2/cloud-ldaps"

// Responses

type ResponseCloudIdentityProviderLdapCreated struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

type ResponseCloudIdentityProviderDefaultMappings struct {
	CloudIdentityProviderDefaultMappingsSubsetUserMappings       CloudIdentityProviderDefaultMappingsSubsetUserMappings       `json:"userMappings"`
	CloudIdentityProviderDefaultMappingsSubsetGroupMappings      CloudIdentityProviderDefaultMappingsSubsetGroupMappings      `json:"groupMappings"`
	CloudIdentityProviderDefaultMappingsSubsetMembershipMappings CloudIdentityProviderDefaultMappingsSubsetMembershipMappings `json:"membershipMappings"`
}

// Subsets & Containers

type CloudIdentityProviderDefaultMappingsSubsetUserMappings struct {
	ObjectClassLimitation string `json:"objectClassLimitation"`
	ObjectClasses         string `json:"objectClasses"`
	SearchBase            string `json:"searchBase"`
	SearchScope           string `json:"searchScope"`
	AdditionalSearchBase  string `json:"additionalSearchBase"`
	UserID                string `json:"userID"`
	Username              string `json:"username"`
	RealName              string `json:"realName"`
	EmailAddress          string `json:"emailAddress"`
	Department            string `json:"department"`
	Building              string `json:"building"`
	Room                  string `json:"room"`
	Phone                 string `json:"phone"`
	Position              string `json:"position"`
	UserUuid              string `json:"userUuid"`
}

type CloudIdentityProviderDefaultMappingsSubsetGroupMappings struct {
	ObjectClassLimitation string `json:"objectClassLimitation"`
	ObjectClasses         string `json:"objectClasses"`
	SearchBase            string `json:"searchBase"`
	SearchScope           string `json:"searchScope"`
	GroupID               string `json:"groupID"`
	GroupName             string `json:"groupName"`
	GroupUuid             string `json:"groupUuid"`
}

type CloudIdentityProviderDefaultMappingsSubsetMembershipMappings struct {
	GroupMembershipMapping string `json:"groupMembershipMapping"`
}

type ResourceCloudLdap struct {
	CloudIdPCommon *CloudIdPCommon    `json:"cloudIdPCommon"`
	Server         *CloudLdapServer   `json:"server"`
	Mappings       *CloudLdapMappings `json:"mappings,omitempty"`
}

type CloudIdPCommon struct {
	ID           string `json:"id,omitempty"`
	ProviderName string `json:"providerName"` // GOOGLE or AZURE
	DisplayName  string `json:"displayName"`
}

type CloudLdapServer struct {
	Enabled                                  bool               `json:"enabled"`
	Keystore                                 *CloudLdapKeystore `json:"keystore"`
	UseWildcards                             bool               `json:"useWildcards"`
	ConnectionType                           string             `json:"connectionType"` // LDAPS or START_TLS
	ServerUrl                                string             `json:"serverUrl"`
	DomainName                               string             `json:"domainName"`
	Port                                     int                `json:"port"`
	ConnectionTimeout                        int                `json:"connectionTimeout"`
	SearchTimeout                            int                `json:"searchTimeout"`
	MembershipCalculationOptimizationEnabled bool               `json:"membershipCalculationOptimizationEnabled,omitempty"`
}

type CloudLdapKeystore struct {
	Type           string `json:"type,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	Subject        string `json:"subject,omitempty"`
	FileName       string `json:"fileName,omitempty"`
	Password       string `json:"password,omitempty"`
	FileBytes      string `json:"fileBytes,omitempty"`
}

type CloudLdapMappings struct {
	UserMappings       CloudIdentityProviderDefaultMappingsSubsetUserMappings       `json:"userMappings"`
	GroupMappings      CloudIdentityProviderDefaultMappingsSubsetGroupMappings      `json:"groupMappings"`
	MembershipMappings CloudIdentityProviderDefaultMappingsSubsetMembershipMappings `json:"membershipMappings"`
}

// CRUD

func (c *Client) GetDefaultCloudIdentityProviderDefaultMappings(providerName string) (*ResponseCloudIdentityProviderDefaultMappings, error) {
	endpoint := fmt.Sprintf("%s/%s/mappings", uriCloudLdaps, providerName)
	var out ResponseCloudIdentityProviderDefaultMappings

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "cloud identity provider default mappings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}

// CreateCloudIdentityProviderLdap creates a new Cloud Identity Provider configuration
func (c *Client) CreateCloudIdentityProviderLdap(config *ResourceCloudLdap) (*ResponseCloudIdentityProviderLdapCreated, error) {
	endpoint := uriCloudLdaps

	var response ResponseCloudIdentityProviderLdapCreated
	resp, err := c.HTTP.DoRequest("POST", endpoint, config, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "cloud identity provider", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetCloudIdentityProviderLdapByID retrieves a specific Cloud Identity Provider LDAP configuration by ID
func (c *Client) GetCloudIdentityProviderLdapByID(id string) (*ResourceCloudLdap, error) {
	endpoint := fmt.Sprintf("%s/%s", uriCloudLdaps, id)

	var out ResourceCloudLdap
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "cloud identity provider LDAP", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateCloudIdentityProviderLdap updates an existing Cloud Identity Provider LDAP configuration
func (c *Client) UpdateCloudIdentityProviderLdap(id string, config *ResourceCloudLdap) (*ResourceCloudLdap, error) {
	endpoint := fmt.Sprintf("%s/%s", uriCloudLdaps, id)

	var out ResourceCloudLdap
	resp, err := c.HTTP.DoRequest("PUT", endpoint, config, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "cloud identity provider LDAP", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// DeleteCloudIdentityProviderLdapByID deletes a Cloud Identity Provider LDAP configuration by ID
func (c *Client) DeleteCloudIdentityProviderLdapByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriCloudLdaps, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDelete, "cloud identity provider LDAP", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
