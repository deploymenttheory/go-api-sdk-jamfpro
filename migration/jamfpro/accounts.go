// accounts.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAccounts = "/JSSResource/accounts"

// XML structure represented in nested Go structs
// Account Response structure

type ResponseAccount struct {
	ID                  int                         `json:"id,omitempty" xml:"id,omitempty"`
	Name                string                      `json:"name" xml:"name"`
	DirectoryUser       bool                        `json:"directory_user,omitempty" xml:"directory_user,omitempty"`
	FullName            string                      `json:"full_name,omitempty" xml:"full_name,omitempty"`
	Email               string                      `json:"email,omitempty" xml:"email,omitempty"`
	EmailAddress        string                      `json:"email_address,omitempty" xml:"email_address,omitempty"`
	Enabled             string                      `json:"enabled,omitempty" xml:"enabled,omitempty"`
	LdapServer          AccountDataSubsetLdapServer `json:"ldap_server,omitempty" xml:"ldap_server,omitempty"` // Added this
	ForcePasswordChange bool                        `json:"force_password_change,omitempty" xml:"force_password_change,omitempty"`
	AccessLevel         string                      `json:"access_level,omitempty" xml:"access_level,omitempty"`
	Password            string                      `json:"password" xml:"password"`
	PrivilegeSet        string                      `json:"privilege_set,omitempty" xml:"privilege_set,omitempty"`
	Site                AccountDataSubsetSite       `json:"site,omitempty" xml:"site,omitempty"`
	Privileges          AccountDataSubsetPrivileges `json:"privileges,omitempty" xml:"privileges,omitempty"`
}

type AccountDataSubsetLdapServer struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name" xml:"name"`
}
type AccountDataSubsetGroup struct {
	ID         int                         `json:"id,omitempty" xml:"id,omitempty"`
	Name       string                      `json:"name" xml:"name"`
	Site       AccountDataSubsetSite       `json:"site,omitempty" xml:"site,omitempty"`
	Privileges AccountDataSubsetPrivileges `json:"privileges,omitempty" xml:"privileges,omitempty"`
}

type AccountDataSubsetSite struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name" xml:"name"`
}

type AccountDataSubsetPrivileges struct {
	JSSObjects    []string `json:"jss_objects,omitempty" xml:"jss_objects>privilege"`
	JSSSettings   []string `json:"jss_settings,omitempty" xml:"jss_settings>privilege"`
	JSSActions    []string `json:"jss_actions,omitempty" xml:"jss_actions>privilege"`
	Recon         []string `json:"recon,omitempty" xml:"recon>privilege"`
	CasperAdmin   []string `json:"casper_admin,omitempty" xml:"casper_admin>privilege"`
	CasperRemote  []string `json:"casper_remote,omitempty" xml:"casper_remote>privilege"`
	CasperImaging []string `json:"casper_imaging,omitempty" xml:"casper_imaging>privilege"`
}

type AccountDataSubsetUser struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type ResponseAccountsList struct {
	Users  Users  `json:"users,omitempty" xml:"users,omitempty"`
	Groups Groups `json:"groups,omitempty" xml:"groups,omitempty"`
}

type Users struct {
	User []AccountUser `json:"user,omitempty" xml:"user,omitempty"`
}

type Groups struct {
	Group []ResponseAccountGroup `json:"group,omitempty" xml:"group,omitempty"`
}

type AccountUser struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name" xml:"name"`
}

type ResponseAccountGroup struct {
	ID           int                         `json:"id,omitempty" xml:"id"`
	Name         string                      `json:"name" xml:"name"`
	AccessLevel  string                      `json:"access_level" xml:"access_level"`
	PrivilegeSet string                      `json:"privilege_set" xml:"privilege_set"`
	Site         AccountDataSubsetSite       `json:"site" xml:"site"`
	Privileges   AccountDataSubsetPrivileges `json:"privileges" xml:"privileges"`
	Members      []AccountDataSubsetUser     `json:"members" xml:"members>user"`
}

// GetAccountByID retrieves the Account by its ID
func (c *Client) GetAccountByID(id int) (*ResponseAccount, error) {
	url := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	var account ResponseAccount
	if err := c.DoRequest("GET", url, nil, nil, &account); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &account, nil
}

// GetAccountByName retrieves the Account by its name
func (c *Client) GetAccountByName(name string) (*ResponseAccount, error) {
	url := fmt.Sprintf("%s/username/%s", uriAPIAccounts, name)

	var account ResponseAccount
	if err := c.DoRequest("GET", url, nil, nil, &account); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &account, nil
}

// GetAccounts retrieves all user accounts
func (c *Client) GetAccounts() (*ResponseAccountsList, error) {
	url := uriAPIAccounts

	var accountsList ResponseAccountsList
	if err := c.DoRequest("GET", url, nil, nil, &accountsList); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &accountsList, nil
}

// GetAccountGroupByID retrieves the Account Group by its ID
func (c *Client) GetAccountGroupByID(id int) (*ResponseAccountGroup, error) {
	url := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	var accountGroup ResponseAccountGroup
	if err := c.DoRequest("GET", url, nil, nil, &accountGroup); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &accountGroup, nil
}

// GetAccountGroupByName retrieves the Account Group by its name
func (c *Client) GetAccountGroupByName(name string) (*ResponseAccountGroup, error) {
	url := fmt.Sprintf("%s/groupname/%s", uriAPIAccounts, name)

	var accountGroup ResponseAccountGroup
	if err := c.DoRequest("GET", url, nil, nil, &accountGroup); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &accountGroup, nil
}

