// classicapi_directory_bindings.go
// Jamf Pro Classic Api - Directory Bindings
// api reference: https://developer.jamf.com/jamf-pro/reference/directorybindings
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// Base URI for Directory Bindings in Jamf Pro API
const uriDirectoryBindings = "/JSSResource/directorybindings"

// List

// Struct to capture the XML response for directory bindings
type ResponseDirectoryBindingsList struct {
	Size             int                         `xml:"size"`
	DirectoryBinding []DirectoryBindingsListItem `xml:"directory_binding"`
}

type DirectoryBindingsListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Response

// Struct to capture the XML response for a single directory binding
type ResponseDirectoryBinding struct {
	ID         int    `xml:"id"`
	Name       string `xml:"name"`
	Priority   int    `xml:"priority"`
	Domain     string `xml:"domain"`
	Username   string `xml:"username"`
	Password   string `xml:"password"`
	ComputerOU string `xml:"computer_ou"`
	Type       string `xml:"type"`
}

// CRUD

// GetDirectoryBindings retrieves a serialized list of directory bindings.
func (c *Client) GetDirectoryBindings() (*ResponseDirectoryBindingsList, error) {
	endpoint := uriDirectoryBindings

	var bindings ResponseDirectoryBindingsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &bindings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "directory bindings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &bindings, nil
}

// GetDirectoryBindingByID retrieves a single directory binding by its ID.
func (c *Client) GetDirectoryBindingByID(id int) (*ResponseDirectoryBinding, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriDirectoryBindings, id)

	var binding ResponseDirectoryBinding
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &binding)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "directory binding", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &binding, nil
}

// GetDirectoryBindingByName retrieves a single directory binding by its name.
func (c *Client) GetDirectoryBindingByName(name string) (*ResponseDirectoryBinding, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDirectoryBindings, name)

	var binding ResponseDirectoryBinding
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &binding)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "directory binding", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &binding, nil
}

// CreateDirectoryBinding creates a new directory binding.
func (c *Client) CreateDirectoryBinding(binding *ResponseDirectoryBinding) (*ResponseDirectoryBinding, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriDirectoryBindings)

	requestBody := struct {
		XMLName xml.Name `xml:"directory_binding"`
		*ResponseDirectoryBinding
	}{
		ResponseDirectoryBinding: binding,
	}

	var createdBinding ResponseDirectoryBinding
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdBinding)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "directory binding", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdBinding, nil
}

// UpdateDirectoryBindingByID updates a directory binding by its ID.
func (c *Client) UpdateDirectoryBindingByID(id int, binding *ResponseDirectoryBinding) (*ResponseDirectoryBinding, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriDirectoryBindings, id)

	requestBody := struct {
		XMLName xml.Name `xml:"directory_binding"`
		*ResponseDirectoryBinding
	}{
		ResponseDirectoryBinding: binding,
	}

	var updatedBinding ResponseDirectoryBinding
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedBinding)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "directory binding", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedBinding, nil
}

// UpdateDirectoryBindingByName updates a directory binding by its name.
func (c *Client) UpdateDirectoryBindingByName(name string, binding *ResponseDirectoryBinding) (*ResponseDirectoryBinding, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDirectoryBindings, name)

	requestBody := struct {
		XMLName xml.Name `xml:"directory_binding"`
		*ResponseDirectoryBinding
	}{
		ResponseDirectoryBinding: binding,
	}

	var updatedBinding ResponseDirectoryBinding
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedBinding)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "directory binding", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedBinding, nil
}

// DeleteDirectoryBindingByID deletes a directory binding by its ID.
func (c *Client) DeleteDirectoryBindingByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriDirectoryBindings, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "directory binding", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDirectoryBindingByName deletes a directory binding by its name.
func (c *Client) DeleteDirectoryBindingByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriDirectoryBindings, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "directory binding", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
