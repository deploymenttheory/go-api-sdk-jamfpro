// classicapi_vpp_assignments.go
// Jamf Pro Classic Api - VPP Assignments
// api reference: https://developer.jamf.com/jamf-pro/reference/vppassignments
// Jamf Pro Classic Api requires the structs to support an XML data structure.

// classicapi_vpp_assignments.go
package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriVPPAssignments = "/JSSResource/vppassignments"

// Struct for the list of VPP assignments
type ResponseVPPAssignmentsList struct {
	VPPAssignments []VPPAssignment `xml:"vpp_assignment"`
}

// Struct for a single VPP assignment
type VPPAssignment struct {
	ID                int    `xml:"id"`
	VPPAdminAccountID int    `xml:"vpp_admin_account_id"`
	Name              string `xml:"name"`
}

// Structs for the detailed VPP assignment response
type ResponseVPPAssignment struct {
	General VPPAssignmentGeneral `xml:"general"`
	IOSApps []VPPApp             `xml:"ios_apps>ios_app"`
	MacApps []VPPApp             `xml:"mac_apps>mac_app"`
	EBooks  []VPPApp             `xml:"ebooks>ebook"`
	Scope   VPPAssignmentScope   `xml:"scope"`
}

type VPPAssignmentGeneral struct {
	ID                  int    `xml:"id"`
	Name                string `xml:"name"`
	VPPAdminAccountID   int    `xml:"vpp_admin_account_id"`
	VPPAdminAccountName string `xml:"vpp_admin_account_name"`
}

type VPPApp struct {
	AdamID int    `xml:"adam_id"`
	Name   string `xml:"name"`
}

// Struct for VPP assignment scope
type VPPAssignmentScope struct {
	AllJSSUsers   bool           `xml:"all_jss_users"`
	JSSUsers      []VPPUser      `xml:"jss_users>user"`
	JSSUserGroups []VPPUserGroup `xml:"jss_user_groups>user_group"`
	Limitations   VPPLimitations `xml:"limitations"`
	Exclusions    VPPExclusions  `xml:"exclusions"`
}

// Struct for VPP user
type VPPUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Struct for VPP user group
type VPPUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Struct for VPP limitations
type VPPLimitations struct {
	UserGroups []VPPUserGroup `xml:"user_groups>user_group"`
}

// Struct for VPP exclusions
type VPPExclusions struct {
	JSSUsers      []VPPUser      `xml:"jss_users>user"`
	UserGroups    []VPPUserGroup `xml:"user_groups>user_group"`
	JSSUserGroups []VPPUserGroup `xml:"jss_user_groups>user_group"`
}

// GetVPPAssignments fetches a list of VPP assignments
func (c *Client) GetVPPAssignments() (*ResponseVPPAssignmentsList, error) {
	endpoint := uriVPPAssignments

	var assignments ResponseVPPAssignmentsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &assignments)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch VPP assignments: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &assignments, nil
}

// GetVPPAssignmentByID fetches a VPP assignment by its ID
func (c *Client) GetVPPAssignmentByID(id int) (*ResponseVPPAssignment, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAssignments, id)

	var assignment ResponseVPPAssignment
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &assignment)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch VPP assignment by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &assignment, nil
}

// CreateVPPAssignment creates a new VPP assignment
func (c *Client) CreateVPPAssignment(assignment *ResponseVPPAssignment) error {
	endpoint := fmt.Sprintf("%s/id/0", uriVPPAssignments) // '0' indicates creation

	// Wrap the assignment with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"vpp_assignment"`
		*ResponseVPPAssignment
	}{
		ResponseVPPAssignment: assignment,
	}

	// Create a dummy struct for the response
	var handleResponse struct{}

	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &handleResponse)
	if err != nil {
		return fmt.Errorf("failed to create VPP assignment: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// UpdateVPPAssignmentByID updates a VPP assignment by its ID
func (c *Client) UpdateVPPAssignmentByID(id int, assignment *ResponseVPPAssignment) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAssignments, id)

	// Wrap the assignment with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"vpp_assignment"`
		*ResponseVPPAssignment
	}{
		ResponseVPPAssignment: assignment,
	}

	// Create a dummy struct for the response
	var handleResponse struct{}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &handleResponse)
	if err != nil {
		return fmt.Errorf("failed to update VPP assignment by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteVPPAssignmentByID deletes a VPP assignment by its ID
func (c *Client) DeleteVPPAssignmentByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAssignments, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete VPP assignment by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
