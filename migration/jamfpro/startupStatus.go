package jamfpro

import (
	"fmt"
)

const uriAPIStartupStatus = "/api/startup-status"

type ResponseStartupStatus struct {
	Step                    string `json:"step"`
	StepCode                string `json:"stepCode"`
	StepParam               string `json:"stepParam"`
	Percentage              int    `json:"percentage"`
	Warning                 string `json:"warning"`
	WarningCode             string `json:"warningCode"`
	WarningParam            string `json:"warningParam"`
	Error                   string `json:"error"`
	ErrorCode               string `json:"errorCode"`
	SetupAssistantNecessary bool   `json:"setupAssistantNecessary"`
}

func (c *Client) GetStartupStatus() (*ResponseStartupStatus, error) {
	uri := uriAPIStartupStatus

	var out ResponseStartupStatus
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get startup status: %v", err)
	}

	return &out, nil
}
