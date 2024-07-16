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
	Size                           int                                       `xml:"size"`
	MobileDeviceExtensionAttribute []MobileDeviceExtensionAttributesListItem `xml:"mobile_device_extension_attribute"`
}

type MobileDeviceExtensionAttributesListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// ResourceMobileExtensionAttributes represents the response structure for a mobile extension attribute.
type ResourceMobileExtensionAttribute struct {
	ID               int                                     `xml:"id"`
	Name             string                                  `xml:"name"`
	Description      string                                  `xml:"description,omitempty"`
	DataType         string                                  `xml:"data_type,omitempty"`
	InputType        MobileExtensionAttributeSubsetInputType `xml:"input_type,omitempty"`
	InventoryDisplay string                                  `xml:"inventory_display,omitempty"`
}

// Subsets

type MobileExtensionAttributeSubsetInputType struct {
	Type string `xml:"type,omitempty"`
}

// CRUD

// GetMobileExtensionAttributes retrieves a serialized list of mobile device extension attributes.
func (c *Client) GetMobileExtensionAttributes() (*ResponseMobileDeviceExtensionAttributesList, error) {
	endpoint := uriMobileDeviceExtensionAttributes

	var extensionAttributes ResponseMobileDeviceExtensionAttributesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &extensionAttributes)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "mobile device extension attributes", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &extensionAttributes, nil
}

// GetMobileExtensionAttributeByID fetches a specific mobile extension attribute by its ID.
func (c *Client) GetMobileExtensionAttributeByID(id int) (*ResourceMobileExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceExtensionAttributes, id)

	var attribute ResourceMobileExtensionAttribute
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device extension attribute", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attribute, nil
}

// GetMobileExtensionAttributeByName fetches a specific mobile extension attribute by its name.
func (c *Client) GetMobileExtensionAttributeByName(name string) (*ResourceMobileExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceExtensionAttributes, name)

	var attribute ResourceMobileExtensionAttribute
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device extension attribute", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attribute, nil
}

// CreateMobileExtensionAttribute creates a new mobile device extension attribute.
func (c *Client) CreateMobileExtensionAttribute(attribute *ResourceMobileExtensionAttribute) (*ResourceMobileExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceExtensionAttributes)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_extension_attribute"`
		*ResourceMobileExtensionAttribute
	}{
		ResourceMobileExtensionAttribute: attribute,
	}

	var responseAttribute ResourceMobileExtensionAttribute
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseAttribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "mobile device extension attribute", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseAttribute, nil
}

// UpdateMobileExtensionAttributeByID updates a mobile extension attribute by its ID.
func (c *Client) UpdateMobileExtensionAttributeByID(id int, attribute *ResourceMobileExtensionAttribute) (*ResourceMobileExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceExtensionAttributes, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_extension_attribute"`
		*ResourceMobileExtensionAttribute
	}{
		ResourceMobileExtensionAttribute: attribute,
	}

	var responseAttribute ResourceMobileExtensionAttribute
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseAttribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "mobile device extension attribute", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseAttribute, nil
}

// UpdateMobileExtensionAttributeByName updates a mobile extension attribute by its name.
func (c *Client) UpdateMobileExtensionAttributeByName(name string, attribute *ResourceMobileExtensionAttribute) (*ResourceMobileExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceExtensionAttributes, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_extension_attribute"`
		*ResourceMobileExtensionAttribute
	}{
		ResourceMobileExtensionAttribute: attribute,
	}

	var responseAttribute ResourceMobileExtensionAttribute
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseAttribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "mobile device extension attribute", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseAttribute, nil
}

// DeleteMobileExtensionAttributeByID deletes a mobile extension attribute by its ID.
func (c *Client) DeleteMobileExtensionAttributeByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriMobileDeviceExtensionAttributes, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device extension attribute", id, err)
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
		return fmt.Errorf(errMsgFailedDeleteByName, "mobile device extension attribute", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
