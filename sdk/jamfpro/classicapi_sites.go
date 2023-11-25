// classicapi_sites.go
// Jamf Pro Classic Api - Sites
// api reference: https://developer.jamf.com/jamf-pro/reference/sites
// Classic API requires the structs to support an XML data structure.
package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriSites = "/JSSResource/sites"

// Structs for the sites

type ResponseSitesList struct {
	Size int        `xml:"size"`
	Site []SiteItem `xml:"site"`
}

type SiteItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type ResponseSite struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// GetSites gets a list of all sites
func (c *Client) GetSites() (*ResponseSitesList, error) {
	endpoint := uriSites

	var sites ResponseSitesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &sites)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Sites: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &sites, nil
}

// GetSiteByID retrieves a site by its ID.
func (c *Client) GetSiteByID(id int) (*ResponseSite, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriSites, id)

	var site ResponseSite
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &site)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Site by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &site, nil
}

// GetSiteByName retrieves a site by its name.
func (c *Client) GetSiteByName(name string) (*ResponseSite, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriSites, name)

	var site ResponseSite
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &site)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Site by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &site, nil
}

// CreateSite creates a new site.
func (c *Client) CreateSite(site *ResponseSite) (*ResponseSite, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriSites) // Using ID 0 for creation as per the pattern

	requestBody := struct {
		XMLName xml.Name `xml:"site"`
		*ResponseSite
	}{
		ResponseSite: site,
	}

	var createdSite ResponseSite
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdSite)
	if err != nil {
		return nil, fmt.Errorf("failed to create Site: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdSite, nil
}

// UpdateSiteByID updates an existing site by its ID.
func (c *Client) UpdateSiteByID(id int, site *ResponseSite) (*ResponseSite, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriSites, id)

	requestBody := struct {
		XMLName xml.Name `xml:"site"`
		*ResponseSite
	}{
		ResponseSite: site,
	}

	var updatedSite ResponseSite
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSite)
	if err != nil {
		return nil, fmt.Errorf("failed to update Site by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSite, nil
}

// UpdateSiteByName updates an existing site by its name.
func (c *Client) UpdateSiteByName(name string, site *ResponseSite) (*ResponseSite, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriSites, name)

	requestBody := struct {
		XMLName xml.Name `xml:"site"`
		*ResponseSite
	}{
		ResponseSite: site,
	}

	var updatedSite ResponseSite
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSite)
	if err != nil {
		return nil, fmt.Errorf("failed to update Site by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSite, nil
}

// DeleteSiteByID deletes a site by its ID.
func (c *Client) DeleteSiteByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriSites, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Site by ID: %v", err)
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
		return fmt.Errorf("failed to delete Site by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
