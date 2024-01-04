// classicapi_allowed_file_extensions.go
// Jamf Pro Classic Api - Allowed File Extensions
// api reference: https://developer.jamf.com/jamf-pro/reference/allowedfileextensions
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAllowedFileExtensions = "/JSSResource/allowedfileextensions"

// List

// Response structure for the list of allowed file extensions
type ResponseAllowedFileExtensionsList struct {
	XMLName               xml.Name                       `xml:"allowed_file_extensions"`
	Size                  int                            `xml:"size"`
	AllowedFileExtensions []ResourceAllowedFileExtension `xml:"allowed_file_extension"`
}

// Resource

type ResourceAllowedFileExtension struct {
	XMLName   xml.Name `xml:"allowed_file_extension"`
	ID        int      `xml:"id"`
	Extension string   `xml:"extension"`
}

// CRUD

// GetAllowedFileExtensions retrieves all allowed file extensions
func (c *Client) GetAllowedFileExtensions() (*ResponseAllowedFileExtensionsList, error) {
	endpoint := uriAPIAllowedFileExtensions

	var allowedExtensionsList ResponseAllowedFileExtensionsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &allowedExtensionsList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "allowed file extension", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &allowedExtensionsList, nil
}

// GetAllowedFileExtensionByID retrieves the allowed file extension by its ID
func (c *Client) GetAllowedFileExtensionByID(id int) (*ResourceAllowedFileExtension, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAllowedFileExtensions, id)

	var extension ResourceAllowedFileExtension
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &extension)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "allowed file extension", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &extension, nil
}

// GetAllowedFileExtensionByName retrieves the allowed file extension by its name
func (c *Client) GetAllowedFileExtensionByName(name string) (*ResourceAllowedFileExtension, error) {
	endpoint := fmt.Sprintf("%s/extension/%s", uriAPIAllowedFileExtensions, name)

	var extension ResourceAllowedFileExtension
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &extension)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "allowed file extension", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &extension, nil
}

// CreateAllowedFileExtension creates a new allowed file extension
func (c *Client) CreateAllowedFileExtension(extension *ResourceAllowedFileExtension) (*ResourceAllowedFileExtension, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriAPIAllowedFileExtensions) // Using 0 as placeholder for creation

	var response ResourceAllowedFileExtension
	resp, err := c.HTTP.DoRequest("POST", endpoint, extension, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "allowed file extension", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteAllowedFileExtensionByID deletes an existing allowed file extension by ID
func (c *Client) DeleteAllowedFileExtensionByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAllowedFileExtensions, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "allowed file extension", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAllowedFileExtensionByNameByID deletes an existing allowed file extension by resolving its name to an ID
func (c *Client) DeleteAllowedFileExtensionByName(name string) error {
	extensionDetail, err := c.GetAllowedFileExtensionByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "allowed file extension", name, err)
	}

	return c.DeleteAllowedFileExtensionByID(extensionDetail.ID)
}
