// jamfproapi_cloud_ldap_keystore.go
// Jamf Pro Api - Cloud LDAP
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v2-cloud-ldaps
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriCloudLdapKeystore = "/api/v1/ldap-keystore/verify"

// Payload

type PayloadCloudLdapVerifyKeystore struct {
	Password  string `json:"password"`
	FileBytes string `json:"fileBytes"`
	FileName  string `json:"fileName"`
}

// Response

type ResponseCloudLdapVerifyKeystore struct {
	Type           string `json:"type"`
	ExpirationDate string `json:"expirationDate"`
	Subject        string `json:"subject"`
	FileName       string `json:"fileName"`
}

// CRUD

func (c *Client) ValidateCloudLdapKeystore(payload PayloadCloudLdapVerifyKeystore) (*ResponseCloudLdapVerifyKeystore, error) {
	endpoint := uriCloudLdapKeystore
	var out ResponseCloudLdapVerifyKeystore
	resp, err := c.HTTP.DoRequest("POST", endpoint, payload, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedValidateCloudLdapKeystore, err)
	}

	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}

	return &out, nil
}
