// jamfproapi_mdm_renewal.go
// Jamf Pro Api - MDM Renewal
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-mdm-renewal-renewal-strategies-clientmanagementid
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriMdmRenewal = "/api/v1/mdm-renewal"

// Resource

// ResourceMdmRenewalStrategy represents a single MDM renewal strategy for a device.
type ResourceMdmRenewalStrategy struct {
	ID                     string `json:"id"`
	MdmRenewalErrorID      string `json:"mdmRenewalErrorId"`
	MdmRenewalStrategyType string `json:"mdmRenewalStrategyType"`
	StrategyTimeStamp      string `json:"strategyTimeStamp,omitempty"`
	MdmRenewalCheckInURL   string `json:"mdmRenewalCheckInUrl,omitempty"`
	MdmRenewalServerURL    string `json:"mdmRenewalServerUrl,omitempty"`
}

// ResourceMdmRenewalDeviceCommonDetails represents the MDM renewal device common details for a device.
type ResourceMdmRenewalDeviceCommonDetails struct {
	ID                                                    string `json:"id,omitempty"`
	ClientManagementID                                    string `json:"clientManagementId"`
	RenewMdmProfileStartDate                              string `json:"renewMdmProfileStartDate,omitempty"`
	MdmProfileNeedsRenewalDueToCaRenewed                  bool   `json:"mdmProfileNeedsRenewalDueToCaRenewed"`
	MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring bool   `json:"mdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring"`
	MdmCheckinURL                                         string `json:"mdmCheckinUrl,omitempty"`
	MdmServerURL                                          string `json:"mdmServerUrl,omitempty"`
}

// ResourceMdmRenewalDeviceCommonDetailsRequest is the PATCH request body for device common details.
type ResourceMdmRenewalDeviceCommonDetailsRequest struct {
	ClientManagementID                                    string `json:"clientManagementId"`
	RenewMdmProfileStartDate                              string `json:"renewMdmProfileStartDate,omitempty"`
	MdmProfileNeedsRenewalDueToCaRenewed                  *bool  `json:"mdmProfileNeedsRenewalDueToCaRenewed,omitempty"`
	MdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring *bool  `json:"mdmProfileNeedsRenewalDueToDeviceIdentityCertExpiring,omitempty"`
	MdmCheckinURL                                         string `json:"mdmCheckinUrl,omitempty"`
	MdmServerURL                                          string `json:"mdmServerUrl,omitempty"`
}

// GetMdmRenewalStrategiesByClientManagementID retrieves the MDM renewal strategies for a device.
func (c *Client) GetMdmRenewalStrategiesByClientManagementID(clientManagementID string) ([]ResourceMdmRenewalStrategy, error) {
	endpoint := fmt.Sprintf("%s/renewal-strategies/%s", uriMdmRenewal, clientManagementID)

	var out []ResourceMdmRenewalStrategy
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mdm renewal strategies", clientManagementID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return out, nil
}

// DeleteMdmRenewalStrategiesByClientManagementID deletes the MDM renewal strategies for a device.
func (c *Client) DeleteMdmRenewalStrategiesByClientManagementID(clientManagementID string) error {
	endpoint := fmt.Sprintf("%s/renewal-strategies/%s", uriMdmRenewal, clientManagementID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDeleteByID, "mdm renewal strategies", clientManagementID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// GetMdmRenewalDeviceCommonDetailsByClientManagementID retrieves the MDM renewal device common details for a device.
func (c *Client) GetMdmRenewalDeviceCommonDetailsByClientManagementID(clientManagementID string) (*ResourceMdmRenewalDeviceCommonDetails, error) {
	endpoint := fmt.Sprintf("%s/device-common-details/%s", uriMdmRenewal, clientManagementID)

	var out ResourceMdmRenewalDeviceCommonDetails
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mdm renewal device common details", clientManagementID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateMdmRenewalDeviceCommonDetails updates (PATCH) the MDM renewal device common details.
func (c *Client) UpdateMdmRenewalDeviceCommonDetails(request *ResourceMdmRenewalDeviceCommonDetailsRequest) error {
	endpoint := fmt.Sprintf("%s/device-common-details", uriMdmRenewal)

	resp, err := c.HTTP.DoRequest("PATCH", endpoint, request, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedUpdate, "mdm renewal device common details", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