// CreateAccount creates a new Jamf Pro Account.
func (c *Client) CreateAccount(account *ResponseAccount) (*ResponseAccount, error) {
	url := uriAPIAccounts

	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if account.Site.ID == 0 {
		account.Site = AccountDataSubsetSite{ID: -1, Name: "None"}
	}

	reqBody := &struct {
		XMLName struct{} `xml:"account"`
		*ResponseAccount
	}{
		ResponseAccount: account,
	}

	var responseAccount ResponseAccount
	if err := c.DoRequest("POST", url, reqBody, nil, &responseAccount); err != nil {
		return nil, fmt.Errorf("failed to create Account: %v", err)
	}

	return &responseAccount, nil
}

// CreateAccountGroup creates a new Jamf Pro Account Group.
func (c *Client) CreateAccountGroup(accountGroup *ResponseAccountGroup) (*ResponseAccountGroup, error) {
	url := fmt.Sprintf("%s/groupid/0", uriAPIAccounts)

	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if accountGroup.Site.ID == 0 {
		accountGroup.Site = AccountDataSubsetSite{ID: -1, Name: "None"}
	}

	reqBody := &struct {
		XMLName struct{} `xml:"group"`
		*ResponseAccountGroup
	}{
		ResponseAccountGroup: accountGroup,
	}

	var responseAccountGroup ResponseAccountGroup
	if err := c.DoRequest("POST", url, reqBody, nil, &responseAccountGroup); err != nil {
		return nil, fmt.Errorf("failed to create Account Group: %v", err)
	}

	return &responseAccountGroup, nil
}

// UpdateAccountByID updates an existing Jamf Pro Account by ID
func (c *Client) UpdateAccountByID(id int, account *ResponseAccount) error {
	url := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if account.Site.ID == 0 {
		account.Site = AccountDataSubsetSite{ID: -1, Name: "None"}
	}

	reqBody := &struct {
		XMLName xml.Name `xml:"account"`
		*ResponseAccount
	}{
		ResponseAccount: account,
	}

	if err := c.DoRequest("PUT", url, reqBody, nil, nil); err != nil {
		return fmt.Errorf("failed to update Account by ID: %v", err)
	}

	return nil
}

// UpdateAccountByName updates an existing Jamf Pro Account by Name
func (c *Client) UpdateAccountByName(name string, account *ResponseAccount) error {
	url := fmt.Sprintf("%s/username/%s", uriAPIAccounts, name)

	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if account.Site.ID == 0 {
		account.Site = AccountDataSubsetSite{ID: -1, Name: "None"}
	}

	reqBody := &struct {
		XMLName xml.Name `xml:"account"`
		*ResponseAccount
	}{
		ResponseAccount: account,
	}

	if err := c.DoRequest("PUT", url, reqBody, nil, nil); err != nil {
		return fmt.Errorf("failed to update Account by Name: %v", err)
	}

	return nil
}

// UpdateAccountGroupByID updates an existing Jamf Pro Account Group by ID
func (c *Client) UpdateAccountGroupByID(id int, accountGroup *ResponseAccountGroup) (*ResponseAccountGroup, error) {
	url := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if accountGroup.Site.ID == 0 {
		accountGroup.Site = AccountDataSubsetSite{ID: -1, Name: "None"}
	}

	reqBody := &struct {
		XMLName xml.Name `xml:"group"`
		*ResponseAccountGroup
	}{
		ResponseAccountGroup: accountGroup,
	}

	var responseAccountGroup ResponseAccountGroup
	if err := c.DoRequestDebug("PUT", url, reqBody, nil, &responseAccountGroup); err != nil {
		return nil, fmt.Errorf("failed to update Account Group by ID: %v", err)
	}

	return &responseAccountGroup, nil
}

// UpdateAccountGroupByName updates an existing Jamf Pro Account Group by Name
func (c *Client) UpdateAccountGroupByName(name string, accountGroup *ResponseAccountGroup) (*ResponseAccountGroup, error) {
	url := fmt.Sprintf("%s/groupname/%s", uriAPIAccounts, name)

	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if accountGroup.Site.ID == 0 {
		accountGroup.Site = AccountDataSubsetSite{ID: -1, Name: "None"}
	}

	reqBody := &struct {
		XMLName xml.Name `xml:"group"`
		*ResponseAccountGroup
	}{
		ResponseAccountGroup: accountGroup,
	}

	var responseAccountGroup ResponseAccountGroup
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseAccountGroup); err != nil {
		return nil, fmt.Errorf("failed to update Account Group by Name: %v", err)
	}

	return &responseAccountGroup, nil
}

// DeleteAccountByID deletes an existing Jamf Pro Account by ID
func (c *Client) DeleteAccountByID(id int) error {
	url := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Account by ID: %v", err)
	}

	return nil
}

// DeleteAccountByName deletes an existing Jamf Pro Account by Name
func (c *Client) DeleteAccountByName(name string) error {
	url := fmt.Sprintf("%s/username/%s", uriAPIAccounts, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Account by Name: %v", err)
	}

	return nil
}

// DeleteAccountGroupByID deletes an existing Jamf Pro Account Group by ID
func (c *Client) DeleteAccountGroupByID(id int) error {
	url := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Account Group by ID: %v", err)
	}

	return nil
}

// DeleteAccountGroupByName deletes an existing Jamf Pro Account Group by Name
func (c *Client) DeleteAccountGroupByName(name string) error {
	url := fmt.Sprintf("%s/groupname/%s", uriAPIAccounts, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Account Group by Name: %v", err)
	}

	return nil
}
