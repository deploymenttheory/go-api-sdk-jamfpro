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

// List

// Struct for the list of VPP assignments
type ResponseVPPAssignmentsList struct {
	VPPAssignments []VPPAssignmentsListItem `xml:"vpp_assignment"`
}

type VPPAssignmentsListItem struct {
	ID                int    `xml:"id"`
	VPPAdminAccountID int    `xml:"vpp_admin_account_id"`
	Name              string `xml:"name"`
}

// Resource

// Structs for the detailed VPP assignment response
type ResourceVPPAssignment struct {
	General VPPAssignmentSubsetGeneral `xml:"general"`
	IOSApps []VPPSubsetVPPApp          `xml:"ios_apps>ios_app"`
	MacApps []VPPSubsetVPPApp          `xml:"mac_apps>mac_app"`
	EBooks  []VPPSubsetVPPApp          `xml:"ebooks>ebook"`
	Scope   VPPAssignmentSubsetScope   `xml:"scope"`
}

// Subsets & Containers

// General

type VPPAssignmentSubsetGeneral struct {
	ID                  int    `xml:"id"`
	Name                string `xml:"name"`
	VPPAdminAccountID   int    `xml:"vpp_admin_account_id"`
	VPPAdminAccountName string `xml:"vpp_admin_account_name"`
}

// Scope

type VPPAssignmentSubsetScope struct {
	AllJSSUsers   bool                                `xml:"all_jss_users"`
	JSSUsers      []VPPSubsetVPPUser                  `xml:"jss_users>user"`
	JSSUserGroups []VPPSubsetVPPUserGroup             `xml:"jss_user_groups>user_group"`
	Limitations   VPPAssignmentSubsetScopeLimitations `xml:"limitations"`
	Exclusions    VPPAssignmentSubsetScopeExclusions  `xml:"exclusions"`
}

type VPPAssignmentSubsetScopeLimitations struct {
	UserGroups []VPPSubsetVPPUserGroup `xml:"user_groups>user_group"`
}

type VPPAssignmentSubsetScopeExclusions struct {
	JSSUsers      []VPPSubsetVPPUser      `xml:"jss_users>user"`
	UserGroups    []VPPSubsetVPPUserGroup `xml:"user_groups>user_group"`
	JSSUserGroups []VPPSubsetVPPUserGroup `xml:"jss_user_groups>user_group"`
}

// Shared

type VPPSubsetVPPApp struct {
	AdamID int    `xml:"adam_id"`
	Name   string `xml:"name"`
}

// Struct for VPP user
type VPPSubsetVPPUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Struct for VPP user group
type VPPSubsetVPPUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// CRUD

// GetVPPAssignments fetches a list of VPP assignments
func (c *Client) GetVPPAssignments() (*ResponseVPPAssignmentsList, error) {
	endpoint := uriVPPAssignments

	var assignments ResponseVPPAssignmentsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &assignments)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "vpp assignments", err)
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
		return nil, fmt.Errorf(errMsgFailedGetByID, "vpp assignment", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &assignment, nil
}

// CreateVPPAssignment creates a new VPP assignment
func (c *Client) CreateVPPAssignment(assignment *ResourceVPPAssignment) error {
	endpoint := fmt.Sprintf("%s/id/0", uriVPPAssignments)

	requestBody := struct {
		XMLName xml.Name `xml:"vpp_assignment"`
		*ResourceVPPAssignment
	}{
		ResourceVPPAssignment: assignment,
	}

	var handleResponse struct{}

	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &handleResponse)
	if err != nil {
		return fmt.Errorf(errMsgFailedCreate, "vpp assignment", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// UpdateVPPAssignmentByID updates a VPP assignment by its ID
func (c *Client) UpdateVPPAssignmentByID(id int, assignment *ResourceVPPAssignment) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAssignments, id)

	requestBody := struct {
		XMLName xml.Name `xml:"vpp_assignment"`
		*ResourceVPPAssignment
	}{
		ResourceVPPAssignment: assignment,
	}

	var handleResponse struct{}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &handleResponse)
	if err != nil {
		return fmt.Errorf(errMsgFailedUpdateByID, "vpp assignment", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteVPPAssignmentByID deletes a VPP assignment by its ID
func (c *Client) DeleteVPPAssignmentByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPAssignments, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "vpp assignment", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
