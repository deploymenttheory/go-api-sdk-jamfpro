package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// URI for Packages in the Jamf Pro Classic API
const uriPackages = "/api/v1/packages"

// ResponsePackagesList struct to capture the JSON response for packages list
type ResponsePackagesList struct {
	TotalCount *int              `json:"totalCount"` // The total count attribute
	Results    []ResourcePackage `json:"results"`    // The package list
}

// ResourcePackage struct to capture individual package items in the list
type ResourcePackage struct {
	ID                   string `json:"id"`                   // The ID element
	PackageName          string `json:"packageName"`          // The package name
	FileName             string `json:"fileName"`             // The file name
	CategoryID           string `json:"categoryId"`           // The category ID
	Info                 string `json:"info"`                 // The info
	Notes                string `json:"notes"`                // The notes
	Priority             int    `json:"priority"`             // The priority
	OSRequirements       string `json:"osRequirements"`       // The OS requirements
	FillUserTemplate     bool   `json:"fillUserTemplate"`     // Fill user template
	Indexed              bool   `json:"indexed"`              // Indexed
	FillExistingUsers    bool   `json:"fillExistingUsers"`    // Fill existing users
	SWU                  bool   `json:"swu"`                  // Software update
	RebootRequired       bool   `json:"rebootRequired"`       // Reboot required
	SelfHealNotify       bool   `json:"selfHealNotify"`       // Self heal notify
	SelfHealingAction    string `json:"selfHealingAction"`    // Self healing action
	OSInstall            bool   `json:"osInstall"`            // OS install
	SerialNumber         string `json:"serialNumber"`         // Serial number
	ParentPackageID      string `json:"parentPackageId"`      // Parent package ID
	BasePath             string `json:"basePath"`             // Base path
	SuppressUpdates      bool   `json:"suppressUpdates"`      // Suppress updates
	CloudTransferStatus  string `json:"cloudTransferStatus"`  // Cloud transfer status
	IgnoreConflicts      bool   `json:"ignoreConflicts"`      // Ignore conflicts
	SuppressFromDock     bool   `json:"suppressFromDock"`     // Suppress from dock
	SuppressEula         bool   `json:"suppressEula"`         // Suppress EULA
	SuppressRegistration bool   `json:"suppressRegistration"` // Suppress registration
	InstallLanguage      string `json:"installLanguage"`      // Install language
	MD5                  string `json:"md5"`                  // MD5
	SHA256               string `json:"sha256"`               // SHA256
	HashType             string `json:"hashType"`             // Hash type
	HashValue            string `json:"hashValue"`            // Hash value
	Size                 string `json:"size"`                 // Size
	OSInstallerVersion   string `json:"osInstallerVersion"`   // OS installer version
	Manifest             string `json:"manifest"`             // Manifest
	ManifestFileName     string `json:"manifestFileName"`     // Manifest file name
	Format               string `json:"format"`               // Format
}

// Response

type ResponsePackageCreatedAndUpdated struct {
	ID int `json:"id,omitempty"` // ID of the created/updated package
}

// CRUD

// GetPackages retrieves a list of packages.
func (c *Client) GetPackages(sort_filter string) (*ResponsePackagesList, error) {
	resp, err := c.DoPaginatedGet(
		uriPackages,
		maxPageSize,
		startingPageNumber,
		sort_filter,
	)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "packages", err)
	}

	var out ResponsePackagesList
	out.TotalCount = &resp.Size

	for _, value := range resp.Results {
		var newObj ResourcePackage
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "packages", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// // GetPackageByID retrieves details of a specific package by its ID.
// func (c *Client) GetPackageByID(id int) (*ResourcePackage, error) {
// 	endpoint := fmt.Sprintf("%s/id/%d", uriPackages, id)

// 	var response ResourcePackage
// 	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
// 	if err != nil {
// 		return nil, fmt.Errorf(errMsgFailedGetByID, "package", id, err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return &response, nil
// }

// // GetPackageByName retrieves details of a specific package by its name.
// func (c *Client) GetPackageByName(name string) (*ResourcePackage, error) {
// 	endpoint := fmt.Sprintf("%s/name/%s", uriPackages, name)

// 	var response ResourcePackage
// 	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
// 	if err != nil {
// 		return nil, fmt.Errorf(errMsgFailedGetByName, "package", name, err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return &response, nil
// }

// // CreatePackage creates a new package in Jamf Pro
// func (c *Client) CreatePackage(pkg ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
// 	endpoint := fmt.Sprintf("%s/id/%d", uriPackages, pkg.ID)

// 	requestBody := struct {
// 		XMLName xml.Name `xml:"package"`
// 		ResourcePackage
// 	}{
// 		ResourcePackage: pkg,
// 	}

// 	var response ResponsePackageCreatedAndUpdated
// 	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
// 	if err != nil {
// 		return nil, fmt.Errorf(errMsgFailedCreate, "package", err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return &response, nil
// }

// // UpdatePackageByID updates an existing package by its ID on the Jamf Pro server
// // and returns the response with the ID of the updated package.
// func (c *Client) UpdatePackageByID(id int, pkg *ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
// 	endpoint := fmt.Sprintf("%s/id/%d", uriPackages, id)

// 	requestBody := struct {
// 		XMLName xml.Name `xml:"package"`
// 		*ResourcePackage
// 	}{
// 		ResourcePackage: pkg,
// 	}

// 	var response ResponsePackageCreatedAndUpdated

// 	// Use PUT method for updating the package
// 	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
// 	if err != nil {
// 		return nil, fmt.Errorf(errMsgFailedUpdateByID, "package", id, err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return &response, nil
// }

// // UpdatePackageByName updates an existing package by its ID on the Jamf Pro server
// // and returns the response with the ID of the updated package.
// func (c *Client) UpdatePackageByName(name string, pkg *ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
// 	endpoint := fmt.Sprintf("%s/name/%s", uriPackages, name)

// 	requestBody := struct {
// 		XMLName xml.Name `xml:"package"`
// 		*ResourcePackage
// 	}{
// 		ResourcePackage: pkg,
// 	}

// 	var response ResponsePackageCreatedAndUpdated

// 	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
// 	if err != nil {
// 		return nil, fmt.Errorf(errMsgFailedUpdateByName, "package", name, err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return &response, nil
// }

// // DeletePackageByID deletes a package by its ID from the Jamf Pro server.
// func (c *Client) DeletePackageByID(id int) error {
// 	endpoint := fmt.Sprintf("%s/id/%d", uriPackages, id)

// 	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
// 	if err != nil {
// 		return fmt.Errorf(errMsgFailedDeleteByID, "package", id, err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return nil
// }

// // DeletePackageByName deletes a package by its name from the Jamf Pro server.
// func (c *Client) DeletePackageByName(name string) error {
// 	endpoint := fmt.Sprintf("%s/name/%s", uriPackages, name)

// 	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
// 	if err != nil {
// 		return fmt.Errorf(errMsgFailedDeleteByName, "package", name, err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return nil
// }
