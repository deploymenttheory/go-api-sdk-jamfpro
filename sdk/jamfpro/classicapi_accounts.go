// classicapi_accounts.go
// Jamf Pro Classic Api - Accounts
// api reference: https://developer.jamf.com/jamf-pro/reference/accounts
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
*/

package jamfpro

import (
	"fmt"
)

const uriAPIAccounts = "/JSSResource/accounts"

// List

type ResponseAccountsList struct {
	Users  []AccountsListSubsetUsers  `xml:"users>user,omitempty"`
	Groups []AccountsListSubsetGroups `xml:"groups>group,omitempty"`
}

// Resource

type ResourceAccount struct {
	ID                  int                     `json:"id,omitempty" xml:"id,omitempty"`
	Name                string                  `json:"name" xml:"name"`
	DirectoryUser       bool                    `json:"directory_user,omitempty" xml:"directory_user,omitempty"`
	FullName            string                  `json:"full_name,omitempty" xml:"full_name,omitempty"`
	Email               string                  `json:"email,omitempty" xml:"email,omitempty"`
	EmailAddress        string                  `json:"email_address,omitempty" xml:"email_address,omitempty"`
	Enabled             string                  `json:"enabled,omitempty" xml:"enabled,omitempty"`
	LdapServer          AccountSubsetLdapServer `json:"ldap_server,omitempty" xml:"ldap_server,omitempty"`
	ForcePasswordChange bool                    `json:"force_password_change,omitempty" xml:"force_password_change,omitempty"`
	AccessLevel         string                  `json:"access_level,omitempty" xml:"access_level,omitempty"`
	Password            string                  `json:"password" xml:"password"`
	PrivilegeSet        string                  `json:"privilege_set,omitempty" xml:"privilege_set,omitempty"`
	Site                SharedResourceSite      `json:"site,omitempty" xml:"site,omitempty"`
	Privileges          AccountSubsetPrivileges `json:"privileges,omitempty" xml:"privileges,omitempty"`
}

// Responses

type ResponseAccountCreatedAndUpdated struct {
	ID int `json:"id,omitempty" xml:"id,omitempty"`
}

// Subsets

type AccountsListSubsetUsers struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type AccountsListSubsetGroups struct {
	ID         int                     `json:"id,omitempty" xml:"id,omitempty"`
	Name       string                  `json:"name" xml:"name"`
	Site       SharedResourceSite      `json:"site,omitempty" xml:"site,omitempty"`
	Privileges AccountSubsetPrivileges `json:"privileges,omitempty" xml:"privileges,omitempty"`
}

type AccountSubsetLdapServer struct {
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

// CRUD

// GetAccounts retrieves a list of all accounts (both users and groups).
func (c *Client) GetAccounts() (*ResponseAccountsList, error) {
	endpoint := uriAPIAccounts

	var accountsList ResponseAccountsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &accountsList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "accounts", err)
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
		return nil, fmt.Errorf(errMsgFailedGetByID, "account", id, err)
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
		return nil, fmt.Errorf(errMsgFailedGetByName, "account", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &account, nil
}

// CreateAccountByID creates an Account using its ID
func (c *Client) CreateAccount(account *ResourceAccount) (*ResponseAccountCreatedAndUpdated, error) {
	// Use a placeholder ID for creating a new account
	placeholderID := 0
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, placeholderID)

	// Check if site is not provided and set default values
	if account.Site.ID == 0 && account.Site.Name == "" {
		account.Site = SharedResourceSite{
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

	var returnedAccount ResponseAccountCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, &returnedAccount)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "account", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &returnedAccount, nil
}

// UpdateAccountByID updates an Account using its ID
func (c *Client) UpdateAccountByID(id int, account *ResourceAccount) (*ResponseAccountCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	if account.Site.ID == 0 && account.Site.Name == "" {
		account.Site = SharedResourceSite{
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

	var updatedAccount ResponseAccountCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedAccount)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "account", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAccount, nil
}

// UpdateAccountByName updates an Account using its name.
func (c *Client) UpdateAccountByName(name string, account *ResourceAccount) (*ResponseAccountCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/username/%s", uriAPIAccounts, name)

	if account.Site.ID == 0 && account.Site.Name == "" {
		account.Site = SharedResourceSite{
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

	var updatedAccount ResponseAccountCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedAccount)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "account", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAccount, nil
}

// DeleteAccountByID deletes an Account using its ID
func (c *Client) DeleteAccountByID(id int) error {
	endpoint := fmt.Sprintf("%s/userid/%d", uriAPIAccounts, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "account", id, err)
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
		return fmt.Errorf(errMsgFailedDeleteByName, "account", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
