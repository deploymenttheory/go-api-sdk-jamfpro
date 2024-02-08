// classicapi_file_share_distribution_points.go
// Jamf Pro Classic Api - File Share Distribution Points
// api reference: https://developer.jamf.com/jamf-pro/reference/distributionpoints
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// URI for Distribution Points in Jamf Pro API
const uriDistributionPoints = "/JSSResource/distributionpoints"

// List

// Struct to capture the XML response for distribution points list
type ResponseDistributionPointsList struct {
	Size              int                       `xml:"size"`
	DistributionPoint DistributionPointListItem `xml:"distribution_point"`
}

type DistributionPointListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// Struct for detailed Distribution Point data
type ResourceFileShareDistributionPoint struct {
	ID                       int    `xml:"id"`
	Name                     string `xml:"name"`
	IPAddress                string `xml:"ipAddress"`
	IP_Address               string `xml:"ip_address"`
	IsMaster                 bool   `xml:"is_master"`
	FailoverPoint            string `xml:"failover_point"`
	FailoverPointURL         string `xml:"failover_point_url"`
	EnableLoadBalancing      bool   `xml:"enable_load_balancing"`
	LocalPath                string `xml:"local_path"`
	SSHUsername              string `xml:"ssh_username"`
	Password                 string `xml:"password"`
	ConnectionType           string `xml:"connection_type"`
	ShareName                string `xml:"share_name"`
	WorkgroupOrDomain        string `xml:"workgroup_or_domain"`
	SharePort                int    `xml:"share_port"`
	ReadOnlyUsername         string `xml:"read_only_username"`
	ReadOnlyPassword         string `xml:"read_only_password"`
	ReadWriteUsername        string `xml:"read_write_username"`
	ReadWritePassword        string `xml:"read_write_password"`
	HTTPDownloadsEnabled     bool   `xml:"http_downloads_enabled"`
	HTTPURL                  string `xml:"http_url"`
	Context                  string `xml:"context"`
	Protocol                 string `xml:"protocol"`
	Port                     int    `xml:"port"`
	NoAuthenticationRequired bool   `xml:"no_authentication_required"`
	UsernamePasswordRequired bool   `xml:"username_password_required"`
	HTTPUsername             string `xml:"http_username"`
	HTTPPassword             string `xml:"http_password"`
}

// CRUD

// GetDistributionPoints retrieves a serialized list of distribution points.
func (c *Client) GetDistributionPoints() (*ResponseDistributionPointsList, error) {
	endpoint := uriDistributionPoints

	var distributionPoints ResponseDistributionPointsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &distributionPoints, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "distribution points", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &distributionPoints, nil
}

// GetDistributionPointByID retrieves a single distribution point by its ID.
func (c *Client) GetDistributionPointByID(id int) (*ResourceFileShareDistributionPoint, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriDistributionPoints, id)

	var distributionPoint ResourceFileShareDistributionPoint
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &distributionPoint, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "distribution point", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &distributionPoint, nil
}

// GetDistributionPointByName retrieves a single distribution point by its name.
func (c *Client) GetDistributionPointByName(name string) (*ResourceFileShareDistributionPoint, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDistributionPoints, name)

	var distributionPoint ResourceFileShareDistributionPoint
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &distributionPoint, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "distribution point", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &distributionPoint, nil
}

// CreateDistributionPoint creates a new distribution point.
func (c *Client) CreateDistributionPoint(dp *ResourceFileShareDistributionPoint) (*ResourceFileShareDistributionPoint, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriDistributionPoints)

	// Wrap the distribution point with the XML root element name
	requestBody := struct {
		XMLName xml.Name `xml:"distribution_point"`
		*ResourceFileShareDistributionPoint
	}{
		ResourceFileShareDistributionPoint: dp,
	}

	var createdDistributionPoint ResourceFileShareDistributionPoint
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdDistributionPoint, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "distribution point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdDistributionPoint, nil
}

// UpdateDistributionPointByID updates a distribution point by its ID.
func (c *Client) UpdateDistributionPointByID(id int, dp *ResourceFileShareDistributionPoint) (*ResourceFileShareDistributionPoint, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriDistributionPoints, id)

	requestBody := struct {
		XMLName xml.Name `xml:"distribution_point"`
		*ResourceFileShareDistributionPoint
	}{
		ResourceFileShareDistributionPoint: dp,
	}

	var updatedDistributionPoint ResourceFileShareDistributionPoint
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedDistributionPoint, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "distribution point", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedDistributionPoint, nil
}

// UpdateDistributionPointByName updates a distribution point by its name.
func (c *Client) UpdateDistributionPointByName(name string, dp *ResourceFileShareDistributionPoint) (*ResourceFileShareDistributionPoint, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDistributionPoints, name)

	requestBody := struct {
		XMLName xml.Name `xml:"distribution_point"`
		*ResourceFileShareDistributionPoint
	}{
		ResourceFileShareDistributionPoint: dp,
	}

	var updatedDistributionPoint ResourceFileShareDistributionPoint
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedDistributionPoint, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "distribution point", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedDistributionPoint, nil
}

// DeleteDistributionPointByID deletes a distribution point by its ID.
func (c *Client) DeleteDistributionPointByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriDistributionPoints, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "distribution point", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDistributionPointByName deletes a distribution point by its name.
func (c *Client) DeleteDistributionPointByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriDistributionPoints, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "distribution point", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
