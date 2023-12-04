// classicapi_vpp_accounts.go
// Jamf Pro Classic Api - VPP Accounts
// api reference: https://developer.jamf.com/jamf-pro/reference/vppaccounts
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriVPPAccounts = "/JSSResource/vppaccounts"

// Structs for VPP Accounts Response
type ResponseVPPAccountsList struct {
	Size     int                  `xml:"size"`
	Accounts []VPPAccountListItem `xml:"vpp_account"`
}

type VPPAccountListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Struct for individual VPP Account
type ResponseVPPAccount struct {
	ID                            int                      `xml:"id"`
	Name                          string                   `xml:"name"`
	Contact                       string                   `xml:"contact"`
	ServiceToken                  string                   `xml:"service_token"`
	AccountName                   string                   `xml:"account_name"`
	ExpirationDate                string                   `xml:"expiration_date"`
	Country                       string                   `xml:"country"`
	AppleID                       string                   `xml:"apple_id"`
	Site                          VPPAccountDataSubsetSite `xml:"site"`
	PopulateCatalogFromVPPContent bool                     `xml:"populate_catalog_from_vpp_content"`
	NotifyDisassociation          bool                     `xml:"notify_disassociation"`
	AutoRegisterManagedUsers      bool                     `xml:"auto_register_managed_users"`
}

type VPPAccountDataSubsetSite struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// GetVPPAccounts retrieves a list of all VPP accounts.
func (c *Client) GetVPPAccounts() (*ResponseVPPAccountsList, error) {
	endpoint := uriVPPAccounts

	var response ResponseVPPAccountsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch VPP accounts: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetVPPAccountByID retrieves a specific VPP account by its ID.
func (c *Client) GetVPPAccountByID(id int) (*ResponseVPPAccount, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAccounts, id)

	var response ResponseVPPAccount
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch VPP account by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// CreateVPPAccount creates a new VPP account.
func (c *Client) CreateVPPAccount(account *ResponseVPPAccount) (*ResponseVPPAccount, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriVPPAccounts) // '0' indicates creation

	// Setting default values for Site if not supplied
	if account.Site.ID == 0 && account.Site.Name == "" {
		account.Site = VPPAccountDataSubsetSite{ID: -1, Name: "None"}
	}

	// Using an anonymous struct for the request body
	requestBody := struct {
		XMLName xml.Name `xml:"vpp_account"`
		*ResponseVPPAccount
	}{
		ResponseVPPAccount: account,
	}

	var response ResponseVPPAccount
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create VPP account: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateVPPAccount updates an existing VPP account.
func (c *Client) UpdateVPPAccount(id int, account *ResponseVPPAccount) (*ResponseVPPAccount, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAccounts, id)

	// Using an anonymous struct for the request body
	requestBody := struct {
		XMLName xml.Name `xml:"vpp_account"`
		*ResponseVPPAccount
	}{
		ResponseVPPAccount: account,
	}

	var response ResponseVPPAccount
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update VPP account: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteVPPAccountByID deletes a specific VPP account by its ID.
func (c *Client) DeleteVPPAccountByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAccounts, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete VPP account by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
