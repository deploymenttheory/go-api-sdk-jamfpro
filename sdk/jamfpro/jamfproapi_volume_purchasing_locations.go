// jamfproapi_volume_purchasing_locations.go
// Jamf Pro Api - Volume Purchasing Locations
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

const uriVolumePurchasingLocations = "/api/v1/volume-purchasing-locations"

// List

// ResponseVolumePurchasingList represents the paginated response for volume purchasing locations.
type ResponseVolumePurchasingList struct {
	TotalCount int                                `json:"totalCount"`
	Results    []ResourceVolumePurchasingLocation `json:"results"`
}

type ResponseVolumePurchasingContentList struct {
	TotalCount int                             `json:"totalCount"`
	Results    []VolumePurchasingSubsetContent `json:"results"`
}

// VolumePurchasingLocationCreateResponse represents the response for creating a volume purchasing location.
type ResponseVolumePurchasingLocationCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResponseVolumePurchasingLocation represents the response structure for a single volume purchasing location.
type ResourceVolumePurchasingLocation struct {
	VolumePurchasingLocationSubsetBody                                 // I don't think this works. See exmaples/volume_purchase_locations/CreateVolumePurchasingLocation.go
	Content                            []VolumePurchasingSubsetContent `json:"content"`
}

// Subsets

// VolumePurchasingLocation represents an individual volume purchasing location.
type VolumePurchasingLocationSubsetBody struct {
	ID                                    string `json:"id,omitempty"`
	Name                                  string `json:"name,omitempty"`
	AppleID                               string `json:"appleId,omitempty"`
	OrganizationName                      string `json:"organizationName,omitempty"`
	TokenExpiration                       string `json:"tokenExpiration,omitempty"`
	CountryCode                           string `json:"countryCode,omitempty"`
	LocationName                          string `json:"locationName,omitempty"`
	ClientContextMismatch                 bool   `json:"clientContextMismatch,omitempty"`
	AutomaticallyPopulatePurchasedContent bool   `json:"automaticallyPopulatePurchasedContent,omitempty"`
	SendNotificationWhenNoLongerAssigned  bool   `json:"sendNotificationWhenNoLongerAssigned,omitempty"`
	AutoRegisterManagedUsers              bool   `json:"autoRegisterManagedUsers,omitempty"`
	SiteID                                string `json:"siteId,omitempty"`
	LastSyncTime                          string `json:"lastSyncTime,omitempty"`
	TotalPurchasedLicenses                int    `json:"totalPurchasedLicenses,omitempty"`
	TotalUsedLicenses                     int    `json:"totalUsedLicenses,omitempty"`
	ServiceToken                          string `json:"serviceToken,omitempty"`
}

// VolumePurchasingContent represents the content associated with a volume purchasing location.
type VolumePurchasingSubsetContent struct {
	Name                 string   `json:"name"`
	LicenseCountTotal    int      `json:"licenseCountTotal"`
	LicenseCountInUse    int      `json:"licenseCountInUse"`
	LicenseCountReported int      `json:"licenseCountReported"`
	IconURL              string   `json:"iconUrl"`
	DeviceTypes          []string `json:"deviceTypes"`
	ContentType          string   `json:"contentType"`
	PricingParam         string   `json:"pricingParam"`
	AdamId               string   `json:"adamId"`
}

// CRUD
// VPP Locations

// GetVolumePurchaseLocations retrieves all volume purchasing locations with optional sorting and filtering.
func (c *Client) GetVolumePurchaseLocations(sort_filter string) (*ResponseVolumePurchasingList, error) {
	resp, err := c.DoPaginatedGet(uriVolumePurchasingLocations, standardPageSize, startingPageNumber, sort_filter)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "vpp locations", err)
	}

	var out ResponseVolumePurchasingList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceVolumePurchasingLocation
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "vpp location", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetVolumePurchasingLocationByID retrieves a specific volume purchasing location by its ID.
func (c *Client) GetVolumePurchasingLocationByID(id string) (*ResourceVolumePurchasingLocation, error) {
	endpoint := fmt.Sprintf("%s/%s", uriVolumePurchasingLocations, id)
	var responseLocation ResourceVolumePurchasingLocation
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseLocation)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "vpp locations", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseLocation, nil
}

// CreateVolumePurchasingLocation creates a new volume purchasing location.
func (c *Client) CreateVolumePurchasingLocation(request *ResourceVolumePurchasingLocation) (*ResponseVolumePurchasingLocationCreate, error) {
	endpoint := uriVolumePurchasingLocations

	var response ResponseVolumePurchasingLocationCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create volume purchasing location: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateVolumePurchasingLocationByID updates a specific volume purchasing location by its ID.
func (c *Client) UpdateVolumePurchasingLocationByID(id string) (*ResourceVolumePurchasingLocation, error) {
	endpoint := fmt.Sprintf("%s/%s", uriVolumePurchasingLocations, id)
	var responseLocation ResourceVolumePurchasingLocation
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, nil, &responseLocation)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "vpp locations", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseLocation, nil
}

// DeleteVolumePurchasingLocationByID deletes a specific volume purchasing location by its ID.
func (c *Client) DeleteVolumePurchasingLocationByID(id string) error {
	endpoint := uriVolumePurchasingLocations
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "vpp location", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// QUERY do we need the stuff below here?

// GetVolumePurchasingContentForLocationByID retrieves the content for a specific volume purchasing location by its ID.
func (c *Client) GetVolumePurchasingContentForLocationByID(id string, sort []string, filter string) (*ResponseVolumePurchasingContentList, error) {
	var allContent []VolumePurchasingSubsetContent

	page := 1
	for {
		params := url.Values{
			"page":      []string{strconv.Itoa(page)},
			"page-size": []string{"100"},
		}

		for _, s := range sort {
			params.Add("sort", s)
		}

		if filter != "" {
			params.Add("filter", filter)
		}

		endpointWithParams := fmt.Sprintf("%s/%s/content?%s", uriVolumePurchasingLocations, id, params.Encode())

		var responseContent ResponseVolumePurchasingContentList
		resp, err := c.HTTP.DoRequest("GET", endpointWithParams, nil, &responseContent)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch volume purchasing content for location ID %s: %v", id, err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		allContent = append(allContent, responseContent.Results...)

		if len(allContent) >= responseContent.TotalCount {
			break
		}

		page++
	}

	return &ResponseVolumePurchasingContentList{
		TotalCount: len(allContent),
		Results:    allContent,
	}, nil
}
