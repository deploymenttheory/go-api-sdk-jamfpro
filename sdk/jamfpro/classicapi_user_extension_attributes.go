// classicapi_user_extension_attributes.go
// Jamf Pro Classic Api - User Extension Attributes
// api reference: https://developer.jamf.com/jamf-pro/reference/userextensionattributes
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriUserExtensionAttributes = "/JSSResource/userextensionattributes"

// Structs for User Extension Attributes

type ResponseUserExtensionAttributesList struct {
	XMLName                 xml.Name `xml:"user_extension_attributes"`
	Size                    int      `xml:"size"`
	UserExtensionAttributes []struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	} `xml:"user_extension_attribute"`
}

// ResponseUserExtensionAttributerepresents a single user extension attribute.
type ResourceUserExtensionAttribute struct {
	XMLName     xml.Name `xml:"user_extension_attribute"`
	ID          int      `xml:"id,omitempty"`
	Name        string   `xml:"name"`
	Description string   `xml:"description"`
	DataType    string   `xml:"data_type"`
	InputType   struct {
		Type string `xml:"type"`
	} `xml:"input_type"`
}

// GetUserExtensionAttributes retrieves a list of all user extension attributes.
func (c *Client) GetUserExtensionAttributes() (*ResponseUserExtensionAttributesList, error) {
	endpoint := uriUserExtensionAttributes

	var extAttributes ResponseUserExtensionAttributesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &extAttributes)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user extension attributes: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &extAttributes, nil
}

// GetUserExtensionAttributeByID retrieves a user extension attribute by its ID.
func (c *Client) GetUserExtensionAttributeByID(id int) (*ResourceUserExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriUserExtensionAttributes, id)

	var userExtAttr ResourceUserExtensionAttribute
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userExtAttr)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user extension attribute by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userExtAttr, nil
}

// GetUserExtensionAttributeByName retrieves a user extension attribute by its name.
func (c *Client) GetUserExtensionAttributeByName(name string) (*ResourceUserExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriUserExtensionAttributes, name)

	var userExtAttr ResourceUserExtensionAttribute
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userExtAttr)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user extension attribute by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userExtAttr, nil
}

// CreateUserExtensionAttribute creates a new user extension attribute.
func (c *Client) CreateUserExtensionAttribute(attribute *ResourceUserExtensionAttribute) (*ResourceUserExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriUserExtensionAttributes) // Using ID 0 for creation

	requestBody := struct {
		XMLName xml.Name `xml:"user_extension_attribute"`
		*ResourceUserExtensionAttribute
	}{
		ResourceUserExtensionAttribute: attribute,
	}

	var createdAttribute ResourceUserExtensionAttribute
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdAttribute)
	if err != nil {
		return nil, fmt.Errorf("failed to create user extension attribute: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdAttribute, nil
}

// UpdateUserExtensionAttributeByID updates a user extension attribute by its ID.
func (c *Client) UpdateUserExtensionAttributeByID(id int, attribute *ResourceUserExtensionAttribute) (*ResourceUserExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriUserExtensionAttributes, id)

	requestBody := struct {
		XMLName xml.Name `xml:"user_extension_attribute"`
		*ResourceUserExtensionAttribute
	}{
		ResourceUserExtensionAttribute: attribute,
	}

	var updatedAttribute ResourceUserExtensionAttribute
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedAttribute)
	if err != nil {
		return nil, fmt.Errorf("failed to update user extension attribute by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAttribute, nil
}

// UpdateUserExtensionAttributeByName updates a user extension attribute by its name.
func (c *Client) UpdateUserExtensionAttributeByName(name string, attribute *ResourceUserExtensionAttribute) (*ResourceUserExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriUserExtensionAttributes, name)

	requestBody := struct {
		XMLName xml.Name `xml:"user_extension_attribute"`
		*ResourceUserExtensionAttribute
	}{
		ResourceUserExtensionAttribute: attribute,
	}

	var updatedAttribute ResourceUserExtensionAttribute
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedAttribute)
	if err != nil {
		return nil, fmt.Errorf("failed to update user extension attribute by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAttribute, nil
}

// DeleteUserExtensionAttributeByID deletes a user extension attribute by its ID.
func (c *Client) DeleteUserExtensionAttributeByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriUserExtensionAttributes, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete user extension attribute by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteUserExtensionAttributeByName deletes a user extension attribute by its name.
func (c *Client) DeleteUserExtensionAttributeByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriUserExtensionAttributes, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete user extension attribute by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
