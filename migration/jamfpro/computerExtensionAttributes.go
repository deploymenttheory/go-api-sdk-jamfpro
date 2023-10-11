package jamfpro

import "fmt"

const uriComputerExtensionAttributes = "/JSSResource/computerextensionattributes"

type ComputerExtensionAttributeInputType struct {
	Type     string   `xml:"type"`
	Platform string   `xml:"platform,omitempty"`
	Script   string   `xml:"script,omitempty"`
	Choices  []string `xml:"popup_choices>choice,omitempty"`
}

type ComputerExtensionAttribute struct {
	Id               int                                 `xml:"id"`
	Name             string                              `xml:"name"`
	Enabled          bool                                `xml:"enabled,omitempty"`
	Description      string                              `xml:"description"`
	DataType         string                              `xml:"data_type"`
	InputType        ComputerExtensionAttributeInputType `xml:"input_type,omitempty"`
	InventoryDisplay string                              `xml:"inventory_display,omitempty"`
	ReconDisplay     string                              `xml:"recon_display,omitempty"`
}

type ComputerExtensionAttributeListResponse struct {
	Size    int                                 `xml:"size"`
	Results []MacOSConfigurationProfileListItem `xml:"computer_extension_attribute"`
}

type ComputerExtensionAttributeListItem struct {
	Id      int    `xml:"id"`
	Name    string `xml:"name"`
	Enabled bool   `xml:"enabled"`
}

func (c *Client) GetComputerExtensionAttribute(id int) (*ComputerExtensionAttribute, error) {
	var out *ComputerExtensionAttribute
	uri := fmt.Sprintf("%s/id/%v", uriComputerExtensionAttributes, id)
	err := c.DoRequest("GET", uri, nil, nil, &out)

	return out, err
}

func (c *Client) GetComputerExtensionAttributeByName(name string) (*ComputerExtensionAttribute, error) {
	var out *ComputerExtensionAttribute
	uri := fmt.Sprintf("%s/name/%v", uriComputerExtensionAttributes, name)
	err := c.DoRequest("GET", uri, nil, nil, &out)

	return out, err
}

func (c *Client) GetComputerExtensionAttributes() (*ComputerExtensionAttributeListResponse, error) {
	out := &ComputerExtensionAttributeListResponse{}
	err := c.DoRequest("GET", uriComputerExtensionAttributes, nil, nil, out)

	return out, err
}

func (c *Client) CreateComputerExtensionAttribute(d *ComputerExtensionAttribute) (int, error) {
	// set defaults
	if d.InputType.Type == "" {
		d.InputType.Type = "Text Field"
	}

	res := &ComputerExtensionAttributeListItem{}
	uri := fmt.Sprintf("%s/id/0", uriComputerExtensionAttributes)
	reqBody := &struct {
		*ComputerExtensionAttribute
		XMLName struct{} `xml:"computer_extension_attribute"`
	}{
		ComputerExtensionAttribute: d,
	}

	err := c.DoRequest("POST", uri, reqBody, nil, res)
	return res.Id, err
}

func (c *Client) UpdateComputerExtensionAttribute(d *ComputerExtensionAttribute) (int, error) {
	res := &ComputerExtensionAttributeListItem{}
	uri := fmt.Sprintf("%s/id/%v", uriComputerExtensionAttributes, d.Id)
	reqBody := &struct {
		*ComputerExtensionAttribute
		XMLName struct{} `xml:"computer_extension_attribute"`
	}{
		ComputerExtensionAttribute: d,
	}

	err := c.DoRequest("PUT", uri, reqBody, nil, res)

	return res.Id, err
}

func (c *Client) DeleteComputerExtensionAttribute(id int) (int, error) {
	attribute := &ComputerExtensionAttributeListItem{}
	uri := fmt.Sprintf("%s/id/%v", uriComputerExtensionAttributes, id)
	err := c.DoRequest("DELETE", uri, nil, nil, attribute)

	return attribute.Id, err
}
