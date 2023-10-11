// Jamf Pro Classic API

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIDockItems = "/JSSResource/dockitems"

// DockItem Response structure
type ResponseDockItem struct {
	ID       int    `xml:"id"`
	Name     string `xml:"name"`
	Type     string `xml:"type"`
	Path     string `xml:"path"`
	Contents string `xml:"contents"`
}

// Response for getting all dock items
type ResponseDockItemsList struct {
	Size      int               `xml:"size"`
	DockItems []DockItemGeneral `xml:"dock_item"`
}

type DockItemGeneral struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Create / Update - Dock Item

type DockItem struct {
	XMLName  xml.Name `xml:"dock_item"`
	Name     string   `xml:"name"`
	Type     string   `xml:"type"`
	Path     string   `xml:"path"`
	Contents string   `xml:"contents"`
}

// GetDockItemByID retrieves the Dock Item by its ID
func (c *Client) GetDockItemByID(id int) (*ResponseDockItem, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIDockItems, id)

	var item ResponseDockItem
	if err := c.DoRequest("GET", url, nil, nil, &item); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &item, nil
}

// GetDockItems retrieves a list of all Dock Items
func (c *Client) GetDockItems() (*ResponseDockItemsList, error) {
	url := uriAPIDockItems

	var dockItemList ResponseDockItemsList

	if err := c.DoRequest("GET", url, nil, nil, &dockItemList); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &dockItemList, nil
}

// GetDockItemByName retrieves the Dock Item by its name
func (c *Client) GetDockItemByName(name string) (*ResponseDockItem, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIDockItems, name)

	var item ResponseDockItem
	if err := c.DoRequest("GET", url, nil, nil, &item); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &item, nil
}

// CreateDockItem creates a new Dock Item
func (c *Client) CreateDockItem(item *DockItem) (*ResponseDockItem, error) {
	url := uriAPIDockItems

	var responseItem ResponseDockItem
	if err := c.DoRequest("POST", url, item, nil, &responseItem); err != nil {
		return nil, fmt.Errorf("failed to create Dock Item: %v", err)
	}

	return &responseItem, nil
}

// UpdateDockItem updates an existing Dock Item By Id
func (c *Client) UpdateDockItemById(id int, item *DockItem) (*ResponseDockItem, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIDockItems, id)

	var responseItem ResponseDockItem
	if err := c.DoRequest("PUT", url, item, nil, &responseItem); err != nil {
		return nil, fmt.Errorf("failed to update Dock Item: %v", err)
	}

	return &responseItem, nil
}

// UpdateDockItemByName updates an existing Dock Item by its name
func (c *Client) UpdateDockItemByName(name string, item *DockItem) (*ResponseDockItem, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIDockItems, name)

	var responseItem ResponseDockItem
	if err := c.DoRequest("PUT", url, item, nil, &responseItem); err != nil {
		return nil, fmt.Errorf("failed to update Dock Item named '%s': %v", name, err)
	}

	return &responseItem, nil
}

// DeleteDockItemByID deletes a Dock Item by its ID
func (c *Client) DeleteDockItemByID(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIDockItems, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Dock Item: %v", err)
	}

	return nil
}

// DeleteDockItemByName deletes a Dock Item by its name
func (c *Client) DeleteDockItemByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIDockItems, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Dock Item named '%s': %v", name, err)
	}

	return nil
}
