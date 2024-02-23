package jamfpro

const uriInventoryInformation = "/api/v1/inventory-information"

type ResponseInventoryInformation struct {
	ManagedComputers   int `json:"managedComputers"`
	UnmanagedComputers int `json:"unmanagedComputers"`
	ManagedDevices     int `json:"managedDevices"`
	UnmanagedDevices   int `json:"unmanagedDevices"`
}

func (c *Client) GetInventoryInformation() (*ResponseInventoryInformation, error) {
	var out *ResponseInventoryInformation
	err := c.DoRequest("GET", uriJamfProInformation, nil, nil, &out)
	return out, err
}
