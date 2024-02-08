// jamfproapi_sso_settings.go
// Jamf Pro Api - SSO Settings
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-sso
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriSsoSettings = "/api/v2/sso"
const uriSsoDependencies = "/api/v2/sso/dependencies"

// Structs

// SSO Settings
// Resource

type ResourceSsoSettings struct {
	SsoForEnrollmentEnabled                        bool                                 `json:"ssoForEnrollmentEnabled"`
	SsoBypassAllowed                               bool                                 `json:"ssoBypassAllowed"`
	SsoEnabled                                     bool                                 `json:"ssoEnabled"`
	SsoForMacOsSelfServiceEnabled                  bool                                 `json:"ssoForMacOsSelfServiceEnabled"`
	TokenExpirationDisabled                        bool                                 `json:"tokenExpirationDisabled"`
	UserAttributeEnabled                           bool                                 `json:"userAttributeEnabled"`
	UserAttributeName                              string                               `json:"userAttributeName"`
	UserMapping                                    string                               `json:"userMapping"`
	EnrollmentSsoForAccountDrivenEnrollmentEnabled bool                                 `json:"enrollmentSsoForAccountDrivenEnrollmentEnabled"`
	EnrollmentSsoConfig                            SsoSettingsSubsetEnrollmentSsoConfig `json:"enrollmentSsoConfig"`
	GroupEnrollmentAccessEnabled                   bool                                 `json:"groupEnrollmentAccessEnabled"`
	GroupAttributeName                             string                               `json:"groupAttributeName"`
	GroupRdnKey                                    string                               `json:"groupRdnKey"`
	GroupEnrollmentAccessName                      string                               `json:"groupEnrollmentAccessName"`
	IdpProviderType                                string                               `json:"idpProviderType"`
	IdpUrl                                         string                               `json:"idpUrl"`
	EntityId                                       string                               `json:"entityId"`
	MetadataFileName                               string                               `json:"metadataFileName"`
	OtherProviderTypeName                          string                               `json:"otherProviderTypeName"`
	FederationMetadataFile                         string                               `json:"federationMetadataFile"`
	MetadataSource                                 string                               `json:"metadataSource"`
	SessionTimeout                                 int                                  `json:"sessionTimeout"`
}

// Subsets

type SsoSettingsSubsetEnrollmentSsoConfig struct {
	Hosts          []string `json:"hosts"`
	ManagementHint string   `json:"managementHint"`
}

// Enrollment Customizations Using SSO
// Resource

type ResponseSsoSubsetEnrollmentCustomizationDependencyList struct {
	Dependencies []SsoSubsetSubsetEnrollmentCustomizationDependency
}

// Subset

type SsoSubsetSubsetEnrollmentCustomizationDependency struct {
	Name              string `json:"name"`
	HumanReadableName string `json:"humanReadableName"`
	Hyperlink         string `json:"hyperlink"`
}

// CRUD

// GetSsoSettings retrieves current Jamf Sso settings
func (c *Client) GetSsoSettings() (*ResourceSsoSettings, error) {
	endpoint := uriSsoSettings
	var out ResourceSsoSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "sso settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateSsoSettings Updates SSO Settings with ResourceSsoSettings struct data
func (c *Client) UpdateSsoSettings(updatedSettings ResourceSsoSettings) (*ResourceSsoSettings, error) {
	endpoint := uriSsoSettings
	var out ResourceSsoSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedSettings, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "sso settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetSsoEnrollmentCustomizationDependencies shows which enrollment customizations are dependent on which sso settings // NOTE I think?
func (c *Client) GetSsoEnrollmentCustomizationDependencies() (*ResponseSsoSubsetEnrollmentCustomizationDependencyList, error) {
	endpoint := uriSsoDependencies
	var out ResponseSsoSubsetEnrollmentCustomizationDependencyList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "sso enrollment customization dependencies", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// QUERY What other endpoints do we need to cover here? It's a bit of a mix mash
