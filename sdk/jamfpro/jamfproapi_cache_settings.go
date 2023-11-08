// classicapi_cache_settings.go
// Jamf Pro Classic Api - Cache Settings
// api reference: Your provided URL (adapt as needed)
// Classic API requires the structs to support a JSON data structure.

package jamfpro

import (
	"encoding/json"
	"fmt"
)

const uriCacheSettings = "/api/v1/cache-settings"

// ResponseCacheSettings represents the JSON response for cache settings.
type ResponseCacheSettings struct {
	ID                         string              `json:"id,omitempty"`
	Name                       string              `json:"name,omitempty"`
	CacheType                  string              `json:"cacheType"`
	TimeToLiveSeconds          int                 `json:"timeToLiveSeconds"`
	TimeToIdleSeconds          int                 `json:"timeToIdleSeconds"`
	DirectoryTimeToLiveSeconds int                 `json:"directoryTimeToLiveSeconds,omitempty"`
	EhcacheMaxBytesLocalHeap   string              `json:"ehcacheMaxBytesLocalHeap,omitempty"`
	CacheUniqueId              string              `json:"cacheUniqueId"`
	Elasticache                bool                `json:"elasticache,omitempty"`
	MemcachedEndpoints         []MemcachedEndpoint `json:"memcachedEndpoints"`
}

// MemcachedEndpoint represents an individual memcached endpoint in the cache settings.
type MemcachedEndpoint struct {
	ID                      string `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	HostName                string `json:"hostName,omitempty"`
	Port                    int    `json:"port,omitempty"`
	Enabled                 bool   `json:"enabled,omitempty"`
	JSSCacheConfigurationID int    `json:"jssCacheConfigurationId,omitempty"`
}

// GetCacheSettings gets the current cache settings.
func (c *Client) GetCacheSettings() (*ResponseCacheSettings, error) {
	endpoint := uriCacheSettings

	var cacheSettings ResponseCacheSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cacheSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cache settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cacheSettings, nil
}

// UpdateCacheSettings updates the cache settings.
func (c *Client) UpdateCacheSettings(settings *ResponseCacheSettings) (*ResponseCacheSettings, error) {
	endpoint := uriCacheSettings

	requestBody, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal cache settings: %v", err)
	}

	var updatedSettings ResponseCacheSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &updatedSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to update cache settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSettings, nil
}

// Please note that the Create and Delete functions are not implemented here as cache settings typically only allow retrieval and update operations.
