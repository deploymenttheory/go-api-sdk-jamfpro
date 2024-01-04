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

	files := map[string]string{
		"file": filePath,
	}
	var uploadResponse ResponseUploadIcon

	resp, err := c.HTTP.DoMultipartRequest("POST", endpoint, nil, files, &uploadResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to upload icon: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &uploadResponse, nil
}

// TODO DownloadIcon downloads an icon by its ID from Jamf Pro and saves it to the specified file path.
// The icon is saved to the path provided in the 'savePath' parameter.
