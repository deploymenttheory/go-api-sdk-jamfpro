// classicapi_buildings.go
// Jamf Pro Classic Api - Buildings
// api reference: https://developer.jamf.com/jamf-pro/reference/buildings
// Classic API requires the structs to support an XML data structure.

package jamfpro

import "fmt"

const uriAPIBuildings = "/JSSResource/buildings"

// For multiple buildings, the structure remains the same.
type ResponseBuildings struct {
	Size      int              `xml:"size"`
	Buildings []BuildingDetail `xml:"building"`
}

type BuildingDetail struct {
	ID             int    `json:"id,omitempty" xml:"id"`
	Name           string `json:"name,omitempty" xml:"name"`
	StreetAddress1 string `json:"streetAddress1,omitempty" xml:"streetAddress1,omitempty"`
	StreetAddress2 string `json:"streetAddress2,omitempty" xml:"streetAddress2,omitempty"`
	City           string `json:"city,omitempty" xml:"city,omitempty"`
	StateProvince  string `json:"stateProvince,omitempty" xml:"stateProvince,omitempty"`
	ZipPostalCode  string `json:"zipPostalCode,omitempty" xml:"zipPostalCode,omitempty"`
	Country        string `json:"country,omitempty" xml:"country,omitempty"`
	Href           string `json:"href,omitempty" xml:"href,omitempty"`
}

// ResponseBuilding is now directly the BuildingDetail
type ResponseBuilding BuildingScope

type BuildingScope struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// GetBuildings retrieves all buildings
func (c *Client) GetBuildings() (*ResponseBuildings, error) {
	endpoint := uriAPIBuildings

	var buildingsList ResponseBuildings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &buildingsList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch buildings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &buildingsList, nil
}

// GetBuildingByID retrieves the building by its ID
func (c *Client) GetBuildingByID(id string) (*BuildingDetail, error) {
	endpoint := fmt.Sprintf("%s/id/%s", uriAPIBuildings, id)

	var building BuildingDetail
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &building)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch building by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &building, nil
}

// GetBuildingByName retrieves the building by its name
func (c *Client) GetBuildingByName(name string) (*Building, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIBuildings, name)

	var building Building
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &building)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch building by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &building, nil
}

// GetBuildingIdByName retrieves the building ID by its name
func (c *Client) GetBuildingIdByName(name string) (string, error) {
	// Fetch the list of all buildings
	buildingsList, err := c.GetBuildings()
	if err != nil {
		return "", err
	}

	// Iterate through the list to find the building with the given name and return its ID
	for _, building := range buildingsList.Buildings {
		if building.Name == name {
			return fmt.Sprintf("%d", building.ID), nil
		}
	}

	// If building with the specified name isn't found, return an error
	return "", fmt.Errorf("building with name %s not found", name)
}

// CreateBuilding creates a new building
func (c *Client) CreateBuilding(building *Building) (*Building, error) {
	endpoint := uriAPIBuildings

	var response Building
	resp, err := c.HTTP.DoRequest("POST", endpoint, building, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create building: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateBuildingByID updates an existing building
func (c *Client) UpdateBuildingByID(building *Building) (*Building, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIBuildings, building.ID)

	var updatedBuilding Building
	resp, err := c.HTTP.DoRequest("PUT", endpoint, building, &updatedBuilding)
	if err != nil {
		return nil, fmt.Errorf("failed to update building: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedBuilding, nil
}

// UpdateBuildingByName updates an existing building by its name
func (c *Client) UpdateBuildingByName(building *Building) (*Building, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIBuildings, building.Name)

	var updatedBuilding Building
	resp, err := c.HTTP.DoRequest("PUT", endpoint, building, &updatedBuilding)
	if err != nil {
		return nil, fmt.Errorf("failed to update building by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedBuilding, nil
}

// DeleteBuildingByID deletes an existing building by its ID
func (c *Client) DeleteBuildingByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIBuildings, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete building by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteBuildingByName deletes an existing building by its name
func (c *Client) DeleteBuildingByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIBuildings, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete building by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
