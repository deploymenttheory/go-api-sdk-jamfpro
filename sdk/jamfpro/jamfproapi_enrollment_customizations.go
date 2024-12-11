// jamfproapi_enrollment_customizations.go
// Jamf Pro Api - Enrollment Customizations
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

const uriEnrollmentCustomizationSettings = "/api/v2/enrollment-customizations"

// Structs

// List

type ResponseEnrollmentCustomizationList struct {
	TotalCount int `json:"totalCount"`
	Results    []ResourceEnrollmentCustomization
}

// Responses

type ResponseEnrollmentCustomizationCreate struct {
	Id   string `json:"id"`
	Href string `json:"href"`
}

type ResponseEnrollmentCustomizationUpload struct {
	Url string `json:"url"`
}

// Resource

type ResourceEnrollmentCustomization struct {
	ID               string                                        `json:"id"`
	SiteID           string                                        `json:"siteId"`
	DisplayName      string                                        `json:"displayName"`
	Description      string                                        `json:"description"`
	BrandingSettings EnrollmentCustomizationSubsetBrandingSettings `json:"enrollmentCustomizationBrandingSettings"`
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
func (c *Client) GetEnrollmentCustomizations(sort_filter string) (*ResponseEnrollmentCustomizationList, error) {
	endpoint := uriEnrollmentCustomizationSettings
	resp, err := c.DoPaginatedGet(
		endpoint,
		standardPageSize,
		startingPageNumber,
		sort_filter,
	)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "enrollment customization", err)
	}

	var out ResponseEnrollmentCustomizationList
	out.TotalCount = resp.Size
	for _, value := range resp.Results {
		var newObj ResourceEnrollmentCustomization
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "enrollment customization", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil

}

// Returns single ResourceEnrollmentCustomization object matching given id
func (c *Client) GetEnrollmentCustomizationByID(id string) (*ResourceEnrollmentCustomization, error) {
	endpoint := fmt.Sprintf("%s/%s", uriEnrollmentCustomizationSettings, id)
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

// Creates new resource enrollment customization
func (c *Client) CreateEnrollmentCustomization(enrollmentCustomization ResourceEnrollmentCustomization) (*ResponseEnrollmentCustomizationCreate, error) {
	endpoint := uriEnrollmentCustomizationSettings
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
	endpoint := fmt.Sprintf("%s/%s", uriEnrollmentCustomizationSettings, id)
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
	endpoint := fmt.Sprintf("%s/%s", uriEnrollmentCustomizationSettings, id)

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
	endpoint := fmt.Sprintf("%s/images", uriEnrollmentCustomizationSettings)

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
