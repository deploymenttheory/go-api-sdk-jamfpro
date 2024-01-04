// jamfproapi_certificate_authority.go
// Jamf Pro Api - Certificate Authority Information
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriCertificateAuthority = "/api/v1/pki/certificate-authority"

// ResponseActiveCertificateAuthority represents the JSON response for the active certificate authority.
type ResponseActiveCertificateAuthority struct {
	SubjectX500Principal string                                    `json:"subjectX500Principal"`
	IssuerX500Principal  string                                    `json:"issuerX500Principal"`
	SerialNumber         string                                    `json:"serialNumber"`
	Version              int                                       `json:"version"`
	NotAfter             int64                                     `json:"notAfter"`
	NotBefore            int64                                     `json:"notBefore"`
	Signature            ActiveCertificateAuthoritySubsetSignature `json:"signature"`
	KeyUsage             []string                                  `json:"keyUsage"`
	KeyUsageExtended     []string                                  `json:"keyUsageExtended"`
	SHA1Fingerprint      string                                    `json:"sha1Fingerprint"`
	SHA256Fingerprint    string                                    `json:"sha256Fingerprint"`
}

type ActiveCertificateAuthoritySubsetSignature struct {
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
		return nil, fmt.Errorf(errMsgFailedGet, "certificate authority", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &certAuth, nil
}
