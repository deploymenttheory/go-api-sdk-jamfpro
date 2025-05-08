// jamfproapi_sso_certificate.go
package jamfpro

import (
	"fmt"
)

const (
	uriSSOCert = "/api/v2/sso/cert"
)

// Resource structures
type ResourceSSOCertKeystore struct {
	Key               string            `json:"key,omitempty"`
	Keys              []ResourceCertKey `json:"keys,omitempty"`
	Type              string            `json:"type,omitempty"`
	KeystoreFileName  string            `json:"keystoreFileName,omitempty"`
	KeystoreSetupType string            `json:"keystoreSetupType,omitempty"`
}

type ResourceCertKey struct {
	ID    string `json:"id,omitempty"`
	Valid bool   `json:"valid"`
}

type ResourceSSOKeystoreDetails struct {
	Keys         []string `json:"keys,omitempty"`
	Issuer       string   `json:"issuer,omitempty"`
	Subject      string   `json:"subject,omitempty"`
	Expiration   string   `json:"expiration,omitempty"`
	SerialNumber int      `json:"serialNumber,omitempty"`
}

type ResourceSSOKeystoreResponse struct {
	Keystore        ResourceSSOCertKeystore     `json:"keystore,omitempty"`
	KeystoreDetails *ResourceSSOKeystoreDetails `json:"keystoreDetails,omitempty"`
}

// GetSSOCertificate gets the certificate currently configured for use with SSO
func (c *Client) GetSSOCertificate() (*ResourceSSOKeystoreResponse, error) {
	endpoint := uriSSOCert

	var certResponse ResourceSSOKeystoreResponse
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &certResponse)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "sso certificate", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &certResponse, nil
}

// CreateSSOCertificate generates a new certificate for signing SSO requests
func (c *Client) CreateSSOCertificate() (*ResourceSSOKeystoreResponse, error) {
	endpoint := uriSSOCert

	var certResponse ResourceSSOKeystoreResponse
	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &certResponse)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "sso certificate", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &certResponse, nil
}

// DeleteSSOCertificate deletes the currently configured certificate used by SSO
func (c *Client) DeleteSSOCertificate() error {
	endpoint := uriSSOCert

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDelete, "sso certificate", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
