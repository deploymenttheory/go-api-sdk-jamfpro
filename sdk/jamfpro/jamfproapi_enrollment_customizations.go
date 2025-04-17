// jamfproapi_enrollment_customizations.go
// Jamf Pro Api - Enrollment Customizations
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const uriEnrollmentCustomizationSettingsV1 = "/api/v1/enrollment-customization"
const uriEnrollmentCustomizationSettingsV2 = "/api/v2/enrollment-customizations"

// Structs

// List

type ResponseEnrollmentCustomizationList struct {
	TotalCount int                               `json:"totalCount"`
	Results    []ResourceEnrollmentCustomization `json:"results"`
}

// Resource

type ResourceEnrollmentCustomization struct {
	ID               string                                        `json:"id"`
	SiteID           string                                        `json:"siteId"`
	DisplayName      string                                        `json:"displayName"`
	Description      string                                        `json:"description"`
	BrandingSettings EnrollmentCustomizationSubsetBrandingSettings `json:"enrollmentCustomizationBrandingSettings"`
}

// Resource structure for LDAP prestage pane
type ResourceEnrollmentCustomizationLDAPPane struct {
	ID                 int                                      `json:"id,omitempty"`
	Type               string                                   `json:"type"`
	DisplayName        string                                   `json:"displayName"`
	Rank               int                                      `json:"rank"`
	Title              string                                   `json:"title"`
	UsernameLabel      string                                   `json:"usernameLabel"`
	PasswordLabel      string                                   `json:"passwordLabel"`
	BackButtonText     string                                   `json:"backButtonText"`
	ContinueButtonText string                                   `json:"continueButtonText"`
	LDAPGroupAccess    []EnrollmentCustomizationLDAPGroupAccess `json:"ldapGroupAccess"`
}

// LDAPGroupAccess represents a single LDAP group access configuration
type EnrollmentCustomizationLDAPGroupAccess struct {
	GroupName    string `json:"groupName"`
	LDAPServerID int    `json:"ldapServerId"`
}

// Resource structure for text prestage pane
type ResourceEnrollmentCustomizationTextPane struct {
	ID                 int    `json:"id,omitempty"`
	Type               string `json:"type"`
	DisplayName        string `json:"displayName"`
	Rank               int    `json:"rank"`
	Title              string `json:"title"`
	Body               string `json:"body"`
	Subtext            string `json:"subtext"`
	BackButtonText     string `json:"backButtonText"`
	ContinueButtonText string `json:"continueButtonText"`
}

// Resource structure for SSO prestage pane
type ResourceEnrollmentCustomizationSSOPane struct {
	ID                             int    `json:"id,omitempty"`
	Type                           string `json:"type"`
	DisplayName                    string `json:"displayName"`
	Rank                           int    `json:"rank"`
	IsGroupEnrollmentAccessEnabled bool   `json:"isGroupEnrollmentAccessEnabled"`
	GroupEnrollmentAccessName      string `json:"groupEnrollmentAccessName"`
	IsUseJamfConnect               bool   `json:"isUseJamfConnect"`
	ShortNameAttribute             string `json:"shortNameAttribute"`
	LongNameAttribute              string `json:"longNameAttribute"`
}

// Responses

type ResponseEnrollmentCustomizationCreate struct {
	Id   string `json:"id"`
	Href string `json:"href"`
}

type ResponseEnrollmentCustomizationUpload struct {
	Url string `json:"url"`
}

// Response for creating SSO prestage pane
type ResponseEnrollmentCustomizationSSOPane struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

// Response for creating text prestage pane
type ResponseEnrollmentCustomizationTextPane struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

// Response for creating LDAP prestage pane
type ResponseEnrollmentCustomizationLDAPPane struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

// ResponsePrestagePanesList represents the structure for a list of prestage panes
type ResponsePrestagePanesList struct {
	Panels []PrestagePaneSummary `json:"panels"`
}

// PrestagePaneSummary represents the summary information for a prestage pane
type PrestagePaneSummary struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
	Rank        int    `json:"rank"`
	Type        string `json:"type"`
}

// Subsets

type EnrollmentCustomizationSubsetBrandingSettings struct {
	TextColor       string `json:"textColor"`
	ButtonColor     string `json:"buttonColor"`
	ButtonTextColor string `json:"buttonTextColor"`
	BackgroundColor string `json:"backgroundColor"`
	IconUrl         string `json:"iconUrl"`
}

// CRUD

// TODO Download an image - https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-images-id

// Returns paginated list of Enrollment Customization
func (c *Client) GetEnrollmentCustomizations(params url.Values) (*ResponseEnrollmentCustomizationList, error) {
	endpoint := uriEnrollmentCustomizationSettingsV2
	resp, err := c.DoPaginatedGet(endpoint, params)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "enrollment customization", err)
	}

	var out ResponseEnrollmentCustomizationList
	out.TotalCount = resp.Size

	// Convert the raw results to JSON and back to properly handle the nested structures
	resultJSON, err := json.Marshal(resp.Results)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal results: %v", err)
	}

	err = json.Unmarshal(resultJSON, &out.Results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal results: %v", err)
	}

	return &out, nil
}

