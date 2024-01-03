// classicapi_account_groups.go
// Jamf Pro Classic Api - Account Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/accounts
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
*/

package jamfpro

import "fmt"

// Resource

type ResourceAccountGroup struct {
	ID           int                       `json:"id,omitempty" xml:"id"`
	Name         string                    `json:"name" xml:"name"`
	AccessLevel  string                    `json:"access_level" xml:"access_level"`
	PrivilegeSet string                    `json:"privilege_set" xml:"privilege_set"`
	Site         SharedResourceSite        `json:"site" xml:"site"`
	Privileges   AccountSubsetPrivileges   `json:"privileges" xml:"privileges"`
	Members      AccountGroupSubsetMembers `json:"members" xml:"members>user"`
}

// Responses

type ResponseAccountGroupCreated struct {
	ID int `json:"id,omitempty" xml:"id,omitempty"`
}

// Subsets

type AccountGroupSubsetMembers []struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// CRUD

// GetAccountGroupByID gets an account group using its ID and returns a response.
func (c *Client) GetAccountGroupByID(id int) (*ResourceAccountGroup, error) {
	endpoint := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	var group ResourceAccountGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "account group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// GetAccountByName retrieves the Account by its name
func (c *Client) GetAccountGroupByName(name string) (*ResourceAccount, error) {
	endpoint := fmt.Sprintf("%s/groupname/%s", uriAPIAccounts, name)

	var account ResourceAccount
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &account)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "account group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &account, nil
}

// CreateAccountGroupByID creates an Account Group using its ID
func (c *Client) CreateAccountGroup(accountGroup *ResourceAccountGroup) (*ResponseAccountGroupCreated, error) {
	// Use a placeholder ID for creating a new account group
	placeholderID := 0
	endpoint := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, placeholderID)

	// Check if site is not provided and set default values
	if accountGroup.Site.ID == 0 && accountGroup.Site.Name == "" {
		accountGroup.Site = SharedResourceSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Define XML requestBody structure
	requestBody := &struct {
		XMLName struct{} `xml:"group"`
		*ResourceAccountGroup
	}{
		ResourceAccountGroup: accountGroup,
	}

	var returnedAccountGroup ResponseAccountGroupCreated
	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, &returnedAccountGroup)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "account group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &returnedAccountGroup, nil
}

// UpdateAccountGroupByID updates an Account Group using its ID
func (c *Client) UpdateAccountGroupByID(id int, accountGroup *ResourceAccountGroup) (*ResourceAccountGroup, error) {
	endpoint := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	if accountGroup.Site.ID == 0 && accountGroup.Site.Name == "" {
		accountGroup.Site = SharedResourceSite{
			ID:   -1,
			Name: "None",
		}
	}

	requestBody := &struct {
		XMLName struct{} `xml:"group"`
		*ResourceAccountGroup
	}{
		ResourceAccountGroup: accountGroup,
	}

	var updatedGroup ResourceAccountGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedGroup)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "account group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// UpdateAccountGroupByName updates an Account Group using its name.
func (c *Client) UpdateAccountGroupByName(name string, accountGroup *ResourceAccountGroup) (*ResourceAccountGroup, error) {
	endpoint := fmt.Sprintf("%s/groupname/%s", uriAPIAccounts, name)

	if accountGroup.Site.ID == 0 && accountGroup.Site.Name == "" {
		accountGroup.Site = SharedResourceSite{
			ID:   -1,
			Name: "None",
		}
	}

	requestBody := &struct {
		XMLName struct{} `xml:"group"`
		*ResourceAccountGroup
	}{
		ResourceAccountGroup: accountGroup,
	}

	var updatedGroup ResourceAccountGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedGroup)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "account group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// DeleteAccountGroupByID deletes an Account Group using its ID.
func (c *Client) DeleteAccountGroupByID(id int) error {
	endpoint := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "account group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAccountGroupByName deletes an Account Group using its name.
func (c *Client) DeleteAccountGroupByName(name string) error {
	endpoint := fmt.Sprintf("%s/groupname/%s", uriAPIAccounts, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "account group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
