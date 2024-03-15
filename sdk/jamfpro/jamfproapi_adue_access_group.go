// jamfproapi_adue_access_group.go
// Jamf Pro Api - Enrollment
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-access-groups
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriAccountDrivenUserEnrollment = "/api/v3/enrollment"

// List

// ResponseAccountDrivenUserEnrollmentAccessGroups represents the structure of the response for a list of access groups

type ResponseAccountDrivenUserEnrollmentAccessGroupsList struct {
	TotalCount int                                              `json:"totalCount"`
	Results    []ResourceAccountDrivenUserEnrollmentAccessGroup `json:"results"`
}

type ResponseAccountDrivenUserEnrollmentAccessGroupCreateAndUpdate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

type ResourceAccountDrivenUserEnrollmentAccessGroup struct {
	ID                                 string `json:"id"`
	GroupID                            string `json:"groupId"`
	LdapServerID                       string `json:"ldapServerId"`
	Name                               string `json:"name"`
	SiteID                             string `json:"siteId"`
	EnterpriseEnrollmentEnabled        bool   `json:"enterpriseEnrollmentEnabled"`
	PersonalEnrollmentEnabled          bool   `json:"personalEnrollmentEnabled"`
	AccountDrivenUserEnrollmentEnabled bool   `json:"accountDrivenUserEnrollmentEnabled"`
	RequireEula                        bool   `json:"requireEula"`
}

// CRUD

// GetAccountDrivenUserEnrollmentAccessGroups fetches all ADUE access groups
func (c *Client) GetAccountDrivenUserEnrollmentAccessGroups(sort_filter string) (*ResponseAccountDrivenUserEnrollmentAccessGroupsList, error) {
	endpoint := uriAccountDrivenUserEnrollment
	resp, err := c.DoPaginatedGet(endpoint, standardPageSize, 0, sort_filter)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "ADUE Access Group List", err)
	}

	var OutStruct ResponseAccountDrivenUserEnrollmentAccessGroupsList
	OutStruct.TotalCount = resp.Size
	for _, value := range resp.Results {
		var newObj ResourceAccountDrivenUserEnrollmentAccessGroup
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "ADUE Access Group List", err)
		}
		OutStruct.Results = append(OutStruct.Results, newObj)
	}

	return &OutStruct, nil
}

// Retrieves AccountDrivenUserEnrollmentAccessGroup from provided ID & returns ResourceAccountDrivenUserEnrollmentAccessGroup
func (c *Client) GetAccountDrivenUserEnrollmentAccessGroupByID(id string) (*ResourceAccountDrivenUserEnrollmentAccessGroup, error) {
	endpoint := fmt.Sprintf("%s/access-groups/%s", uriAccountDrivenUserEnrollment, id)

	var ADUEGroup ResourceAccountDrivenUserEnrollmentAccessGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ADUEGroup)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "ADUE Access Group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ADUEGroup, nil
}

// GetAccountDrivenUserEnrollmentAccessGroupByName retrieves an Account Driven User Enrollment Access Group by its name
func (c *Client) GetAccountDrivenUserEnrollmentAccessGroupByName(name string) (*ResourceAccountDrivenUserEnrollmentAccessGroup, error) {
	accessGroupsList, err := c.GetAccountDrivenUserEnrollmentAccessGroups("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "ADUE access group", err)
	}

	for _, group := range accessGroupsList.Results {
		if group.Name == name {
			return &group, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "ADUE access group", name, errMsgNoName)
}

// Creates Account Driven User Enrollment Access Group from ResourceScript struct
func (c *Client) CreateAccountDrivenUserEnrollmentAccessGroup(script *ResourceAccountDrivenUserEnrollmentAccessGroup) (*ResponseAccountDrivenUserEnrollmentAccessGroupCreateAndUpdate, error) {
	endpoint := uriScripts
	var out ResponseAccountDrivenUserEnrollmentAccessGroupCreateAndUpdate

	resp, err := c.HTTP.DoRequest("POST", endpoint, script, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "ADUE access group", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateAccountDrivenUserEnrollmentAccessGroupByID updates an ADUE access group by resource ID
func (c *Client) UpdateAccountDrivenUserEnrollmentAccessGroupByID(id string, groupUpdate *ResourceAccountDrivenUserEnrollmentAccessGroup) (*ResourceAccountDrivenUserEnrollmentAccessGroup, error) {
	endpoint := fmt.Sprintf("%s/access-groups/%s", uriAccountDrivenUserEnrollment, id)
	var out ResourceAccountDrivenUserEnrollmentAccessGroup

	resp, err := c.HTTP.DoRequest("PUT", endpoint, groupUpdate, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "ADUE Access Group", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateAccountDrivenUserEnrollmentAccessGroupByName updates an ADUE access group by resource name
func (c *Client) UpdateAccountDrivenUserEnrollmentAccessGroupByName(targetName string, groupUpdate *ResourceAccountDrivenUserEnrollmentAccessGroup) (*ResourceAccountDrivenUserEnrollmentAccessGroup, error) {
	target, err := c.GetAccountDrivenUserEnrollmentAccessGroupByName(targetName)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "ADUE access group", targetName, err)
	}

	resp, err := c.UpdateAccountDrivenUserEnrollmentAccessGroupByID(target.ID, groupUpdate)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "ADUE access group", targetName, err)
	}

	return resp, nil
}

// DeleteAccountDrivenUserEnrollmentAccessGroupByID deletes an ADUE access group with given id
func (c *Client) DeleteAccountDrivenUserEnrollmentAccessGroupByID(id string) error {
	endpoint := fmt.Sprintf("%s/access-groups/%s", uriAccountDrivenUserEnrollment, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)

	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDeleteByID, "ADUE access group", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAccountDrivenUserEnrollmentAccessGroupByName deletes an ADUE access group with given name, leverages GetAccountDrivenUserEnrollmentAccessGroupByName
func (c *Client) DeleteAccountDrivenUserEnrollmentAccessGroupByName(targetName string) error {
	target, err := c.GetAccountDrivenUserEnrollmentAccessGroupByName(targetName)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "ADUE access group", targetName, err)
	}

	err = c.DeleteAccountDrivenUserEnrollmentAccessGroupByID(target.ID)

	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "ADUE access group", targetName, err)
	}

	return nil
}
