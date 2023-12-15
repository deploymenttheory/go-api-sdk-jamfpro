// classicapi_accounts.go
// Jamf Pro Classic Api - Accounts
// api reference: https://developer.jamf.com/jamf-pro/reference/accounts
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"fmt"
)

const uriAPIAccounts = "/JSSResource/accounts"

// ResponseAccountsList represents a serialized list of accounts
type ResponseAccountsList struct {
	Users  *AccountsListSubsetUsers  `xml:"users,omitempty"`
	Groups *AccountsListSubsetGroups `xml:"groups,omitempty"`
}

type AccountsListSubsetUsers struct {
	User []AccountsListSubsetUserItem `xml:"user,omitempty"`
}

type AccountsListSubsetGroups struct {
	Group []AccountsListSubsetGroupItem `xml:"group,omitempty"`
}

type AccountsListSubsetUserItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type AccountsListSubsetGroupItem struct {
	ID         int                     `json:"id,omitempty" xml:"id,omitempty"`
	Name       string                  `json:"name" xml:"name"`
	Site       AccountSubsetSite       `json:"site,omitempty" xml:"site,omitempty"`
	Privileges AccountSubsetPrivileges `json:"privileges,omitempty" xml:"privileges,omitempty"`
}

// ResponseAccount represents an account object
type ResourceAccount struct {
	ID                  int                     `json:"id,omitempty" xml:"id,omitempty"`
	Name                string                  `json:"name" xml:"name"`
	DirectoryUser       bool                    `json:"directory_user,omitempty" xml:"directory_user,omitempty"`
	FullName            string                  `json:"full_name,omitempty" xml:"full_name,omitempty"`
	Email               string                  `json:"email,omitempty" xml:"email,omitempty"`
	EmailAddress        string                  `json:"email_address,omitempty" xml:"email_address,omitempty"`
	Enabled             string                  `json:"enabled,omitempty" xml:"enabled,omitempty"`
	LdapServer          AccountSubsetLdapServer `json:"ldap_server,omitempty" xml:"ldap_server,omitempty"` // Added this
	ForcePasswordChange bool                    `json:"force_password_change,omitempty" xml:"force_password_change,omitempty"`
	AccessLevel         string                  `json:"access_level,omitempty" xml:"access_level,omitempty"`
	Password            string                  `json:"password" xml:"password"`
	PrivilegeSet        string                  `json:"privilege_set,omitempty" xml:"privilege_set,omitempty"`
	Site                AccountSubsetSite       `json:"site,omitempty" xml:"site,omitempty"`
	Privileges          AccountSubsetPrivileges `json:"privileges,omitempty" xml:"privileges,omitempty"`
}

type AccountSubsetLdapServer struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name" xml:"name"`
}

type AccountSubsetSite struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name" xml:"name"`
}

type AccountSubsetPrivileges struct {
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

// type Users struct {
// 	User []AccountUser `json:"user,omitempty" xml:"user,omitempty"`
// }

// type Groups struct {
// 	Group []ResponseAccountGroup `json:"group,omitempty" xml:"group,omitempty"`
// }

// type AccountUser struct {
// 	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
// 	Name string `json:"name" xml:"name"`
// }

type ResponseAccountGroup struct {
	ID           int                     `json:"id,omitempty" xml:"id"`
	Name         string                  `json:"name" xml:"name"`
	AccessLevel  string                  `json:"access_level" xml:"access_level"`
	PrivilegeSet string                  `json:"privilege_set" xml:"privilege_set"`
	Site         AccountSubsetSite       `json:"site" xml:"site"`
	Privileges   AccountSubsetPrivileges `json:"privileges" xml:"privileges"`
	Members      []AccountDataSubsetUser `json:"members" xml:"members>user"`
}

// GetAccounts retrieves a list of all accounts (both users and groups).
func (c *Client) GetAccounts() (*ResponseAccountsList, error) {
	endpoint := uriAPIAccounts

	var accountsList ResponseAccountsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &accountsList)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &accountsList, nil
}

