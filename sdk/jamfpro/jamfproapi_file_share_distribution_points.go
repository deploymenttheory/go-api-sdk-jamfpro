package jamfpro

import (
	"fmt"
)

const uriDistributionPoints = "/api/v1/distribution-points"

type ResponseDistributionPointList struct {
	TotalCount int                                  `json:"totalCount"`
	Results    []ResourceFileShareDistributionPoint `json:"results"`
}

type ResourceFileShareDistributionPoint struct {
	ShareName                 string `json:"shareName,omitempty"`
	Workgroup                 string `json:"workgroup,omitempty"`
	Port                      int    `json:"port,omitempty"`
	ReadWriteUsername         string `json:"readWriteUsername,omitempty"`
	ReadWritePassword         string `json:"readWritePassword,omitempty"`
	ReadOnlyUsername          string `json:"readOnlyUsername,omitempty"`
	ReadOnlyPassword          string `json:"readOnlyPassword,omitempty"`
	ID                        string `json:"id,omitempty"`
	Name                      string `json:"name"`
	ServerName                string `json:"serverName"`
	Principal                 bool   `json:"principal,omitempty"`
	BackupDistributionPointID string `json:"backupDistributionPointId,omitempty"`
	SSHUsername               string `json:"sshUsername,omitempty"`
	SSHPassword               string `json:"sshPassword,omitempty"`
	LocalPathToShare          string `json:"localPathToShare,omitempty"`
	FileSharingConnectionType string `json:"fileSharingConnectionType"`
	HTTPSEnabled              bool   `json:"httpsEnabled"`
	HTTPSPort                 int    `json:"httpsPort"`
	HTTPSContext              string `json:"httpsContext,omitempty"`
	HTTPSSecurityType         string `json:"httpsSecurityType"`
	HTTPSUsername             string `json:"httpsUsername,omitempty"`
	HTTPSPassword             string `json:"httpsPassword,omitempty"`
	EnableLoadBalancing       bool   `json:"enableLoadBalancing,omitempty"`
}

type ResponseFileShareDistributionPointCreatedAndUpdated struct {
	ID   string `json:"id,omitempty"`
	Href string `json:"href,omitempty"`
}

func (c *Client) GetDistributionPoints() (*ResponseDistributionPointList, error) {
	endpoint := uriDistributionPoints

	var distributionPoints ResponseDistributionPointList

	_, err := c.HTTP.DoRequest("GET", endpoint, nil, &distributionPoints)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "file share distribution point", err)
	}

	return &distributionPoints, nil
}

// GetDistributionPointByID retrieves a single distribution point by its ID.
func (c *Client) CreateDistributionPoint(payload *ResourceFileShareDistributionPoint) (*ResponseFileShareDistributionPointCreatedAndUpdated, error) {
	endpoint := uriDistributionPoints
	var createdResponseObject ResponseFileShareDistributionPointCreatedAndUpdated

	resp, err := c.HTTP.DoRequest("POST", endpoint, payload, &createdResponseObject)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "file share distribution point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdResponseObject, nil
	// return nil, nil
}

func (c *Client) GetDistributionPointByID(id string) (*ResourceFileShareDistributionPoint, error) {
	var out ResourceFileShareDistributionPoint
	endpoint := fmt.Sprintf("%s/%s", uriDistributionPoints, id)

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf(errMsgFailedGetByID, "file share distribution point", id, err)
		return nil, err
	}

	return &out, nil
}

func (c *Client) UpdateDistributionPointByID(id string, distributionPointUpdate *ResourceFileShareDistributionPoint) (*ResourceFileShareDistributionPoint, error) {
	endpoint := fmt.Sprintf("%s/%s", uriDistributionPoints, id)

	var response ResourceFileShareDistributionPoint
	resp, err := c.HTTP.DoRequest("PUT", endpoint, distributionPointUpdate, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "file share distribution point", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

func (c *Client) DeleteDistributionPointByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriDistributionPoints, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "file share distribution point", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
