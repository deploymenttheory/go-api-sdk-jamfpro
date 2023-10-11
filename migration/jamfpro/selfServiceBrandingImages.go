package jamfpro

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
)

const uriSelfServiceBrandingImages = "/api/self-service/branding/images"

// UploadBrandingImage uploads a branding image to the Jamf Pro server.
func (c *Client) UploadBrandingImage(imageName string, imageData []byte) error {
	// Base64 encode the image data
	encodedData := base64.StdEncoding.EncodeToString(imageData)

	// Create the multipart payload
	var b bytes.Buffer
	boundary := "---011000010111000001101001"
	b.WriteString(fmt.Sprintf("--%s\r\nContent-Disposition: form-data; name=\"file\"\r\n\r\ndata:image/png;name=%s;base64,", boundary, imageName))
	b.WriteString(encodedData)
	b.WriteString(fmt.Sprintf("\r\n--%s--\r\n\r\n", boundary))

	// Create the request
	req, err := c.createRequest("POST", uriSelfServiceBrandingImages, &url.Values{}, b.Bytes())
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Dump the request to a file for debugging
	err = DumpRequestToFile(req, "request_dump.txt")
	if err != nil {
		fmt.Println("Error dumping request to file:", err)
	} else {
		fmt.Println("Request details written to request_dump.txt")
	}

	// Use DoRequest method to send the request and handle the response.
	err = c.DoRequest("POST", uriSelfServiceBrandingImages, b.Bytes(), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// Inside your jamf package

// CreateUploadBrandingImageRequest creates but does not send the request for branding image upload.
func (c *Client) CreateUploadBrandingImageRequest(imageName string, imageData []byte) (*http.Request, error) {
	// Base64 encode the image data
	encodedData := base64.StdEncoding.EncodeToString(imageData)

	// Create the multipart payload
	var b bytes.Buffer
	boundary := "---011000010111000001101001"
	b.WriteString(fmt.Sprintf("--%s\r\nContent-Disposition: form-data; name=\"file\"\r\n\r\ndata:image/png;name=%s;base64,", boundary, imageName))
	b.WriteString(encodedData)
	b.WriteString(fmt.Sprintf("\r\n--%s--\r\n\r\n", boundary))

	// Create the request
	req, err := c.createRequest("POST", uriSelfServiceBrandingImages, &url.Values{}, b.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Print the request headers
	PrintRequestHeaders(req)

	return req, nil
}
