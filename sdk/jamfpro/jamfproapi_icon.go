// jamfproapi_icon.go
// Jamf Pro Api - Icons
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-jcds-files
// Jamf Pro API requires the structs to support an JSON data structure.
// Ref: https://grahamrpugh.com/2023/08/21/introducing-jcds2.html
// Ref: https://aws.github.io/aws-sdk-go-v2/docs/sdk-utilities/s3/

package jamfpro

import (
	"fmt"
	"net/http"
)

// Constants for API endpoints
const uriIcon = "/api/v1/icon"

// ResponseIconUpload represents the response from the icon upload endpoint
type ResponseIconUpload struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
}

// GetIconByID retrieves icon metadata by its ID
func (c *Client) GetIconByID(id int) (*ResponseIconUpload, error) {
	endpoint := fmt.Sprintf("%s/%d", uriIcon, id)
	var iconResource ResponseIconUpload

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &iconResource)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "icon", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &iconResource, nil
}

// UploadIcon uploads an icon image file using the custom multipart format
func (c *Client) UploadIcon(filePath string) (*ResponseIconUpload, error) {
	files := map[string][]string{
		"file": {filePath},
	}

	formFields := map[string]string{}
	contentTypes := map[string]string{
		"file": "image/png",
	}
	headersMap := map[string]http.Header{}

	var response ResponseIconUpload
	resp, err := c.HTTP.DoMultiPartRequest(http.MethodPost, uriIcon, files, formFields, contentTypes, headersMap, "raw", &response)

	if err != nil {
		return nil, fmt.Errorf("failed to upload icon: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
