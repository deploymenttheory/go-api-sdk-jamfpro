// jamfproapi_apns_client_push_status.go
// Jamf Pro Api - APNS Client Push Status
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-apns-client-push-status
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
)

const uriApnsClientPushStatus = "/api/v1/apns-client-push-status"

// List

// ResponseApnsClientPushStatusList represents the search results containing APNS client push status records.
type ResponseApnsClientPushStatusList struct {
	TotalCount int                            `json:"totalCount"`
	Results    []ResourceApnsClientPushStatus `json:"results"`
}

// Resource

// ResourceApnsClientPushStatus represents information about a client with push notifications disabled.
type ResourceApnsClientPushStatus struct {
	DeviceType   string `json:"deviceType,omitempty"`
	ClientID     string `json:"clientId,omitempty"`
	DisabledAt   string `json:"disabledAt,omitempty"`
	ManagementID string `json:"managementId,omitempty"`
}

// ResponseApnsPushEnableRequest represents the status of an enable-all-clients push request.
type ResponseApnsPushEnableRequest struct {
	RequestedTime string `json:"requestedTime,omitempty"`
	Status        string `json:"status,omitempty"`
	ProcessedTime string `json:"processedTime,omitempty"`
}

// ResourceApnsEnablePushRequest is the request body to enable push notifications for a single client.
type ResourceApnsEnablePushRequest struct {
	ManagementID string `json:"managementId"`
}

// GetApnsClientPushStatus retrieves the list of clients with push notifications disabled.
func (c *Client) GetApnsClientPushStatus(params url.Values) (*ResponseApnsClientPushStatusList, error) {
	endpoint := uriApnsClientPushStatus
	if params != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var out ResponseApnsClientPushStatusList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "apns client push status", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetApnsClientPushEnableAllClientsStatus retrieves the status of the most recent enable-all-clients request.
func (c *Client) GetApnsClientPushEnableAllClientsStatus() (*ResponseApnsPushEnableRequest, error) {
	endpoint := fmt.Sprintf("%s/enable-all-clients/status", uriApnsClientPushStatus)

	var out ResponseApnsPushEnableRequest
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "apns client push enable-all-clients status", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// EnableApnsPushForAllClients queues a request to re-enable push notifications for all disabled clients.
func (c *Client) EnableApnsPushForAllClients() error {
	endpoint := fmt.Sprintf("%s/enable-all-clients", uriApnsClientPushStatus)

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedAction, "enable apns push for all clients", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// EnableApnsPushForClient re-enables push notifications for a single client by management id.
func (c *Client) EnableApnsPushForClient(request *ResourceApnsEnablePushRequest) error {
	endpoint := fmt.Sprintf("%s/enable-client", uriApnsClientPushStatus)

	resp, err := c.HTTP.DoRequest("POST", endpoint, request, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedAction, "enable apns push for client", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
