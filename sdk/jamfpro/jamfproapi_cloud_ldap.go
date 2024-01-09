// jamfproapi_cloud_ldap.go
// Jamf Pro Api - Cloud LDAP
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v2-cloud-ldaps
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

// TODO - Figure out if we need this.

const uriCloudLdaps = "/api/v2/cloud-ldaps"

// Responses

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
	GroupMembershipMapping string `json:"memberOf"`
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
