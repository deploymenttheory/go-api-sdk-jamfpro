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
func (c *Client) CreateFileAttachments(resource, idType, id string, files map[string]string) (*http.Response, error) {
	// Construct the full endpoint URL
	endpoint := fmt.Sprintf("%s/%s/%s/%s", uriFileUploads, resource, idType, id)

	// Add the FORCE_IPA_UPLOAD query parameter if the resource is 'mobiledeviceapplicationsipa'
	if resource == "mobiledeviceapplicationsipa" {
		endpoint += "?FORCE_IPA_UPLOAD=true"
	} else {
		endpoint += "?FORCE_IPA_UPLOAD=false"
	}

	// Call DoMultipartRequest with the method, endpoint, and files
	resp, err := c.HTTP.DoMultipartRequest("POST", endpoint, nil, files, nil) // Pass nil instead of pre-initialized http.Response
	if err != nil {
		return nil, fmt.Errorf("failed to upload file attachments: %v", err)
	}

	// Ensure response body is closed after function call
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Return the response from DoMultipartRequest
	return resp, nil
}
