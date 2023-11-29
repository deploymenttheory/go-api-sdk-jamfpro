// classicapi_file_uploads.go
// Jamf Pro Classic Api - File Uploads
// api reference: https://developer.jamf.com/jamf-pro/reference/uploadfiles
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"fmt"
	"net/http"
)

const uriFileUploads = "/JSSResource/fileuploads"

// CreateFileAttachments uploads file attachments to a specific resource in Jamf Pro.
// The function assumes that the file paths are provided as a map where the keys are the form field names.
func (c *Client) CreateFileAttachments(resource, idType string, id string, files map[string]string) (*http.Response, error) {
	endpoint := fmt.Sprintf("%s/%s/%s/%s", uriFileUploads, resource, idType, id)

	// Initialize the response struct
	var response http.Response

	// Call DoMultipartRequest with the method, endpoint, and the files
	resp, err := c.HTTP.DoMultipartRequest("POST", endpoint, nil, files, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file attachments: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Return the http.Response pointer
	return resp, nil
}
