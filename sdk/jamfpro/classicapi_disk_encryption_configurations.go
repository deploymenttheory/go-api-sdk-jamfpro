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

// Struct to capture the XML response for disk encryption configurations
type ResponseDiskEncryptionConfigurationsList struct {
	Size                        int                                 `xml:"size"`
	DiskEncryptionConfiguration []DiskEncryptionConfigurationDetail `xml:"disk_encryption_configuration"`
}

type DiskEncryptionConfigurationDetail struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type ResponseDiskEncryptionConfiguration struct {
	ID                    int    `xml:"id"`
	Name                  string `xml:"name"`
	KeyType               string `xml:"key_type"`
	FileVaultEnabledUsers string `xml:"file_vault_enabled_users"`
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

// GetDiskEncryptionConfigurations retrieves a serialized list of disk encryption configurations.
func (c *Client) GetDiskEncryptionConfigurations() (*ResponseDiskEncryptionConfigurationsList, error) {
	endpoint := uriDiskEncryptionConfigurations

	var configurations ResponseDiskEncryptionConfigurationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &configurations)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Disk Encryption Configurations: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &configurations, nil
}

// GetDiskEncryptionConfigurationByID retrieves a single disk encryption configuration by its ID.
func (c *Client) GetDiskEncryptionConfigurationByID(configID int) (*ResponseDiskEncryptionConfiguration, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriDiskEncryptionConfigurations, configID)

	var configuration ResponseDiskEncryptionConfiguration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &configuration)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Disk Encryption Configuration by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &configuration, nil
}

// GetDiskEncryptionConfigurationByName retrieves a disk encryption configuration by its name.
func (c *Client) GetDiskEncryptionConfigurationByName(configName string) (*ResponseDiskEncryptionConfiguration, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDiskEncryptionConfigurations, configName)

	var configuration ResponseDiskEncryptionConfiguration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &configuration)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Disk Encryption Configuration by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &configuration, nil
}

// CreateDiskEncryptionConfiguration creates a new disk encryption configuration.
func (c *Client) CreateDiskEncryptionConfiguration(config *DiskEncryptionConfiguration) (*ResponseDiskEncryptionConfiguration, error) {
	// When creating a new configuration, the ID in the URL should be 0
	endpoint := fmt.Sprintf("%s/id/0", uriDiskEncryptionConfigurations)

	// Wrap the configuration with the XML root element name
	requestBody := struct {
		XMLName xml.Name `xml:"disk_encryption_configuration"`
		*DiskEncryptionConfiguration
	}{
		DiskEncryptionConfiguration: config,
	}

	var createdConfig ResponseDiskEncryptionConfiguration
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Disk Encryption Configuration: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdConfig, nil
}

// UpdateDiskEncryptionConfigurationByID updates a disk encryption configuration by its ID.
func (c *Client) UpdateDiskEncryptionConfigurationByID(configID int, config *DiskEncryptionConfiguration) (*DiskEncryptionConfiguration, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriDiskEncryptionConfigurations, configID)

	requestBody := struct {
		XMLName xml.Name `xml:"disk_encryption_configuration"`
		*DiskEncryptionConfiguration
	}{
		DiskEncryptionConfiguration: config,
	}

	var updatedConfig DiskEncryptionConfiguration
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to update Disk Encryption Configuration by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedConfig, nil
}

// UpdateDiskEncryptionConfigurationByName updates a disk encryption configuration by its name.
func (c *Client) UpdateDiskEncryptionConfigurationByName(configName string, config *DiskEncryptionConfiguration) (*DiskEncryptionConfiguration, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriDiskEncryptionConfigurations, configName)

	requestBody := struct {
		XMLName xml.Name `xml:"disk_encryption_configuration"`
		*DiskEncryptionConfiguration
	}{
		DiskEncryptionConfiguration: config,
	}

	var updatedConfig DiskEncryptionConfiguration
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to update Disk Encryption Configuration by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedConfig, nil
}

// DeleteDiskEncryptionConfigurationByID deletes a disk encryption configuration by its ID.
func (c *Client) DeleteDiskEncryptionConfigurationByID(configID int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriDiskEncryptionConfigurations, configID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Disk Encryption Configuration by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDiskEncryptionConfigurationByName deletes a disk encryption configuration by its name.
func (c *Client) DeleteDiskEncryptionConfigurationByName(configName string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriDiskEncryptionConfigurations, configName)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Disk Encryption Configuration by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
