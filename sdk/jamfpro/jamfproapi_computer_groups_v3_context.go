package jamfpro

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var errGroupResponseMissing = errors.New("computer group request returned no HTTP response")

var (
	ErrSmartComputerGroupNotFound      = errors.New("smart computer group not found")
	ErrSmartComputerGroupNameAmbiguous = errors.New("multiple smart computer groups have the requested name")
	ErrComputerGroupPlatformIDNotFound = errors.New("Platform identifier not found for computer group")
)

type groupRequest struct {
	method   string
	endpoint string
	body     any
	out      any
}

func (c *Client) doGroupRequest(ctx context.Context, request groupRequest) error {
	resp, err := c.HTTP.DoRequestNoRetryWithContext(
		ctx, request.method, request.endpoint, request.body, request.out,
	)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("%s %s: %w", request.method, request.endpoint, err)
	}
	if resp == nil {
		return errGroupResponseMissing
	}
	return nil
}

// GetSmartComputerGroupsV3WithContext retrieves one page of smart computer groups.
func (c *Client) GetSmartComputerGroupsV3WithContext(
	ctx context.Context, params url.Values,
) (*ResponseSmartComputerGroupsListV3, error) {
	var result ResponseSmartComputerGroupsListV3
	endpoint := uriAPIV3ComputerGroups + "/smart-groups"
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}
	if err := c.doGroupRequest(ctx, groupRequest{method: http.MethodGet, endpoint: endpoint, out: &result}); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSmartComputerGroupByIDV3WithContext retrieves a smart computer group by numeric Jamf Pro ID.
