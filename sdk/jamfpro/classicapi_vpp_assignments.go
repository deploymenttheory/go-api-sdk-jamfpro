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
	VPPAssignments []struct {
		ID                int    `xml:"id"`
		VPPAdminAccountID int    `xml:"vpp_admin_account_id"`
		Name              string `xml:"name"`
	} `xml:"vpp_assignment"`
}

// Structs for the detailed VPP assignment response
type ResourceVPPAssignment struct {
	General struct {
		ID                  int    `xml:"id"`
		Name                string `xml:"name"`
		VPPAdminAccountID   int    `xml:"vpp_admin_account_id"`
		VPPAdminAccountName string `xml:"vpp_admin_account_name"`
	} `xml:"general"`
	IOSApps []ResourceVPPSubsetVPPApp `xml:"ios_apps>ios_app"`
	MacApps []ResourceVPPSubsetVPPApp `xml:"mac_apps>mac_app"`
	EBooks  []ResourceVPPSubsetVPPApp `xml:"ebooks>ebook"`
	Scope   struct {
		AllJSSUsers   bool                            `xml:"all_jss_users"`
		JSSUsers      []ResourceVPPSubsetVPPUser      `xml:"jss_users>user"`
		JSSUserGroups []ResourceVPPSubsetVPPUserGroup `xml:"jss_user_groups>user_group"`
		Limitations   struct {
			UserGroups []ResourceVPPSubsetVPPUserGroup `xml:"user_groups>user_group"`
		} `xml:"limitations"`
		Exclusions struct {
			JSSUsers      []ResourceVPPSubsetVPPUser      `xml:"jss_users>user"`
			UserGroups    []ResourceVPPSubsetVPPUserGroup `xml:"user_groups>user_group"`
			JSSUserGroups []ResourceVPPSubsetVPPUserGroup `xml:"jss_user_groups>user_group"`
		} `xml:"exclusions"`
	} `xml:"scope"`
}

type ResourceVPPSubsetVPPApp struct {
	AdamID int    `xml:"adam_id"`
	Name   string `xml:"name"`
}

// Struct for VPP user
type ResourceVPPSubsetVPPUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Struct for VPP user group
type ResourceVPPSubsetVPPUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
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
func (c *Client) GetVPPAssignmentByID(id int) (*ResourceVPPAssignment, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAssignments, id)

	var assignment ResourceVPPAssignment
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
func (c *Client) CreateVPPAssignment(assignment *ResourceVPPAssignment) error {
	endpoint := fmt.Sprintf("%s/id/0", uriVPPAssignments) // '0' indicates creation

	// Wrap the assignment with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"vpp_assignment"`
		*ResourceVPPAssignment
	}{
		ResourceVPPAssignment: assignment,
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
func (c *Client) UpdateVPPAssignmentByID(id int, assignment *ResourceVPPAssignment) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAssignments, id)

	// Wrap the assignment with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"vpp_assignment"`
		*ResourceVPPAssignment
	}{
		ResourceVPPAssignment: assignment,
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
