// classicapi_mobile_device_extension_attributes.go
// Jamf Pro Classic Api - Mobile Extension Attributes
// API reference: https://developer.jamf.com/jamf-pro/reference/mobiledeviceextensionattributes
// Jamf Pro Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMobileDeviceExtensionAttributes = "/JSSResource/mobiledeviceextensionattributes"

// ResponseMobileDeviceExtensionAttributesList represents the response for a list of mobile device extension attributes.
type ResponseMobileDeviceExtensionAttributesList struct {
	Size                           int                                  `xml:"size"`
	MobileDeviceExtensionAttribute []MobileDeviceExtensionAttributeItem `xml:"mobile_device_extension_attribute"`
}

// MobileDeviceExtensionAttributeItem represents a single mobile device extension attribute item.
type MobileDeviceExtensionAttributeItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResponseMobileExtensionAttributes represents the response structure for a mobile extension attribute.
type ResponseMobileExtensionAttributes struct {
	ID               int                               `xml:"id"`
	Name             string                            `xml:"name"`
	Description      string                            `xml:"description,omitempty"`
	DataType         string                            `xml:"data_type,omitempty"`
	InputType        MobileExtensionAttributeInputType `xml:"input_type,omitempty"`
	InventoryDisplay string                            `xml:"inventory_display,omitempty"`
}

// MobileExtensionAttributeInputType represents the input type of the mobile extension attribute.
type MobileExtensionAttributeInputType struct {
	Type string `xml:"type,omitempty"`
}

// GetMobileExtensionAttributes retrieves a serialized list of mobile device extension attributes.
func (c *Client) GetMobileExtensionAttributes() (*ResponseMobileDeviceExtensionAttributesList, error) {
	endpoint := uriMobileDeviceExtensionAttributes

	var extensionAttributes ResponseMobileDeviceExtensionAttributesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &extensionAttributes)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device extension attributes: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &extensionAttributes, nil
}

// GetMobileExtensionAttributeByID fetches a specific mobile extension attribute by its ID.
func (c *Client) GetMobileExtensionAttributeByID(id int) (*ResponseMobileExtensionAttributes, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceExtensionAttributes, id)

	var attribute ResponseMobileExtensionAttributes
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attribute)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile extension attribute by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attribute, nil
}

// GetMobileExtensionAttributeByName fetches a specific mobile extension attribute by its name.
func (c *Client) GetMobileExtensionAttributeByName(name string) (*ResponseMobileExtensionAttributes, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceExtensionAttributes, name)

	var attribute ResponseMobileExtensionAttributes
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attribute)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile extension attribute by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attribute, nil
}

// CreateMobileExtensionAttribute creates a new mobile device extension attribute.
func (c *Client) CreateMobileExtensionAttribute(attribute *ResponseMobileExtensionAttributes) (*ResponseMobileExtensionAttributes, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceExtensionAttributes)

	// Wrap the attribute with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_extension_attribute"`
		*ResponseMobileExtensionAttributes
	}{
		ResponseMobileExtensionAttributes: attribute,
	}

	var responseAttribute ResponseMobileExtensionAttributes
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseAttribute)
	if err != nil {
		return nil, fmt.Errorf("failed to create mobile device extension attribute: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseAttribute, nil
}

// UpdateMobileExtensionAttributeByID updates a mobile extension attribute by its ID.
func (c *Client) UpdateMobileExtensionAttributeByID(id int, attribute *ResponseMobileExtensionAttributes) (*ResponseMobileExtensionAttributes, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceExtensionAttributes, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_extension_attribute"`
		*ResponseMobileExtensionAttributes
	}{
		ResponseMobileExtensionAttributes: attribute,
	}

	var responseAttribute ResponseMobileExtensionAttributes
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseAttribute)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile extension attribute by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseAttribute, nil
}

// UpdateMobileExtensionAttributeByName updates a mobile extension attribute by its name.
func (c *Client) UpdateMobileExtensionAttributeByName(name string, attribute *ResponseMobileExtensionAttributes) (*ResponseMobileExtensionAttributes, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceExtensionAttributes, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_extension_attribute"`
		*ResponseMobileExtensionAttributes
	}{
		ResponseMobileExtensionAttributes: attribute,
	}

	var responseAttribute ResponseMobileExtensionAttributes
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseAttribute)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile extension attribute by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseAttribute, nil
}

// DeleteMobileExtensionAttributeByID deletes a mobile extension attribute by its ID.
func (c *Client) DeleteMobileExtensionAttributeByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceExtensionAttributes, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile extension attribute by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileExtensionAttributeByName deletes a mobile extension attribute by its name.
func (c *Client) DeleteMobileExtensionAttributeByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceExtensionAttributes, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile extension attribute by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
