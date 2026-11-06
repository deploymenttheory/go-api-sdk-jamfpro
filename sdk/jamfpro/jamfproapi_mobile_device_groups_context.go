package jamfpro

import (
	"context"
	"net/http"
	"net/url"
)

// GetSmartMobileDeviceGroupByIDV1WithContext reads a group with typed API errors and caller cancellation.
func (c *Client) GetSmartMobileDeviceGroupByIDV1WithContext(ctx context.Context, id string) (*ResourceSmartMobileDeviceGroupV1, error) {
	var result ResourceSmartMobileDeviceGroupV1
	if err := c.doGroupRequest(ctx, groupRequest{
		method: http.MethodGet, endpoint: uriAPIV1MobileDeviceGroups + "/smart-groups/" + url.PathEscape(id), out: &result,
	}); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteSmartMobileDeviceGroupByIDV1WithContext deletes a group with typed API errors and caller cancellation.
func (c *Client) DeleteSmartMobileDeviceGroupByIDV1WithContext(ctx context.Context, id string) error {
	return c.doGroupRequest(ctx, groupRequest{
		method: http.MethodDelete, endpoint: uriAPIV1MobileDeviceGroups + "/smart-groups/" + url.PathEscape(id),
	})
}
