// jamfproapi_account_driven_user_enrollment_token_settings.go
// Jamf Pro Api - Enrollment
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-history
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriUserEnrollmentTokenSettings = "/api/v1/adue-session-token-settings"
const uriEnrollmentSettings = "/api/v4/enrollment"

// Resource

type ResourceADUETokenSettings struct {
	Enabled                   bool `json:"enabled"`
	ExpirationIntervalDays    int  `json:"expirationIntervalDays,omitempty"`
	ExpirationIntervalSeconds int  `json:"expirationIntervalSeconds,omitempty"`
}

type EnrollmentSubsetCertificateSettings struct {
	InstallSingleProfile                bool                               `json:"installSingleProfile"`
	SigningMdmProfileEnabled            bool                               `json:"signingMdmProfileEnabled"`
	MdmSigningCertificate               *EnrollmentSubsetCertificate       `json:"mdmSigningCertificate"`
	MdmSigningCertificateDetails        EnrollmentSubsetCertificateDetails `json:"mdmSigningCertificateDetails"`
	RestrictReenrollment                bool                               `json:"restrictReenrollment"`

	//
	// Sunsetting Re-enrollment options from this resource
	// Use jamfpro_reenrollment resource instead
	//
	// FlushLocationInformation            bool                               `json:"flushLocationInformation"`
	// FlushLocationHistoryInformation     bool                               `json:"flushLocationHistoryInformation"`
	// FlushPolicyHistory                  bool                               `json:"flushPolicyHistory"`
	// FlushExtensionAttributes            bool                               `json:"flushExtensionAttributes"`
	// FlushSoftwareUpdatePlans            bool                               `json:"flushSoftwareUpdatePlans"`
	// FlushMdmCommandsOnReenroll          string                             `json:"flushMdmCommandsOnReenroll"`

	MacOsEnterpriseEnrollmentEnabled    bool                               `json:"macOsEnterpriseEnrollmentEnabled"`
	ManagementUsername                  string                             `json:"managementUsername"`
	ManagementPasswordSet               bool                               `json:"managementPasswordSet"`
	PasswordType                        *string                            `json:"passwordType"`
	RandomPasswordLength                int                                `json:"randomPasswordLength"`
	CreateManagementAccount             bool                               `json:"createManagementAccount"`
	HideManagementAccount               bool                               `json:"hideManagementAccount"`
	AllowSshOnlyManagementAccount       bool                               `json:"allowSshOnlyManagementAccount"`
	EnsureSshRunning                    bool                               `json:"ensureSshRunning"`
	LaunchSelfService                   bool                               `json:"launchSelfService"`
	SignQuickAdd                        bool                               `json:"signQuickAdd"`
	DeveloperCertificateIdentity        *EnrollmentSubsetCertificate       `json:"developerCertificateIdentity"`
	DeveloperCertificateIdentityDetails EnrollmentSubsetCertificateDetails `json:"developerCertificateIdentityDetails"`
	IosEnterpriseEnrollmentEnabled      bool                               `json:"iosEnterpriseEnrollmentEnabled"`
	IosPersonalEnrollmentEnabled        bool                               `json:"iosPersonalEnrollmentEnabled"`
	PersonalDeviceEnrollmentType        string                             `json:"personalDeviceEnrollmentType"`
	AccountDrivenUserEnrollmentEnabled  bool                               `json:"accountDrivenUserEnrollmentEnabled"`
}

type EnrollmentSubsetCertificate struct {
	Filename string `json:"filename"`
	Md5Sum   string `json:"md5Sum"`
}

type EnrollmentSubsetCertificateDetails struct {
	Subject      string `json:"subject"`
	SerialNumber string `json:"serialNumber"`
}

// CRUD

func (c *Client) GetADUESessionTokenSettings() (*ResourceADUETokenSettings, error) {
	endpoint := uriUserEnrollmentTokenSettings
	var out ResourceADUETokenSettings

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "ADUE token settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}

func (c *Client) UpdateADUESessionTokenSettings(updatedSettings ResourceADUETokenSettings) (*ResourceADUETokenSettings, error) {
	endpoint := uriUserEnrollmentTokenSettings
	var out ResourceADUETokenSettings

	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedSettings, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "ADUE token settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetEnrollmentSettings retrieves the enrollment settings from the Jamf Pro server.
func (c *Client) GetEnrollmentSettings() (*EnrollmentSubsetCertificateSettings, error) {
	endpoint := uriEnrollmentSettings
	var out EnrollmentSubsetCertificateSettings

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "enrollment settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
