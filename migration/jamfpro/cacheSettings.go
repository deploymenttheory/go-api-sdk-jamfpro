package jamfpro

const uriCacheSettings = "/api/v1/cache-settings"

type CacheSettings struct {
	Id                         string                                     `json:"id"`
	Name                       string                                     `json:"name"`
	CacheType                  string                                     `json:"cacheType"`
	TimeToLiveSeconds          int32                                      `json:"timeToLiveSeconds"`
	TimeToIdleSeconds          int32                                      `json:"timeToIdleSeconds"`
	DirectoryTimeToLiveSeconds int32                                      `json:"directoryTimeToLiveSeconds"`
	EhcacheMaxBytesLocalHeap   string                                     `json:"ehcacheMaxBytesLocalHeap"`
	CacheUniqueId              string                                     `json:"cacheUniqueId"`
	Elasticache                bool                                       `json:"elasticache"`
	MemcachedEndpoints         []CacheSettingsDataSubsetMemcachedEndpoint `json:"memcachedEndpoints"`
}

type CacheSettingsDataSubsetMemcachedEndpoint struct {
	Id                      string `json:"id"`
	Name                    string `json:"name"`
	HostName                string `json:"hostName"`
	Port                    int    `json:"port"`
	Enabled                 bool   `json:"enabled"`
	JssCacheConfigurationId int    `json:"jssCacheConfigurationId"`
}

type ResponseCacheSettings struct {
	TotalCount *int            `json:"totalCount,omitempty"`
	Results    []CacheSettings `json:"results,omitempty"`
}

func (c *Client) GetCacheSettings() (*CacheSettings, error) {
	var out CacheSettings

	err := c.DoRequest("GET", uriCacheSettings, nil, nil, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *Client) UpdateCacheSettings(settings *CacheSettings) (*CacheSettings, error) {
	var updatedSettings CacheSettings

	err := c.DoRequest("PUT", uriCacheSettings, settings, nil, &updatedSettings)
	if err != nil {
		return nil, err
	}

	return &updatedSettings, nil
}
