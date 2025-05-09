// jamfproapi_sso_settings_v3.go
// Jamf Pro Api - SSO Settings
// api reference: https://developer.jamf.com/jamf-pro/reference/put_v3-sso
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriSsoSettings = "/api/v3/sso"
const uriSsoEnrollmentCustomizationDependencies = "/api/v3/sso/dependencies"

// Structs

// SSO Settings
// Resource
type ResourceSsoSettings struct {
	SsoEnabled                                     bool                 `json:"ssoEnabled"`
	ConfigurationType                              string               `json:"configurationType"` // enum: SAML, OIDC, OIDC_WITH_SAML
	OidcSettings                                   *OidcSettings        `json:"oidcSettings"`
	SamlSettings                                   *SamlSettings        `json:"samlSettings"`
	SsoBypassAllowed                               bool                 `json:"ssoBypassAllowed"`
	SsoForEnrollmentEnabled                        bool                 `json:"ssoForEnrollmentEnabled"`
	SsoForMacOsSelfServiceEnabled                  bool                 `json:"ssoForMacOsSelfServiceEnabled"`
	EnrollmentSsoForAccountDrivenEnrollmentEnabled bool                 `json:"enrollmentSsoForAccountDrivenEnrollmentEnabled"`
	GroupEnrollmentAccessEnabled                   bool                 `json:"groupEnrollmentAccessEnabled"`
	GroupEnrollmentAccessName                      string               `json:"groupEnrollmentAccessName"`
	EnrollmentSsoConfig                            *EnrollmentSsoConfig `json:"enrollmentSsoConfig,omitempty"`
}

// OIDC Settings
type OidcSettings struct {
	UserMapping string `json:"userMapping"` // enum: USERNAME, EMAIL
}

// SAML Settings
type SamlSettings struct {
	IdpUrl                  string `json:"idpUrl,omitempty"`
	EntityId                string `json:"entityId,omitempty"`
	MetadataSource          string `json:"metadataSource,omitempty"`  // enum: URL, FILE, UNKNOWN
	UserMapping             string `json:"userMapping,omitempty"`     // enum: USERNAME, EMAIL
	IdpProviderType         string `json:"idpProviderType,omitempty"` // enum: ADFS, OKTA, GOOGLE, SHIBBOLETH, ONELOGIN, PING, CENTRIFY, AZURE, OTHER
	GroupRdnKey             string `json:"groupRdnKey"`
	UserAttributeName       string `json:"userAttributeName"`
	GroupAttributeName      string `json:"groupAttributeName,omitempty"`
	UserAttributeEnabled    bool   `json:"userAttributeEnabled"`
	MetadataFileName        string `json:"metadataFileName,omitempty"`
	OtherProviderTypeName   string `json:"otherProviderTypeName"`
	FederationMetadataFile  string `json:"federationMetadataFile,omitempty"`
	TokenExpirationDisabled bool   `json:"tokenExpirationDisabled"`
	SessionTimeout          int    `json:"sessionTimeout,omitempty"`
}

// Enrollment SSO Config
type EnrollmentSsoConfig struct {
	Hosts          []string `json:"hosts,omitempty"`
	ManagementHint string   `json:"managementHint,omitempty"`
}

// SSO Enrollment Customization Dependencies
type ResponseSsoEnrollmentCustomizationDependencies struct {
	Dependencies []EnrollmentCustomizationDependency `json:"dependencies"`
}

// Enrollment Customization Dependency
type EnrollmentCustomizationDependency struct {
	Name              string `json:"name"`
	Hyperlink         string `json:"hyperlink"`
	HumanReadableName string `json:"humanReadableName"`
}

// CRUD

// GetSsoSettings retrieves current Jamf SSO settings
func (c *Client) GetSsoSettings() (*ResourceSsoSettings, error) {
	endpoint := uriSsoSettings
	var out ResourceSsoSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
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
	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedSettings, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "sso settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetSsoEnrollmentCustomizationDependencies retrieves current SSO Enrollment Customization dependencies
func (c *Client) GetSsoEnrollmentCustomizationDependencies() (*ResponseSsoEnrollmentCustomizationDependencies, error) {
	endpoint := uriSsoEnrollmentCustomizationDependencies
	var out ResponseSsoEnrollmentCustomizationDependencies
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "sso dependencies", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
