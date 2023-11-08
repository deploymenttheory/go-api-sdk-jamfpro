// jamfproapi_certificate_authority.go
// Jamf Pro Classic Api - Certificate Authority Information
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"io"
)

const uriCertificateAuthority = "/api/v1/pki/certificate-authority"

// ResponseActiveCertificateAuthority represents the JSON response for the active certificate authority.
type ResponseActiveCertificateAuthority struct {
	SubjectX500Principal string    `json:"subjectX500Principal"`
	IssuerX500Principal  string    `json:"issuerX500Principal"`
	SerialNumber         string    `json:"serialNumber"`
	Version              int       `json:"version"`
	NotAfter             int64     `json:"notAfter"`
	NotBefore            int64     `json:"notBefore"`
	Signature            Signature `json:"signature"`
	KeyUsage             []string  `json:"keyUsage"`
	KeyUsageExtended     []string  `json:"keyUsageExtended"`
	SHA1Fingerprint      string    `json:"sha1Fingerprint"`
	SHA256Fingerprint    string    `json:"sha256Fingerprint"`
}

// Signature represents the signature part of a certificate authority.
type Signature struct {
	Algorithm    string `json:"algorithm"`
	AlgorithmOID string `json:"algorithmOid"`
	Value        string `json:"value"`
}

// GetActiveCertificateAuthority retrieves the active certificate authority details.
func (c *Client) GetActiveCertificateAuthority() (*ResponseActiveCertificateAuthority, error) {
	endpoint := fmt.Sprintf("%s/active", uriCertificateAuthority)

	var certAuth ResponseActiveCertificateAuthority
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &certAuth)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch active certificate authority: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &certAuth, nil
}

// GetActiveCertificateAuthorityInDER retrieves the active certificate authority in DER format.
func (c *Client) GetActiveCertificateAuthorityInDER() (string, error) {
	endpoint := fmt.Sprintf("%s/active/der", uriCertificateAuthority)

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, nil)
	if err != nil {
		return "", fmt.Errorf("failed to fetch active certificate authority DER: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Assuming the response body is the raw DER string
	derBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read DER data: %v", err)
	}

	return string(derBytes), nil
}
