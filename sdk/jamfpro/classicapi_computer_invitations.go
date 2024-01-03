// classicapi_computer_invitations.go
// Jamf Pro Classic Api - Computer Invitations
// api reference: https://developer.jamf.com/jamf-pro/reference/computerinvitations
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriComputerInvitations = "/JSSResource/computerinvitations"

// List

type ResponseComputerInvitationsList struct {
	Size               int                          `xml:"size"`
	ComputerInvitation []ComputerInvitationListItem `xml:"computer_invitation"`
}

type ComputerInvitationListItem struct {
	ID                  int    `xml:"id,omitempty"`
	Invitation          int64  `xml:"invitation,omitempty"`
	InvitationType      string `xml:"invitation_type,omitempty"`
	ExpirationDate      string `xml:"expiration_date,omitempty"`
	ExpirationDateUTC   string `xml:"expiration_date_utc,omitempty"`
	ExpirationDateEpoch int64  `xml:"expiration_date_epoch,omitempty"`
}

// Resource

type ResourceComputerInvitation struct {
	ID                          int                                     `xml:"id,omitempty"`
	Invitation                  string                                  `xml:"invitation,omitempty"`
	InvitationStatus            string                                  `xml:"invitation_status,omitempty"`
	InvitationType              string                                  `xml:"invitation_type,omitempty"`
	ExpirationDate              string                                  `xml:"expiration_date,omitempty"`
	ExpirationDateUTC           string                                  `xml:"expiration_date_utc,omitempty"`
	ExpirationDateEpoch         int64                                   `xml:"expiration_date_epoch,omitempty"`
	SSHUsername                 string                                  `xml:"ssh_username,omitempty"`
	SSHPassword                 string                                  `xml:"ssh_password,omitempty"`
	MultipleUsersAllowed        bool                                    `xml:"multiple_users_allowed,omitempty"`
	TimesUsed                   int                                     `xml:"times_used,omitempty"`
	CreateAccountIfDoesNotExist bool                                    `xml:"create_account_if_does_not_exist,omitempty"`
	HideAccount                 bool                                    `xml:"hide_account,omitempty"`
	LockDownSSH                 bool                                    `xml:"lock_down_ssh,omitempty"`
	InvitedUserUUID             string                                  `xml:"invited_user_uuid,omitempty"`
	EnrollIntoSite              ComputerInvitationSubsetEnrollIntoState `xml:"enroll_into_site,omitempty"`
	KeepExistingSiteMembership  bool                                    `xml:"keep_existing_site_membership,omitempty"`
	Site                        SharedResourceSite                      `xml:"site,omitempty"`
}

// Subsets & Containers

type ComputerInvitationSubsetEnrollIntoState struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// CRUD

// GetComputerInvitations retrieves a list of all computer invitations.
func (c *Client) GetComputerInvitations() (*ResponseComputerInvitationsList, error) {
	endpoint := uriComputerInvitations

	var invitations ResponseComputerInvitationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &invitations)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Computer Invitations: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &invitations, nil
}

// Duplicate function ???
// GetComputerInvitationByID retrieves a computer invitation by its ID.
func (c *Client) GetComputerInvitationByID(invitationID int) (*ResourceComputerInvitation, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerInvitations, invitationID)

	var invitation ResourceComputerInvitation
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &invitation)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Invitation by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &invitation, nil
}

// GetComputerInvitationsByName retrieves a computer invitation by its invitation Name.
func (c *Client) GetComputerInvitationByInvitationID(invitationID int) (*ResourceComputerInvitation, error) {
	endpoint := fmt.Sprintf("%s/invitation/%d", uriComputerInvitations, invitationID)

	var invitation ResourceComputerInvitation
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &invitation)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Invitation by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &invitation, nil
}

// CreateComputerInvitation creates a new computer invitation.
func (c *Client) CreateComputerInvitation(invitation *ResourceComputerInvitation) (*ResourceComputerInvitation, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriComputerInvitations)

	if invitation.Site.ID == 0 && invitation.Site.Name == "" {
		invitation.Site.ID = -1
		invitation.Site.Name = "none"
	}

	requestBody := struct {
		XMLName xml.Name `xml:"computer_invitation"`
		*ResourceComputerInvitation
	}{
		ResourceComputerInvitation: invitation,
	}

	var createdInvitation ResourceComputerInvitation
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdInvitation)
	if err != nil {
		return nil, fmt.Errorf("failed to create Computer Invitation: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdInvitation, nil
}

// DeleteComputerInvitationByID deletes a computer invitation by its ID.
func (c *Client) DeleteComputerInvitationByID(invitationID int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerInvitations, invitationID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Computer Invitation by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
