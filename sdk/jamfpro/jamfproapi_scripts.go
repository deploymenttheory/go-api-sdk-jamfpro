package jamfpro

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

const uriScripts = "/api/v1/scripts"

// Struct which represents Script object JSON from Pro API
type ResourceScript struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CategoryName   string `json:"categoryName,omitempty"`
	CategoryId     string `json:"categoryId,omitempty"`
	Info           string `json:"info,omitempty"`
	Notes          string `json:"notes,omitempty"`
	OSRequirements string `json:"osRequirements,omitempty"`
	Priority       string `json:"priority,omitempty"`
	ScriptContents string `json:"scriptContents,omitempty"`
	Parameter4     string `json:"parameter4,omitempty"`
	Parameter5     string `json:"parameter5,omitempty"`
	Parameter6     string `json:"parameter6,omitempty"`
	Parameter7     string `json:"parameter7,omitempty"`
	Parameter8     string `json:"parameter8,omitempty"`
	Parameter9     string `json:"parameter9,omitempty"`
	Parameter10    string `json:"parameter10,omitempty"`
	Parameter11    string `json:"parameter11,omitempty"`
}

// Struct for paginated response for scripts
type ResponseScriptsList struct {
	Size    int              `json:"totalCount"`
	Results []ResourceScript `json:"results"`
}

// Response format struct for create function
type ResponseScriptCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Gets full list of scripts & handles pagination
func (c *Client) GetScripts() (*ResponseScriptsList, error) {
	resp, err := c.DoPaginatedGet(
		uriScripts,
		standardPageSize,
		startingPageNumber,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch scripts %v", err)
	}

	var out ResponseScriptsList
	out.Size = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceScript
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf("failed to map struct, %v", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil

}

// Retrieves script from provided ID & returns ResourceScript
func (c *Client) GetScriptByID(id string) (*ResourceScript, error) {
	endpoint := fmt.Sprintf("%s/%s", uriScripts, id)
	var script ResourceScript
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &script)
	if err != nil {
		log.Fatalf("Failed to get script %s", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &script, nil
}

// Retrieves script by Name by leveraging GetScripts(), returns ResourceScript
func (c *Client) GetScriptByName(name string) (*ResourceScript, error) {
	scripts, err := c.GetScripts()
	if err != nil {
		return nil, fmt.Errorf("failed to get script by name, %v", err)
	}

	for _, value := range scripts.Results {
		if value.Name == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf("failed to locate script by name %v", name)
}

// Creates script from ResourceScript struct
func (c *Client) CreateScript(script *ResourceScript) (*ResponseScriptCreate, error) {
	endpoint := uriScripts
	var ResponseScriptCreate ResponseScriptCreate

	resp, err := c.HTTP.DoRequest("POST", endpoint, script, &ResponseScriptCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create script, %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &ResponseScriptCreate, nil
}

// Updates script from provided ResourceScript - only updates provided keys
func (c *Client) UpdateScriptByID(id string, script *ResourceScript) (*ResourceScript, error) {
	endpoint := fmt.Sprintf("%s/%s", uriScripts, id)
	var updatedScript ResourceScript
	resp, err := c.HTTP.DoRequest("PUT", endpoint, script, &updatedScript)

	if err != nil {
		return nil, fmt.Errorf("failed to update script, %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &updatedScript, nil

}

// Leverages UpdateScriptByID and GetScripts to update script from provided ResourceScript
func (c *Client) UpdateScriptByName(name string, script *ResourceScript) (*ResourceScript, error) {

	target, err := c.GetScriptByName(name)

	if err != nil {
		return nil, fmt.Errorf("failed to get script by id, %v", err)
	}

	target_id := target.ID
	resp, err := c.UpdateScriptByID(target_id, script)

	if err != nil {
		return nil, fmt.Errorf("failed to update by id, %v", err)
	}

	return resp, nil
}

// Deletes script with provided ID
func (c *Client) DeleteScriptByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriScripts, id)
	var response interface{}
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to delete script %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}

// Leverages DeleteScriptByID and GetScripts to delete script by Name
func (c *Client) DeleteScriptByName(name string) error {
	target, err := c.GetScriptByName(name)
	if err != nil {
		return fmt.Errorf("failed to get script by name, %v", err)
	}

	target_id := target.ID

	err = c.DeleteScriptByID(target_id)
	if err != nil {
		return fmt.Errorf("failed to delete script by found id, %v", err)
	}

	return fmt.Errorf("an error occured")
}
