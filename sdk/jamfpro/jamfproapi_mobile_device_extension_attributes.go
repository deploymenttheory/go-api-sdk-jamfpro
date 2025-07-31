// jamfproapi_mobile_device_extension_attributes.go
// Jamf Pro  Api - Mobile Device Extension Attributes
// API reference: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriMobileDeviceExtensionAttributes = "/api/v1/mobile-device-extension-attributes"

// ResponseMobileDeviceExtensionAttributesList represents the paginated response for mobile device extension attributes
type ResponseMobileDeviceExtensionAttributesList struct {
	TotalCount int                                      `json:"totalCount"`
	Results    []ResourceMobileDeviceExtensionAttribute `json:"results"`
}

// ResourceMobileDeviceExtensionAttribute represents a mobile device extension attribute
type ResourceMobileDeviceExtensionAttribute struct {
	ID                            string   `json:"id"`
	Name                          string   `json:"name"`
	Description                   string   `json:"description,omitempty"`
	DataType                      string   `json:"dataType"`
	InventoryDisplayType          string   `json:"inventoryDisplayType"`
	InputType                     string   `json:"inputType"`
	PopupMenuChoices              []string `json:"popupMenuChoices,omitempty"`
	LDAPAttributeMapping          string   `json:"ldapAttributeMapping,omitempty"`
	LDAPExtensionAttributeAllowed *bool    `json:"ldapExtensionAttributeAllowed,omitempty"`
}

// ResponseMobileDeviceExtensionAttributeCreated represents the response for a created mobile device extension attribute
type ResponseMobileDeviceExtensionAttributeCreated struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// GetMobileDeviceExtensionAttributes retrieves all mobile device extension attributes with pagination
func (c *Client) GetMobileDeviceExtensionAttributes(params url.Values) (*ResponseMobileDeviceExtensionAttributesList, error) {
	resp, err := c.DoPaginatedGet(uriMobileDeviceExtensionAttributes, params)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "mobile device extension attributes", err)
	}

	var out ResponseMobileDeviceExtensionAttributesList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceMobileDeviceExtensionAttribute
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "mobile device extension attribute", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetMobileDeviceExtensionAttributeByID retrieves a mobile device extension attribute by its ID
func (c *Client) GetMobileDeviceExtensionAttributeByID(id string) (*ResourceMobileDeviceExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/%s", uriMobileDeviceExtensionAttributes, id)
	var attribute ResourceMobileDeviceExtensionAttribute
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attribute)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device extension attribute", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attribute, nil
}

// GetMobileDeviceExtensionAttributeByName retrieves a mobile device extension attribute by its name
func (c *Client) GetMobileDeviceExtensionAttributeByName(name string) (*ResourceMobileDeviceExtensionAttribute, error) {
	attributes, err := c.GetMobileDeviceExtensionAttributes(nil)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "mobile device extension attributes", err)
	}

	for _, value := range attributes.Results {
		if value.Name == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device extension attribute", name, errMsgNoName)
}

// CreateMobileDeviceExtensionAttribute creates a new mobile device extension attribute
func (c *Client) CreateMobileDeviceExtensionAttribute(attribute *ResourceMobileDeviceExtensionAttribute) (*ResponseMobileDeviceExtensionAttributeCreated, error) {
	endpoint := uriMobileDeviceExtensionAttributes
	var response ResponseMobileDeviceExtensionAttributeCreated

	resp, err := c.HTTP.DoRequest("POST", endpoint, attribute, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "mobile device extension attribute", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateMobileDeviceExtensionAttributeByID updates a mobile device extension attribute by its ID
func (c *Client) UpdateMobileDeviceExtensionAttributeByID(id string, attribute *ResourceMobileDeviceExtensionAttribute) (*ResourceMobileDeviceExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/%s", uriMobileDeviceExtensionAttributes, id)
	var updatedAttribute ResourceMobileDeviceExtensionAttribute
	resp, err := c.HTTP.DoRequest("PUT", endpoint, attribute, &updatedAttribute)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "mobile device extension attribute", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &updatedAttribute, nil
}

// UpdateMobileDeviceExtensionAttributeByName updates a mobile device extension attribute by its name
func (c *Client) UpdateMobileDeviceExtensionAttributeByName(name string, attribute *ResourceMobileDeviceExtensionAttribute) (*ResourceMobileDeviceExtensionAttribute, error) {
	target, err := c.GetMobileDeviceExtensionAttributeByName(name)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device extension attribute", name, err)
	}

	return c.UpdateMobileDeviceExtensionAttributeByID(target.ID, attribute)
}

// DeleteMobileDeviceExtensionAttributeByID deletes a mobile device extension attribute by its ID
func (c *Client) DeleteMobileDeviceExtensionAttributeByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriMobileDeviceExtensionAttributes, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device extension attribute", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceExtensionAttributeByName deletes a mobile device extension attribute by its name
func (c *Client) DeleteMobileDeviceExtensionAttributeByName(name string) error {
	target, err := c.GetMobileDeviceExtensionAttributeByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "mobile device extension attribute", name, err)
	}

	return c.DeleteMobileDeviceExtensionAttributeByID(target.ID)
}

// DeleteMultipleMobileDeviceExtensionAttributeByID deletes multiple mobile device extension attributes by their IDs
func (c *Client) DeleteMultipleMobileDeviceExtensionAttributeByID(ids []string) error {
	endpoint := fmt.Sprintf("%s/delete-multiple", uriMobileDeviceExtensionAttributes)

	// Create the request body
	requestBody := struct {
		IDs []string `json:"ids"`
	}{
		IDs: ids,
	}

	// Send the request
	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteMultiple, "mobile device extension attributes", ids, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}
