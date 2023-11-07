// classicapi_accounts.go
// Jamf Pro Classic Api - Accounts
// api reference: https://developer.jamf.com/jamf-pro/reference/accounts
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"fmt"
)

const uriAPIAccounts = "/JSSResource/accounts"

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

// Get Account List
type ResponseAccountsList struct {
	Users  *AccountDataSubsetUsers  `xml:"users,omitempty"`
	Groups *AccountDataSubsetGroups `xml:"groups,omitempty"`
}

type AccountDataSubsetUsers struct {
	User []AccountDataSubsetUserItem `xml:"user,omitempty"`
}

type AccountDataSubsetUserItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type AccountDataSubsetGroups struct {
	Group []AccountDataSubsetGroup `xml:"group,omitempty"`
}

// GetAccounts retrieves a list of all accounts (both users and groups).
func (c *Client) GetAccounts() (*ResponseAccountsList, error) {
	endpoint := fmt.Sprintf("%s", uriAPIAccounts)

	var accountsList ResponseAccountsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &accountsList)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &accountsList, nil
}

// GetAccountByID retrieves the Account by its ID
func (c *Client) GetAccountByID(id int) (*ResponseAccount, error) {
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	var account ResponseAccount
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &account)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &account, nil
}

// GetAccountByName retrieves the Account by its name
func (c *Client) GetAccountByName(name string) (*ResponseAccount, error) {
	endpoint := fmt.Sprintf("%s/username/%s", uriAPIAccounts, name)

	var account ResponseAccount
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &account)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &account, nil
}

// GetAccountGroupByID gets an account group using its ID and returns a response.
func (c *Client) GetAccountGroupByID(id int) (*ResponseAccountGroup, error) {
	endpoint := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	var group ResponseAccountGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)
	if err != nil {
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// GetAccountByName retrieves the Account by its name
func (c *Client) GetAccountGroupByName(name string) (*ResponseAccount, error) {
	endpoint := fmt.Sprintf("%s/groupname/%s", uriAPIAccounts, name)

	var account ResponseAccount
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &account)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &account, nil
}

// CreateAccountByID creates an Account using its ID
func (c *Client) CreateAccountByID(account *ResponseAccount) (*ResponseAccount, error) {
	// Use a placeholder ID for creating a new account
	placeholderID := 0
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, placeholderID)

	// Check if site is not provided and set default values
	if account.Site.ID == 0 && account.Site.Name == "" {
		account.Site = AccountDataSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Define XML requestBody structure
	requestBody := &struct {
		XMLName struct{} `xml:"account"`
		*ResponseAccount
	}{
		ResponseAccount: account,
	}

	var returnedAccount ResponseAccount
	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, &returnedAccount)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &returnedAccount, nil
}

// CreateAccountGroupByID creates an Account Group using its ID
func (c *Client) CreateAccountGroupByID(accountGroup *ResponseAccountGroup) (*ResponseAccountGroup, error) {
	// Use a placeholder ID for creating a new account group
	placeholderID := 0
	endpoint := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, placeholderID)

	// Check if site is not provided and set default values
	if accountGroup.Site.ID == 0 && accountGroup.Site.Name == "" {
		accountGroup.Site = AccountDataSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Define XML requestBody structure
	requestBody := &struct {
		XMLName struct{} `xml:"group"`
		*ResponseAccountGroup
	}{
		ResponseAccountGroup: accountGroup,
	}

	var returnedAccountGroup ResponseAccountGroup
	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, &returnedAccountGroup)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &returnedAccountGroup, nil
}

// UpdateAccountByID updates an Account using its ID
func (c *Client) UpdateAccountByID(id int, account *ResponseAccount) (*ResponseAccount, error) {
	// Construct the endpoint URL using the provided account ID
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	// Check if site is not provided and set default values
	if account.Site.ID == 0 && account.Site.Name == "" {
		account.Site = AccountDataSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Define XML requestBody structure
	requestBody := &struct {
		XMLName struct{} `xml:"account"`
		*ResponseAccount
	}{
		ResponseAccount: account,
	}

	var updatedAccount ResponseAccount
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedAccount)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAccount, nil
}

// UpdateAccountByName updates an Account using its name.
func (c *Client) UpdateAccountByName(name string, account *ResponseAccount) (*ResponseAccount, error) {
	// Construct the endpoint URL using the provided account name
	endpoint := fmt.Sprintf("%s/username/%s", uriAPIAccounts, name)

	// Check if site is not provided and set default values
	if account.Site.ID == 0 && account.Site.Name == "" {
		account.Site = AccountDataSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Define XML requestBody structure
	requestBody := &struct {
		XMLName struct{} `xml:"account"`
		*ResponseAccount
	}{
		ResponseAccount: account,
	}

	var updatedAccount ResponseAccount
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedAccount)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAccount, nil
}

// UpdateAccountGroupByID updates an Account Group using its ID
func (c *Client) UpdateAccountGroupByID(id int, group *ResponseAccountGroup) (*ResponseAccountGroup, error) {
	// Construct the endpoint URL using the provided group ID
	endpoint := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	// Check if site is not provided and set default values
	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site = AccountDataSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Define XML requestBody structure
	requestBody := &struct {
		XMLName struct{} `xml:"group"`
		*ResponseAccountGroup
	}{
		ResponseAccountGroup: group,
	}

	var updatedGroup ResponseAccountGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedGroup)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// UpdateAccountGroupByName updates an Account Group using its name.
func (c *Client) UpdateAccountGroupByName(name string, group *ResponseAccountGroup) (*ResponseAccountGroup, error) {
	// Construct the endpoint URL using the provided group name
	endpoint := fmt.Sprintf("%s/groupname/%s", uriAPIAccounts, name)

	// Check if site is not provided and set default values
	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site = AccountDataSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Define XML requestBody structure
	requestBody := &struct {
		XMLName struct{} `xml:"group"`
		*ResponseAccountGroup
	}{
		ResponseAccountGroup: group,
	}

	var updatedGroup ResponseAccountGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedGroup)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// DeleteAccountByID deletes an Account using its ID
func (c *Client) DeleteAccountByID(id int) error {
	// Construct the endpoint URL using the provided account ID
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAccountByName deletes an Account using its name.
func (c *Client) DeleteAccountByName(name string) error {
	// Construct the endpoint URL using the provided account name
	endpoint := fmt.Sprintf("%s/username/%s", uriAPIAccounts, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAccountGroupByID deletes an Account Group using its ID.
func (c *Client) DeleteAccountGroupByID(id int) error {
	// Construct the endpoint URL using the provided group ID
	endpoint := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAccountGroupByName deletes an Account Group using its name.
func (c *Client) DeleteAccountGroupByName(name string) error {
	// Construct the endpoint URL using the provided group name
	endpoint := fmt.Sprintf("%s/groupname/%s", uriAPIAccounts, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		fmt.Printf("Failed to execute request", "Error", err)
		return err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
