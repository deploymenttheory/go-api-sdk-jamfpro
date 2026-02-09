// classicapi_file_uploads.go
// Jamf Pro Classic Api - File Uploads
// api reference: https://developer.jamf.com/jamf-pro/reference/uploadfiles
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"fmt"
	"net/http"
)

// Constants for API endpoints and valid resources
const uriClassicFileUploads = "/JSSResource/fileuploads"

// ValidFileUploadResources contains the list of valid resources for file uploads
var ValidFileUploadResources = []string{
	"computers",
	"mobiledevices",
	"enrollmentprofiles",
	"printers",
	"peripherals",
	"policies",
	"ebooks",
	"mobiledeviceapplications",
	"icon",
	"mobiledeviceapplicationsipa",
	"diskencryptionconfigurations",
}

// ResourceIDType represents the type of identifier being used (id or name)
type ResourceIDType string

const (
	ResourceIDTypeID   ResourceIDType = "id"
	ResourceIDTypeName ResourceIDType = "name"
)

// CreateFileAttachment uploads a file to a specific resource in Jamf Pro
func (c *Client) CreateFileAttachment(resource string, idType ResourceIDType, identifier string, filePath string, forceIpaUpload bool) error {
	// Validate resource
	validResource := false
	for _, r := range ValidFileUploadResources {
		if r == resource {
			validResource = true
			break
		}
	}
	if !validResource {
		return fmt.Errorf("invalid resource type: %s", resource)
	}

	// Validate idType
	if idType != ResourceIDTypeID && idType != ResourceIDTypeName {
		return fmt.Errorf("invalid ID type: %s", idType)
	}

	// For peripherals, only ID is supported
	if resource == "peripherals" && idType == ResourceIDTypeName {
		return fmt.Errorf("peripherals resource only supports ID type")
	}

	// Construct endpoint
	endpoint := fmt.Sprintf("%s/%s/%s/%s", uriClassicFileUploads, resource, idType, identifier)

	// Add query parameter for IPA upload if specified
	if forceIpaUpload && resource == "mobiledeviceapplicationsipa" {
		endpoint = fmt.Sprintf("%s?FORCE_IPA_UPLOAD=true", endpoint)
	}

	// Setup file upload
	files := map[string][]string{
		"name": {filePath},
	}

	formFields := map[string]string{}
	contentTypes := map[string]string{}
	headersMap := map[string]http.Header{}

	// Perform the upload
	var response any // Classic API typically doesn't return a response for file uploads
	resp, err := c.HTTP.DoMultiPartRequest(
		http.MethodPost,
		endpoint,
		files,
		formFields,
		contentTypes,
		headersMap,
		"byte",
		&response,
	)

	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
