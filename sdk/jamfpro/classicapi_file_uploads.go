// classicapi_file_uploads.go
// Jamf Pro Classic Api - File Uploads
// api reference: https://developer.jamf.com/jamf-pro/reference/uploadfiles
// Classic API requires the structs to support an XML data structure.

package jamfpro

const uriFileUploads = "/JSSResource/fileuploads"

// CRUD

// CreateFileAttachments uploads file attachments to a specific resource in Jamf Pro.
// The function assumes that the file paths are provided as a map where the keys are the form field names.
// func (c *Client) CreateFileAttachments(resource, idType, id string, filePaths map[string]string) (*http.Response, error) {
// 	endpoint := fmt.Sprintf("%s/%s/%s/%s", uriFileUploads, resource, idType, id)

// 	if resource == "mobiledeviceapplicationsipa" {
// 		endpoint += "?FORCE_IPA_UPLOAD=true"
// 	} else {
// 		endpoint += "?FORCE_IPA_UPLOAD=false"
// 	}

// 	// Use the file paths directly
// 	files := filePaths

// 	// Include form fields if needed (currently none based on docs)
// 	formFields := map[string]string{}

// 	// No custom content types for this request
// 	contentTypes := map[string]string{}

// 	// No additional headers for this request
// 	headersMap := map[string]http.Header{}

// 	var response ResponsePackageCreatedAndUpdated
// 	resp, err := c.HTTP.DoMultiPartRequest("POST", endpoint, files, formFields, contentTypes, headersMap, &response)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to upload package: %v", err)
// 	}
// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return &response, nil
// }
