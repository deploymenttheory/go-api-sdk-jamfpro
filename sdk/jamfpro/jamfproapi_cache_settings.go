// jamfapi_cache_settings.go
// Jamf Pro Api - Cache Settings
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-cache-settings
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"encoding/json"
	"fmt"
)

const uriCacheSettings = "/api/v1/cache-settings"

// Resource

// ResponseCacheSettings represents the JSON response for cache settings.
type ResourceCacheSettings struct {
	ID                         string                                  `json:"id,omitempty"`
	Name                       string                                  `json:"name,omitempty"`
	CacheType                  string                                  `json:"cacheType"`
	TimeToLiveSeconds          int                                     `json:"timeToLiveSeconds"`
	TimeToIdleSeconds          int                                     `json:"timeToIdleSeconds"`
	DirectoryTimeToLiveSeconds int                                     `json:"directoryTimeToLiveSeconds,omitempty"`
	EhcacheMaxBytesLocalHeap   string                                  `json:"ehcacheMaxBytesLocalHeap,omitempty"`
	CacheUniqueId              string                                  `json:"cacheUniqueId"`
	Elasticache                bool                                    `json:"elasticache,omitempty"`
	MemcachedEndpoints         []CacheSettingsSubsetMemcachedEndpoints `json:"memcachedEndpoints"`
}

// Subsets

type CacheSettingsSubsetMemcachedEndpoints struct {
	ID                      string `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	HostName                string `json:"hostName,omitempty"`
	Port                    int    `json:"port,omitempty"`
	Enabled                 bool   `json:"enabled,omitempty"`
	JSSCacheConfigurationID int    `json:"jssCacheConfigurationId,omitempty"`
}

// GetCacheSettings gets the current cache settings.
func (c *Client) GetCacheSettings() (*ResourceCacheSettings, error) {
	endpoint := uriCacheSettings

	var cacheSettings ResourceCacheSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cacheSettings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "cache settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cacheSettings, nil
}

// UpdateCacheSettings updates the cache settings.
func (c *Client) UpdateCacheSettings(cacheSettingsUpdate *ResourceCacheSettings) (*ResourceCacheSettings, error) {
	endpoint := uriCacheSettings

	requestBody, err := json.Marshal(cacheSettingsUpdate)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedJsonMarshal, "cache settings", err)
	}

	var updatedSettings ResourceCacheSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedSettings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "cache settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSettings, nil
}

// Please note that the Create and Delete functions are not implemented here as cache settings typically only allow retrieval and update operations.
