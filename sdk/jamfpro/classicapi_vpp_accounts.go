// classicapi_vpp_accounts.go
// Jamf Pro Classic Api - VPP Accounts
// api reference: https://developer.jamf.com/jamf-pro/reference/vppaccounts
// Jamf Pro Classic Api requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriVPPAccounts = "/JSSResource/vppaccounts"

// List

// Structs for VPP Accounts Response
type ResponseVPPAccountsList struct {
	Size     int                   `xml:"size"`
	Accounts []VPPAccountsListItem `xml:"vpp_account"`
}

type VPPAccountsListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// Struct for individual VPP Account
type ResourceVPPAccount struct {
	ID                            int                `xml:"id"`
	Name                          string             `xml:"name"`
	Contact                       string             `xml:"contact"`
	ServiceToken                  string             `xml:"service_token"`
	AccountName                   string             `xml:"account_name"`
	ExpirationDate                string             `xml:"expiration_date"`
	Country                       string             `xml:"country"`
	AppleID                       string             `xml:"apple_id"`
	Site                          SharedResourceSite `xml:"site"`
	PopulateCatalogFromVPPContent bool               `xml:"populate_catalog_from_vpp_content"`
	NotifyDisassociation          bool               `xml:"notify_disassociation"`
	AutoRegisterManagedUsers      bool               `xml:"auto_register_managed_users"`
}

// CRUD

// GetVPPAccounts retrieves a list of all VPP accounts.
func (c *Client) GetVPPAccounts() (*ResponseVPPAccountsList, error) {
	endpoint := uriVPPAccounts

	var response ResponseVPPAccountsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "vpp accounts", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetVPPAccountByID retrieves a specific VPP account by its ID.
func (c *Client) GetVPPAccountByID(id int) (*ResourceVPPAccount, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAccounts, id)

	var response ResourceVPPAccount
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "vpp account", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// CreateVPPAccount creates a new VPP account.
func (c *Client) CreateVPPAccount(account *ResourceVPPAccount) (*ResourceVPPAccount, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriVPPAccounts)

	requestBody := struct {
		XMLName xml.Name `xml:"vpp_account"`
		*ResourceVPPAccount
	}{
		ResourceVPPAccount: account,
	}

	var response ResourceVPPAccount
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "vpp account", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateVPPAccount updates an existing VPP account.
func (c *Client) UpdateVPPAccountByID(id int, account *ResourceVPPAccount) (*ResourceVPPAccount, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAccounts, id)

	requestBody := struct {
		XMLName xml.Name `xml:"vpp_account"`
		*ResourceVPPAccount
	}{
		ResourceVPPAccount: account,
	}

	var response ResourceVPPAccount
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "vpp account", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteVPPAccountByID deletes a specific VPP account by its ID.
func (c *Client) DeleteVPPAccountByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriVPPAccounts, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "vpp account", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
