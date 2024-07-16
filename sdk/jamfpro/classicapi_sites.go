// classicapi_sites.go
// Jamf Pro Classic Api - Sites
// api reference: https://developer.jamf.com/jamf-pro/reference/sites
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedSharedResourceSite
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriSites = "/JSSResource/sites"

// Structs for the sites

// List

type ResponseSitesList struct {
	Size int                  `xml:"size"`
	Site []SharedResourceSite `xml:"site"`
}

// No Resource as using a shared one: SharedSharedResourceSite

// CRUD

// GetSites gets a list of all sites
func (c *Client) GetSites() (*ResponseSitesList, error) {
	endpoint := uriSites

	var sites ResponseSitesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &sites)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "sites", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &sites, nil
}

// GetSiteByID retrieves a site by its ID.
func (c *Client) GetSiteByID(id int) (*SharedResourceSite, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriSites, id)

	var site SharedResourceSite
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &site)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "site", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &site, nil
}

// GetSiteByName retrieves a site by its name.
func (c *Client) GetSiteByName(name string) (*SharedResourceSite, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriSites, name)

	var site SharedResourceSite
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &site)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "site", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &site, nil
}

// CreateSite creates a new site.
func (c *Client) CreateSite(site *SharedResourceSite) (*SharedResourceSite, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriSites)

	requestBody := struct {
		XMLName xml.Name `xml:"site"`
		*SharedResourceSite
	}{
		SharedResourceSite: site,
	}

	var createdSite SharedResourceSite
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdSite)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "site", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdSite, nil
}

// UpdateSiteByID updates an existing site by its ID.
func (c *Client) UpdateSiteByID(id int, site *SharedResourceSite) (*SharedResourceSite, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriSites, id)

	requestBody := struct {
		XMLName xml.Name `xml:"site"`
		*SharedResourceSite
	}{
		SharedResourceSite: site,
	}

	var updatedSite SharedResourceSite
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSite)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "site", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSite, nil
}

// UpdateSiteByName updates an existing site by its name.
func (c *Client) UpdateSiteByName(name string, site *SharedResourceSite) (*SharedResourceSite, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriSites, name)

	requestBody := struct {
		XMLName xml.Name `xml:"site"`
		*SharedResourceSite
	}{
		SharedResourceSite: site,
	}

	var updatedSite SharedResourceSite
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSite)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "site", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSite, nil
}

// DeleteSiteByID deletes a site by its ID.
func (c *Client) DeleteSiteByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriSites, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "site", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteSiteByName deletes a site by its name.
func (c *Client) DeleteSiteByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriSites, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "site", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
