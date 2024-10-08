// jamfproapi_onboarding.go
// Jamf Pro Api - Onboarding
// api reference: https://developer.jamf.com/jamf-pro/reference/put_v1-onboarding
// Jamf Pro Api requires the structs to support an JSON data structure.
// Ref: https://grahamrpugh.com/2024/05/16/jamf-new-packages-api-endpoint.html

package jamfpro

import (
	"fmt"
	"net/url"
)

const uriOnboardingSettings = "/api/v1/onboarding"

// List

// ResponseEligiblilityForOnboardingList represents the paginated response for eligible apps
type ResponseEligiblilityForOnboardingList struct {
	TotalCount int                                     `json:"totalCount"`
	Results    []ResourceEligiblilityForOnboardingList `json:"results"`
}

// ResourceEligiblilityForOnboardingList represents an individual eligible app item
type ResourceEligiblilityForOnboardingList struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ScopeDescription string `json:"scopeDescription"`
	SiteDescription  string `json:"siteDescription"`
}

// Response

// ResponseOnboardingSettings represents the structure of the onboarding settings response
type ResponseOnboardingSettings struct {
	AccountDrivenUserEnrollmentEnabled bool   `json:"accountDrivenUserEnrollmentEnabled"`
	IdpEnabled                         bool   `json:"idpEnabled"`
	ApnsCertificateExpirationDate      string `json:"apnsCertificateExpirationDate"`
	CloudServicesEnabled               bool   `json:"cloudServicesEnabled"`
	EnrollCustomizationEnabled         bool   `json:"enrollCustomizationEnabled"`
	PrestageEnrollmentEnabled          bool   `json:"prestageEnrollmentEnabled"`
}

// ResponseUpdateOnboardingSettings represents the response body after updating onboarding settings
type ResponseUpdateOnboardingSettings struct {
	ID              string                   `json:"id"`
	Enabled         bool                     `json:"enabled"`
	OnboardingItems []OnboardingItemResponse `json:"onboardingItems"`
}

// OnboardingItemResponse represents an item in the onboarding configuration response
type OnboardingItemResponse struct {
	ID                    string `json:"id"`
	EntityID              string `json:"entityId"`
	EntityName            string `json:"entityName"`
	ScopeDescription      string `json:"scopeDescription"`
	SiteDescription       string `json:"siteDescription"`
	SelfServiceEntityType string `json:"selfServiceEntityType"`
	Priority              int    `json:"priority"`
}

// Resource

// ResourceUpdateOnboardingSettings represents the request body for updating onboarding settings
type ResourceUpdateOnboardingSettings struct {
	Enabled                                        bool                          `json:"enabled"`
	DisplayName                                    string                        `json:"displayName"`
	SsoForEnrollmentEnabled                        bool                          `json:"ssoForEnrollmentEnabled"`
	SsoBypassAllowed                               bool                          `json:"ssoBypassAllowed"`
	SsoEnabled                                     bool                          `json:"ssoEnabled"`
	SsoForMacOsSelfServiceEnabled                  bool                          `json:"ssoForMacOsSelfServiceEnabled"`
	TokenExpirationDisabled                        bool                          `json:"tokenExpirationDisabled"`
	UserAttributeEnabled                           bool                          `json:"userAttributeEnabled"`
	UserAttributeName                              string                        `json:"userAttributeName"`
	UserMapping                                    string                        `json:"userMapping"`
	EnrollmentSsoForAccountDrivenEnrollmentEnabled bool                          `json:"enrollmentSsoForAccountDrivenEnrollmentEnabled"`
	GroupEnrollmentAccessEnabled                   bool                          `json:"groupEnrollmentAccessEnabled"`
	GroupAttributeName                             string                        `json:"groupAttributeName"`
	GroupRdnKey                                    string                        `json:"groupRdnKey"`
	GroupEnrollmentAccessName                      string                        `json:"groupEnrollmentAccessName"`
	IdpProviderType                                string                        `json:"idpProviderType"`
	OtherProviderTypeName                          string                        `json:"otherProviderTypeName"`
	MetadataSource                                 string                        `json:"metadataSource"`
	SessionTimeout                                 int                           `json:"sessionTimeout"`
	DeviceType                                     string                        `json:"deviceType"`
	OnboardingItems                                []SubsetOnboardingItemRequest `json:"onboardingItems"`
}

// SubsetOnboardingItemRequest represents an item in the onboarding configuration request
type SubsetOnboardingItemRequest struct {
	SelfServiceEntityType string `json:"selfServiceEntityType"`
	ID                    string `json:"id,omitempty"`
	EntityID              string `json:"entityId"`
	Priority              int    `json:"priority"`
}

