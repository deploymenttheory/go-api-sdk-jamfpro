// jamfproapi_volume_purchasing_locations.go
// Jamf Pro Api - Volume Purchasing Locations
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	uriVolumePurchasingLocations         = "/api/v1/volume-purchasing-locations" // Endpoint for volume purchasing locations
	volumePurchasingLocationsMaxPageSize = 2000                                  // Maximum number of items per page
)

// ResponseVolumePurchasingList represents the paginated response for volume purchasing locations.
type ResponseVolumePurchasingList struct {
	TotalCount int                        `json:"totalCount"`
	Results    []VolumePurchasingLocation `json:"results"`
}

// VolumePurchasingLocation represents an individual volume purchasing location.
type VolumePurchasingLocation struct {
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

// ResponseVolumePurchasingLocation represents the response structure for a single volume purchasing location.
type ResponseVolumePurchasingLocation struct {
	VolumePurchasingLocation
	Content []VolumePurchasingContent `json:"content"`
}

// VolumePurchasingContent represents the content associated with a volume purchasing location.
type VolumePurchasingContent struct {
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

// ResponseVolumePurchasingContentList represents the paginated response for volume purchasing content.
type ResponseVolumePurchasingContentList struct {
	TotalCount int                       `json:"totalCount"`
	Results    []VolumePurchasingContent `json:"results"`
}

// VolumePurchasingLocationCreateResponse represents the response for creating a volume purchasing location.
type VolumePurchasingLocationCreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// GetVolumePurchaseLocations retrieves all volume purchasing locations with optional sorting and filtering.
func (c *Client) GetVolumePurchaseLocations(sort []string, filter string) (*ResponseVolumePurchasingList, error) {
	var allLocations []VolumePurchasingLocation

	page := 1 // Pagination starts from page 1
	for {
		params := url.Values{
			"page":      []string{strconv.Itoa(page)},
			"page-size": []string{strconv.Itoa(volumePurchasingLocationsMaxPageSize)},
		}

		// Append sort parameters
		for _, s := range sort {
			params.Add("sort", s)
		}

		// Append filter parameter if provided
		if filter != "" {
			params.Add("filter", filter)
		}

		endpointWithParams := fmt.Sprintf("%s?%s", uriVolumePurchasingLocations, params.Encode())

		// Fetch the volume purchasing locations for the current page
		var responseLocations ResponseVolumePurchasingList
		resp, err := c.HTTP.DoRequest("GET", endpointWithParams, nil, &responseLocations)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch volume purchasing locations: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		// Add the fetched locations to the total list
		allLocations = append(allLocations, responseLocations.Results...)

		// Check if all locations have been fetched
		if len(allLocations) >= responseLocations.TotalCount {
			break
		}

		// Increment page number for the next iteration
		page++
	}

	// Return the combined list of all volume purchasing locations
	return &ResponseVolumePurchasingList{
		TotalCount: len(allLocations),
		Results:    allLocations,
	}, nil
}

// GetVolumePurchasingLocationByID retrieves a specific volume purchasing location by its ID.
func (c *Client) GetVolumePurchasingLocationByID(id string) (*ResponseVolumePurchasingLocation, error) {
	// Construct the endpoint URL using the provided ID
	endpoint := fmt.Sprintf("%s/%s", uriVolumePurchasingLocations, id)

	var responseLocation ResponseVolumePurchasingLocation
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseLocation)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch volume purchasing location by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseLocation, nil
}

// GetVolumePurchasingContentForLocationByID retrieves the content for a specific volume purchasing location by its ID.
func (c *Client) GetVolumePurchasingContentForLocationByID(id string, sort []string, filter string) (*ResponseVolumePurchasingContentList, error) {
	var allContent []VolumePurchasingContent

	page := 1
	for {
		params := url.Values{
			"page":      []string{strconv.Itoa(page)},
			"page-size": []string{"100"}, // Adjust the page size as per requirement
		}

		// Append sort parameters
		for _, s := range sort {
			params.Add("sort", s)
		}

		// Append filter parameter if provided
		if filter != "" {
			params.Add("filter", filter)
		}

		endpointWithParams := fmt.Sprintf("%s/%s/content?%s", uriVolumePurchasingLocations, id, params.Encode())

		// Fetch the content for the current page
		var responseContent ResponseVolumePurchasingContentList
		resp, err := c.HTTP.DoRequest("GET", endpointWithParams, nil, &responseContent)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch volume purchasing content for location ID %s: %v", id, err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		// Add the fetched content to the total list
		allContent = append(allContent, responseContent.Results...)

		// Check if all content has been fetched
		if len(allContent) >= responseContent.TotalCount {
			break
		}

		// Increment page number for the next iteration
		page++
	}

	// Return the combined list of all content for the given volume purchasing location
	return &ResponseVolumePurchasingContentList{
		TotalCount: len(allContent),
		Results:    allContent,
	}, nil
}

// CreateVolumePurchasingLocation creates a new volume purchasing location.
func (c *Client) CreateVolumePurchasingLocation(request *VolumePurchasingLocation) (*VolumePurchasingLocationCreateResponse, error) {
	endpoint := uriVolumePurchasingLocations

	var response VolumePurchasingLocationCreateResponse
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create volume purchasing location: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
