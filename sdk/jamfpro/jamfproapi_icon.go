// jamfproapi_upload_icon.go
// Jamf Pro Api - Upload Icon
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v1-icon
// This endpoint uploads an icon and returns its URL and ID.

package jamfpro

import (
	"fmt"
	"io"
	"net/url"
	"os"
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
		"file": filePath, // 'file' is the form field name for the file uploads from jamf docs
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

// DownloadIcon downloads an icon by its ID from Jamf Pro and saves it to the specified file path.
// The icon is saved to the path provided in the 'savePath' parameter.
func (c *Client) DownloadIcon(iconID int, savePath string, res string, scale string) error {
	// Construct the endpoint with query parameters
	params := url.Values{}
	if res != "" {
		params.Add("res", res)
	}
	if scale != "" {
		params.Add("scale", scale)
	}
	queryString := params.Encode()
	endpoint := fmt.Sprintf("%s/download/%d?%s", uriUploadIcon, iconID, queryString)

	// Create the file to write to
	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Initialize an empty struct since we are interested in the response body (the file content)
	var placeholder struct{}

	// Call DoRequest with the method, endpoint, nil for body since this is a GET request, and the placeholder for the response struct
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &placeholder)
	if err != nil {
		return fmt.Errorf("failed to download icon: %v", err)
	}
	defer resp.Body.Close()

	// Write the response body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	// Ensure the file is written to disk
	err = file.Sync()
	if err != nil {
		return fmt.Errorf("failed to sync file: %v", err)
	}

	// Return nil if no error occurred, indicating success
	return nil
}
