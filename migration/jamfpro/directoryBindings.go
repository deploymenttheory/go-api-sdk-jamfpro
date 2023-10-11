// directoryBindings.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIDirectoryBindings = "/JSSResource/directorybindings"

// DirectoryBinding structure
type ResponseDirectoryBinding struct {
	ID         int    `json:"id,omitempty" xml:"id,omitempty"`
	Name       string `json:"name" xml:"name"`
	Priority   int    `json:"priority,omitempty" xml:"priority,omitempty"`
	Domain     string `json:"domain,omitempty" xml:"domain,omitempty"`
	Username   string `json:"username,omitempty" xml:"username,omitempty"`
	Password   string `json:"password,omitempty" xml:"password,omitempty"`
	ComputerOU string `json:"computer_ou,omitempty" xml:"computer_ou,omitempty"`
	Type       string `json:"type,omitempty" xml:"type,omitempty"`
}

type ResponseDirectoryBindingsList struct {
	Size             int                  `json:"size" xml:"size"`
	DirectoryBinding DirectoryBindingList `json:"directory_binding" xml:"directory_binding"`
}

type DirectoryBindingList struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// DirectoryBinding structure to represent the XML request body
type DirectoryBinding struct {
	XMLName xml.Name `xml:"directory_binding"`
	ResponseDirectoryBinding
}

// GetDirectoryBindingByID retrieves the Directory Binding by its ID
func (c *Client) GetDirectoryBindingByID(id int) (*ResponseDirectoryBinding, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIDirectoryBindings, id)

	var binding ResponseDirectoryBinding
	if err := c.DoRequest("GET", url, nil, nil, &binding); err != nil {
		return nil, fmt.Errorf("failed to get directory binding by ID: %v", err)
	}

	return &binding, nil
}

// GetDirectoryBindingByName retrieves the Directory Binding by its name
func (c *Client) GetDirectoryBindingByName(name string) (*ResponseDirectoryBinding, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIDirectoryBindings, name)

	var binding ResponseDirectoryBinding
	if err := c.DoRequest("GET", url, nil, nil, &binding); err != nil {
		return nil, fmt.Errorf("failed to get directory binding by name: %v", err)
	}

	return &binding, nil
}

// GetDirectoryBindings retrieves all directory bindings
func (c *Client) GetDirectoryBindings() ([]ResponseDirectoryBindingsList, error) {
	url := uriAPIDirectoryBindings

	var bindings []ResponseDirectoryBindingsList
	if err := c.DoRequest("GET", url, nil, nil, &bindings); err != nil {
		return nil, fmt.Errorf("failed to get all directory bindings: %v", err)
	}

	return bindings, nil
}

// CreateDirectoryBindingByID creates a new directory binding using the given ID
func (c *Client) CreateDirectoryBindingByID(id int, binding *DirectoryBinding) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIDirectoryBindings, id)

	if err := c.DoRequest("POST", url, binding, nil, nil); err != nil {
		return fmt.Errorf("failed to create directory binding by ID: %v", err)
	}

	return nil
}

// UpdateDirectoryBindingByID updates an existing directory binding using the given ID
func (c *Client) UpdateDirectoryBindingByID(id int, binding *DirectoryBinding) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIDirectoryBindings, id)

	if err := c.DoRequest("PUT", url, binding, nil, nil); err != nil {
		return fmt.Errorf("failed to update directory binding by ID: %v", err)
	}

	return nil
}

// UpdateDirectoryBindingByName updates an existing directory binding using the given name
func (c *Client) UpdateDirectoryBindingByName(name string, binding *DirectoryBinding) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIDirectoryBindings, name)

	if err := c.DoRequestDebug("PUT", url, binding, nil, nil); err != nil {
		return fmt.Errorf("failed to update directory binding by name: %v", err)
	}

	return nil
}

// DeleteDirectoryBindingByID deletes an existing directory binding using the given ID
func (c *Client) DeleteDirectoryBindingByID(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIDirectoryBindings, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete directory binding by ID: %v", err)
	}

	return nil
}

// DeleteDirectoryBindingByName deletes an existing directory binding using the given name
func (c *Client) DeleteDirectoryBindingByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIDirectoryBindings, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete directory binding by name: %v", err)
	}

	return nil
}