// GetOnboardingSettings retrieves the current onboarding settings configuration
func (c *Client) GetOnboardingSettings() (*ResponseOnboardingSettings, error) {
	var onboardingSettings ResponseOnboardingSettings

	resp, err := c.HTTP.DoRequest("GET", uriOnboardingSettings, nil, &onboardingSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch onboarding settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &onboardingSettings, nil
}

// UpdateOnboardingSettings updates the onboarding configuration
func (c *Client) UpdateOnboardingSettings(request ResourceUpdateOnboardingSettings) (*ResponseUpdateOnboardingSettings, error) {
	endpoint := uriOnboardingSettings

	var response ResponseUpdateOnboardingSettings

	resp, err := c.HTTP.DoRequest("PUT", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update onboarding settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetEligibleAppsForOnboarding retrieves a list of applications that are eligible to be used in an onboarding configuration
func (c *Client) GetEligibleAppsForOnboarding(sort, filter string) (*ResponseEligiblilityForOnboardingList, error) {
	const maxPageSize = 200

	var allResults []ResourceEligiblilityForOnboardingList
	var totalCount int
	page := 0

	for {
		endpoint := fmt.Sprintf("%s/eligible-apps", uriOnboardingSettings)

		u, err := url.Parse(endpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to parse URL: %v", err)
		}

		query := u.Query()
		query.Set("page", fmt.Sprintf("%d", page))
		query.Set("page-size", fmt.Sprintf("%d", maxPageSize))
		if sort != "" {
			query.Set("sort", sort)
		}
		if filter != "" {
			query.Set("filter", filter)
		}
		u.RawQuery = query.Encode()

		var paginatedResponse ResponseEligiblilityForOnboardingList
		resp, err := c.HTTP.DoRequest("GET", u.String(), nil, &paginatedResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch eligible apps for onboarding: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		totalCount = paginatedResponse.TotalCount
		allResults = append(allResults, paginatedResponse.Results...)

		if len(paginatedResponse.Results) < maxPageSize {
			break
		}
		page++
	}

	return &ResponseEligiblilityForOnboardingList{
		TotalCount: totalCount,
		Results:    allResults,
	}, nil
}

// GetEligibleConfigurationProfilesForOnboarding retrieves a list of configuration profiles that are eligible to be used in an onboarding configuration
func (c *Client) GetEligibleConfigurationProfilesForOnboarding(sort, filter string) (*ResponseEligiblilityForOnboardingList, error) {
	const maxPageSize = 200

	var allResults []ResourceEligiblilityForOnboardingList
	var totalCount int
	page := 0

	for {
		endpoint := fmt.Sprintf("%s/eligible-configuration-profiles", uriOnboardingSettings)

		u, err := url.Parse(endpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to parse URL: %v", err)
		}

		query := u.Query()
		query.Set("page", fmt.Sprintf("%d", page))
		query.Set("page-size", fmt.Sprintf("%d", maxPageSize))
		if sort != "" {
			query.Set("sort", sort)
		}
		if filter != "" {
			query.Set("filter", filter)
		}
		u.RawQuery = query.Encode()

		var paginatedResponse ResponseEligiblilityForOnboardingList
		resp, err := c.HTTP.DoRequest("GET", u.String(), nil, &paginatedResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch eligible configuration profiles for onboarding: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		totalCount = paginatedResponse.TotalCount
		allResults = append(allResults, paginatedResponse.Results...)

		if len(paginatedResponse.Results) < maxPageSize {
			break
		}
		page++
	}

	return &ResponseEligiblilityForOnboardingList{
		TotalCount: totalCount,
		Results:    allResults,
	}, nil
}

// GetEligiblePoliciesForOnboarding retrieves a list of configuration profiles that are eligible to be used in an onboarding configuration
func (c *Client) GetEligiblePoliciesForOnboarding(sort, filter string) (*ResponseEligiblilityForOnboardingList, error) {
	const maxPageSize = 200

	var allResults []ResourceEligiblilityForOnboardingList
	var totalCount int
	page := 0

	for {
		endpoint := fmt.Sprintf("%s/eligible-policies", uriOnboardingSettings)

		u, err := url.Parse(endpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to parse URL: %v", err)
		}

		query := u.Query()
		query.Set("page", fmt.Sprintf("%d", page))
		query.Set("page-size", fmt.Sprintf("%d", maxPageSize))
		if sort != "" {
			query.Set("sort", sort)
		}
		if filter != "" {
			query.Set("filter", filter)
		}
		u.RawQuery = query.Encode()

		var paginatedResponse ResponseEligiblilityForOnboardingList
		resp, err := c.HTTP.DoRequest("GET", u.String(), nil, &paginatedResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch eligible configuration profiles for onboarding: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		totalCount = paginatedResponse.TotalCount
		allResults = append(allResults, paginatedResponse.Results...)

		if len(paginatedResponse.Results) < maxPageSize {
			break
		}
		page++
	}

	return &ResponseEligiblilityForOnboardingList{
		TotalCount: totalCount,
		Results:    allResults,
	}, nil
}
