// jamfproapi_adcs_settings.go
// Jamf Pro Api - AD CS Settings
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/http"
)

const uriAdcsSettingsV1 = "/api/v1/pki/adcs-settings"

// ResourceAdcsSettingsV1 represents the writable AD CS configuration payload.
type ResourceAdcsSettingsV1 struct {
	ID                string                     `json:"id,omitempty"`
	DisplayName       string                     `json:"displayName,omitempty"`
	CAName            string                     `json:"caName,omitempty"`
	FQDN              string                     `json:"fqdn,omitempty"`
	AdcsURL           string                     `json:"adcsUrl,omitempty"`
	ServerCert        *ResourceAdcsCertificateV1 `json:"serverCert,omitempty"`
	ClientCert        *ResourceAdcsCertificateV1 `json:"clientCert,omitempty"`
	RevocationEnabled *bool                      `json:"revocationEnabled,omitempty"`
	APIClientID       string                     `json:"apiClientId,omitempty"`
	Outbound          *bool                      `json:"outbound,omitempty"`
}

// ResourceAdcsCertificateV1 bundles the file metadata and base64-encoded certificate.
type ResourceAdcsCertificateV1 struct {
	Filename string `json:"filename,omitempty"`
	Data     []byte `json:"data,omitempty"`
	Password string `json:"password,omitempty"`
}

// ResponseAdcsSettingsCreatedV1 captures the identifier returned after creating a configuration.
type ResponseAdcsSettingsCreatedV1 struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResponseAdcsSettingsV1 models the read-only fields returned for an AD CS configuration.
type ResponseAdcsSettingsV1 struct {
	ID                            string                     `json:"id"`
	DisplayName                   string                     `json:"displayName"`
	CAName                        string                     `json:"caName"`
	FQDN                          string                     `json:"fqdn"`
	AdcsURL                       string                     `json:"adcsUrl"`
	ServerCert                    *ResponseAdcsCertificateV1 `json:"serverCert,omitempty"`
	ClientCert                    *ResponseAdcsCertificateV1 `json:"clientCert,omitempty"`
	RevocationEnabled             bool                       `json:"revocationEnabled"`
	APIClientID                   string                     `json:"apiClientId"`
	Outbound                      bool                       `json:"outbound"`
	ConnectorLastCheckInTimestamp string                     `json:"connectorLastCheckInTimestamp"`
}

// ResponseAdcsCertificateV1 surfaces certificate details that Jamf Pro stores for AD CS.
type ResponseAdcsCertificateV1 struct {
	Filename       string `json:"filename"`
	SerialNumber   string `json:"serialNumber"`
	Subject        string `json:"subject"`
	Issuer         string `json:"issuer"`
	ExpirationDate string `json:"expirationDate"`
}

// CreateAdcsSettingsV1 creates an inbound or outbound AD CS configuration.
func (c *Client) CreateAdcsSettingsV1(settings *ResourceAdcsSettingsV1) (*ResponseAdcsSettingsCreatedV1, error) {
	endpoint := uriAdcsSettingsV1

	var response ResponseAdcsSettingsCreatedV1
	resp, err := c.HTTP.DoRequest("POST", endpoint, settings, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "AD CS settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetAdcsSettingsByIDV1 fetches a single AD CS configuration.
func (c *Client) GetAdcsSettingsByIDV1(id string) (*ResponseAdcsSettingsV1, error) {
	endpoint := fmt.Sprintf("%s/%s", uriAdcsSettingsV1, id)

	var settings ResponseAdcsSettingsV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &settings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "AD CS settings", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &settings, nil
}

// UpdateAdcsSettingsByIDV1 updates an AD CS configuration using merge-patch semantics.
func (c *Client) UpdateAdcsSettingsByIDV1(id string, settings *ResourceAdcsSettingsV1) error {
	endpoint := fmt.Sprintf("%s/%s", uriAdcsSettingsV1, id)

	resp, err := c.HTTP.DoRequest("PATCH", endpoint, settings, nil)
	if err != nil && resp == nil {
		return fmt.Errorf(errMsgFailedUpdateByID, "AD CS settings", id, err)
	}

	if resp == nil {
		return fmt.Errorf("failed to update AD CS settings %s: nil response", id)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to update AD CS settings %s: unexpected status code %d", id, resp.StatusCode)
	}

	return nil
}

// DeleteAdcsSettingsByIDV1 removes an AD CS configuration when it is no longer in use.
func (c *Client) DeleteAdcsSettingsByIDV1(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriAdcsSettingsV1, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "AD CS settings", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
