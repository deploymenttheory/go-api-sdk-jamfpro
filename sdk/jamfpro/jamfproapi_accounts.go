// jamfproapi_accounts.go
// Jamf Pro Api - Accounts (Local Admin Accounts)
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-accounts
// Jamf Pro API requires the structs to support a JSON data structure.
//
// Note: This is the modern Jamf Pro API (/api/v1/accounts) accounts resource. The Classic API
// accounts resource (/JSSResource/accounts) lives in classicapi_accounts.go.

package jamfpro

import (
	"fmt"
	"net/url"
)

const uriAccountsV1 = "/api/v1/accounts"

// List

// ResponseAccountsListV1 represents the search results for Jamf Pro API user accounts.
type ResponseAccountsListV1 struct {
	TotalCount int                   `json:"totalCount"`
	Results    []ResourceUserAccount `json:"results"`
}

// Resource

// ResourceUserAccount represents a Jamf Pro API user (local admin) account.
type ResourceUserAccount struct {
	ID                        string `json:"id,omitempty"`
	PlainPassword             string `json:"plainPassword,omitempty"`
	Username                  string `json:"username,omitempty"`
	Realname                  string `json:"realname,omitempty"`
	Email                     string `json:"email,omitempty"`
	Phone                     string `json:"phone,omitempty"`
	LdapServerID              int    `json:"ldapServerId"`
	DistinguishedName         string `json:"distinguishedName,omitempty"`
	SiteID                    int    `json:"siteId"`
	AccessLevel               string `json:"accessLevel,omitempty"`
	PrivilegeLevel            string `json:"privilegeLevel,omitempty"`
	LastPasswordChange        string `json:"lastPasswordChange,omitempty"`
	ChangePasswordOnNextLogin bool   `json:"changePasswordOnNextLogin,omitempty"`
	FailedLoginAttempts       int    `json:"failedLoginAttempts,omitempty"`
	AccountStatus             string `json:"accountStatus,omitempty"`
	AccountType               string `json:"accountType,omitempty"`
}

// GetAccountsV1 retrieves the list of Jamf Pro API user accounts.
func (c *Client) GetAccountsV1(params url.Values) (*ResponseAccountsListV1, error) {
	endpoint := uriAccountsV1
	if params != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var out ResponseAccountsListV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "accounts", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetAccountByIDV1 retrieves a single Jamf Pro API user account by ID.
func (c *Client) GetAccountByIDV1(id string) (*ResourceUserAccount, error) {
	endpoint := fmt.Sprintf("%s/%s", uriAccountsV1, id)

	var out ResourceUserAccount
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "account", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateAccountV1 creates a new Jamf Pro API user account.
func (c *Client) CreateAccountV1(account *ResourceUserAccount) (*ResourceUserAccount, error) {
	endpoint := uriAccountsV1

	var out ResourceUserAccount
	resp, err := c.HTTP.DoRequest("POST", endpoint, account, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "account", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateAccountByIDV1 updates an existing Jamf Pro API user account by ID.
func (c *Client) UpdateAccountByIDV1(id string, account *ResourceUserAccount) (*ResourceUserAccount, error) {
	endpoint := fmt.Sprintf("%s/%s", uriAccountsV1, id)

	var out ResourceUserAccount
	resp, err := c.HTTP.DoRequest("PUT", endpoint, account, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "account", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// DeleteAccountByIDV1 deletes a Jamf Pro API user account by ID.
func (c *Client) DeleteAccountByIDV1(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriAccountsV1, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDeleteByID, "account", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
