// shared_schema.go
// Jamf Pro Api - The swagger schema for jamf pro and classic api
// api reference: https://instance-name.jamfcloud.com/api/doc/
// The swagger schema is a shared schema for all the api endpoints and written in a JSON format.

package jamfpro

import (
	"bytes"
	"fmt"
	"os"
)

const uriSharedSchema = "/api/schema/"

// DownloadJamfProSchemaToFile retrieves the schema and saves it to a specified file
func (c *Client) DownloadJamfProSchemaToFile(filePath string) error {
	endpoint := uriSharedSchema

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Perform the API request using the DoDownloadRequest method
	resp, err := c.HTTP.DoDownloadRequest("GET", endpoint, file)
	if err != nil {
		return fmt.Errorf("failed to download schema: %w", err)
	}
	defer resp.Body.Close()

	return nil
}

// DownloadJamfProSchema retrieves the schema and returns it as a byte slice
func (c *Client) DownloadJamfProSchema() ([]byte, error) {
	endpoint := uriSharedSchema

	// Create a buffer to hold the downloaded data
	var buffer bytes.Buffer

	// Perform the API request using the DoDownloadRequest method
	resp, err := c.HTTP.DoDownloadRequest("GET", endpoint, &buffer)
	if err != nil {
		return nil, fmt.Errorf("failed to download schema: %w", err)
	}
	defer resp.Body.Close()

	return buffer.Bytes(), nil
}
