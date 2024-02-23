package jamfpro

import (
	"fmt"
)

const uriTeamViewerRemoteAdmin = "/api/preview/remote-administration-configurations/team-viewer"

type ResponseTeamViewerRemoteAdministrationConfiguration struct {
	TotalCount int                                           `json:"totalCount"`
	Results    []TeamViewerRemoteAdministrationConfiguration `json:"results"`
}

type ResponseCreateTeamViewerConfiguration struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

type ResponseTeamViewerRemoteAdminStatus struct {
	ConnectionVerificationResult string `json:"connectionVerificationResult"`
}

type ResponseTeamViewerSessionStatus struct {
	SessionState string `json:"sessionState"`
	Online       bool   `json:"online"`
}

type TeamViewerRemoteAdministrationConfiguration struct {
	ID             string `json:"id"`
	SiteID         string `json:"siteId"`
	DisplayName    string `json:"displayName"`
	Enabled        bool   `json:"enabled"`
	SessionTimeout int    `json:"sessionTimeout"`
}

type CreateTeamViewerConfiguration struct {
	Enabled        bool   `json:"enabled"`
	SiteID         string `json:"siteId"`
	DisplayName    string `json:"displayName"`
	ScriptToken    string `json:"scriptToken"`
	SessionTimeout int    `json:"sessionTimeout"`
}

type UpdateTeamViewerConfiguration struct {
	DisplayName    string `json:"displayName"`
	Enabled        bool   `json:"enabled"`
	SessionTimeout int    `json:"sessionTimeout"`
	Token          string `json:"token"`
}

type TeamViewerSession struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	Description   string `json:"description"`
	SupporterLink string `json:"supporterLink"`
	EndUserLink   string `json:"endUserLink"`
	DeviceID      string `json:"deviceId"`
	DeviceName    string `json:"deviceName"`
	DeviceType    string `json:"deviceType"`
	State         string `json:"state"`
	CreatorID     string `json:"creatorId"`
	CreatorName   string `json:"creatorName"`
	CreatedAt     string `json:"createdAt"`
}

type CreateTeamViewerSessionRequest struct {
	DeviceID    string `json:"deviceId"`
	DeviceType  string `json:"deviceType"`
	Description string `json:"description"`
}

type ResponseCreateTeamViewerSession struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

func (c *Client) GetTeamViewerRemoteAdministrationConfigurationIdByName(name string) (string, error) {
	var id string
	configurations, err := c.GetTeamViewerRemoteAdministrationConfigurations()
	if err != nil {
		return "", err
	}

	for _, v := range configurations.Results {
		if v.DisplayName == name {
			id = v.ID
			break
		}
	}
	return id, err
}

func (c *Client) GetTeamViewerRemoteAdministrationConfigurationByName(name string) (*TeamViewerRemoteAdministrationConfiguration, error) {
	allConfigurationsResponse, err := c.GetTeamViewerRemoteAdministrationConfigurations()
	if err != nil {
		return nil, err
	}

	for _, config := range allConfigurationsResponse.Results {
		if config.DisplayName == name {
			return &config, nil
		}
	}

	return nil, fmt.Errorf("TeamViewer Remote Administration Configuration with name '%s' not found", name)
}

func (c *Client) GetTeamViewerRemoteAdministrationConfigurations() (*ResponseTeamViewerRemoteAdministrationConfiguration, error) {
	uri := fmt.Sprintf("%s?page=0&page-size=100&sort=id%%3Aasc", uriTeamViewerRemoteAdmin)

	var out ResponseTeamViewerRemoteAdministrationConfiguration
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get TeamViewer Remote Administration Configurations: %v", err)
	}

	return &out, nil
}

func (c *Client) GetTeamViewerRemoteAdministrationConfigurationByID(configID string) (*TeamViewerRemoteAdministrationConfiguration, error) {
	uri := fmt.Sprintf("%s/%s", uriTeamViewerRemoteAdmin, configID)

	var out TeamViewerRemoteAdministrationConfiguration
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get TeamViewer Remote Administration Configuration by ID: %v", err)
	}

	return &out, nil
}

func (c *Client) CreateTeamViewerRemoteAdministrationConfiguration(displayName *string, siteID *string, scriptToken *string, enabled *bool, sessionTimeout *int) (*TeamViewerRemoteAdministrationConfiguration, error) {

	in := struct {
		DisplayName    *string `json:"displayName"`
		SiteID         *string `json:"siteId"`
		ScriptToken    *string `json:"scriptToken"`
		Enabled        *bool   `json:"enabled"`
		SessionTimeout *int    `json:"sessionTimeout"`
	}{
		DisplayName:    displayName,
		SiteID:         siteID,
		ScriptToken:    scriptToken,
		Enabled:        enabled,
		SessionTimeout: sessionTimeout,
	}

	var out *TeamViewerRemoteAdministrationConfiguration

	err := c.DoRequest("POST", uriTeamViewerRemoteAdmin, in, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to create TeamViewer Remote Administration Configuration: %v", err)
	}
	return out, nil
}

func (c *Client) UpdateTeamViewerRemoteAdministrationConfiguration(id string, config UpdateTeamViewerConfiguration) (*TeamViewerRemoteAdministrationConfiguration, error) {
	uri := fmt.Sprintf("%s/%s", uriTeamViewerRemoteAdmin, id)

	var out TeamViewerRemoteAdministrationConfiguration
	err := c.DoRequest("PATCH", uri, config, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to update TeamViewer Remote Administration Configuration: %v", err)
	}

	return &out, nil
}

func (c *Client) DeleteTeamViewerRemoteAdministrationConfiguration(id string) error {
	uri := fmt.Sprintf("%s/%s", uriTeamViewerRemoteAdmin, id)

	err := c.DoRequest("DELETE", uri, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete TeamViewer Remote Administration Configuration with ID %s: %v", id, err)
	}

	return nil
}

func (c *Client) GetTeamViewerRemoteAdminStatus(configID string) (*ResponseTeamViewerRemoteAdminStatus, error) {
	uri := fmt.Sprintf("%s/%s/status", uriTeamViewerRemoteAdmin, configID)

	var out ResponseTeamViewerRemoteAdminStatus
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get TeamViewer Remote Administration connection status: %v", err)
	}

	return &out, nil
}

func (c *Client) GetTeamViewerSessionStatusByID(configurationId string, sessionId string) (*ResponseTeamViewerSessionStatus, error) {
	uri := fmt.Sprintf("%s/%s/sessions/%s/status", uriTeamViewerRemoteAdmin, configurationId, sessionId)

	var out ResponseTeamViewerSessionStatus
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get TeamViewer Session status by ID: %v", err)
	}

	return &out, nil
}

func (c *Client) CreateTeamViewerSession(configurationId string, session CreateTeamViewerSessionRequest) (*ResponseCreateTeamViewerSession, error) {
	uri := fmt.Sprintf("%s/%s/sessions", uriTeamViewerRemoteAdmin, configurationId)

	var out ResponseCreateTeamViewerSession
	err := c.DoRequest("POST", uri, session, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to create TeamViewer session: %v", err)
	}

	return &out, nil
}

func (c *Client) CloseTeamViewerSession(configurationId string, sessionId string) error {
	uri := fmt.Sprintf("%s/%s/sessions/%s/close", uriTeamViewerRemoteAdmin, configurationId, sessionId)

	err := c.DoRequest("POST", uri, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to close TeamViewer session: %v", err)
	}

	return nil
}
