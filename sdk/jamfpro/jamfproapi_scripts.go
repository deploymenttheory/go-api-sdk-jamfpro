package jamfpro

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

const uriScripts = "/api/v1/scripts"

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

type ResponseScriptsList struct {
	Size    int              `json:"totalCount"`
	Results []ResourceScript `json:"results"`
}

type ResponseScriptCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

func (c *Client) GetScripts() (*ResponseScriptsList, error) {
	resp, err := c.DoPaginatedGet(
		uriScripts,
		100,
		0,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch scripts %v", err)
	}

	var out ResponseScriptsList
	out.Size = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceScript
		mapstructure.Decode(value, &newObj)
		out.Results = append(out.Results, newObj)
	}

	return &out, nil

}

func (c *Client) GetScriptsByID(id int) (*ResourceScript, error) {
	endpoint := fmt.Sprintf("%s/%d", uriScripts, id)
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

func (c *Client) GetScriptsByName(name string) (*ResourceScript, error) {
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

	return &ResponseScriptCreate, err
}

func (c *Client) UpdateScriptById(id string, script *ResourceScript) (*ResourceScript, error) {
	endpoint := fmt.Sprintf("%s/%s", uriScripts, id)
	var NewScript ResourceScript
	resp, err := c.HTTP.DoRequest("PUT", endpoint, script, &NewScript)

	if err != nil {
		return nil, fmt.Errorf("failed to update script, %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &NewScript, nil

}

func (c *Client) UpdateScriptByName(name string, script *ResourceScript) (*ResourceScript, error) {

	target, err := c.GetScriptsByName(name)

	if err != nil {
		return nil, fmt.Errorf("failed to get script by id, %v", err)
	}

	target_id := target.ID

	resp, err := c.UpdateScriptById(target_id, script)

	if err != nil {
		return nil, fmt.Errorf("failed to update by id, %v", err)
	}

	return &resp, nil
}
