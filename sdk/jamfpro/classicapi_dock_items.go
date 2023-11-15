// classicapi_dock_items.go
// Jamf Pro Classic Api - Dock Items
// api reference: https://developer.jamf.com/jamf-pro/reference/dockitems
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"fmt"
)

// URI for Dock Items in Jamf Pro API
const uriDockItems = "/JSSResource/dockitems"

// Struct to capture the XML response for dock items list
type ResponseDockItemsList struct {
	Size     int        `xml:"size"`
	DockItem []DockItem `xml:"dock_item"`
}

type DockItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

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
