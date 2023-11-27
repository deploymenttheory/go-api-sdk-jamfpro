// jamfapi_buildings.go
// Jamf Pro Api - Buildings
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-buildings
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const uriBuildingsV1 = "/api/v1/buildings"

// ResponseBuildings represents the structure of the response for the buildings list.
type ResponseBuildingsList struct {
	TotalCount *int               `json:"totalCount"`
	Results    []ResponseBuilding `json:"results"`
}

// ResponseBuilding represents the structure of each building item in the response.
type ResponseBuilding struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	StreetAddress1 string `json:"streetAddress1"`
	StreetAddress2 string `json:"streetAddress2"`
	City           string `json:"city"`
	StateProvince  string `json:"stateProvince"`
	ZipPostalCode  string `json:"zipPostalCode"`
	Country        string `json:"country"`
}

// ResponseBuildingResourceHistoryList represents the structure of the response for the building resource history list.
type ResponseBuildingResourceHistoryList struct {
	TotalCount *int                              `json:"totalCount"`
	Results    []ResponseBuildingResourceHistory `json:"results"`
}

// ResponseBuildingResourceHistory represents the structure of each resource history item in the response.
type ResponseBuildingResourceHistory struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// ResponseBuildingCreate represents the response structure for creating a building.
type ResponseBuildingCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// GetBuildings retrieves all building information with optional sorting.
func (c *Client) GetBuildings(sort []string) (*ResponseBuildingsList, error) {
	const maxPageSize = 2000 // Assuming 2000 is a suitable limit for this API
	var allBuildings []ResponseBuilding

	page := 0
	for {
		// Construct the endpoint with query parameters for the current page
		params := url.Values{
			"page":      []string{strconv.Itoa(page)},
			"page-size": []string{strconv.Itoa(maxPageSize)},
		}
		if len(sort) > 0 {
			params.Add("sort", url.QueryEscape(strings.Join(sort, ",")))
		}
		endpointWithParams := fmt.Sprintf("%s?%s", uriBuildingsV1, params.Encode())

		// Fetch the buildings for the current page
		var responseBuildings ResponseBuildingsList
		resp, err := c.HTTP.DoRequest("GET", endpointWithParams, nil, &responseBuildings)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch buildings: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		// Add the fetched buildings to the total list
		allBuildings = append(allBuildings, responseBuildings.Results...)

		// Check if all buildings have been fetched
		if responseBuildings.TotalCount == nil || len(allBuildings) >= *responseBuildings.TotalCount {
			break
		}

		// Increment page number for the next iteration
		page++
	}

	// Create an int variable for the total count and assign its address to TotalCount
	totalCount := len(allBuildings)

	// Return the combined list of all buildings
	return &ResponseBuildingsList{
		TotalCount: &totalCount,
		Results:    allBuildings,
	}, nil
}

// GetBuildingByID retrieves a single building information by its ID.
func (c *Client) GetBuildingByID(id string) (*ResponseBuilding, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf("%s/%s", uriBuildingsV1, id)

	var building ResponseBuilding
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &building)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch building with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &building, nil
}

// GetBuildingByNameByID retrieves a single building information by its name using GetBuildingByID.
func (c *Client) GetBuildingByNameByID(name string) (*ResponseBuilding, error) {
	// Fetch all buildings
	buildings, err := c.GetBuildings(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch buildings: %v", err)
	}

	// Search for the building with the given name
	var buildingID string
	for _, building := range buildings.Results {
		if building.Name == name {
			buildingID = building.ID
			break
		}
	}

	if buildingID == "" {
		return nil, fmt.Errorf("no building found with the name %s", name)
	}

	// Use the found ID to get the full details of the building
	return c.GetBuildingByID(buildingID)
}

// GetBuildingResourceHistoryByID retrieves the resource history of a specific building by its ID.
func (c *Client) GetBuildingResourceHistoryByID(id string) (*ResponseBuildingResourceHistoryList, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf("%s/%s/history", uriBuildingsV1, id)

	var history ResponseBuildingResourceHistoryList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &history)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch building resource history with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &history, nil
}

// CreateBuilding creates a new building in Jamf Pro
func (c *Client) CreateBuilding(building *ResponseBuilding) (*ResponseBuildingCreate, error) {
	endpoint := uriBuildingsV1

	var responseBuildingCreate ResponseBuildingCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, building, &responseBuildingCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create building: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseBuildingCreate, nil
}

// UpdateBuildingByID updates a building's information in Jamf Pro by its ID.
func (c *Client) UpdateBuildingByID(id string, buildingUpdate *ResponseBuilding) (*ResponseBuilding, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBuildingsV1, id)

	var responseUpdate ResponseBuilding
	resp, err := c.HTTP.DoRequest("PUT", endpoint, buildingUpdate, &responseUpdate)
	if err != nil {
		return nil, fmt.Errorf("failed to update building with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseUpdate, nil
}

// CreateBuildingResourceHistoryByID updates the resource history of a building in Jamf Pro by its ID.
func (c *Client) CreateBuildingResourceHistoryByID(id string, historyUpdate *ResponseBuildingResourceHistory) (*ResponseBuildingResourceHistory, error) {
	endpoint := fmt.Sprintf("%s/%s/history", uriBuildingsV1, id)

	var updatedHistory ResponseBuildingResourceHistory
	resp, err := c.HTTP.DoRequest("POST", endpoint, historyUpdate, &updatedHistory)
	if err != nil {
		return nil, fmt.Errorf("failed to update building resource history with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedHistory, nil
}

// UpdateBuildingByNameByID updates a building's information in Jamf Pro by its name.
func (c *Client) UpdateBuildingByNameByID(name string, buildingUpdate *ResponseBuilding) (*ResponseBuilding, error) {
	// Fetch all buildings to find the one with the given name
	buildings, err := c.GetBuildings(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch buildings: %v", err)
	}

	// Search for the building with the given name
	var buildingID string
	for _, building := range buildings.Results {
		if building.Name == name {
			buildingID = building.ID
			break
		}
	}

	if buildingID == "" {
		return nil, fmt.Errorf("no building found with the name %s", name)
	}

	// Use the found ID to update the building with the new details
	return c.UpdateBuildingByID(buildingID, buildingUpdate)
}

// DeleteBuildingByID deletes a building in Jamf Pro by its ID.
func (c *Client) DeleteBuildingByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriBuildingsV1, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete building with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
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
		return fmt.Errorf("failed to delete multiple buildings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteBuildingByNameByID deletes a building in Jamf Pro by its name.
func (c *Client) DeleteBuildingByNameByID(name string) error {
	// Fetch all buildings to find the one with the given name
	buildings, err := c.GetBuildings(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch buildings: %v", err)
	}

	// Search for the building with the given name
	var buildingID string
	for _, building := range buildings.Results {
		if building.Name == name {
			buildingID = building.ID
			break
		}
	}

	if buildingID == "" {
		return fmt.Errorf("no building found with the name %s", name)
	}

	// Use the found ID to delete the building
	return c.DeleteBuildingByID(buildingID)
}
