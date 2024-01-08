// jamfproapi_departments.go
// Jamf Pro Api - Departments
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-departments
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// Responses

const uriDepartments = "/api/v1/departments"

type ResponseDepartmentsList struct {
	TotalCount int                  `json:"totalCount"`
	Results    []ResourceDepartment `json:"results"`
}

type ResponseDepartmentCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

type ResourceDepartment struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Returns all departments in list
func (c *Client) GetDepartments(sort_filter string) (*ResponseDepartmentsList, error) {
	endpoint := uriDepartments
	resp, err := c.DoPaginatedGet(
		endpoint,
		standardPageSize,
		startingPageNumber,
		sort_filter,
	)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "departments", err)
	}

	var out ResponseDepartmentsList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceDepartment
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "department", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// Returns ResourceDepartment with supplied id
func (c *Client) GetDepartmentByID(id string) (*ResourceDepartment, error) {
	endpoint := fmt.Sprintf("%s/%v", uriDepartments, id)
	var out ResourceDepartment
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "department", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// Returns ResourceDepartment with supplied name, leverages GetDepartments
func (c *Client) GetDepartmentByName(name string) (*ResourceDepartment, error) {
	depts, err := c.GetDepartments("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "departments", err)
	}

	for _, value := range depts.Results {
		if value.Name == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "department", name, errMsgNoName)
}

// Creates a new department, returns ResponseDepartmentCreate
func (c *Client) CreateDepartment(departmentName string) (*ResponseDepartmentCreate, error) {
	endpoint := uriDepartments
	var out ResponseDepartmentCreate

	payload := struct {
		Name string `json:"name"`
	}{
		Name: departmentName,
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, payload, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "department", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}

// Updates department name (only attr it has besides id) with given id
func (c *Client) UpdateDepartmentByID(id string, newName string) (*ResourceDepartment, error) {
	endpoint := fmt.Sprintf("%s/%s", uriDepartments, id)
	var out ResourceDepartment
	payload := struct {
		Name string `json:"name"`
	}{
		Name: newName,
	}
	resp, err := c.HTTP.DoRequest("PUT", endpoint, payload, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "department", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// Updates department name (only attr it has besides id) with given name, leverages GetDepartmentByName, UpdateDepartmentByID
func (c *Client) UpdateDepartmentByName(targetName, newName string) (*ResourceDepartment, error) {
	target, err := c.GetDepartmentByName(targetName)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "department", targetName, err)
	}

	target_id := target.ID
	resp, err := c.UpdateDepartmentByID(target_id, newName)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "department", targetName, err)
	}

	return resp, nil

}

// Deletes department with given id
func (c *Client) DeleteDepartmentByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriDepartments, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)

	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDeleteByID, "department", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}

// Deletes department with given name, leverages GetDepartmentByName
func (c *Client) DeleteDepartmentByName(targetName string) error {
	target, err := c.GetDepartmentByName(targetName)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "department", targetName, err)
	}

	target_id := target.ID
	err = c.DeleteDepartmentByID(target_id)

	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "department", targetName, err)
	}

	return nil
}
