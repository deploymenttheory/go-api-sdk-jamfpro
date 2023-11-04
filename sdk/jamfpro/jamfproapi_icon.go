// jamfproapi_upload_icon.go
// Jamf Pro Api - Upload Icon
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v1-icon
// This endpoint uploads an icon and returns its URL and ID.

package jamfpro

import (
	"fmt"
)

const uriUploadIcon = "/api/v1/icon"

// ResponseUploadIcon is the response structure for uploading icons.
type ResponseUploadIcon struct {
	URL string `json:"url"`
	ID  int    `json:"id"`
}

// UploadIcon uploads an icon file to Jamf Pro and returns the icon URL and ID.
func (c *Client) UploadIcon(filePath string) (*ResponseUploadIcon, error) {
	endpoint := uriUploadIcon

	// Construct the files map
	files := map[string]string{
		"file": filePath, // 'file' is the form field name for the file upload according to the provided example
	}

	// Initialize the response struct
	var uploadResponse ResponseUploadIcon

	// Call DoMultipartRequest with the method, endpoint, files, and the response struct
	resp, err := c.HTTP.DoMultipartRequest("POST", endpoint, nil, files, &uploadResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to upload icon: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Return the response struct pointer
	return &uploadResponse, nil
}
