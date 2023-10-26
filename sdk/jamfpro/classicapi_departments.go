// classicapi_departments.go
// Jamf Pro Classic Api - Departments
// api reference: https://developer.jamf.com/jamf-pro/reference/departments
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriDepartments = "/JSSResource/departments"

// Response structure for the list of departments

type ResponseDepartmentsList struct {
	TotalCount int              `xml:"size"`
	Results    []DepartmentItem `xml:"department"`
}

type DepartmentItem struct {
	Id   int    `xml:"id,omitempty"`
	Name string `xml:"name"`
}
type DepartmentScope struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// GetDepartments retrieves all departments
func (c *Client) GetDepartments() (*ResponseDepartmentsList, error) {
	endpoint := uriDepartments

	var departmentsList ResponseDepartmentsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &departmentsList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch departments: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &departmentsList, nil
}

// GetDepartmentByID retrieves the department by its ID
func (c *Client) GetDepartmentByID(id int) (*DepartmentItem, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriDepartments, id)

	var department DepartmentItem
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &department)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch department by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &department, nil
}

// GetDepartmentByName retrieves the department by its name
func (c *Client) GetDepartmentByName(name string) (*DepartmentItem, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDepartments, name)

	var department DepartmentItem
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &department)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch department by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &department, nil
}

// GetDepartmentIdByName retrieves the department ID by its name
func (c *Client) GetDepartmentIdByName(name string) (int, error) {
	departmentsList, err := c.GetDepartments()
	if err != nil {
		return 0, err
	}

	for _, dept := range departmentsList.Results {
		if dept.Name == name {
			return dept.Id, nil
		}
	}
	return 0, fmt.Errorf("department with name %s not found", name)
}

// CreateDepartment creates a new department
func (c *Client) CreateDepartment(departmentName string) (*DepartmentItem, error) {
	endpoint := uriDepartments

	// Wrap the department with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"department"`
		DepartmentItem
	}{
		DepartmentItem: DepartmentItem{
			Name: departmentName,
		},
	}

	var response DepartmentItem
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create department: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateDepartmentByID updates an existing department
func (c *Client) UpdateDepartmentByID(id int, departmentName string) (*DepartmentItem, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriDepartments, id)

	requestBody := struct {
		XMLName xml.Name `xml:"department"`
		DepartmentItem
	}{
		DepartmentItem: DepartmentItem{
			Name: departmentName,
		},
	}

	var updatedDepartment DepartmentItem
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedDepartment)
	if err != nil {
		return nil, fmt.Errorf("failed to update department: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedDepartment, nil
}

// UpdateDepartmentByName updates an existing department by its name
func (c *Client) UpdateDepartmentByName(oldName string, newName string) (*DepartmentItem, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDepartments, oldName)

	requestBody := struct {
		XMLName xml.Name `xml:"department"`
		DepartmentItem
	}{
		DepartmentItem: DepartmentItem{
			Name: newName,
		},
	}

	var updatedDepartment DepartmentItem
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedDepartment)
	if err != nil {
		return nil, fmt.Errorf("failed to update department by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedDepartment, nil
}

// DeleteDepartmentByID deletes an existing department by its ID
func (c *Client) DeleteDepartmentByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriDepartments, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete department by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDepartmentByName deletes an existing department by its name
func (c *Client) DeleteDepartmentByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriDepartments, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete department by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
