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

// CRUD

// CreateFileAttachments uploads file attachments to a specific resource in Jamf Pro.
// The function assumes that the file paths are provided as a map where the keys are the form field names.
func (c *Client) CreateFileAttachments(resource, idType, id string, files map[string]string) (*http.Response, error) {
	endpoint := fmt.Sprintf("%s/%s/%s/%s", uriFileUploads, resource, idType, id)

	if resource == "mobiledeviceapplicationsipa" {
		endpoint += "?FORCE_IPA_UPLOAD=true"
	} else {
		endpoint += "?FORCE_IPA_UPLOAD=false"
	}

	resp, err := c.HTTP.DoMultipartRequest("POST", endpoint, nil, files, nil)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "attachment", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return resp, nil
}
