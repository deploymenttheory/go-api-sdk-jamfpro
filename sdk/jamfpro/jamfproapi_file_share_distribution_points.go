package jamfpro

import (
	"encoding/json"
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
		return nil, fmt.Errorf(errMsgFailedGet, "distribution point", err)
	}

	return &distributionPoints, nil
}

// GetDistributionPointByID retrieves a single distribution point by its ID.
func (c *Client) CreateDistributionPoint(payload *ResourceFileShareDistributionPoint) (*ResponseFileShareDistributionPointCreatedAndUpdated, error) {
	endpoint := uriDistributionPoints
	var createdResponseObject ResponseFileShareDistributionPointCreatedAndUpdated

	//TODO: Remove
	fmt.Println("DEBUBDEBUGDBEUG")
	// Marshal and log the payload
	payloadBytes, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %v", err)
	}
	fmt.Printf("Distribution Point Payload:\n%s\n", string(payloadBytes))

	resp, err := c.HTTP.DoRequest("POST", endpoint, payload, &createdResponseObject)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "distribution point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdResponseObject, nil
	// return nil, nil
}

// func (c *Client) getDistributonPointsByName() {

// }

// func (c *Client) CreateDistributionPoint() {

// }

// func (c *Client) updateDistributionPointByID() {

// }

// func (c *Client) updateDistributionPointByName() {

// }

// func (c *Client) deleteDistributionPointByID() {

// }

// func (c *Client) DeleteDistributionPointByName() {

// }
