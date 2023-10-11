package jamfpro

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"os"
)

const uriAPIIcon = "/api/v1/icon"

type ResponseIcon struct {
	URL string `json:"url"`
	ID  int    `json:"id"`
}

func (c *Client) GetIconByID(iconID int) (*ResponseIcon, error) {
	uri := fmt.Sprintf("%s/%d", uriAPIIcon, iconID)

	var out ResponseIcon
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get icon by ID: %v", err)
	}

	return &out, nil
}

// UploadIcon uploads an icon and returns the response containing the URL and ID of the uploaded icon
func (c *Client) UploadIcon(iconPath string, filename string) (*ResponseIcon, error) {
	// Read the icon file
	iconData, err := os.ReadFile(iconPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read icon file: %v", err)
	}

	// Base64 encode the icon data
	encodedIcon, err := Base64Encode(iconData)
	if err != nil {
		return nil, fmt.Errorf("failed to encode icon: %v", err)
	}

	// Prepare multipart/form-data body
	body := &bytes.Buffer{}
	boundary := "---011000010111000001101001"
	writer := multipart.NewWriter(body)
	writer.SetBoundary(boundary) // Setting the same boundary

	imageContentType := GetImageContentType(iconPath) // func in utilities.go
	partHeader := textproto.MIMEHeader{}
	partHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename))
	partHeader.Set("Content-Type", imageContentType)

	part, err := writer.CreatePart(partHeader)
	if err != nil {
		return nil, err
	}
	part.Write([]byte(encodedIcon)) // Writing the base64 encoded data

	writer.Close()

	// Use the DoRequest method to send the request
	var iconResponse ResponseIcon
	err = c.DoRequestDebug("POST", uriAPIIcon, body, nil, &iconResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to upload icon: %v", err)
	}

	return &iconResponse, nil
}

// DownloadSelfServiceIconByID downloads an icon by its ID and returns the icon's binary data.
func (c *Client) DownloadSelfServiceIconByID(iconID int, res string, scale string) ([]byte, error) {
	apiEndpoint := fmt.Sprintf("/api/v1/icon/download/%d?res=%s&scale=%s", iconID, res, scale)

	// Call the DoRequestDebug function which is structured to handle binary responses
	var response []byte
	err := c.DoRequestDebug("GET", apiEndpoint, nil, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to download icon: %v", err)
	}

	return response, nil
}
