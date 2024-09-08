package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriComputerExtensionAttributes = "/api/v1/computer-extension-attributes"

// ResponseComputerExtensionAttributesList represents the paginated response for computer extension attributes
type ResponseComputerExtensionAttributesList struct {
	TotalCount int                                  `json:"totalCount"`
	Results    []ResourceComputerExtensionAttribute `json:"results"`
}

// ResourceComputerExtensionAttribute represents a computer extension attribute
type ResourceComputerExtensionAttribute struct {
	ID                            string   `json:"id"`
	Name                          string   `json:"name"`
	Description                   string   `json:"description"`
	DataType                      string   `json:"dataType"`
	Enabled                       bool     `json:"enabled"`
	InventoryDisplayType          string   `json:"inventoryDisplayType"`
	InputType                     string   `json:"inputType"`
	ScriptContents                string   `json:"scriptContents"`
	PopupMenuChoices              []string `json:"popupMenuChoices"`
	LDAPAttributeMapping          string   `json:"ldapAttributeMapping"`
	LDAPExtensionAttributeAllowed bool     `json:"ldapExtensionAttributeAllowed"`
}

// ResponseComputerExtensionAttributeCreated represents the response for a created computer extension attribute
type ResponseComputerExtensionAttributeCreated struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// GetComputerExtensionAttributes retrieves all computer extension attributes with pagination
func (c *Client) GetComputerExtensionAttributes(sortFilter string) (*ResponseComputerExtensionAttributesList, error) {
	resp, err := c.DoPaginatedGet(
		uriComputerExtensionAttributes,
		standardPageSize,
		startingPageNumber,
		sortFilter,
	)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "computer extension attributes", err)
	}

	var out ResponseComputerExtensionAttributesList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceComputerExtensionAttribute
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "computer extension attribute", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetComputerExtensionAttributeByID retrieves a computer extension attribute by its ID
func (c *Client) GetComputerExtensionAttributeByID(id string) (*ResourceComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/%s", uriComputerExtensionAttributes, id)
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

// GetComputerExtensionAttributeByName retrieves a computer extension attribute by its name
func (c *Client) GetComputerExtensionAttributeByName(name string) (*ResourceComputerExtensionAttribute, error) {
	attributes, err := c.GetComputerExtensionAttributes("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "computer extension attributes", err)
	}

	for _, value := range attributes.Results {
		if value.Name == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "computer extension attribute", name, errMsgNoName)
}

// CreateComputerExtensionAttribute creates a new computer extension attribute
func (c *Client) CreateComputerExtensionAttribute(attribute *ResourceComputerExtensionAttribute) (*ResponseComputerExtensionAttributeCreated, error) {
	endpoint := uriComputerExtensionAttributes
	var response ResponseComputerExtensionAttributeCreated

	resp, err := c.HTTP.DoRequest("POST", endpoint, attribute, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "computer extension attribute", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateComputerExtensionAttributeByID updates a computer extension attribute by its ID
func (c *Client) UpdateComputerExtensionAttributeByID(id string, attribute *ResourceComputerExtensionAttribute) (*ResourceComputerExtensionAttribute, error) {
	endpoint := fmt.Sprintf("%s/%s", uriComputerExtensionAttributes, id)
	var updatedAttribute ResourceComputerExtensionAttribute
	resp, err := c.HTTP.DoRequest("PUT", endpoint, attribute, &updatedAttribute)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "computer extension attribute", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &updatedAttribute, nil
}

// UpdateComputerExtensionAttributeByName updates a computer extension attribute by its name
func (c *Client) UpdateComputerExtensionAttributeByName(name string, attribute *ResourceComputerExtensionAttribute) (*ResourceComputerExtensionAttribute, error) {
	target, err := c.GetComputerExtensionAttributeByName(name)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "computer extension attribute", name, err)
	}

	return c.UpdateComputerExtensionAttributeByID(target.ID, attribute)
}

// DeleteComputerExtensionAttributeByID deletes a computer extension attribute by its ID
func (c *Client) DeleteComputerExtensionAttributeByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriComputerExtensionAttributes, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "computer extension attribute", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteComputerExtensionAttributeByName deletes a computer extension attribute by its name
func (c *Client) DeleteComputerExtensionAttributeByName(name string) error {
	target, err := c.GetComputerExtensionAttributeByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "computer extension attribute", name, err)
	}

	return c.DeleteComputerExtensionAttributeByID(target.ID)
}
