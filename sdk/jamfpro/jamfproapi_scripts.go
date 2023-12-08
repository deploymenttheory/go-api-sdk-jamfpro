package jamfpro

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

const uriProScripts = "/api/v1/scripts"

type ProScript struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CategoryName   string `json:"categoryName,omitempty"`
	CategoryId     string `json:"CategoryID,omitempty"`
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

type ResponseProScriptsList struct {
	Size    int         `json:"totalCount"`
	Results []ProScript `json:"results"`
}

func (c *Client) GetProScripts() (*ResponseProScriptsList, error) {
	resp, err := c.DoPaginatedGet(
		uriProScripts,
		100,
		0,
	)

	if err != nil {
		return nil, fmt.Errorf("fail %v", err)
	}

	var out ResponseProScriptsList
	out.Size = resp.Size

	for _, value := range resp.Results {
		var newObj ProScript
		mapstructure.Decode(value, &newObj)
		out.Results = append(out.Results, newObj)
	}

	return &out, nil

}

func (c *Client) GetProScriptByID(id int) (*ProScript, error) {
	endpoint := fmt.Sprintf("%s/%d", uriProScripts, id)
	fmt.Println(endpoint)
	var script ProScript
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &script)
	if err != nil {
		log.Fatalf("Failed to get script %s", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &script, nil
}
