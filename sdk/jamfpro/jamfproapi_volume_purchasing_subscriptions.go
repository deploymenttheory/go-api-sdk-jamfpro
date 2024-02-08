// jamfproapi_volume_purchasing_subscriptions.go
// Jamf Pro Api - Volume Purchasing Subscriptions
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
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
func (c *Client) GetVolumePurchasingSubscriptions(sort_filter string) (*ResponseVolumePurchasingSubscriptionsList, error) {
	resp, err := c.DoPaginatedGet(
		uriVolumePurchasingSubscriptions,
		maxPageSize,
		startingPageNumber,
		sort_filter,
	)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "volume purchasing subscriptions", err)
	}

	var out ResponseVolumePurchasingSubscriptionsList
	out.TotalCount = &resp.Size

	for _, value := range resp.Results {
		var newObj ResourceVolumePurchasingSubscription
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "volume purchasing subscription", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetVolumePurchasingSubscriptionByID retrieves a single volume purchasing subscription by its ID
func (c *Client) GetVolumePurchasingSubscriptionByID(id string) (*ResourceVolumePurchasingSubscription, error) {
	endpoint := fmt.Sprintf("%s/%s", uriVolumePurchasingSubscriptions, id)

	var subscription ResourceVolumePurchasingSubscription
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &subscription, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "volume purchasing subscription", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &subscription, nil
}

// GetVolumePurchasingSubscriptionByNameByID fetches a volume purchasing subscription by its display name and retrieves its details using its ID.
func (c *Client) GetVolumePurchasingSubscriptionByName(name string) (*ResourceVolumePurchasingSubscription, error) {
	vpSubcriptions, err := c.GetVolumePurchasingSubscriptions("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "volume purchasing subscriptions", err)
	}

	for _, value := range vpSubcriptions.Results {
		if value.Name == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "volume purchasing subscription", name, errMsgNoName)

}

// CreateVolumePurchasingSubscription creates a new volume purchasing subscription
func (c *Client) CreateVolumePurchasingSubscription(subscription *ResourceVolumePurchasingSubscription) (*ResourceVolumePurchasingSubscription, error) {
	endpoint := uriVolumePurchasingSubscriptions

	// Default the SiteId to "-1" if not provided
	if subscription.SiteId == "" {
		subscription.SiteId = "-1"
	}

	var createdSubscription ResourceVolumePurchasingSubscription
	resp, err := c.HTTP.DoRequest("POST", endpoint, subscription, &createdSubscription, c.HTTP.Logger)
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
	resp, err := c.HTTP.DoRequest("PUT", endpoint, subscription, &updatedSubscription, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to update volume purchasing subscription with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSubscription, nil
}

// UpdateVolumePurchasingSubscriptionByNameByID updates a volume purchasing subscription by its display name
func (c *Client) UpdateVolumePurchasingSubscriptionByName(name string, updateData *ResourceVolumePurchasingSubscription) (*ResourceVolumePurchasingSubscription, error) {
	target, err := c.GetVolumePurchasingSubscriptionByName(name)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "volume purchasing subscription", name, err)
	}

	target_id := target.Id
	resp, err := c.UpdateVolumePurchasingSubscriptionByID(target_id, updateData)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "volume purchasing subscription", name, err)
	}
	return resp, nil
}

// DeleteVolumePurchasingSubscriptionByID deletes a volume purchasing subscription by its ID
func (c *Client) DeleteVolumePurchasingSubscriptionByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriVolumePurchasingSubscriptions, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
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
	target, err := c.GetVolumePurchasingSubscriptionByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "volume purchasing subscription", name, err)
	}

	target_id := target.Id
	err = c.DeleteVolumePurchasingSubscriptionByID(target_id)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "volume purchasing subscription", name, err)
	}

	return nil

}
