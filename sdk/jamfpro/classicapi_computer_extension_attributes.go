// classicapi_computer_extension_attributes.go
// Jamf Pro Classic Api - Computer Extension Attributes
// api reference: https://developer.jamf.com/jamf-pro/reference/computerextensionattributes
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriComputerExtensionAttributes = "/JSSResource/computerextensionattributes"

// Structs for the computer extension attributes

// List

type ResponseComputerExtensionAttributesList struct {
	Size    int                                   `xml:"size"`
	Results []ComputerExtenstionAttributeListItem `xml:"computer_extension_attribute"`
}

type ComputerExtenstionAttributeListItem struct {
	ID      int    `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

// Resource

type ResourceComputerExtensionAttribute struct {
	ID               int                                       `xml:"id"`
	Name             string                                    `xml:"name"`
	Enabled          bool                                      `xml:"enabled,omitempty"`
	Description      string                                    `xml:"description,omitempty"`
	DataType         string                                    `xml:"data_type,omitempty"`
	InputType        ComputerExtensionAttributeSubsetInputType `xml:"input_type"`
	InventoryDisplay string                                    `xml:"inventory_display,omitempty"`
	ReconDisplay     string                                    `xml:"recon_display,omitempty"`
}

// Subsets

type ComputerExtensionAttributeSubsetInputType struct {
	Type     string   `xml:"type"`
	Platform string   `xml:"platform,omitempty"`
	Script   string   `xml:"script,omitempty"`
	Choices  []string `xml:"popup_choices>choice,omitempty"`
}

// CRUD

// GetComputerExtensionAttributes gets a list of all computer extension attributes
func (c *Client) GetComputerExtensionAttributes() (*ResponseComputerExtensionAttributesList, error) {
	endpoint := uriComputerExtensionAttributes

	var attributes ResponseComputerExtensionAttributesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attributes)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "computer extension attributes", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attributes, nil
}

// GetComputerExtensionAttributeByID retrieves a computer extension attribute by its ID.
func (c *Client) GetComputerExtensionAttributeByID(id int) (*ResourceComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerExtensionAttributes, id)

	var attribute ResourceComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "computer extension attribute", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attribute, nil
}

// GetComputerExtensionAttributeByName retrieves a computer extension attribute by its name.
func (c *Client) GetComputerExtensionAttributeByName(name string) (*ResourceComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerExtensionAttributes, name)

	var attribute ResourceComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "computer extension attribute", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attribute, nil
}

// CreateComputerExtensionAttribute creates a new computer extension attribute.
func (c *Client) CreateComputerExtensionAttribute(attribute *ResourceComputerExtensionAttribute) (*ResourceComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriComputerExtensionAttributes)

	requestBody := struct {
		XMLName xml.Name `xml:"computer_extension_attribute"`
		*ResourceComputerExtensionAttribute
	}{
		ResourceComputerExtensionAttribute: attribute,
	}

	var createdAttribute ResourceComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdAttribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "computer extension attribute", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdAttribute, nil
}

// UpdateComputerExtensionAttributeByID updates an existing computer extension attribute by its ID.
func (c *Client) UpdateComputerExtensionAttributeByID(id int, attribute *ResourceComputerExtensionAttribute) (*ResourceComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerExtensionAttributes, id)

	requestBody := struct {
		XMLName xml.Name `xml:"computer_extension_attribute"`
		*ResourceComputerExtensionAttribute
	}{
		ResourceComputerExtensionAttribute: attribute,
	}

	var updatedAttribute ResourceComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedAttribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "computer extension attribute", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAttribute, nil
}

// UpdateComputerExtensionAttributeByName updates a computer extension attribute by its name.
func (c *Client) UpdateComputerExtensionAttributeByName(name string, attribute *ResourceComputerExtensionAttribute) (*ResourceComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerExtensionAttributes, name)

	requestBody := struct {
		XMLName xml.Name `xml:"computer_extension_attribute"`
		*ResourceComputerExtensionAttribute
	}{
		ResourceComputerExtensionAttribute: attribute,
	}

	var updatedAttribute ResourceComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedAttribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "computer extension attribute", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAttribute, nil
}

// DeleteComputerExtensionAttributeByID deletes a computer extension attribute by its ID.
func (c *Client) DeleteComputerExtensionAttributeByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriComputerExtensionAttributes, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "computer extension attribute", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteComputerExtensionAttributeByNameByID deletes a computer extension attribute by its name.
// func (c *Client) DeleteComputerExtensionAttributeByNameByID(name string) error {
// 	attributes, err := c.GetComputerExtensionAttributes()
// 	if err != nil {
// 		return fmt.Errorf(errMsgFailedDeleteByName, "computer extension attribute", name, err)
// 	}

// 	var attributeID int
// 	for _, attribute := range attributes.Results {
// 		if attribute.Name == name {
// 			attributeID = attribute.ID
// 			break
// 		}
// 	}

// 	if attributeID == 0 {
// 		return fmt.Errorf(errMsgFailedDeleteByName, "computer extension attribute", name, err)
// 	}

// 	return c.DeleteComputerExtensionAttributeByID(attributeID)
// }
