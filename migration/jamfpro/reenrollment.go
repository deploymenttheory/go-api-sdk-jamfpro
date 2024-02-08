package jamfpro

import (
	"fmt"
)

const uriAPIReEnrollment = "/api/v1/reenrollment"

type ResponseReEnrollment struct {
	IsFlushPolicyHistoryEnabled              bool   `json:"isFlushPolicyHistoryEnabled"`
	IsFlushLocationInformationEnabled        bool   `json:"isFlushLocationInformationEnabled"`
	IsFlushLocationInformationHistoryEnabled bool   `json:"isFlushLocationInformationHistoryEnabled"`
	IsFlushExtensionAttributesEnabled        bool   `json:"isFlushExtensionAttributesEnabled"`
	FlushMDMQueue                            string `json:"flushMDMQueue"`
}

type ReEnrollmentHistoryResult struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

type ResponseReEnrollmentHistory struct {
	TotalCount int                         `json:"totalCount"`
	Results    []ReEnrollmentHistoryResult `json:"results"`
}

func (c *Client) GetReEnrollment() (*ResponseReEnrollment, error) {
	uri := uriAPIReEnrollment

	var out ResponseReEnrollment
	err := c.DoRequest("GET", uri, nil, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to get re-enrollment data: %v", err)
	}

	return &out, nil
}

func (c *Client) UpdateReEnrollment(d *ResponseReEnrollment) (*ResponseReEnrollment, error) {
	uri := uriAPIReEnrollment

	updatedReEnrollment := &ResponseReEnrollment{}

	err := c.DoRequest("PUT", uri, d, nil, updatedReEnrollment)
	if err != nil {
		return nil, fmt.Errorf("failed to update re-enrollment data: %v", err)
	}

	return updatedReEnrollment, nil
}

func (c *Client) GetReEnrollmentHistory(page, pageSize int, sort string) (*ResponseReEnrollmentHistory, error) {
	uri := fmt.Sprintf("%s/history?page=%d&page-size=%d&sort=%s", uriAPIReEnrollment, page, pageSize, sort)

	var out ResponseReEnrollmentHistory
	err := c.DoRequest("GET", uri, nil, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to get re-enrollment history: %v", err)
	}

	return &out, nil
}
