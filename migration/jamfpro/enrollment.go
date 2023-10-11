package jamfpro

import (
	"fmt"
)

const uriEnrollment = "/api/v3/enrollment"

type CertificateDetails struct {
	Filename         string `json:"filename"`
	Md5Sum           string `json:"md5Sum"`
	Subject          string `json:"subject,omitempty"`
	SerialNumber     string `json:"serialNumber,omitempty"`
	IdentityKeystore string `json:"identityKeystore,omitempty"`
	KeystorePassword string `json:"keystorePassword,omitempty"`
}

type ResponseEnrollmentSettings struct {
	InstallSingleProfile                bool                `json:"installSingleProfile"`
	SigningMdmProfileEnabled            bool                `json:"signingMdmProfileEnabled"`
	MdmSigningCertificate               *CertificateDetails `json:"mdmSigningCertificate"`
	RestrictReenrollment                bool                `json:"restrictReenrollment"`
	FlushLocationInformation            bool                `json:"flushLocationInformation"`
	FlushLocationHistoryInformation     bool                `json:"flushLocationHistoryInformation"`
	FlushPolicyHistory                  bool                `json:"flushPolicyHistory"`
	FlushExtensionAttributes            bool                `json:"flushExtensionAttributes"`
	FlushMdmCommandsOnReenroll          string              `json:"flushMdmCommandsOnReenroll"`
	MacOsEnterpriseEnrollmentEnabled    bool                `json:"macOsEnterpriseEnrollmentEnabled"`
	ManagementUsername                  string              `json:"managementUsername"`
	CreateManagementAccount             bool                `json:"createManagementAccount"`
	HideManagementAccount               bool                `json:"hideManagementAccount"`
	AllowSshOnlyManagementAccount       bool                `json:"allowSshOnlyManagementAccount"`
	EnsureSshRunning                    bool                `json:"ensureSshRunning"`
	LaunchSelfService                   bool                `json:"launchSelfService"`
	SignQuickAdd                        bool                `json:"signQuickAdd"`
	DeveloperCertificateIdentity        *CertificateDetails `json:"developerCertificateIdentity"`
	DeveloperCertificateIdentityDetails CertificateDetails  `json:"developerCertificateIdentityDetails"`
	MdmSigningCertificateDetails        CertificateDetails  `json:"mdmSigningCertificateDetails"`
	IosEnterpriseEnrollmentEnabled      bool                `json:"iosEnterpriseEnrollmentEnabled"`
	IosPersonalEnrollmentEnabled        bool                `json:"iosPersonalEnrollmentEnabled"`
	PersonalDeviceEnrollmentType        string              `json:"personalDeviceEnrollmentType"`
	AccountDrivenUserEnrollmentEnabled  bool                `json:"accountDrivenUserEnrollmentEnabled"`
}

func (c *Client) GetEnrollmentAndReenrollmentSettings() (*ResponseEnrollmentSettings, error) {
	uri := uriEnrollment

	var out ResponseEnrollmentSettings
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get enrollment and reenrollment settings: %v", err)
	}

	return &out, nil
}

func (c *Client) UpdateEnrollmentAndReenrollmentSettings(settings *ResponseEnrollmentSettings) (*ResponseEnrollmentSettings, error) {
	uri := uriEnrollment

	var out ResponseEnrollmentSettings
	err := c.DoRequest("PUT", uri, settings, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to update enrollment and re-enrollment settings: %v", err)
	}

	return &out, nil
}
