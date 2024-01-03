// classicapi_dock_items.go
// Jamf Pro Classic Api - Dock Items
// api reference: https://developer.jamf.com/jamf-pro/reference/dockitems
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// URI for Dock Items in Jamf Pro API
const uriDockItems = "/JSSResource/dockitems"

// List

// Struct to capture the XML response for dock items list
type ResponseDockItemsList struct {
	Size      int                `xml:"size"`
	DockItems []DockItemListItem `xml:"dock_item"`
}

type DockItemListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// Struct to capture the response for a single Dock Item
type ResourceDockItem struct {
	ID       int    `xml:"id" json:"id"`
	Name     string `xml:"name" json:"name"`
	Type     string `xml:"type" json:"type"`
	Path     string `xml:"path" json:"path"`
	Contents string `xml:"contents" json:"contents"`
}

// CRUD

// GetDockItems retrieves a serialized list of dock items.
func (c *Client) GetDockItems() (*ResponseDockItemsList, error) {
	endpoint := uriDockItems

	var dockItems ResponseDockItemsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &dockItems)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Dock Items: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &dockItems, nil
}

// GetDockItemsByID retrieves a single dock item by its ID.
func (c *Client) GetDockItemByID(id int) (*ResourceDockItem, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriDockItems, id)

	var dockItem ResourceDockItem
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &dockItem)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Dock Item by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &dockItem, nil
}

// GetDockItemsByName retrieves a single dock item by its name.
func (c *Client) GetDockItemByName(name string) (*ResourceDockItem, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDockItems, name)

	var dockItem ResourceDockItem
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &dockItem)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Dock Item by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &dockItem, nil
}

// CreateDockItems creates a new dock item.
func (c *Client) CreateDockItem(dockItem *ResourceDockItem) (*ResourceDockItem, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriDockItems) // Typically, APIs use '0' or a similar placeholder for creation endpoints

	requestBody := struct {
		XMLName xml.Name `xml:"dock_item"`
		*ResourceDockItem
	}{
		ResourceDockItem: dockItem,
	}

	var createdDockItem ResourceDockItem
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdDockItem)
	if err != nil {
		return nil, fmt.Errorf("failed to create Dock Item: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdDockItem, nil
}

// UpdateDockItemByID updates a dock item by its ID.
func (c *Client) UpdateDockItemByID(id int, dockItem *ResourceDockItem) (*ResourceDockItem, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriDockItems, id)

	requestBody := struct {
		XMLName xml.Name `xml:"dock_item"`
		*ResourceDockItem
	}{
		ResourceDockItem: dockItem,
	}

	var updatedDockItem ResourceDockItem
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedDockItem)
	if err != nil {
		return nil, fmt.Errorf("failed to update Dock Item by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedDockItem, nil
}

// UpdateDockItemByName updates a dock item by its name.
func (c *Client) UpdateDockItemByName(name string, dockItem *ResourceDockItem) (*ResourceDockItem, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDockItems, name)

	requestBody := struct {
		XMLName xml.Name `xml:"dock_item"`
		*ResourceDockItem
	}{
		ResourceDockItem: dockItem,
	}

	var updatedDockItem ResourceDockItem
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedDockItem)
	if err != nil {
		return nil, fmt.Errorf("failed to update Dock Item by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedDockItem, nil
}

// DeleteDockItemsByID deletes a dock item by its ID.
func (c *Client) DeleteDockItemByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriDockItems, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Dock Item by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDockItemsByName deletes a dock item by its name.
func (c *Client) DeleteDockItemByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriDockItems, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Dock Item by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
