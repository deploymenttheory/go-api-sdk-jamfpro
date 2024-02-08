package jamfpro

const uriSSOCert = "/api/v2/sso/cert"

type KeyDetails struct {
	ID    string `json:"id"`
	Valid bool   `json:"valid"`
}

type Keystore struct {
	Key               string       `json:"key"`
	Keys              []KeyDetails `json:"keys"`
	Type              string       `json:"type"`
	KeystoreSetupType string       `json:"keystoreSetupType"`
	KeystoreFileName  string       `json:"keystoreFileName"`
}

type KeystoreDetails struct {
	Keys         []string `json:"keys"`
	SerialNumber int64    `json:"serialNumber"`
	Subject      string   `json:"subject"`
	Issuer       string   `json:"issuer"`
	Expiration   string   `json:"expiration"`
}

type SSOCertificateResponse struct {
	Keystore        Keystore        `json:"keystore"`
	KeystoreDetails KeystoreDetails `json:"keystoreDetails"`
}

// Get SSO Certificate
func (c *Client) GetSSOCertificate() (*SSOCertificateResponse, error) {
	// Check and refresh auth token if needed
	if err := c.refreshAuthToken(); err != nil {
		return nil, err
	}

	var out *SSOCertificateResponse
	headers := map[string]string{"accept": "application/json"}

	// Make the request using the client's DoRequest method
	err := c.DoRequest("GET", uriSSOCert, headers, nil, &out, c.HTTP.Logger)
	return out, err
}
