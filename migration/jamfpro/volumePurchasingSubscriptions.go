// volumePurchasingSubscriptions.go
// Jamf Pro Api

package jamfpro

import (
	"fmt"
)

const uriVolumePurchasingSubscriptions = "/api/v1/volume-purchasing-subscriptions"

type ResponseVolumePurchasingSubscription struct {
	TotalCount int                            `json:"totalCount"`
	Results    []VolumePurchasingSubscription `json:"results"`
}

type VolumePurchasingSubscription struct {
	Name               string              `json:"name"` // Required
	Enabled            bool                `json:"enabled,omitempty"`
	Triggers           []string            `json:"triggers,omitempty"`
	LocationIds        []string            `json:"locationIds,omitempty"`
	InternalRecipients []InternalRecipient `json:"internalRecipients,omitempty"`
	ExternalRecipients []ExternalRecipient `json:"externalRecipients,omitempty"`
	SiteId             string              `json:"siteId,omitempty"`
	ID                 string              `json:"id,omitempty"`
}

type InternalRecipient struct {
	AccountId string `json:"accountId"` // Required
	Frequency string `json:"frequency,omitempty"`
}

type ExternalRecipient struct {
	Name  string `json:"name"`  // Required
	Email string `json:"email"` // Required
}

// GetVolumePurchasingSubscriptionByID retrieves the Volume Purchasing Subscription by its ID
func (c *Client) GetVolumePurchasingSubscriptionByID(id int) (*ResponseVolumePurchasingSubscription, error) {
	url := fmt.Sprintf("%s/%d", uriVolumePurchasingSubscriptions, id)

	var subscription ResponseVolumePurchasingSubscription
	if err := c.DoRequest("GET", url, nil, nil, &subscription); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &subscription, nil
}

// GetVolumePurchasingSubscriptions retrieves all Volume Purchasing Subscriptions
func (c *Client) GetVolumePurchasingSubscriptions() (*ResponseVolumePurchasingSubscription, error) {
	url := uriVolumePurchasingSubscriptions

	var subscriptions ResponseVolumePurchasingSubscription
	if err := c.DoRequest("GET", url, nil, nil, &subscriptions); err != nil {
		return nil, fmt.Errorf("failed to fetch all Volume Purchasing Subscriptions: %v", err)
	}

	return &subscriptions, nil
}

// CreateVolumePurchasingSubscription creates a new Volume Purchasing Subscription
func (c *Client) CreateVolumePurchasingSubscription(subscription *VolumePurchasingSubscription) (*ResponseVolumePurchasingSubscription, error) {
	url := uriVolumePurchasingSubscriptions

	// Construct the request body
	reqBody := &struct {
		*VolumePurchasingSubscription
	}{
		VolumePurchasingSubscription: subscription,
	}

	// Execute the request
	var responseSubscription ResponseVolumePurchasingSubscription
	if err := c.DoRequestDebug("POST", url, reqBody, nil, &responseSubscription); err != nil {
		return nil, fmt.Errorf("failed to create Volume Purchasing Subscription: %v", err)
	}

	return &responseSubscription, nil
}