func (c *Client) GetSmartComputerGroupByIDV3WithContext(
	ctx context.Context, id string,
) (*ResourceSmartComputerGroupV3, error) {
	var result ResourceSmartComputerGroupV3
	if err := c.doGroupRequest(ctx, groupRequest{
		method: http.MethodGet, endpoint: uriAPIV3ComputerGroups + "/smart-groups/" + url.PathEscape(id), out: &result,
	}); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateSmartComputerGroupV3WithContext creates a smart computer group without transport replay.
func (c *Client) CreateSmartComputerGroupV3WithContext(
	ctx context.Context, input ResourceSmartComputerGroupV3,
) (*ResponseSmartComputerGroupCreateV3, error) {
	var result ResponseSmartComputerGroupCreateV3
	if err := c.doGroupRequest(ctx, groupRequest{
		method: http.MethodPost, endpoint: uriAPIV3ComputerGroups + "/smart-groups", body: input, out: &result,
	}); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateSmartComputerGroupByIDV3WithContext replaces a smart group's definition without transport replay.
func (c *Client) UpdateSmartComputerGroupByIDV3WithContext(
	ctx context.Context, id string, input ResourceSmartComputerGroupV3,
) (*ResourceSmartComputerGroupV3, error) {
	var result ResourceSmartComputerGroupV3
	if err := c.doGroupRequest(ctx, groupRequest{
		method: http.MethodPut, endpoint: uriAPIV3ComputerGroups + "/smart-groups/" + url.PathEscape(id), body: input, out: &result,
	}); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteSmartComputerGroupByIDV3WithContext deletes a smart group within the caller's deadline.
func (c *Client) DeleteSmartComputerGroupByIDV3WithContext(ctx context.Context, id string) error {
	return c.doGroupRequest(ctx, groupRequest{
		method: http.MethodDelete, endpoint: uriAPIV3ComputerGroups + "/smart-groups/" + url.PathEscape(id),
	})
}

// GetGroupsV2WithContext retrieves one page of unified groups, including their Platform and Jamf Pro IDs.
func (c *Client) GetGroupsV2WithContext(
	ctx context.Context, params url.Values,
) (*ResponseGroupSearchResultV2, error) {
	var result ResponseGroupSearchResultV2
	endpoint := uriGroupsV2
	if len(params) > 0 {
		endpoint += "?" + params.Encode()
	}
	if err := c.doGroupRequest(ctx, groupRequest{method: http.MethodGet, endpoint: endpoint, out: &result}); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSmartComputerGroupByNameV3WithContext resolves a case-insensitive name across every result page.
func (c *Client) GetSmartComputerGroupByNameV3WithContext(
	ctx context.Context, name string,
) (*ResponseSmartComputerGroupListItemV3, error) {
	var found *ResponseSmartComputerGroupListItemV3
	for page := 0; ; page++ {
		result, err := c.GetSmartComputerGroupsV3WithContext(ctx, url.Values{
			"page": {strconv.Itoa(page)}, "page-size": {"100"},
		})
		if err != nil {
			return nil, err
		}
		for _, group := range result.Results {
			if strings.EqualFold(group.Name, name) {
				if found != nil {
					return nil, fmt.Errorf("%w: %q", ErrSmartComputerGroupNameAmbiguous, name)
				}
				copy := group
				found = &copy
			}
		}
		if len(result.Results) == 0 || (page+1)*100 >= result.TotalCount {
			break
		}
	}
	if found == nil {
		return nil, fmt.Errorf("%w: %q", ErrSmartComputerGroupNotFound, name)
	}
	return found, nil
}

// GetComputerSmartGroupPlatformIDByJamfProIDWithContext bridges the two identifier spaces.
func (c *Client) GetComputerSmartGroupPlatformIDByJamfProIDWithContext(
	ctx context.Context, id string,
) (string, error) {
	for page := 0; ; page++ {
		result, err := c.GetGroupsV2WithContext(ctx, url.Values{
			"filter": {`groupType=="COMPUTER" and isSmart=="true"`},
			"page":   {strconv.Itoa(page)}, "page-size": {"100"},
		})
		if err != nil {
			return "", err
		}
		for _, group := range result.Results {
			if group.GroupJamfProId == id && group.GroupPlatformId != "" && group.GroupType == "COMPUTER" {
				return group.GroupPlatformId, nil
			}
		}
		if len(result.Results) == 0 || (page+1)*100 >= result.TotalCount {
			return "", fmt.Errorf("%w: %s", ErrComputerGroupPlatformIDNotFound, id)
		}
	}
}

// UpdateGroupDescriptionByIDV2WithContext patches only the description using the Platform ID.
// An empty description is sent explicitly so the existing description can be cleared.
func (c *Client) UpdateGroupDescriptionByIDV2WithContext(ctx context.Context, platformID, description string) error {
	input := struct {
		GroupDescription string `json:"groupDescription"`
	}{GroupDescription: description}
	return c.doGroupRequest(ctx, groupRequest{
		method: http.MethodPatch, endpoint: uriGroupsV2 + "/" + url.PathEscape(platformID), body: input,
	})
}

// CreateComputerGroupWithContext sends a complete Classic group definition in a single request.
func (c *Client) CreateComputerGroupWithContext(
	ctx context.Context, group *ResourceComputerGroup,
) (*ResponseComputerGroupreatedAndUpdated, error) {
	input := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResourceComputerGroup
	}{ResourceComputerGroup: group}
	var result ResponseComputerGroupreatedAndUpdated
	if err := c.doGroupRequest(ctx, groupRequest{
		method: http.MethodPost, endpoint: uriComputerGroups + "/id/0", body: &input, out: &result,
	}); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateComputerGroupByIDWithContext sends name, site, and all criteria atomically through Classic.
func (c *Client) UpdateComputerGroupByIDWithContext(
	ctx context.Context, id string, group *ResourceComputerGroup,
) (*ResponseComputerGroupreatedAndUpdated, error) {
	input := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResourceComputerGroup
	}{ResourceComputerGroup: group}
	var result ResponseComputerGroupreatedAndUpdated
	if err := c.doGroupRequest(ctx, groupRequest{
		method: http.MethodPut, endpoint: uriComputerGroups + "/id/" + url.PathEscape(id), body: &input, out: &result,
	}); err != nil {
		return nil, err
	}
	return &result, nil
}
