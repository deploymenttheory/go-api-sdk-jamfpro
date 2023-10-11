package jamfpro

import "fmt"

const uriBuildings = "/api/v1/buildings"

type ResponseBuildings struct {
	TotalCount *int       `json:"totalCount,omitempty"`
	Results    []Building `json:"results,omitempty"`
}

type Building struct {
	Id             *string `json:"id,omitempty"` // The response type to be returned is a string
	Name           *string `json:"name,omitempty"`
	StreetAddress1 *string `json:"streetAddress1,omitempty"`
	StreetAddress2 *string `json:"streetAddress2,omitempty"`
	City           *string `json:"city,omitempty"`
	StateProvince  *string `json:"stateProvince,omitempty"`
	ZipPostalCode  *string `json:"zipPostalCode,omitempty"`
	Country        *string `json:"country,omitempty"`
	Href           *string `json:"href,omitempty"`
}

type BuildingScope struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

func (c *Client) GetBuildingIdByName(name string) (string, error) {
	var id string
	d, err := c.GetBuildings()
	if err != nil {
		return "", err
	}

	for _, v := range d.Results {
		if *v.Name == name {
			id = *v.Id
			break
		}
	}
	return id, err
}

func (c *Client) GetBuildingByName(name string) (*Building, error) {
	id, err := c.GetBuildingIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetBuilding(id)
}

func (c *Client) GetBuilding(id string) (*Building, error) {
	var out *Building
	uri := fmt.Sprintf("%s/%s", uriBuildings, id)
	err := c.DoRequest("GET", uri, nil, nil, &out)

	return out, err
}

func (c *Client) GetBuildings() (*ResponseBuildings, error) {
	out := &ResponseBuildings{}
	err := c.DoRequest("GET", uriBuildings, nil, nil, &out)

	return out, err
}

func (c *Client) CreateBuilding(name, sa1, sa2, city, sp, zpc, country *string) (*Building, error) {
	in := struct {
		Name           *string `json:"name"`
		StreetAddress1 *string `json:"streetAddress1"`
		StreetAddress2 *string `json:"streetAddress2"`
		City           *string `json:"city"`
		StateProvince  *string `json:"stateProvince"`
		ZipPostalCode  *string `json:"zipPostalCode"`
		Country        *string `json:"country"`
	}{
		Name:           name,
		StreetAddress1: sa1,
		StreetAddress2: sa2,
		City:           city,
		StateProvince:  sp,
		ZipPostalCode:  zpc,
		Country:        country,
	}

	var out *Building

	err := c.DoRequest("POST", uriBuildings, in, nil, &out)
	return out, err
}

func (c *Client) UpdateBuilding(d *Building) (*Building, error) {
	var out *Building
	uri := fmt.Sprintf("%s/%s", uriBuildings, *d.Id)

	in := struct {
		Name           *string `json:"name"`
		StreetAddress1 *string `json:"streetAddress1"`
		StreetAddress2 *string `json:"streetAddress2"`
		City           *string `json:"city"`
		StateProvince  *string `json:"stateProvince"`
		ZipPostalCode  *string `json:"zipPostalCode"`
		Country        *string `json:"country"`
	}{
		Name:           d.Name,
		StreetAddress1: d.StreetAddress1,
		StreetAddress2: d.StreetAddress2,
		City:           d.City,
		StateProvince:  d.StateProvince,
		ZipPostalCode:  d.ZipPostalCode,
		Country:        d.Country,
	}

	err := c.DoRequest("PUT", uri, in, nil, &out)
	return out, err
}

func (c *Client) DeleteBuilding(name string) error {
	id, err := c.GetBuildingIdByName(name)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s/%s", uriBuildings, id)
	return c.DoRequest("DELETE", uri, nil, nil, nil)
}