// Returns single ResourceEnrollmentCustomization object matching given id
func (c *Client) GetEnrollmentCustomizationByID(id string) (*ResourceEnrollmentCustomization, error) {
	endpoint := fmt.Sprintf("%s/%s", uriEnrollmentCustomizationSettingsV2, id)
	var out ResourceEnrollmentCustomization
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "enrollment customization", id, err)

	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetEnrollmentCustomizationByName retrieves an enrollment customization by its display name
func (c *Client) GetEnrollmentCustomizationByName(name string) (*ResourceEnrollmentCustomization, error) {
	customizations, err := c.GetEnrollmentCustomizations(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get enrollment customizations list: %v", err)
	}

	var targetID string
	for _, customization := range customizations.Results {
		if customization.DisplayName == name {
			targetID = customization.ID
			break
		}
	}

	if targetID == "" {
		return nil, fmt.Errorf("no enrollment customization found with name: %s", name)
	}

	return c.GetEnrollmentCustomizationByID(targetID)
}

// Creates new resource enrollment customization
func (c *Client) CreateEnrollmentCustomization(enrollmentCustomization ResourceEnrollmentCustomization) (*ResponseEnrollmentCustomizationCreate, error) {
	endpoint := uriEnrollmentCustomizationSettingsV2
	var out ResponseEnrollmentCustomizationCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, enrollmentCustomization, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "enrollment customization", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}

// Updates resource enrollment customization by id
func (c *Client) UpdateEnrollmentCustomizationByID(id string, updatedCustomization ResourceEnrollmentCustomization) (*ResourceEnrollmentCustomization, error) {
	endpoint := fmt.Sprintf("%s/%s", uriEnrollmentCustomizationSettingsV2, id)
	var out ResourceEnrollmentCustomization
	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedCustomization, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "enrollment customization settings", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// Deletes resource enrollment customization by id
func (c *Client) DeleteEnrollmentCustomizationByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriEnrollmentCustomizationSettingsV2, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)

	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "enrollment customization settings", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}

