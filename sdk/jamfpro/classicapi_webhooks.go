// classicapi_webhooks.go
// Jamf Pro Classic Api - Webhooks
// api reference: https://developer.jamf.com/jamf-pro/reference/webhooks
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriWebhooks = "/JSSResource/webhooks"

// Structs for Webhooks Response
type ResponseWebhooksList struct {
	Size     int               `xml:"size"`
	Webhooks []WebhookListItem `xml:"webhook"`
}

type WebhookListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Struct for individual Webhook
type ResponseWebhook struct {
	ID                          int                                     `xml:"id"`
	Name                        string                                  `xml:"name"`
	Enabled                     bool                                    `xml:"enabled"`
	URL                         string                                  `xml:"url"`
	ContentType                 string                                  `xml:"content_type"`
	Event                       string                                  `xml:"event"`
	ConnectionTimeout           int                                     `xml:"connection_timeout"`
	ReadTimeout                 int                                     `xml:"read_timeout"`
	AuthenticationType          string                                  `xml:"authentication_type"`
	Username                    string                                  `xml:"username"`
	Password                    string                                  `xml:"password"`
	EnableDisplayFieldsForGroup bool                                    `xml:"enable_display_fields_for_group_object"`
	DisplayFields               []ResponseWebhookDataSubsetDisplayField `xml:"display_fields>display_field"`
	SmartGroupID                int                                     `xml:"smart_group_id"`
}

type ResponseWebhookDataSubsetDisplayField struct {
	Size         int                                       `xml:"size"`
	DisplayField ResponseWebhookDataSubsetDisplayFieldItem `xml:"display_field"`
}

type ResponseWebhookDataSubsetDisplayFieldItem struct {
	Name string `xml:"name"`
}

// GetWebhooks retrieves a list of all webhooks.
func (c *Client) GetWebhooks() (*ResponseWebhooksList, error) {
	endpoint := uriWebhooks

	var response ResponseWebhooksList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch webhooks: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetWebhookByID retrieves a specific webhook by its ID.
func (c *Client) GetWebhookByID(id int) (*ResponseWebhook, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriWebhooks, id)

	var response ResponseWebhook
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch webhook by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetWebhookByName retrieves a specific webhook by its name.
func (c *Client) GetWebhookByName(name string) (*ResponseWebhook, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriWebhooks, name)

	var response ResponseWebhook
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch webhook by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// CreateWebhook creates a new webhook.
func (c *Client) CreateWebhook(webhook *ResponseWebhook) (*ResponseWebhook, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriWebhooks) // '0' indicates creation

	// Using an anonymous struct for the request body
	requestBody := struct {
		XMLName xml.Name `xml:"webhook"`
		*ResponseWebhook
	}{
		ResponseWebhook: webhook,
	}

	var response ResponseWebhook
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create webhook: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateWebhookByID updates a specific webhook by its ID.
func (c *Client) UpdateWebhookByID(id int, webhook *ResponseWebhook) (*ResponseWebhook, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriWebhooks, id)

	// Using an anonymous struct for the request body
	requestBody := struct {
		XMLName xml.Name `xml:"webhook"`
		*ResponseWebhook
	}{
		ResponseWebhook: webhook,
	}

	var response ResponseWebhook
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update webhook by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateWebhookByName updates a specific webhook by its name.
func (c *Client) UpdateWebhookByName(name string, webhook *ResponseWebhook) (*ResponseWebhook, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriWebhooks, name)

	// Using an anonymous struct for the request body
	requestBody := struct {
		XMLName xml.Name `xml:"webhook"`
		*ResponseWebhook
	}{
		ResponseWebhook: webhook,
	}

	var response ResponseWebhook
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update webhook by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteWebhookByID deletes a specific webhook by its ID.
func (c *Client) DeleteWebhookByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriWebhooks, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete webhook by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete webhook by ID, status code: %d", resp.StatusCode)
	}

	return nil
}

// DeleteWebhookByName deletes a specific webhook by its name.
func (c *Client) DeleteWebhookByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriWebhooks, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete webhook by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete webhook by name, status code: %d", resp.StatusCode)
	}

	return nil
}
