// jamfproapi_locales.go
// Jamf Pro Api - Locales
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
// Jamf Pro API requires the structs to support a JSON data structure.
package jamfpro

import "fmt"

const uriLocales = "/api/v1/locales"

// ResourceLocale represents a single locale in Jamf Pro
type ResourceLocale struct {
	Description string `json:"description"`
	Identifier  string `json:"identifier"`
}

// GetLocales retrieves all available locales from Jamf Pro
func (c *Client) GetLocales() ([]ResourceLocale, error) {
	endpoint := uriLocales

	var locales []ResourceLocale
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &locales)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "locales", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return locales, nil
}
