// classicapi_removeable_mac_addresses.go
// Jamf Pro Classic Api - Removable Mac Addresses
// API reference: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddresses
// Jamf Pro Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriRemovableMacAddresses = "/JSSResource/removablemacaddresses"

// Structs for Removable MAC Addresses List
type ResponseRemovableMacAddressesList struct {
	Size         int                   `xml:"size"`
	RemovableMac []RemovableMacAddress `xml:"removable_mac_address"`
}

type RemovableMacAddress struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Struct for individual Removable MAC Address
type ResponseRemovableMACAddress struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// GetRemovableMACAddresses retrieves a list of all removable MAC addresses.
func (c *Client) GetRemovableMACAddresses() (*ResponseRemovableMacAddressesList, error) {
	endpoint := uriRemovableMacAddresses

	var macAddressesList ResponseRemovableMacAddressesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macAddressesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all removable MAC addresses: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macAddressesList, nil
}

// GetRemovableMACAddressByID retrieves the details of a removable MAC address by its ID.
func (c *Client) GetRemovableMACAddressByID(id int) (*ResponseRemovableMACAddress, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriRemovableMacAddresses, id)

	var macAddressDetails ResponseRemovableMACAddress
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macAddressDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch removable MAC address by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macAddressDetails, nil
}

// GetRemovableMACAddressByName retrieves the details of a removable MAC address by its name.
func (c *Client) GetRemovableMACAddressByName(name string) (*ResponseRemovableMACAddress, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriRemovableMacAddresses, name)

	var macAddressDetails ResponseRemovableMACAddress
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macAddressDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch removable MAC address by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macAddressDetails, nil
}

// CreateRemovableMACAddress creates a new removable MAC address.
func (c *Client) CreateRemovableMACAddress(macAddress *RemovableMacAddress) (*RemovableMacAddress, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriRemovableMacAddresses, macAddress.ID)

	// Wrap the removable MAC address with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"removable_mac_address"`
		*RemovableMacAddress
	}{
		RemovableMacAddress: macAddress,
	}

	var responseMacAddress RemovableMacAddress
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseMacAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to create removable MAC address: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseMacAddress, nil
}

// UpdateRemovableMACAddressByID updates an existing removable MAC address by its ID.
func (c *Client) UpdateRemovableMACAddressByID(id int, macAddress *RemovableMacAddress) (*RemovableMacAddress, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriRemovableMacAddresses, id)

	requestBody := struct {
		XMLName xml.Name `xml:"removable_mac_address"`
		*RemovableMacAddress
	}{
		RemovableMacAddress: macAddress,
	}

	var responseMacAddress RemovableMacAddress
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseMacAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to update removable MAC address by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseMacAddress, nil
}

// UpdateRemovableMACAddressByName updates an existing removable MAC address by its name.
func (c *Client) UpdateRemovableMACAddressByName(name string, macAddress *RemovableMacAddress) (*RemovableMacAddress, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriRemovableMacAddresses, name)

	requestBody := struct {
		XMLName xml.Name `xml:"removable_mac_address"`
		*RemovableMacAddress
	}{
		RemovableMacAddress: macAddress,
	}

	var responseMacAddress RemovableMacAddress
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseMacAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to update removable MAC address by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseMacAddress, nil
}

// DeleteRemovableMACAddressByID deletes a removable MAC address by its ID.
func (c *Client) DeleteRemovableMACAddressByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriRemovableMacAddresses, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete removable MAC address by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteRemovableMACAddressByName deletes a removable MAC address by its name.
func (c *Client) DeleteRemovableMACAddressByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriRemovableMacAddresses, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete removable MAC address by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
