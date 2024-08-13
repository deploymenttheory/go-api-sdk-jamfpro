// jamfproapi_dss_declarations.go
// Jamf Pro Api - Declaration from DSS
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-dss-declarations-id
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriDSSDeclarations = "/api/v1/dss-declarations"

// Responses
type ResponseDSSDeclaration struct {
	Declarations []ResourceDSSDeclaration `json:"declarations"`
}

// Resource
type ResourceDSSDeclaration struct {
	UUID        string `json:"uuid"`
	PayloadJson string `json:"payloadJson"`
	Type        string `json:"type"`
	Group       string `json:"group"`
}

// GetDSSDeclarationByUUID retrieves a DSS declaration by UUID.
func (c *Client) GetDSSDeclarationByUUID(uuid string) (*ResponseDSSDeclaration, error) {
	endpoint := fmt.Sprintf("%s/%v", uriDSSDeclarations, uuid)
	var out ResponseDSSDeclaration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "DSS declaration", uuid, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
