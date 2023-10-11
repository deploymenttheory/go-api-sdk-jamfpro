package jamfpro

import (
	"errors"
	"fmt"
)

const uriScripts = "/api/v1/scripts"

type Script struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name"`
	CategoryID     string `json:"categoryId"`
	CategoryName   string `json:"categoryName"`
	Info           string `json:"info"`
	Notes          string `json:"notes"`
	Priority       string `json:"priority"`
	Parameter4     string `json:"parameter4"`
	Parameter5     string `json:"parameter5"`
	Parameter6     string `json:"parameter6"`
	Parameter7     string `json:"parameter7"`
	Parameter8     string `json:"parameter8"`
	Parameter9     string `json:"parameter9"`
	Parameter10    string `json:"parameter10"`
	Parameter11    string `json:"parameter11"`
	OsRequirements string `json:"osRequirements"`
	ScriptContents string `json:"scriptContents"`
}

type ScriptsListResponse struct {
	Count   int      `json:"totalCount,omitempty"`
	Results []Script `json:"results,omitempty"`
}

type ScriptCreateResponse struct {
	ID   string `json:"id,omitempty"`
	Href string `json:"href,omitempty"`
}

func (c *Client) GetScriptByName(name string) (*Script, error) {

	var script *Script
	d, err := c.GetScripts()
	if err != nil {
		return script, err
	}

	for _, v := range d.Results {
		if v.Name == name {
			return &v, err
		}
	}
	return script, errors.New("Script not found")
}

func (c *Client) GetScript(id string) (*Script, error) {

	var out *Script
	err := c.DoRequest("GET", fmt.Sprintf("%s/%v", uriScripts, id), nil, nil, &out)

	return out, err
}

func (c *Client) GetScripts() (*ScriptsListResponse, error) {

	out := &ScriptsListResponse{}
	err := c.DoRequest("GET", uriScripts, nil, nil, out)

	return out, err
}

func (c *Client) CreateScript(d *Script) (string, error) {

	if d.Priority == "" {
		d.Priority = "AFTER"
	}

	if d.CategoryName == "" && d.CategoryID == "" {
		d.CategoryID = "-1"
		d.CategoryName = "NONE"
	}

	resp := &ScriptCreateResponse{}
	err := c.DoRequest("POST", uriScripts, d, nil, resp)

	return resp.ID, err
}

func (c *Client) UpdateScript(d *Script) (*Script, error) {

	script := &Script{}
	err := c.DoRequest("PUT", fmt.Sprintf("%s/%v", uriScripts, d.ID), d, nil, script)

	return script, err
}

func (c *Client) DeleteScript(id string) (string, error) {

	group := &Script{}
	err := c.DoRequest("DELETE", fmt.Sprintf("%s/%v", uriScripts, id), nil, nil, group)

	return group.ID, err
}
