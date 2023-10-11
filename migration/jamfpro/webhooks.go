// webhooks.go
// Jamf Pro Classic Api
// API requires the structs to support JSON and xml.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIWebhooks = "/JSSResource/webhooks"

type ResponseWebhook struct {
	ID                                int            `json:"id" xml:"id"`
	Name                              string         `json:"name" xml:"name"`
	Enabled                           bool           `json:"enabled" xml:"enabled"`
	URL                               string         `json:"url" xml:"url"`
	ContentType                       string         `json:"content_type" xml:"content_type"`
	Event                             string         `json:"event" xml:"event"`
	ConnectionTimeout                 int            `json:"connection_timeout" xml:"connection_timeout"`
	ReadTimeout                       int            `json:"read_timeout" xml:"read_timeout"`
	AuthenticationType                string         `json:"authentication_type" xml:"authentication_type"`
	Username                          string         `json:"username" xml:"username"`
	Password                          string         `json:"password" xml:"password"`
	EnableDisplayFieldsForGroupObject bool           `json:"enable_display_fields_for_group_object" xml:"enable_display_fields_for_group_object"`
	DisplayFields                     []DisplayField `json:"display_fields" xml:"display_fields"`
	SmartGroupID                      int            `json:"smart_group_id" xml:"smart_group_id"`
}

type DisplayField struct {
	Size int    `json:"size" xml:"size"`
	Name string `json:"name" xml:"name"`
}

type ResponseWebhookList struct {
	Size    int           `json:"size" xml:"size"`
	Webhook []WebhookItem `json:"webhook" xml:"webhook"`
}

type WebhookItem struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// GetWebhookByID gets a webhook by its id
func (c *Client) GetWebhookByID(id int) (*ResponseWebhook, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIWebhooks, id)

	var webhook ResponseWebhook
	if err := c.DoRequest("GET", url, nil, nil, &webhook); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &webhook, nil
}

// GetWebhookByName gets a webhook by its name
func (c *Client) GetWebhookByName(name string) (*ResponseWebhook, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIWebhooks, name)

	var webhook ResponseWebhook
	if err := c.DoRequest("GET", url, nil, nil, &webhook); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &webhook, nil
}

// GetWebhooks gets a list of all webhooks
func (c *Client) GetWebhooks() ([]ResponseWebhookList, error) {
	url := uriAPIWebhooks

	var webhookList []ResponseWebhookList
	if err := c.DoRequest("GET", url, nil, nil, &webhookList); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return webhookList, nil
}

// CreateWebhook creates a new Jamf Pro Webhook.
func (c *Client) CreateWebhook(webhook *ResponseWebhook) (*ResponseWebhook, error) {
	url := fmt.Sprintf("%s/id/0", uriAPIWebhooks)

	reqBody := &struct {
		XMLName xml.Name `xml:"webhook"`
		*ResponseWebhook
	}{
		ResponseWebhook: webhook,
	}

	var responseWebhook ResponseWebhook
	if err := c.DoRequest("POST", url, reqBody, nil, &responseWebhook); err != nil {
		return nil, fmt.Errorf("failed to create webhook: %v", err)
	}

	return &responseWebhook, nil
}

// UpdateWebhookByName updates an existing Jamf Pro Webhook by Name.
func (c *Client) UpdateWebhookByName(name string, webhook *ResponseWebhook) (*ResponseWebhook, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIWebhooks, name)

	reqBody := &struct {
		XMLName xml.Name `xml:"webhook"`
		*ResponseWebhook
	}{
		ResponseWebhook: webhook,
	}

	var responseWebhook ResponseWebhook
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseWebhook); err != nil {
		return nil, fmt.Errorf("failed to update webhook by Name: %v", err)
	}

	return &responseWebhook, nil
}

// UpdateWebhookByID updates an existing Jamf Pro Webhook by ID.
func (c *Client) UpdateWebhookByID(id int, webhook *ResponseWebhook) (*ResponseWebhook, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIWebhooks, id)

	reqBody := &struct {
		XMLName xml.Name `xml:"webhook"`
		*ResponseWebhook
	}{
		ResponseWebhook: webhook,
	}

	var responseWebhook ResponseWebhook
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseWebhook); err != nil {
		return nil, fmt.Errorf("failed to update webhook by ID: %v", err)
	}

	return &responseWebhook, nil
}

// DeleteWebhookByID deletes an existing Jamf Pro Webhook by ID.
func (c *Client) DeleteWebhookByID(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIWebhooks, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete webhook by ID: %v", err)
	}

	return nil
}

// DeleteWebhookByName deletes an existing Jamf Pro Webhook by Name.
func (c *Client) DeleteWebhookByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIWebhooks, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete webhook by Name: %v", err)
	}

	return nil
}
