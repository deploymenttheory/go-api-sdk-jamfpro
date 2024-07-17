// classicapi_disk_encryption_configurations.go
// Jamf Pro Classic Api - Disk Encryption Configurations
// api reference: https://developer.jamf.com/jamf-pro/reference/diskencryptionconfigurations
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// URI for Disk Encryption Configurations in Jamf Pro API
const uriDiskEncryptionConfigurations = "/JSSResource/diskencryptionconfigurations"

// Responses & Lists

// Struct to capture the XML response for disk encryption configurations
type ResponseDiskEncryptionConfigurationsList struct {
	Size                        int                                    `xml:"size"`
	DiskEncryptionConfiguration []DiskEncryptionConfigurationsListItem `xml:"disk_encryption_configuration"`
}

type DiskEncryptionConfigurationsListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type ResponseDiskEncryptionConfigurationCreatedAndUpdated struct {
	ID int `xml:"id"`
}

// Resource

// DiskEncryptionConfiguration represents the top-level XML structure for creating/updating a Disk Encryption Configuration.
type ResourceDiskEncryptionConfiguration struct {
	ID                       int                                                  `xml:"id" json:"id"`
	Name                     string                                               `xml:"name"`
	KeyType                  string                                               `xml:"key_type"`
	FileVaultEnabledUsers    string                                               `xml:"file_vault_enabled_users"`
	InstitutionalRecoveryKey *DiskEncryptionConfigurationInstitutionalRecoveryKey `xml:"institutional_recovery_key,omitempty"`
}

// Subsets & Containers

type DiskEncryptionConfigurationInstitutionalRecoveryKey struct {
	Key             string `xml:"key"`
	CertificateType string `xml:"certificate_type"`
	Password        string `xml:"password"`
	Data            string `xml:"data"`
}

// CRUD

// GetDiskEncryptionConfigurations retrieves a serialized list of disk encryption configurations.
func (c *Client) GetDiskEncryptionConfigurations() (*ResponseDiskEncryptionConfigurationsList, error) {
	endpoint := uriDiskEncryptionConfigurations

	var configurations ResponseDiskEncryptionConfigurationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &configurations)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "disk encryption configurations", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &configurations, nil
}

// GetDiskEncryptionConfigurationByID retrieves a single disk encryption configuration by its ID.
func (c *Client) GetDiskEncryptionConfigurationByID(id string) (*ResourceDiskEncryptionConfiguration, error) {
	endpoint := fmt.Sprintf("%s/id/%s", uriDiskEncryptionConfigurations, id)

	var configuration ResourceDiskEncryptionConfiguration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &configuration)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "disk encryption configuration", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &configuration, nil
}

// GetDiskEncryptionConfigurationByName retrieves a disk encryption configuration by its name.
func (c *Client) GetDiskEncryptionConfigurationByName(name string) (*ResourceDiskEncryptionConfiguration, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDiskEncryptionConfigurations, name)

	var configuration ResourceDiskEncryptionConfiguration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &configuration)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "disk encryption configuration", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &configuration, nil
}

// CreateDiskEncryptionConfiguration creates a new disk encryption configuration.
func (c *Client) CreateDiskEncryptionConfiguration(config *ResourceDiskEncryptionConfiguration) (*ResponseDiskEncryptionConfigurationCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriDiskEncryptionConfigurations)

	requestBody := struct {
		XMLName xml.Name `xml:"disk_encryption_configuration"`
		*ResourceDiskEncryptionConfiguration
	}{
		ResourceDiskEncryptionConfiguration: config,
	}

	var createdConfig ResponseDiskEncryptionConfigurationCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdConfig)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "disk encryption configuration", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdConfig, nil
}

// UpdateDiskEncryptionConfigurationByID updates a disk encryption configuration by its ID.
func (c *Client) UpdateDiskEncryptionConfigurationByID(id string, config *ResourceDiskEncryptionConfiguration) (*ResponseDiskEncryptionConfigurationCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/%s", uriDiskEncryptionConfigurations, id)

	requestBody := struct {
		XMLName xml.Name `xml:"disk_encryption_configuration"`
		*ResourceDiskEncryptionConfiguration
	}{
		ResourceDiskEncryptionConfiguration: config,
	}

	var updatedConfig ResponseDiskEncryptionConfigurationCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedConfig)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "disk encryption configuration", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedConfig, nil
}

// UpdateDiskEncryptionConfigurationByName updates a disk encryption configuration by its name.
func (c *Client) UpdateDiskEncryptionConfigurationByName(name string, config *ResourceDiskEncryptionConfiguration) (*ResourceDiskEncryptionConfiguration, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDiskEncryptionConfigurations, name)

	requestBody := struct {
		XMLName xml.Name `xml:"disk_encryption_configuration"`
		*ResourceDiskEncryptionConfiguration
	}{
		ResourceDiskEncryptionConfiguration: config,
	}

	var updatedConfig ResourceDiskEncryptionConfiguration
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedConfig)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "disk encryption configuration", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedConfig, nil
}

// DeleteDiskEncryptionConfigurationByID deletes a disk encryption configuration by its ID.
func (c *Client) DeleteDiskEncryptionConfigurationByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriDiskEncryptionConfigurations, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "disk encryption configuration", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDiskEncryptionConfigurationByName deletes a disk encryption configuration by its name.
func (c *Client) DeleteDiskEncryptionConfigurationByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriDiskEncryptionConfigurations, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "disk encryption configuration", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
