// classicapi_webhooks.go
// Jamf Pro Classic Api - Webhooks
// api reference: https://developer.jamf.com/jamf-pro/reference/webhooks
// Jamf Pro Classic Api requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedAdvancedSearchContainerDisplayField
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriWebhooks = "/JSSResource/webhooks"

// List

// Structs for Webhooks Response
type ResponseWebhooksList struct {
	Size     int                `xml:"size"`
	Webhooks []WebhooksListItem `xml:"webhook"`
}

type WebhooksListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// Struct for individual Webhook
type ResourceWebhook struct {
	ID                          int            `xml:"id"`
	Name                        string         `xml:"name"`
	Enabled                     bool           `xml:"enabled"`
	URL                         string         `xml:"url,omitempty"`
	ContentType                 string         `xml:"content_type,omitempty"`
	Event                       string         `xml:"event,omitempty"`
	ConnectionTimeout           int            `xml:"connection_timeout,omitempty"`
	ReadTimeout                 int            `xml:"read_timeout,omitempty"`
	AuthenticationType          string         `xml:"authentication_type,omitempty"`
	Username                    string         `xml:"username,omitempty"`
	Password                    string         `xml:"password,omitempty"`
	EnableDisplayFieldsForGroup bool           `xml:"enable_display_fields_for_group_object,omitempty"`
	DisplayFields               []DisplayField `xml:"display_fields>display_field,omitempty"`
	SmartGroupID                int            `xml:"smart_group_id,omitempty"`
}

// Subsets & Containers

// CRUD

// GetWebhooks retrieves a list of all webhooks.
func (c *Client) GetWebhooks() (*ResponseWebhooksList, error) {
	endpoint := uriWebhooks

	var response ResponseWebhooksList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "webhooks", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetWebhookByID retrieves a specific webhook by its ID.
func (c *Client) GetWebhookByID(id int) (*ResourceWebhook, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriWebhooks, id)

	var response ResourceWebhook
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "webhook", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetWebhookByName retrieves a specific webhook by its name.
func (c *Client) GetWebhookByName(name string) (*ResourceWebhook, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriWebhooks, name)

	var response ResourceWebhook
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "webhook", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// CreateWebhook creates a new webhook.
func (c *Client) CreateWebhook(webhook *ResourceWebhook) (*ResourceWebhook, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriWebhooks)

	requestBody := struct {
		XMLName xml.Name `xml:"webhook"`
		*ResourceWebhook
	}{
		ResourceWebhook: webhook,
	}

	var response ResourceWebhook
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "webhook", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateWebhookByID updates a specific webhook by its ID.
func (c *Client) UpdateWebhookByID(id int, webhook *ResourceWebhook) (*ResourceWebhook, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriWebhooks, id)

	requestBody := struct {
		XMLName xml.Name `xml:"webhook"`
		*ResourceWebhook
	}{
		ResourceWebhook: webhook,
	}

	var response ResourceWebhook
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "webhook", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateWebhookByName updates a specific webhook by its name.
func (c *Client) UpdateWebhookByName(name string, webhook *ResourceWebhook) (*ResourceWebhook, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriWebhooks, name)

	requestBody := struct {
		XMLName xml.Name `xml:"webhook"`
		*ResourceWebhook
	}{
		ResourceWebhook: webhook,
	}

	var response ResourceWebhook
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "webhook", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteWebhookByID deletes a specific webhook by its ID.
func (c *Client) DeleteWebhookByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriWebhooks, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "webhook", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteWebhookByName deletes a specific webhook by its name.
func (c *Client) DeleteWebhookByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriWebhooks, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "webhook", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
