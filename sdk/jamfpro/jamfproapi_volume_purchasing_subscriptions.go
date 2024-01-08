// jamfproapi_volume_purchasing_subscriptions.go
// Jamf Pro Api - Volume Purchasing Subscriptions
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
// Jamf Pro API requires the structs to support an JSON data structure.

// TODO Refactor this - pagination etc

package jamfpro

import (
	"fmt"
	"strconv"
)

const uriVolumePurchasingSubscriptions = "/api/v1/volume-purchasing-subscriptions"

// List

type ResponseVolumePurchasingSubscriptionsList struct {
	TotalCount *int                                   `json:"totalCount,omitempty"`
	Results    []ResourceVolumePurchasingSubscription `json:"results,omitempty"`
}

// Resource

type ResourceVolumePurchasingSubscription struct {
	Id                 string                                                 `json:"id,omitempty"`
	Name               string                                                 `json:"name"`
	Enabled            bool                                                   `json:"enabled,omitempty"`
	Triggers           []string                                               `json:"triggers,omitempty"`
	LocationIds        []string                                               `json:"locationIds,omitempty"`
	InternalRecipients []VolumePurchasingSubscriptionSubsetInternalRecipients `json:"internalRecipients,omitempty"`
	ExternalRecipients []VolumePurchasingSubscriptionSubsetExternalRecipients `json:"externalRecipients,omitempty"`
	SiteId             string                                                 `json:"siteId,omitempty"`
}

// Subsets

type VolumePurchasingSubscriptionSubsetInternalRecipients struct {
	AccountId string `json:"accountId,omitempty"`
	Frequency string `json:"frequency,omitempty"`
}

type VolumePurchasingSubscriptionSubsetExternalRecipients struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// CRUD

// GetVolumePurchasingSubscriptions retrieves all volume purchasing subscriptions
func (c *Client) GetVolumePurchasingSubscriptions() (*ResponseVolumePurchasingSubscriptionsList, error) {
	var subscriptionsList ResponseVolumePurchasingSubscriptionsList
	resp, err := c.HTTP.DoRequest("GET", uriVolumePurchasingSubscriptions, nil, &subscriptionsList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "volume purchasing subscriptions", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &subscriptionsList, nil
}

// GetVolumePurchasingSubscriptionByID retrieves a single volume purchasing subscription by its ID
func (c *Client) GetVolumePurchasingSubscriptionByID(id string) (*ResourceVolumePurchasingSubscription, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf("%s/%s", uriVolumePurchasingSubscriptions, id)

	var subscription ResourceVolumePurchasingSubscription
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &subscription)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "volume purchasing subscription", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &subscription, nil
}

// GetVolumePurchasingSubscriptionByNameByID fetches a volume purchasing subscription by its display name and retrieves its details using its ID.
func (c *Client) GetVolumePurchasingSubscriptionByNameByID(name string) (*ResourceVolumePurchasingSubscription, error) {
	subscriptionsList, err := c.GetVolumePurchasingSubscriptions()
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "volume purchasing subscription", name, err)
	}

	// Search for the subscription with the given name
	for _, subscription := range subscriptionsList.Results {
		if subscription.Name == name {
			// Assuming the ID in the struct is a string and needs to be converted to int
			subscriptionID, convErr := strconv.Atoi(subscription.Id)
			if convErr != nil {
				return nil, fmt.Errorf("failed to convert subscription ID '%s' to integer: %v", subscription.Id, convErr)
			}
			return c.GetVolumePurchasingSubscriptionByID(strconv.Itoa(subscriptionID))
		}
	}

	return nil, fmt.Errorf("no volume purchasing subscription found with the name '%s'", name)
}

// CreateVolumePurchasingSubscription creates a new volume purchasing subscription
func (c *Client) CreateVolumePurchasingSubscription(subscription *ResourceVolumePurchasingSubscription) (*ResourceVolumePurchasingSubscription, error) {
	endpoint := uriVolumePurchasingSubscriptions

	// Default the SiteId to "-1" if not provided
	if subscription.SiteId == "" {
		subscription.SiteId = "-1"
	}

	var createdSubscription ResourceVolumePurchasingSubscription
	resp, err := c.HTTP.DoRequest("POST", endpoint, subscription, &createdSubscription)
	if err != nil {
		return nil, fmt.Errorf("failed to create volume purchasing subscription: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdSubscription, nil
}

// UpdateVolumePurchasingSubscriptionByID updates a volume purchasing subscription by its ID
func (c *Client) UpdateVolumePurchasingSubscriptionByID(id string, subscription *ResourceVolumePurchasingSubscription) (*ResourceVolumePurchasingSubscription, error) {
	endpoint := fmt.Sprintf("%s/%s", uriVolumePurchasingSubscriptions, id)

	var updatedSubscription ResourceVolumePurchasingSubscription
	resp, err := c.HTTP.DoRequest("PUT", endpoint, subscription, &updatedSubscription)
	if err != nil {
		return nil, fmt.Errorf("failed to update volume purchasing subscription with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSubscription, nil
}

// UpdateVolumePurchasingSubscriptionByNameByID updates a volume purchasing subscription by its display name
func (c *Client) UpdateVolumePurchasingSubscriptionByNameByID(name string, updateData *ResourceVolumePurchasingSubscription) (*ResourceVolumePurchasingSubscription, error) {
	subscriptionsList, err := c.GetVolumePurchasingSubscriptions()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all volume purchasing subscriptions: %v", err)
	}

	for _, subscription := range subscriptionsList.Results {
		if subscription.Name == name {
			return c.UpdateVolumePurchasingSubscriptionByID(subscription.Id, updateData)
		}
	}

	return nil, fmt.Errorf("no volume purchasing subscription found with the name '%s'", name)
}

// DeleteVolumePurchasingSubscriptionByID deletes a volume purchasing subscription by its ID
func (c *Client) DeleteVolumePurchasingSubscriptionByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriVolumePurchasingSubscriptions, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete volume purchasing subscription with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteVolumePurchasingSubscriptionByName finds a subscription by name and deletes it by its ID
func (c *Client) DeleteVolumePurchasingSubscriptionByName(name string) error {
	subscriptionsList, err := c.GetVolumePurchasingSubscriptions()
	if err != nil {
		return fmt.Errorf("failed to fetch all volume purchasing subscriptions: %v", err)
	}

	for _, subscription := range subscriptionsList.Results {
		if subscription.Name == name {
			return c.DeleteVolumePurchasingSubscriptionByID(subscription.Id)
		}
	}

	return fmt.Errorf("no volume purchasing subscription found with the name '%s'", name)
}
