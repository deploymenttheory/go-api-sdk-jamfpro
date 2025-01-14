// util_pagination.go
// Api documentaton: https://developer.jamf.com/developer-guide/docs/api-style-guide#query-parameters
package jamfpro

import (
	"fmt"
	"log"
)

type StandardPaginatedResponse struct {
	Size    int           `json:"totalCount"`
	Results []interface{} `json:"results"`
}

// DoPaginatedGet performs a paginated GET request to a specified endpoint in the Jamf Pro API.
//
// This method is designed to fetch data in a paginated manner from Jamf Pro's RESTful API. It constructs
// the API endpoint using the provided parameters, handles the pagination logic, and accumulates the results
// into a single response structure. It's particularly useful for endpoints where the response is expected to
// contain a large number of items that might be paginated by the server.
//
// Parameters:
//   - endpoint_root: The root URL of the API endpoint. This is the base URL to which pagination and sorting
//     parameters will be appended.
//   - maxPageSize: Maximum number of items to be fetched in each paginated request. If set to 0, defaults to 200.
//   - startingPageNumber: The page number from which to start the paginated fetching.
//   - sort_filter: A string specifying the sorting criteria. It follows the format
//     'sort=<field_name>[:sort_direction][,<secondary_sort_field_name>[:sort_direction]]*'. The default sort
//     direction is 'asc' (Ascending). Use 'desc' for Descending ordering. Additional sort parameters are
//     supported and determine the order of results that have equivalent values for previous sort parameters.
//
// The method returns a pointer to a StandardPaginatedResponse containing the aggregated results from all
// fetched pages, or an error if the fetch operation fails at any point.
//
// Example usage:
// client.GetSelfServiceBrandingMacOS("sort=id:desc")
//
// Note:
// The method appends the results from each page to a slice and breaks the loop when the total number of items
// fetched matches the reported total count from the server, or when a fetched page contains fewer items than
// the maximum page size, indicating that it is the last page.
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
		endpoint := fmt.Sprintf("%s?page=%d&page-size=%d%s", endpoint_root, startingPageNumber, maxPageSize, sort_filter)
		log.Printf("DEBUG: Fetching from endpoint: %s", endpoint)
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
