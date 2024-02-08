package jamfpro

import (
	"fmt"
)

const uriDashboard = "/api/v1/dashboard"

type ErrorDetail struct {
	HTTPStatusCode string `json:"httpStatusCode"`
	Description    string `json:"description"`
	ID             string `json:"id"`
}

type SetupTaskOption struct {
	Available bool        `json:"available"`
	Error     ErrorDetail `json:"error"`
}

type MetricsDetail struct {
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
	Tag     string `json:"tag"`
}

type Details struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type FeatureOption struct {
	ID       string          `json:"id"`
	Title    string          `json:"title"`
	Subtitle string          `json:"subtitle"`
	Info     string          `json:"info"`
	Enabled  bool            `json:"enabled"`
	Metrics  []MetricsDetail `json:"metrics"`
	Details  []Details       `json:"details"`
	Error    ErrorDetail     `json:"error"`
}

type ResponseDashboard struct {
	SetupTaskOptions map[string]SetupTaskOption `json:"setupTaskOptions"`
	FeatureOptions   map[string][]FeatureOption `json:"featureOptions"`
}

func (c *Client) GetDashboard() (*ResponseDashboard, error) {
	uri := uriDashboard

	var out ResponseDashboard
	err := c.DoRequest("GET", uri, nil, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to get dashboard details: %v", err)
	}

	return &out, nil
}
