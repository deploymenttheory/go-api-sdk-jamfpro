// jamfproapi_jcds2.go
// Jamf Pro Api - Jamf Cloud Distribution Service (JCDS)
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-jcds-files
// Jamf Pro API requires the structs to support an JSON data structure.
// Ref: https://grahamrpugh.com/2023/08/21/introducing-jcds2.html
// Ref: https://aws.github.io/aws-sdk-go-v2/docs/sdk-utilities/s3/

package jamfpro

// Constants for API endpoints
const uriIcon = "/api/v1/icon"

// ResponseIconUpload represents the response from the icon upload endpoint
type ResponseIconUpload struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// UploadIcon uploads an icon image file using the custom multipart format
// func (c *Client) UploadIcon(filePath string) (*ResponseIconUpload, error) {
// 	// Read and encode the file
// 	fileContent, err := os.ReadFile(filePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read file: %v", err)
// 	}

// 	// Encode to base64
// 	base64Data := base64.StdEncoding.EncodeToString(fileContent)
// 	fileName := filepath.Base(filePath)

// 	// Use custom boundary matching the example
// 	boundary := "-----011000010111000001101001"

// 	var response ResponseIconUpload
// 	_, err = c.HTTP.DoImageMultiPartUpload(
// 		http.MethodPost,
// 		uriIcon,
// 		fileName,
// 		base64Data,
// 		boundary,
// 		&response,
// 	)

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to upload icon: %v", err)
// 	}

// 	return &response, nil
// }