// GetAccountByID retrieves the Account by its ID
func (c *Client) GetAccountByID(id int) (*ResourceAccount, error) {
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	var account ResourceAccount
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &account)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &account, nil
}

// GetAccountByName retrieves the Account by its name
func (c *Client) GetAccountByName(name string) (*ResourceAccount, error) {
	endpoint := fmt.Sprintf("%s/username/%s", uriAPIAccounts, name)

	var account ResourceAccount
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &account)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
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
		fmt.Printf("Failed to execute request: %v\n", err)
		return nil, err
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
		fmt.Printf("Failed to execute request: %v\n", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &account, nil
}

// CreateAccountByID creates an Account using its ID
func (c *Client) CreateAccountByID(account *ResourceAccount) (*ResourceAccount, error) {
	// Use a placeholder ID for creating a new account
	placeholderID := 0
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, placeholderID)

	// Check if site is not provided and set default values
	if account.Site.ID == 0 && account.Site.Name == "" {
		account.Site = AccountSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Define XML requestBody structure
	requestBody := &struct {
		XMLName struct{} `xml:"account"`
		*ResourceAccount
	}{
		ResourceAccount: account,
	}

	var returnedAccount ResourceAccount
	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, &returnedAccount)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
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
		accountGroup.Site = AccountSubsetSite{
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
		fmt.Printf("Failed to execute request: %v\n", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &returnedAccountGroup, nil
}

// UpdateAccountByID updates an Account using its ID
func (c *Client) UpdateAccountByID(id int, account *ResourceAccount) (*ResourceAccount, error) {
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	if account.Site.ID == 0 && account.Site.Name == "" {
		account.Site = AccountSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	requestBody := &struct {
		XMLName struct{} `xml:"account"`
		*ResourceAccount
	}{
		ResourceAccount: account,
	}

	var updatedAccount ResourceAccount
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedAccount)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAccount, nil
}

// UpdateAccountByName updates an Account using its name.
func (c *Client) UpdateAccountByName(name string, account *ResourceAccount) (*ResourceAccount, error) {
	endpoint := fmt.Sprintf("%s/username/%s", uriAPIAccounts, name)

	if account.Site.ID == 0 && account.Site.Name == "" {
		account.Site = AccountSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	requestBody := &struct {
		XMLName struct{} `xml:"account"`
		*ResourceAccount
	}{
		ResourceAccount: account,
	}

	var updatedAccount ResourceAccount
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedAccount)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAccount, nil
}

// UpdateAccountGroupByID updates an Account Group using its ID
func (c *Client) UpdateAccountGroupByID(id int, group *ResponseAccountGroup) (*ResponseAccountGroup, error) {
	endpoint := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site = AccountSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	requestBody := &struct {
		XMLName struct{} `xml:"group"`
		*ResponseAccountGroup
	}{
		ResponseAccountGroup: group,
	}

	var updatedGroup ResponseAccountGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedGroup)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// UpdateAccountGroupByName updates an Account Group using its name.
func (c *Client) UpdateAccountGroupByName(name string, group *ResponseAccountGroup) (*ResponseAccountGroup, error) {
	endpoint := fmt.Sprintf("%s/groupname/%s", uriAPIAccounts, name)

	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site = AccountSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	requestBody := &struct {
		XMLName struct{} `xml:"group"`
		*ResponseAccountGroup
	}{
		ResponseAccountGroup: group,
	}

	var updatedGroup ResponseAccountGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedGroup)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
		return nil, err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// DeleteAccountByID deletes an Account using its ID
func (c *Client) DeleteAccountByID(id int) error {
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
		return err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAccountByName deletes an Account using its name.
func (c *Client) DeleteAccountByName(name string) error {
	endpoint := fmt.Sprintf("%s/username/%s", uriAPIAccounts, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
		return err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAccountGroupByID deletes an Account Group using its ID.
func (c *Client) DeleteAccountGroupByID(id int) error {
	endpoint := fmt.Sprintf("%s/groupid/%d", uriAPIAccounts, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		fmt.Printf("Failed to execute request: %v\n", err)
		return err
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
		fmt.Printf("Failed to execute request: %v\n", err)
		return err
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
