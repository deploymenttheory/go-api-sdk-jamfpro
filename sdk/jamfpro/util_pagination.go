// util_pagination.go
// Api documentaton: https://developer.jamf.com/developer-guide/docs/api-style-guide#query-parameters
package jamfpro

import (
	"fmt"
	"net/url"
	"strconv"
)

type ResponsePaginated struct {
	Size    int   `json:"totalCount"`
	Results []any `json:"results"`
}

// DoPaginatedGet retrieves paginated results from a Jamf Pro API endpoint.
//
// It constructs the request with optional sorting and pagination parameters, performs repeated GETs
// until all pages are fetched (based on reported total size or final page), and accumulates results.
//
// Parameters:
//   - endpoint_root: Base URL for the API resource.
//   - params: Query parameters including optional "page", "page-size", and "sort".
//
// Returns a combined StandardPaginatedResponse or an error if a request fails.
func (c *Client) DoPaginatedGet(endpoint_root string, params url.Values) (*ResponsePaginated, error) {

	if params == nil {
		params = url.Values{}
	}

	if params.Get("page") == "" {
		// Some warning log should be here
		params.Add("page", "0")
	}

	if params.Get("page-size") == "" {
		// and here
		params.Add("page-size", "100")
	}

	var outStruct ResponsePaginated
	var targetObjectAccumulator ResponsePaginated
	var outData []any
	var page, err = strconv.Atoi(params.Get("page"))

	if err != nil {
		return nil, fmt.Errorf("error converting page number: %v", err)
	}

	for {
		targetObjectAccumulator = ResponsePaginated{}
		encodedParams := params.Encode()
		endpoint := fmt.Sprintf("%s?%s", endpoint_root, encodedParams)

		resp, err := c.HTTP.DoRequest(
			"GET",
			endpoint,
			nil,
			&targetObjectAccumulator,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to fetch page %v", err)
		}

		if resp != nil {
			defer resp.Body.Close()
		}

		outData = append(outData, targetObjectAccumulator.Results...)

		if len(outData) >= targetObjectAccumulator.Size ||
			len(targetObjectAccumulator.Results) == 0 {
			break
		}

		page++
		params.Del("page")
		params.Add("page", strconv.Itoa(page))
	}

	outStruct.Size = targetObjectAccumulator.Size
	outStruct.Results = outData

	return &outStruct, nil

}
