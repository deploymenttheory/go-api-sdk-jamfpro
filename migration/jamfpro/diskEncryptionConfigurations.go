// diskEncryptionConfigurations.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIDiskEncryptionConfigurations = "/JSSResource/diskencryptionconfigurations"

// Top-level struct for individual configurations
type ResponseDiskEncryptionConfiguration struct {
	ID                    int    `json:"id,omitempty" xml:"id,omitempty"`
	Name                  string `json:"name" xml:"name"`
	KeyType               string `json:"key_type" xml:"key_type"`
	FileVaultEnabledUsers string `json:"file_vault_enabled_users" xml:"file_vault_enabled_users"`
}

// Struct for listing all configurations
type ResponseDiskEncryptionConfigurationsList struct {
	Size                        int                             `json:"size" xml:"size"`
	DiskEncryptionConfiguration DiskEncryptionConfigurationList `json:"disk_encryption_configuration" xml:"disk_encryption_configuration"`
}

type DiskEncryptionConfigurationList struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// DiskEncryptionConfiguration represents the top-level XML structure for creating/updating a Disk Encryption Configuration.
type DiskEncryptionConfiguration struct {
	XMLName                  xml.Name                                                       `xml:"disk_encryption_configuration"`
	Name                     string                                                         `xml:"name"`
	KeyType                  string                                                         `xml:"key_type"`
	FileVaultEnabledUsers    string                                                         `xml:"file_vault_enabled_users"`
	InstitutionalRecoveryKey *DiskEncryptionConfigurationDataSubsetInstitutionalRecoveryKey `xml:"institutional_recovery_key,omitempty"`
}

// DiskEncryptionConfigurationDataSubsetInstitutionalRecoveryKey represents the XML structure for Institutional Recovery Key.
type DiskEncryptionConfigurationDataSubsetInstitutionalRecoveryKey struct {
	Key             string `xml:"key"`
	CertificateType string `xml:"certificate_type"`
	Password        string `xml:"password"`
	Data            string `xml:"data"`
}

//--- diskEncryptionConfigurations CRUD Functions ---//

// GetDiskEncryptionConfigurationByID retrieves the Disk Encryption Configuration by its ID
func (c *Client) GetDiskEncryptionConfigurationByID(id int) (*ResponseDiskEncryptionConfiguration, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIDiskEncryptionConfigurations, id)

	var config ResponseDiskEncryptionConfiguration
	if err := c.DoRequest("GET", url, nil, nil, &config); err != nil {
		return nil, fmt.Errorf("failed to get disk encryption configuration by ID: %v", err)
	}

	return &config, nil
}

// GetDiskEncryptionConfigurations retrieves a list of all Disk Encryption Configurations
func (c *Client) GetDiskEncryptionConfigurations() ([]ResponseDiskEncryptionConfigurationsList, error) {
	url := uriAPIDiskEncryptionConfigurations

	var configList []ResponseDiskEncryptionConfigurationsList
	if err := c.DoRequest("GET", url, nil, nil, &configList); err != nil {
		return nil, fmt.Errorf("failed to get all disk encryption configurations: %v", err)
	}

	return configList, nil
}

// GetDiskEncryptionConfigurationByName retrieves the Disk Encryption Configuration by its name
func (c *Client) GetDiskEncryptionConfigurationByName(name string) (*ResponseDiskEncryptionConfiguration, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIDiskEncryptionConfigurations, name)

	var config ResponseDiskEncryptionConfiguration
	if err := c.DoRequest("GET", url, nil, nil, &config); err != nil {
		return nil, fmt.Errorf("failed to get disk encryption configuration by name: %v", err)
	}

	return &config, nil
}

// CreateDiskEncryptionConfiguration creates a new Disk Encryption Configuration
func (c *Client) CreateDiskEncryptionConfiguration(config *DiskEncryptionConfiguration) error {
	url := fmt.Sprintf("%s/id/0", uriAPIDiskEncryptionConfigurations)

	if err := c.DoRequest("POST", url, config, nil, nil); err != nil {
		return fmt.Errorf("failed to create disk encryption configuration: %v", err)
	}

	return nil
}

// UpdateDiskEncryptionConfigurationByID updates an existing Disk Encryption Configuration by its ID
func (c *Client) UpdateDiskEncryptionConfigurationByID(id int, config *DiskEncryptionConfiguration) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIDiskEncryptionConfigurations, id)

	if err := c.DoRequest("PUT", url, config, nil, nil); err != nil {
		return fmt.Errorf("failed to update disk encryption configuration by ID: %v", err)
	}

	return nil
}

// UpdateDiskEncryptionConfigurationByName updates an existing Disk Encryption Configuration by its name
func (c *Client) UpdateDiskEncryptionConfigurationByName(name string, config *DiskEncryptionConfiguration) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIDiskEncryptionConfigurations, name)

	if err := c.DoRequest("PUT", url, config, nil, nil); err != nil {
		return fmt.Errorf("failed to update disk encryption configuration by name: %v", err)
	}

	return nil
}

// DeleteDiskEncryptionConfigurationByID deletes a Disk Encryption Configuration by its ID
func (c *Client) DeleteDiskEncryptionConfigurationByID(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIDiskEncryptionConfigurations, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete disk encryption configuration by ID: %v", err)
	}

	return nil
}

// DeleteDiskEncryptionConfigurationByName deletes a Disk Encryption Configuration by its name
func (c *Client) DeleteDiskEncryptionConfigurationByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIDiskEncryptionConfigurations, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete disk encryption configuration by name: %v", err)
	}

	return nil
}
