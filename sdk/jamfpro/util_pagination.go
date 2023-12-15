package jamfpro

import (
	"fmt"
)

type StandardPaginatedResponse struct {
	Size    int           `json:"totalCount"`
	Results []interface{} `json:"results"`
}

func (c *Client) DoPaginatedGet(
	endpoint_root string,
	maxPageSize, startingPageNumber int,
	sort_filter string,
) (*StandardPaginatedResponse, error) {

	if maxPageSize == 0 {
		maxPageSize = 200
	}

	var OutStruct StandardPaginatedResponse
	var TargetObjectAccumulator StandardPaginatedResponse
	var OutData []interface{}
	var page = startingPageNumber

	for {
		endpoint := fmt.Sprintf("%s?page=%d&page-size=%d%s", endpoint_root, maxPageSize, startingPageNumber, sort_filter)
		fmt.Println(endpoint)
		resp, err := c.HTTP.DoRequest(
			"GET",
			endpoint,
			nil,
			&TargetObjectAccumulator,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to fetch obj %v", err)
		}

		if resp != nil {
			defer resp.Body.Close()
		}

		OutData = append(OutData, TargetObjectAccumulator.Results...)

		if len(OutData) >= TargetObjectAccumulator.Size ||
			len(TargetObjectAccumulator.Results) < maxPageSize ||
			len(TargetObjectAccumulator.Results) == 0 {
			break
		}

		page++

	}

	OutStruct.Size = TargetObjectAccumulator.Size
	OutStruct.Results = OutData

	return &OutStruct, nil

}
