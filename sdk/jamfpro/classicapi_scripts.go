// classicapi_scripts.go
// Jamf Pro Classic Api - Scripts
// api reference: https://developer.jamf.com/jamf-pro/reference/scripts
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriScripts = "/JSSResource/scripts"

type ResponseScript struct {
	ID                    int        `xml:"id"`
	Name                  string     `xml:"name"`
	Category              string     `xml:"category,omitempty"`
	Filename              string     `xml:"filename,omitempty"`
	Info                  string     `xml:"info,omitempty"`
	Notes                 string     `xml:"notes,omitempty"`
	Priority              string     `xml:"priority,omitempty"`
	Parameters            Parameters `xml:"parameters"`
	OSRequirements        string     `xml:"os_requirements,omitempty"`
	ScriptContents        string     `xml:"script_contents,omitempty"`
	ScriptContentsEncoded string     `xml:"script_contents_encoded,omitempty"`
}

type Parameters struct {
	Parameter4  string `xml:"parameter4,omitempty"`
	Parameter5  string `xml:"parameter5,omitempty"`
	Parameter6  string `xml:"parameter6,omitempty"`
	Parameter7  string `xml:"parameter7,omitempty"`
	Parameter8  string `xml:"parameter8,omitempty"`
	Parameter9  string `xml:"parameter9,omitempty"`
	Parameter10 string `xml:"parameter10,omitempty"`
	Parameter11 string `xml:"parameter11,omitempty"`
}

// Scripts List Structs
type ResponseScriptsList struct {
	Size   int          `xml:"size"`
	Script []ScriptItem `xml:"script"`
}

type ScriptItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// GetScripts retrieves a list of all scripts.
func (c *Client) GetScripts() (*ResponseScriptsList, error) {
	endpoint := uriScripts

	var scriptsList ResponseScriptsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &scriptsList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all scripts: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &scriptsList, nil
}

// GetScriptsByID retrieves the details of a script by its ID.
func (c *Client) GetScriptsByID(id int) (*ResponseScript, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriScripts, id)

	var scriptDetails ResponseScript
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &scriptDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch script by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &scriptDetails, nil
}

// GetScriptsByName retrieves the details of a script by its name.
func (c *Client) GetScriptsByName(name string) (*ResponseScript, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriScripts, name)

	var scriptDetails ResponseScript
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &scriptDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch script by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &scriptDetails, nil
}

// CreateScriptByID creates a new script by its ID.
func (c *Client) CreateScriptByID(script *ResponseScript) (*ResponseScript, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriScripts, script.ID)

	// Wrap the script with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"script"`
		*ResponseScript
	}{
		ResponseScript: script,
	}

	var responseScript ResponseScript
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseScript)
	if err != nil {
		return nil, fmt.Errorf("failed to create script by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseScript, nil
}

// UpdateScriptByID updates an existing script by its ID.
func (c *Client) UpdateScriptByID(script *ResponseScript) (*ResponseScript, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriScripts, script.ID)

	// Wrap the script with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"script"`
		*ResponseScript
	}{
		ResponseScript: script,
	}

	var updatedScript ResponseScript
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedScript)
	if err != nil {
		return nil, fmt.Errorf("failed to update script by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedScript, nil
}

// UpdateScriptByName updates an existing script by its name.
func (c *Client) UpdateScriptByName(script *ResponseScript) (*ResponseScript, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriScripts, script.Name)

	// Wrap the script with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"script"`
		*ResponseScript
	}{
		ResponseScript: script,
	}

	var updatedScript ResponseScript
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedScript)
	if err != nil {
		return nil, fmt.Errorf("failed to update script by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedScript, nil
}

// DeleteScriptByID deletes a script by its ID.
func (c *Client) DeleteScriptByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriScripts, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete script by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteScriptByName deletes a script by its name.
func (c *Client) DeleteScriptByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriScripts, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete script by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
