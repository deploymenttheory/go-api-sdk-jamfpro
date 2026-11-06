package jamfpro

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GetSmartComputerGroupsV2WithContext retrieves one page within the caller's deadline.
func (c *Client) GetSmartComputerGroupsV2WithContext(ctx context.Context, params url.Values) (*ResponseSmartComputerGroupsListV2, error) {
	endpoint := uriAPIV2ComputerGroups + "/smart-groups"
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}
	var result ResponseSmartComputerGroupsListV2
	if err := c.doGroupRequest(ctx, groupRequest{method: http.MethodGet, endpoint: endpoint, out: &result}); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSmartComputerGroupByIDV2WithContext retrieves a group and preserves typed API errors.
func (c *Client) GetSmartComputerGroupByIDV2WithContext(ctx context.Context, id string) (*ResourceSmartComputerGroupV2, error) {
	var result ResourceSmartComputerGroupV2
	if err := c.doGroupRequest(ctx, groupRequest{
		method: http.MethodGet, endpoint: uriAPIV2ComputerGroups + "/smart-groups/" + url.PathEscape(id), out: &result,
	}); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSmartComputerGroupByNameV2WithContext finds a case-insensitive exact name across every page.
func (c *Client) GetSmartComputerGroupByNameV2WithContext(ctx context.Context, name string) (*ResponseSmartComputerGroupListItemV2, error) {
	const pageSize = 100
	var match *ResponseSmartComputerGroupListItemV2
	for page := 0; ; page++ {
		params := url.Values{"page": {strconv.Itoa(page)}, "page-size": {strconv.Itoa(pageSize)}, "sort": {"id:asc"}}
		groups, err := c.GetSmartComputerGroupsV2WithContext(ctx, params)
		if err != nil {
			return nil, err
		}
		for _, group := range groups.Results {
			if !strings.EqualFold(group.Name, name) {
				continue
			}
			if match != nil && match.ID != group.ID {
				return nil, fmt.Errorf("%w: %q", ErrSmartComputerGroupNameAmbiguous, name)
			}
			candidate := group
			match = &candidate
		}
		if (page+1)*pageSize >= groups.TotalCount || len(groups.Results) == 0 {
			break
		}
	}
	if match == nil {
		return nil, fmt.Errorf("%w: %q", ErrSmartComputerGroupNotFound, name)
	}
	return match, nil
}

// CreateSmartComputerGroupV2WithContext creates a group without replaying the write.
func (c *Client) CreateSmartComputerGroupV2WithContext(ctx context.Context, input ResourceSmartComputerGroupV2) (*ResponseSmartComputerGroupCreateV2, error) {
	var result ResponseSmartComputerGroupCreateV2
	if err := c.doGroupRequest(ctx, groupRequest{
		method: http.MethodPost, endpoint: uriAPIV2ComputerGroups + "/smart-groups", body: input, out: &result,
	}); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateSmartComputerGroupByIDV2WithContext updates a group without replaying the write.
func (c *Client) UpdateSmartComputerGroupByIDV2WithContext(ctx context.Context, id string, input ResourceSmartComputerGroupV2) (*ResourceSmartComputerGroupV2, error) {
	var result ResourceSmartComputerGroupV2
	if err := c.doGroupRequest(ctx, groupRequest{
		method: http.MethodPut, endpoint: uriAPIV2ComputerGroups + "/smart-groups/" + url.PathEscape(id), body: input, out: &result,
	}); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteSmartComputerGroupByIDV2WithContext deletes a group within the caller's deadline.
func (c *Client) DeleteSmartComputerGroupByIDV2WithContext(ctx context.Context, id string) error {
	return c.doGroupRequest(ctx, groupRequest{
		method: http.MethodDelete, endpoint: uriAPIV2ComputerGroups + "/smart-groups/" + url.PathEscape(id),
	})
}
