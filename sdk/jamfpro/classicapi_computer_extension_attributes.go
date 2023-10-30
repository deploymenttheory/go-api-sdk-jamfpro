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

type ResponseComputerExtensionAttributesList struct {
	Size    int                              `xml:"size"`
	Results []ComputerExtensionAttributeItem `xml:"computer_extension_attribute"`
}

type ComputerExtensionAttributeItem struct {
	ID      int    `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	Enabled bool   `xml:"enabled,omitempty"`
}

type ResponseComputerExtensionAttribute struct {
	ID               int                                 `xml:"id"`
	Name             string                              `xml:"name"`
	Enabled          bool                                `xml:"enabled,omitempty"`
	Description      string                              `xml:"description,omitempty"`
	DataType         string                              `xml:"data_type,omitempty"`
	InputType        ComputerExtensionAttributeInputType `xml:"input_type"`
	InventoryDisplay string                              `xml:"inventory_display,omitempty"`
	ReconDisplay     string                              `xml:"recon_display,omitempty"`
}

type ComputerExtensionAttributeInputType struct {
	Type     string   `xml:"type"`
	Platform string   `xml:"platform,omitempty"`
	Script   string   `xml:"script,omitempty"`
	Choices  []string `xml:"popup_choices>choice,omitempty"`
}

// GetComputerExtensionAttributes gets a list of all computer extension attributes
func (c *Client) GetComputerExtensionAttributes() (*ResponseComputerExtensionAttributesList, error) {
	endpoint := uriComputerExtensionAttributes

	var attributes ResponseComputerExtensionAttributesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attributes)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Computer Extension Attributes: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attributes, nil
}

// GetComputerExtensionAttributeByID retrieves a computer extension attribute by its ID.
func (c *Client) GetComputerExtensionAttributeByID(id int) (*ResponseComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerExtensionAttributes, id)

	var attribute ResponseComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attribute)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Extension Attribute by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attribute, nil
}

// GetComputerExtensionAttributeByName retrieves a computer extension attribute by its name.
func (c *Client) GetComputerExtensionAttributeByName(name string) (*ResponseComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerExtensionAttributes, name)

	var attribute ResponseComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &attribute)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Extension Attribute by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &attribute, nil
}

// CreateComputerExtensionAttribute creates a new computer extension attribute.
func (c *Client) CreateComputerExtensionAttribute(attribute *ResponseComputerExtensionAttribute) (*ResponseComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriComputerExtensionAttributes) // Using ID 0 for creation as per the pattern

	// Wrap the attribute request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"computer_extension_attribute"`
		*ResponseComputerExtensionAttribute
	}{
		ResponseComputerExtensionAttribute: attribute,
	}

	var createdAttribute ResponseComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdAttribute)
	if err != nil {
		return nil, fmt.Errorf("failed to create Computer Extension Attribute: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdAttribute, nil
}

// UpdateComputerExtensionAttributeByID updates an existing computer extension attribute by its ID.
func (c *Client) UpdateComputerExtensionAttributeByID(id int, attribute *ResponseComputerExtensionAttribute) (*ResponseComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerExtensionAttributes, id)

	// Wrap the attribute request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"computer_extension_attribute"`
		*ResponseComputerExtensionAttribute
	}{
		ResponseComputerExtensionAttribute: attribute,
	}

	var updatedAttribute ResponseComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedAttribute)
	if err != nil {
		return nil, fmt.Errorf("failed to update Computer Extension Attribute by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAttribute, nil
}

// UpdateComputerExtensionAttributeByName updates a computer extension attribute by its name.
func (c *Client) UpdateComputerExtensionAttributeByName(name string, attribute *ResponseComputerExtensionAttribute) (*ResponseComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerExtensionAttributes, name)

	// Wrap the attribute request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"computer_extension_attribute"`
		*ResponseComputerExtensionAttribute
	}{
		ResponseComputerExtensionAttribute: attribute,
	}

	var updatedAttribute ResponseComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedAttribute)
	if err != nil {
		return nil, fmt.Errorf("failed to update Computer Extension Attribute by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedAttribute, nil
}

// DeleteComputerExtensionAttributeByID deletes a computer extension attribute by its ID.
func (c *Client) DeleteComputerExtensionAttributeByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerExtensionAttributes, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Computer Extension Attribute by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteComputerExtensionAttributeByNameByID deletes a computer extension attribute by its name.
// there is no url directly for deletion by resource name. so it is resolved in a two step process.
func (c *Client) DeleteComputerExtensionAttributeByNameByID(name string) error {
	// Step 1: Fetch all extension attributes to find the ID for the given name
	attributes, err := c.GetComputerExtensionAttributes()
	if err != nil {
		return fmt.Errorf("failed to fetch Computer Extension Attributes: %v", err)
	}

	var attributeID int
	for _, attribute := range attributes.Results {
		if attribute.Name == name {
			attributeID = attribute.ID
			break
		}
	}

	if attributeID == 0 {
		return fmt.Errorf("no Computer Extension Attribute found with name: %s", name)
	}

	// Step 2: Use the discovered ID to delete the attribute
	return c.DeleteComputerExtensionAttributeByID(attributeID)
}
