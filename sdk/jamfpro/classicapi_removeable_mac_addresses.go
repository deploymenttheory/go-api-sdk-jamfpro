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
	Size         int                           `xml:"size"`
	RemovableMac []ResourceRemovableMacAddress `xml:"removable_mac_address"`
}

type ResourceRemovableMacAddress struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// CRUD

// GetRemovableMACAddresses retrieves a list of all removable MAC addresses.
func (c *Client) GetRemovableMACAddresses() (*ResponseRemovableMacAddressesList, error) {
	endpoint := uriRemovableMacAddresses

	var macAddressesList ResponseRemovableMacAddressesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macAddressesList, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "removeable macaddresses", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macAddressesList, nil
}

// GetRemovableMACAddressByID retrieves the details of a removable MAC address by its ID.
func (c *Client) GetRemovableMACAddressByID(id int) (*ResourceRemovableMacAddress, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriRemovableMacAddresses, id)

	var macAddressDetails ResourceRemovableMacAddress
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macAddressDetails, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "removeable macaddress", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macAddressDetails, nil
}

// GetRemovableMACAddressByName retrieves the details of a removable MAC address by its name.
func (c *Client) GetRemovableMACAddressByName(name string) (*ResourceRemovableMacAddress, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriRemovableMacAddresses, name)

	var macAddressDetails ResourceRemovableMacAddress
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macAddressDetails, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "removeable macaddress", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macAddressDetails, nil
}

// CreateRemovableMACAddress creates a new removable MAC address.
func (c *Client) CreateRemovableMACAddress(macAddress *ResourceRemovableMacAddress) (*ResourceRemovableMacAddress, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriRemovableMacAddresses, macAddress.ID)

	// Wrap the removable MAC address with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"removable_mac_address"`
		*ResourceRemovableMacAddress
	}{
		ResourceRemovableMacAddress: macAddress,
	}

	var responseMacAddress ResourceRemovableMacAddress
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseMacAddress, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "removeable macaddress", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseMacAddress, nil
}

// UpdateRemovableMACAddressByID updates an existing removable MAC address by its ID.
func (c *Client) UpdateRemovableMACAddressByID(id int, macAddress *ResourceRemovableMacAddress) (*ResourceRemovableMacAddress, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriRemovableMacAddresses, id)

	requestBody := struct {
		XMLName xml.Name `xml:"removable_mac_address"`
		*ResourceRemovableMacAddress
	}{
		ResourceRemovableMacAddress: macAddress,
	}

	var responseMacAddress ResourceRemovableMacAddress
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseMacAddress, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "removeable macaddress", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseMacAddress, nil
}

// UpdateRemovableMACAddressByName updates an existing removable MAC address by its name.
func (c *Client) UpdateRemovableMACAddressByName(name string, macAddress *ResourceRemovableMacAddress) (*ResourceRemovableMacAddress, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriRemovableMacAddresses, name)

	requestBody := struct {
		XMLName xml.Name `xml:"removable_mac_address"`
		*ResourceRemovableMacAddress
	}{
		ResourceRemovableMacAddress: macAddress,
	}

	var responseMacAddress ResourceRemovableMacAddress
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseMacAddress, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "removeable macaddress", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseMacAddress, nil
}

// DeleteRemovableMACAddressByID deletes a removable MAC address by its ID.
func (c *Client) DeleteRemovableMACAddressByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriRemovableMacAddresses, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "removeable macaddress", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteRemovableMACAddressByName deletes a removable MAC address by its name.
func (c *Client) DeleteRemovableMACAddressByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriRemovableMacAddresses, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "removeable macaddress", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
