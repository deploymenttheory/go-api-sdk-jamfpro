// jamfproapi_oidc.go
// Jamf Pro Api - OIDC
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-oidc-public-key
// https://developer.jamf.com/jamf-pro/reference/get_v1-oidc-direct-idp-login-url
// Jamf Pro Api requires the structs to support an JSON data structure.

package jamfpro

import "fmt"

const uriOIDC = "/api/v1/oidc"

// ResponseOIDCDirectIdPLoginURL represents the response structure for the OIDC Direct IdP Login URL.
type ResponseOIDCDirectIdPLoginURL struct {
	URL string `json:"url"`
}

// ResponseOIDCPublicKey represents the response structure for the OIDC public key.
type ResponseOIDCPublicKey struct {
	Keys []ResourceOIDCKey `json:"keys"`
}

// ResourceOIDCKey represents a single key in the OIDC public key response.
type ResourceOIDCKey struct {
	Kty string `json:"kty"`
	E   string `json:"e"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
	Iat int64  `json:"iat"`
	N   string `json:"n"`
}

// ResourceOIDCRedirectURL represents the request body for getting the OIDC redirect URL.
type ResourceOIDCRedirectURL struct {
	OriginalURL  string `json:"originalUrl"`
	EmailAddress string `json:"emailAddress"`
}

// ResponseOIDCRedirectURL represents the response structure for the OIDC redirect URL.
type ResponseOIDCRedirectURL struct {
	RedirectURL string `json:"redirectUrl"`
}

// GetDirectURLForOIDCLogin retrieves the direct IdP login URL for OIDC.
func (c *Client) GetDirectURLForOIDCLogin() (*ResponseOIDCDirectIdPLoginURL, error) {
	endpoint := fmt.Sprintf("%s/direct-idp-login-url", uriOIDC)

	var response ResponseOIDCDirectIdPLoginURL
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get OIDC direct IdP login URL: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetPublicKeyOfOIDCKeystore retrieves the public key of the keystore used for signing OIDC messages as a JWT.
func (c *Client) GetPublicKeyOfOIDCKeystore() (*ResponseOIDCPublicKey, error) {
	endpoint := fmt.Sprintf("%s/public-key", uriOIDC)

	var response ResponseOIDCPublicKey
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get OIDC public key: %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GenerateKeystoreForOIDCMessages generates a new certificate used for signing OIDC messages.
func (c *Client) GenerateKeystoreForOIDCMessages() error {
	endpoint := fmt.Sprintf("%s/generate-certificate", uriOIDC)

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to generate OIDC certificate: %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}

// SetRedirectURLForOIDCLogon provides the URL to redirect for OIDC login based on the original URL and email address.
func (c *Client) SetRedirectURLForOIDCLogon(request *ResourceOIDCRedirectURL) (*ResponseOIDCRedirectURL, error) {
	endpoint := fmt.Sprintf("%s/dispatch", uriOIDC)

	var response ResponseOIDCRedirectURL
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get OIDC redirect URL: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
