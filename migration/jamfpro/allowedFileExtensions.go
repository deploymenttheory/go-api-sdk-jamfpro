// allowedFileExtensions.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAllowedFileExtensions = "/JSSResource/allowedfileextensions"

// XML structure represented in nested Go structs
// AllowedFileExtension Response structure
type ResponseAllowedFileExtension struct {
	ID        int    `json:"id" xml:"id"`
	Extension string `json:"extension" xml:"extension"`
}

// Response structure for the list of allowed file extensions
type AllowedFileExtensionsList struct {
	Size                  int                    `json:"size" xml:"size"`
	AllowedFileExtensions []AllowedFileExtension `json:"allowed_file_extension" xml:"allowed_file_extension"`
}

// Create / Update - Account structure. XML only as api endpoint only accepts XML for puts
type AllowedFileExtension struct {
	XMLName   xml.Name `xml:"allowed_file_extension"`
	ID        int      `xml:"id,omitempty"`
	Extension string   `xml:"extension"`
}

// GetAllAllowedFileExtensions retrieves all allowed file extensions
func (c *Client) GetAllowedFileExtensions() (*AllowedFileExtensionsList, error) {
	url := uriAPIAllowedFileExtensions

	var allowedExtensionsList AllowedFileExtensionsList
	if err := c.DoRequest("GET", url, nil, nil, &allowedExtensionsList); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &allowedExtensionsList, nil
}

// GetAllowedFileExtensionByID retrieves the allowed file extension by its ID
func (c *Client) GetAllowedFileExtensionByID(id int) (*ResponseAllowedFileExtension, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIAllowedFileExtensions, id)

	var extension ResponseAllowedFileExtension
	if err := c.DoRequest("GET", url, nil, nil, &extension); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &extension, nil
}

// GetAllowedFileExtensionByName retrieves the allowed file extension by its name
func (c *Client) GetAllowedFileExtensionByName(extension string) (*ResponseAllowedFileExtension, error) {
	url := fmt.Sprintf("%s/extension/%s", uriAPIAllowedFileExtensions, extension)

	var ext ResponseAllowedFileExtension
	if err := c.DoRequest("GET", url, nil, nil, &ext); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &ext, nil
}

// CreateAllowedFileExtension creates a new allowed file extension
func (c *Client) CreateAllowedFileExtension(extension *AllowedFileExtension) error {
	url := fmt.Sprintf("%s/id/0", uriAPIAllowedFileExtensions) // Set ID to 0 to let Jamf assign an ID

	if err := c.DoRequest("POST", url, extension, nil, nil); err != nil {
		return fmt.Errorf("failed to create allowed file extension: %v", err)
	}

	return nil
}

// UpdateAllowedFileExtensionByID updates an existing allowed file extension by ID
func (c *Client) UpdateAllowedFileExtensionByID(id int, extension *AllowedFileExtension) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIAllowedFileExtensions, id)

	if err := c.DoRequestDebug("PUT", url, extension, nil, nil); err != nil {
		return fmt.Errorf("failed to update allowed file extension by ID: %v", err)
	}

	return nil
}

// UpdateAllowedFileExtensionByName updates an existing allowed file extension by Name
func (c *Client) UpdateAllowedFileExtensionByName(extensionName string, extension *ResponseAllowedFileExtension) error {
	url := fmt.Sprintf("%s/extension/%s", uriAPIAllowedFileExtensions, extensionName)

	if err := c.DoRequest("PUT", url, extension, nil, nil); err != nil {
		return fmt.Errorf("failed to update allowed file extension by Name: %v", err)
	}

	return nil
}

// DeleteAllowedFileExtensionByID deletes an existing allowed file extension by ID
func (c *Client) DeleteAllowedFileExtensionByID(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIAllowedFileExtensions, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete allowed file extension by ID: %v", err)
	}

	return nil
}

// DeleteAllowedFileExtensionByName deletes an existing allowed file extension by Name
func (c *Client) DeleteAllowedFileExtensionByName(extensionName string) error {
	url := fmt.Sprintf("%s/extension/%s", uriAPIAllowedFileExtensions, extensionName)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete allowed file extension by Name: %v", err)
	}

	return nil
}
