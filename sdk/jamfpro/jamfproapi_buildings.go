// jamfapi_buildings.go
// Jamf Pro Api - Buildings
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-buildings
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriBuildings = "/api/v1/buildings"

// List

// ResponseBuildings represents the structure of the response for the buildings list.
type ResponseBuildingsList struct {
	TotalCount int                `json:"totalCount"`
	Results    []ResourceBuilding `json:"results"`
}

// Responses

// ResponseBuildingResourceHistoryList represents the structure of the response for the building resource history list.
type ResponseBuildingResourceHistoryList struct {
	Size    int                               `json:"totalCount"`
	Results []ResourceBuildingResourceHistory `json:"results"`
}

// ResponseBuildingCreate represents the response structure for creating a building.
type ResponseBuildingCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

// ResponseBuilding represents the structure of each building item in the response.
type ResourceBuilding struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	StreetAddress1 string `json:"streetAddress1"`
	StreetAddress2 string `json:"streetAddress2"`
	City           string `json:"city"`
	StateProvince  string `json:"stateProvince"`
	ZipPostalCode  string `json:"zipPostalCode"`
	Country        string `json:"country"`
}

// ResponseBuildingResourceHistory represents the structure of each resource history item in the response.
type ResourceBuildingResourceHistory struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// CRUD

// GetBuildings retrieves all building information with optional sorting.
func (c *Client) GetBuildings(params url.Values) (*ResponseBuildingsList, error) {
	resp, err := c.DoPaginatedGet(uriBuildings, params)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "buildings", err)
	}

	var out ResponseBuildingsList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceBuilding
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "building", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetBuildingByID retrieves a single building information by its ID.
func (c *Client) GetBuildingByID(id string) (*ResourceBuilding, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBuildings, id)

	var building ResourceBuilding
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &building)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "building", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &building, nil
}

// GetBuildingByNameByID retrieves a single building information by its name using GetBuildingByID.
func (c *Client) GetBuildingByName(name string) (*ResourceBuilding, error) {
	buildings, err := c.GetBuildings(nil)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "building", err)
	}

	for _, value := range buildings.Results {
		if value.Name == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "building", name, errMsgNoName)
}

// CreateBuilding creates a new building in Jamf Pro
func (c *Client) CreateBuilding(building *ResourceBuilding) (*ResponseBuildingCreate, error) {
	endpoint := uriBuildings

	var responseBuildingCreate ResponseBuildingCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, building, &responseBuildingCreate)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "building", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseBuildingCreate, nil
}

// UpdateBuildingByID updates a building's information in Jamf Pro by its ID.
func (c *Client) UpdateBuildingByID(id string, buildingUpdate *ResourceBuilding) (*ResourceBuilding, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBuildings, id)

	var updatedBuilding ResourceBuilding
	resp, err := c.HTTP.DoRequest("PUT", endpoint, buildingUpdate, &updatedBuilding)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "building", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedBuilding, nil
}

// UpdateBuildingByNameByID updates a building's information in Jamf Pro by its name.
func (c *Client) UpdateBuildingByName(name string, buildingUpdate *ResourceBuilding) (*ResourceBuilding, error) {
	target, err := c.GetBuildingByName(name)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "building", name, err)
	}

	target_id := target.ID
	resp, err := c.UpdateBuildingByID(target_id, buildingUpdate)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "building", name, err)
	}

	return resp, nil

}

// DeleteBuildingByID deletes a building in Jamf Pro by its ID.
func (c *Client) DeleteBuildingByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriBuildings, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "buidling", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteBuildingByNameByID deletes a building in Jamf Pro by its name.
func (c *Client) DeleteBuildingByName(name string) error {
	target, err := c.GetBuildingByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "building", name, err)
	}

	target_id := target.ID
	err = c.DeleteBuildingByID(target_id)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "building", name, err)
	}

	return nil
}

// DeleteMultipleBuildingsByID deletes multiple buildings in Jamf Pro by their IDs.
func (c *Client) DeleteMultipleBuildingsByID(ids []string) error {
	endpoint := "/api/v1/buildings/delete-multiple"

	payload := struct {
		IDs []string `json:"ids"`
	}{
		IDs: ids,
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, payload, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteMultiple, "buildings", ids, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// GetBuildingResourceHistoryByID retrieves the resource history of a specific building by its ID.
func (c *Client) GetBuildingResourceHistoryByID(id string, params url.Values) (*ResponseBuildingResourceHistoryList, error) {
	endpoint := fmt.Sprintf("%s/%s/history", uriBuildings, id)

	resp, err := c.DoPaginatedGet(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "building histories", err)
	}

	var out ResponseBuildingResourceHistoryList
	out.Size = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceBuildingResourceHistory
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "buidling histories", err)
		}
		out.Results = append(out.Results, newObj)

	}

	return &out, nil
}

// CreateBuildingResourceHistoryByID updates the resource history of a building in Jamf Pro by its ID.
func (c *Client) CreateBuildingResourceHistoryByID(id string, historyUpdate *ResourceBuildingResourceHistory) (*ResourceBuildingResourceHistory, error) {
	endpoint := fmt.Sprintf("%s/%s/history", uriBuildings, id)

	var updatedHistory ResourceBuildingResourceHistory
	resp, err := c.HTTP.DoRequest("POST", endpoint, historyUpdate, &updatedHistory)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "building histories", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedHistory, nil
}
