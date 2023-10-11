// apiRolePrivileges.go
// Jamf Pro Api
// Jamf Pro Api requires the structs to support JSON.

package jamfpro

import (
	"fmt"
)

const uriAPIRolePrivileges = "/api/v1/api-role-privileges"

type ResponseAPIRolePrivileges struct {
	Privileges []string `json:"privileges"`
}

func (c *Client) GetApiRolePrivileges() (*ResponseAPIRolePrivileges, error) {
	uri := uriAPIRolePrivileges

	var out ResponseAPIRolePrivileges
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get API role privileges: %v", err)
	}

	return &out, nil
}

func (c *Client) SearchApiRolePrivileges(name string, limit int) (*ResponseAPIRolePrivileges, error) {
	uri := fmt.Sprintf("%s/search?name=%s&limit=%d", uriAPIRolePrivileges, name, limit)

	var out ResponseAPIRolePrivileges
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to search API role privileges: %v", err)
	}

	return &out, nil
}
