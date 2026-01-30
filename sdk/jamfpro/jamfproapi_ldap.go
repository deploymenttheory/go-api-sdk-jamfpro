// jamfproapi_ldap.go
// Jamf Pro Api - LDAP
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-groups
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-servers
// Jamf Pro Api requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	uriLdapGroupsV1  = "/api/v1/ldap/groups"
	uriLdapServersV1 = "/api/v1/ldap/servers"
)

// ResponseLdapGroupSearchResultsV1 models the search payload for LDAP groups.
type ResponseLdapGroupSearchResultsV1 struct {
	TotalCount int                   `json:"totalCount"`
	Results    []ResourceLdapGroupV1 `json:"results"`
}

// ResourceLdapGroupV1 represents a Jamf Pro LDAP group definition.
type ResourceLdapGroupV1 struct {
	ID                string `json:"id"`
	UUID              string `json:"uuid"`
	LdapServerID      int    `json:"ldapServerId"`
	Name              string `json:"name"`
	DistinguishedName string `json:"distinguishedName"`
}

// ResourceLdapServerV1 represents a Jamf Pro LDAP server summary.
type ResourceLdapServerV1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetLdapGroupsV1 retrieves LDAP groups whose names contain the supplied search text.
func (c *Client) GetLdapGroupsV1(search string) (*ResponseLdapGroupSearchResultsV1, error) {
	endpoint := uriLdapGroupsV1
	trimmedSearch := strings.TrimSpace(search)

	if trimmedSearch != "" {
		u, err := url.Parse(endpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to build ldap groups query: %v", err)
		}

		query := u.Query()
		query.Set("q", trimmedSearch)
		u.RawQuery = query.Encode()
		endpoint = u.String()
	}

	var groups ResponseLdapGroupSearchResultsV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &groups)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "ldap groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &groups, nil
}

// GetLdapServersV1 retrieves every active LDAP or cloud identity provider server definition.
func (c *Client) GetLdapServersV1() ([]ResourceLdapServerV1, error) {
	var servers []ResourceLdapServerV1

	resp, err := c.HTTP.DoRequest("GET", uriLdapServersV1, nil, &servers)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "ldap servers", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return servers, nil
}