// UploadEnrollmentCustomizationsImage uploads an enrollment image file using the custom multipart format
func (c *Client) UploadEnrollmentCustomizationsImage(filePath string) (*ResponseEnrollmentCustomizationUpload, error) {
	endpoint := fmt.Sprintf("%s/images", uriEnrollmentCustomizationSettingsV2)

	files := map[string][]string{
		"file": {filePath},
	}

	formFields := map[string]string{}
	contentTypes := map[string]string{
		"file": "image/png",
	}
	headersMap := map[string]http.Header{}

	var response ResponseEnrollmentCustomizationUpload
	resp, err := c.HTTP.DoMultiPartRequest(
		http.MethodPost,
		endpoint,
		files,
		formFields,
		contentTypes,
		headersMap,
		"byte",
		&response,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to upload icon: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetPrestagePanes retrieves all prestage panes for a specific enrollment customization
func (c *Client) GetPrestagePanes(customizationResourceID string) (*ResponsePrestagePanesList, error) {
	endpoint := fmt.Sprintf("%s/%s/all", uriEnrollmentCustomizationSettingsV1, customizationResourceID)

	var out ResponsePrestagePanesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to get prestage panes for enrollment customization %s: %v", customizationResourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateTextPrestagePane creates a text prestage pane for an enrollment customization
func (c *Client) CreateTextPrestagePane(customizationResourceID string, textPane ResourceEnrollmentCustomizationTextPane) (*ResponseEnrollmentCustomizationTextPane, error) {
	endpoint := fmt.Sprintf("%s/%s/text", uriEnrollmentCustomizationSettingsV1, customizationResourceID)

	// Ensure type is set to "text"
	textPane.Type = "text"

	var out ResponseEnrollmentCustomizationTextPane
	resp, err := c.HTTP.DoRequest("POST", endpoint, textPane, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to create text prestage pane for enrollment customization %s: %v", customizationResourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateLDAPPrestagePane creates a LDAP prestage pane for an enrollment customization
func (c *Client) CreateLDAPPrestagePane(customizationID string, ldapPane ResourceEnrollmentCustomizationLDAPPane) (*ResponseEnrollmentCustomizationLDAPPane, error) {
	endpoint := fmt.Sprintf("%s/%s/ldap", uriEnrollmentCustomizationSettingsV1, customizationID)

	// Ensure type is set to "ldap"
	ldapPane.Type = "ldap"

	var out ResponseEnrollmentCustomizationLDAPPane
	resp, err := c.HTTP.DoRequest("POST", endpoint, ldapPane, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to create LDAP prestage pane for enrollment customization %s: %v", customizationID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateSSOPrestagePane creates a SSO prestage pane for an enrollment customization
func (c *Client) CreateSSOPrestagePane(customizationID string, ssoPane ResourceEnrollmentCustomizationSSOPane) (*ResponseEnrollmentCustomizationSSOPane, error) {
	// Use the correct endpoint with singular "customization"
	endpoint := fmt.Sprintf("/api/v1/enrollment-customization/%s/sso", customizationID)

	// Ensure type is set to "sso"
	ssoPane.Type = "sso"

	var out ResponseEnrollmentCustomizationSSOPane
	resp, err := c.HTTP.DoRequest("POST", endpoint, ssoPane, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to create SSO prestage pane for enrollment customization %s: %v", customizationID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateTextPrestagePaneByID updates a text prestage pane for an enrollment customization
func (c *Client) UpdateTextPrestagePaneByID(customizationResourceID string, panelID string, textPane ResourceEnrollmentCustomizationTextPane) (*ResourceEnrollmentCustomizationTextPane, error) {
	endpoint := fmt.Sprintf("%s/%s/text/%s", uriEnrollmentCustomizationSettingsV1, customizationResourceID, panelID)

	// Ensure type is set to "text"
	textPane.Type = "text"

	var out ResourceEnrollmentCustomizationTextPane
	resp, err := c.HTTP.DoRequest("PUT", endpoint, textPane, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to update text prestage pane %s for enrollment customization %s: %v", panelID, customizationResourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateLDAPPrestagePaneByID updates an LDAP prestage pane for an enrollment customization
func (c *Client) UpdateLDAPPrestagePaneByID(customizationResourceID string, panelID string, ldapPane ResourceEnrollmentCustomizationLDAPPane) (*ResourceEnrollmentCustomizationLDAPPane, error) {
	endpoint := fmt.Sprintf("%s/%s/ldap/%s", uriEnrollmentCustomizationSettingsV1, customizationResourceID, panelID)

	// Ensure type is set to "ldap"
	ldapPane.Type = "ldap"

	var out ResourceEnrollmentCustomizationLDAPPane
	resp, err := c.HTTP.DoRequest("PUT", endpoint, ldapPane, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to update LDAP prestage pane %s for enrollment customization %s: %v", panelID, customizationResourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateSSOPrestagePaneByID updates an SSO prestage pane for an enrollment customization
func (c *Client) UpdateSSOPrestagePaneByID(customizationResourceID string, panelID string, ssoPane ResourceEnrollmentCustomizationSSOPane) (*ResourceEnrollmentCustomizationSSOPane, error) {
	endpoint := fmt.Sprintf("%s/%s/sso/%s", uriEnrollmentCustomizationSettingsV1, customizationResourceID, panelID)

	// Ensure type is set to "sso"
	ssoPane.Type = "sso"

	var out ResourceEnrollmentCustomizationSSOPane
	resp, err := c.HTTP.DoRequest("PUT", endpoint, ssoPane, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to update SSO prestage pane %s for enrollment customization %s: %v", panelID, customizationResourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetTextPrestagePaneByID gets a text prestage pane by enrollment customization ID and pane ID
func (c *Client) GetTextPrestagePaneByID(customizationResourceID string, paneID string) (*ResourceEnrollmentCustomizationTextPane, error) {
	endpoint := fmt.Sprintf("%s/%s/text/%s", uriEnrollmentCustomizationSettingsV1, customizationResourceID, paneID)

	var out ResourceEnrollmentCustomizationTextPane
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to get text prestage pane %s for enrollment customization %s: %v", paneID, customizationResourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetLDAPPrestagePaneByID gets an LDAP prestage pane by enrollment customization ID and pane ID
func (c *Client) GetLDAPPrestagePaneByID(customizationResourceID string, paneID string) (*ResourceEnrollmentCustomizationLDAPPane, error) {
	endpoint := fmt.Sprintf("%s/%s/ldap/%s", uriEnrollmentCustomizationSettingsV1, customizationResourceID, paneID)

	var out ResourceEnrollmentCustomizationLDAPPane
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to get LDAP prestage pane %s for enrollment customization %s: %v", paneID, customizationResourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetSSOPrestagePaneByID gets a SSO prestage pane by enrollment customization ID and pane ID
func (c *Client) GetSSOPrestagePaneByID(customizationResourceID string, paneID string) (*ResourceEnrollmentCustomizationSSOPane, error) {
	endpoint := fmt.Sprintf("/api/v1/enrollment-customization/%s/sso/%s", customizationResourceID, paneID)

	var out ResourceEnrollmentCustomizationSSOPane
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to get SSO prestage pane %s for enrollment customization %s: %v", paneID, customizationResourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// DeleteSSOPrestagePane deletes a SSO prestage pane for an enrollment customization
func (c *Client) DeleteSSOPrestagePane(customizationResourceID string, paneID string) error {
	endpoint := fmt.Sprintf("/api/v1/enrollment-customization/%s/sso/%s", customizationResourceID, paneID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)

	if err != nil {
		return fmt.Errorf("failed to delete SSO prestage pane %s for enrollment customization %s: %v", paneID, customizationResourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeletePrestagePane deletes a text, LDAP or SSO prestage pane from an enrollment customization
func (c *Client) DeletePrestagePane(customizationResourceID string, paneID string) error {
	endpoint := fmt.Sprintf("%s/%s/all/%s", uriEnrollmentCustomizationSettingsV1, customizationResourceID, paneID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)

	if err != nil {
		return fmt.Errorf("failed to delete text prestage pane %s for enrollment customization %s: %v", paneID, customizationResourceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
