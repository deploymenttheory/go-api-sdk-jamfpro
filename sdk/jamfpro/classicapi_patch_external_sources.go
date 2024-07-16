// classicapi_patch_external_sources.go
// Jamf Pro Classic Api  - Patch External Sources
// api reference: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
// Jamf Pro Classic Api requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSelfServiceIcon
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// Constant for the Patch External Sources endpoint
const uriPatchExternalSources = "/JSSResource/patchexternalsources"

// Data Models

// ResponsePatchExternalSourcesList represents the list of patch external sources.
type ResponsePatchExternalSourcesList struct {
	PatchExternalSources []ResponsePatchExternalSourcesListItem `xml:"patch_external_source" json:"patch_external_source"`
}

// ResponsePatchExternalSourcesListItem represents a single external source item.
type ResponsePatchExternalSourcesListItem struct {
	ID   int    `xml:"id" json:"id"`
	Name string `xml:"name,omitempty" json:"name"`
}

// ResourcePatchExternalSource represents the root element of the patch external source.
type ResourcePatchExternalSource struct {
	HostName   string `xml:"host_name,omitempty" json:"host_name"`
	SSLEnabled bool   `xml:"ssl_enabled" json:"ssl_enabled"`
	Port       int    `xml:"port" json:"port"`
	ID         int    `xml:"id" json:"id"`
	Name       string `xml:"name,omitempty" json:"name"`
}

// Client Functions

// GetPatchExternalSources retrieves all patch external sources.
func (c *Client) GetPatchExternalSources() (*ResponsePatchExternalSourcesList, error) {
	endpoint := uriPatchExternalSources

	var externalSources ResponsePatchExternalSourcesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &externalSources)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "patch external sources", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &externalSources, nil
}

// GetPatchExternalSourceByID retrieves a specific patch external source by its ID.
func (c *Client) GetPatchExternalSourceByID(id int) (*ResourcePatchExternalSource, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPatchExternalSources, id)

	var externalSource ResourcePatchExternalSource
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &externalSource)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "patch external source", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &externalSource, nil
}

// GetPatchExternalSourceByName retrieves a specific patch external source by its name.
func (c *Client) GetPatchExternalSourceByName(name string) (*ResourcePatchExternalSource, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPatchExternalSources, name)

	var externalSource ResourcePatchExternalSource
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &externalSource)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "patch external source", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &externalSource, nil
}

// CreateExternalPatchSource creates a new external patch source on the Jamf Pro server.
func (c *Client) CreateExternalPatchSource(patchSource *ResourcePatchExternalSource) (*ResourcePatchExternalSource, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriPatchExternalSources)

	requestBody := struct {
		XMLName xml.Name `xml:"patch_external_source"`
		*ResourcePatchExternalSource
	}{
		ResourcePatchExternalSource: patchSource,
	}

	var responseSource ResourcePatchExternalSource
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseSource)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "patch external source", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseSource, nil
}

// UpdateExternalPatchSourceByID updates an existing external patch source by its ID on the Jamf Pro server.
func (c *Client) UpdateExternalPatchSourceByID(id int, patchSource *ResourcePatchExternalSource) (*ResourcePatchExternalSource, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriPatchExternalSources)

	requestBody := struct {
		XMLName xml.Name `xml:"patch_external_source"`
		*ResourcePatchExternalSource
	}{
		ResourcePatchExternalSource: patchSource,
	}

	var responseSource ResourcePatchExternalSource
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseSource)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "patch external source", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseSource, nil
}

// UpdateExternalPatchSourceByName updates an existing external patch source by its name on the Jamf Pro server.
func (c *Client) UpdateExternalPatchSourceByName(name string, patchSource *ResourcePatchExternalSource) (*ResourcePatchExternalSource, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPatchExternalSources, name)

	requestBody := struct {
		XMLName xml.Name `xml:"patch_external_source"`
		*ResourcePatchExternalSource
	}{
		ResourcePatchExternalSource: patchSource,
	}

	var responseSource ResourcePatchExternalSource
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseSource)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "patch external source", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseSource, nil
}

// DeleteExternalPatchSourceByID deletes an external patch source by its ID from the Jamf Pro server.
func (c *Client) DeleteExternalPatchSourceByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriPatchExternalSources, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "patch external source", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
