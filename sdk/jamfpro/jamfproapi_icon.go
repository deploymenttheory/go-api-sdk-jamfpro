// jamfproapi_upload_icon.go
// Jamf Pro Api - Upload Icon
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v1-icon
// This endpoint uploads an icon and returns its URL and ID.

package jamfpro

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const uriUploadIcon = "/api/v1/icon"

// Response

// ResponseUploadIcon is the response structure for uploading icons.
type ResponseUploadIcon struct {
	URL string `json:"url"`
	ID  int    `json:"id"`
}

// CRUD

// UploadIcon uploads an icon file to Jamf Pro and returns the icon URL and ID.
func (c *Client) UploadIcon(filePath string) (*ResponseUploadIcon, error) {
	endpoint := uriUploadIcon

	// Create a map for the file to be uploaded
	files := map[string]string{
		"file": filePath,
	}

	// Include form fields if needed (currently none based on docs)
	formFields := map[string]string{}

	// No custom content types for this request
	contentTypes := map[string]string{}

	// No additional headers for this request
	headersMap := map[string]http.Header{}

	var response ResponseUploadIcon
	resp, err := c.HTTP.DoMultiPartRequest("POST", endpoint, files, formFields, contentTypes, headersMap, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to upload package: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DownloadIcon downloads an icon by its ID from Jamf Pro and saves it to the specified file path.
// The icon is saved to the path provided in the 'savePath' parameter.
func (c *Client) DownloadIcon(iconID int, savePath string, res string, scale string) error {
	params := url.Values{}
	if res != "" {
		params.Add("res", res)
	}
	if scale != "" {
		params.Add("scale", scale)
	}
	queryString := params.Encode()
	endpoint := fmt.Sprintf("%s/download/%d?%s", uriUploadIcon, iconID, queryString)

	var placeholder struct{}

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &placeholder)
	if err != nil {
		return fmt.Errorf("failed to download icon: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("received non-success status code: %d", resp.StatusCode)
	}

	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	err = file.Sync()
	if err != nil {
		return fmt.Errorf("failed to sync file: %v", err)
	}

	return nil
}
