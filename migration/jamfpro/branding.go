//go:build ignore

package jamfpro

import (
	"fmt"
)

const uriBrandingImages = "/api/v1/branding-images/download/"

// DownloadSelfServiceBrandingImage downloads the branding image by its ID
func (c *Client) DownloadSelfServiceBrandingImage(imageID string) ([]byte, error) {
	uri := fmt.Sprintf("%s%s", uriBrandingImages, imageID)

	// Set the custom header
	c.ExtraHeader = map[string]string{
		"accept": "image/*",
	}

	var out []byte
	err := c.DoRequest("GET", uri, nil, nil, &out)

	// Clear the custom header after the request is made
	c.ExtraHeader = nil

	return out, err
}
