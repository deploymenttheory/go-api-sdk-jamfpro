// jamfproapi_self_service.go
// Jamf Pro Api - Self Service BrandingImages
// api reference: https://developer.jamf.com/jamf-pro/reference/post_self-service-branding-images
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
	"net/http"
)

// Constants for API endpoints
const uriSelfService = "/api/self-service/branding/images"

// Response

// ResponseSelfServiceBrandingImage represents the response from the branding upload endpoint
type ResponseSelfServiceBrandingImage struct {
	URL string `json:"url"`
}

// UploadSelfServiceBrandingImage uploads an branding image file using the custom multipart format
func (c *Client) UploadSelfServiceBrandingImage(filePath string) (*ResponseSelfServiceBrandingImage, error) {
	files := map[string][]string{
		"file": {filePath},
	}

	formFields := map[string]string{}
	contentTypes := map[string]string{
		"file": "image/png",
	}
	headersMap := map[string]http.Header{}

	var response ResponseSelfServiceBrandingImage
	resp, err := c.HTTP.DoMultiPartRequest(http.MethodPost, uriSelfService, files, formFields, contentTypes, headersMap, "byte", &response)

	if err != nil {
		return nil, fmt.Errorf("failed to upload branding image: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
