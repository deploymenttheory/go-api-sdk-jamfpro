// classicapi_allowed_file_extensions.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAllowedFileExtensions = "/JSSResource/allowedfileextensions"

// Response structure for the list of allowed file extensions
type ResponseAllowedFileExtensionsList struct {
	XMLName               xml.Name                     `xml:"allowed_file_extensions"`
	Size                  int                          `xml:"size"`
	AllowedFileExtensions []AllowedFileExtensionDetail `xml:"allowed_file_extension"`
}

type AllowedFileExtensionDetail struct {
	XMLName   xml.Name `xml:"allowed_file_extension"`
	ID        int      `xml:"id"`
	Extension string   `xml:"extension"`
}

// GetAllowedFileExtensions retrieves all allowed file extensions
func (c *Client) GetAllowedFileExtensions() (*ResponseAllowedFileExtensionsList, error) {
	endpoint := uriAPIAllowedFileExtensions

	var allowedExtensionsList ResponseAllowedFileExtensionsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &allowedExtensionsList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch allowed file extensions: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &allowedExtensionsList, nil
}

// GetAllowedFileExtensionByID retrieves the allowed file extension by its ID
func (c *Client) GetAllowedFileExtensionByID(id int) (*AllowedFileExtensionDetail, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAllowedFileExtensions, id)

	var extension AllowedFileExtensionDetail
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &extension)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch allowed file extension by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &extension, nil
}

// GetAllowedFileExtensionByName retrieves the allowed file extension by its name
func (c *Client) GetAllowedFileExtensionByName(extensionName string) (*AllowedFileExtensionDetail, error) {
	endpoint := fmt.Sprintf("%s/extension/%s", uriAPIAllowedFileExtensions, extensionName)

	var extension AllowedFileExtensionDetail
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &extension)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch allowed file extension by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &extension, nil
}

// CreateAllowedFileExtension creates a new allowed file extension
func (c *Client) CreateAllowedFileExtension(extension *AllowedFileExtensionDetail) (*AllowedFileExtensionDetail, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriAPIAllowedFileExtensions) // Using 0 as placeholder for creation

	var response AllowedFileExtensionDetail
	resp, err := c.HTTP.DoRequest("POST", endpoint, extension, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create allowed file extension: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateAllowedFileExtensionByID - doesn't exist. api doesn't support update

// DeleteAllowedFileExtensionByID deletes an existing allowed file extension by ID
func (c *Client) DeleteAllowedFileExtensionByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAllowedFileExtensions, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete allowed file extension by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAllowedFileExtensionByNameByID deletes an existing allowed file extension by resolving its name to an ID
func (c *Client) DeleteAllowedFileExtensionByName(extensionName string) error {
	extensionDetail, err := c.GetAllowedFileExtensionByName(extensionName)
	if err != nil {
		return fmt.Errorf("failed to resolve allowed file extension name to ID: %v", err)
	}

	// Now use the DeleteAllowedFileExtensionByID function
	return c.DeleteAllowedFileExtensionByID(extensionDetail.ID)
}
