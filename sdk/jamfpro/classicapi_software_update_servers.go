// classicapi_software_update_servers.go
// Jamf Pro Classic Api - Software Update Servers
// api reference: https://developer.jamf.com/jamf-pro/reference/softwareupdateservers
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriSoftwareUpdateServers = "/JSSResource/softwareupdateservers"

// Structs for Software Update Servers Response
type ResponseSoftwareUpdateServersList struct {
	XMLName xml.Name `xml:"software_update_servers"`
	Size    int      `xml:"size"`
	Servers []struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	} `xml:"software_update_server"`
}

// Struct for individual Software Update Server
type ResourceSoftwareUpdateServer struct {
	ID            int    `xml:"id"`
	Name          string `xml:"name"`
	IPAddress     string `xml:"ip_address"`
	Port          int    `xml:"port"`
	SetSystemWide bool   `xml:"set_system_wide"`
}

// GetSoftwareUpdateServers retrieves a list of all software update servers.
func (c *Client) GetSoftwareUpdateServers() (*ResponseSoftwareUpdateServersList, error) {
	endpoint := uriSoftwareUpdateServers

	var response ResponseSoftwareUpdateServersList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch software update servers: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetSoftwareUpdateServersByID retrieves a specific software update server by its ID.
func (c *Client) GetSoftwareUpdateServersByID(id int) (*ResourceSoftwareUpdateServer, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriSoftwareUpdateServers, id)

	var response ResourceSoftwareUpdateServer
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch software update server by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetSoftwareUpdateServersByName retrieves a specific software update server by its name.
func (c *Client) GetSoftwareUpdateServersByName(name string) (*ResourceSoftwareUpdateServer, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriSoftwareUpdateServers, name)

	var response ResourceSoftwareUpdateServer
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch software update server by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// CreateSoftwareUpdateServer creates a new software update server.
func (c *Client) CreateSoftwareUpdateServer(server *ResourceSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriSoftwareUpdateServers) // '0' indicates creation

	requestBody := struct {
		XMLName xml.Name `xml:"software_update_server"`
		*ResourceSoftwareUpdateServer
	}{
		ResourceSoftwareUpdateServer: server,
	}

	var response ResourceSoftwareUpdateServer
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create software update server: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateSoftwareUpdateServerByID updates a software update server by its ID.
func (c *Client) UpdateSoftwareUpdateServerByID(id int, server *ResourceSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriSoftwareUpdateServers, id)

	requestBody := struct {
		XMLName xml.Name `xml:"software_update_server"`
		*ResourceSoftwareUpdateServer
	}{
		ResourceSoftwareUpdateServer: server,
	}

	var response ResourceSoftwareUpdateServer
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update software update server by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateSoftwareUpdateServerByName updates a software update server by its name.
func (c *Client) UpdateSoftwareUpdateServerByName(name string, server *ResourceSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriSoftwareUpdateServers, name)

	requestBody := struct {
		XMLName xml.Name `xml:"software_update_server"`
		*ResourceSoftwareUpdateServer
	}{
		ResourceSoftwareUpdateServer: server,
	}

	var response ResourceSoftwareUpdateServer
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update software update server by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteSoftwareUpdateServerByID deletes a software update server by its ID.
func (c *Client) DeleteSoftwareUpdateServerByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriSoftwareUpdateServers, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete software update server by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteSoftwareUpdateServerByName deletes a software update server by its name.
func (c *Client) DeleteSoftwareUpdateServerByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriSoftwareUpdateServers, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete software update server by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
