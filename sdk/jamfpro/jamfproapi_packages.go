// jamfproapi_packages.go
// Jamf Pro Api - Packages
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v1-packages
// Jamf Pro Api requires the structs to support an JSON data structure.
// Ref: https://grahamrpugh.com/2024/05/16/jamf-new-packages-api-endpoint.html

package jamfpro

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
)

// URI for Packages in the Jamf Pro Classic API
const uriPackages = "/api/v1/packages"

// ResponsePackagesList struct to capture the JSON response for packages list
type ResponsePackagesList struct {
	TotalCount int               `json:"totalCount"` // The total count attribute
	Results    []ResourcePackage `json:"results"`    // The package list
}

// ResourcePackage struct to capture individual package items in the list
type ResourcePackage struct {
	ID                   string `json:"id"`
	PackageName          string `json:"packageName"`
	FileName             string `json:"fileName,omitempty"`
	CategoryID           string `json:"categoryId,omitempty"`
	Info                 string `json:"info,omitempty"`
	Notes                string `json:"notes,omitempty"`
	Priority             int    `json:"priority,omitempty"`
	OSRequirements       string `json:"osRequirements,omitempty"`
	FillUserTemplate     *bool  `json:"fillUserTemplate,omitempty"`
	Indexed              *bool  `json:"indexed,omitempty"`
	FillExistingUsers    *bool  `json:"fillExistingUsers,omitempty"`
	SWU                  *bool  `json:"swu,omitempty"`
	RebootRequired       *bool  `json:"rebootRequired,omitempty"`
	SelfHealNotify       *bool  `json:"selfHealNotify,omitempty"`
	SelfHealingAction    string `json:"selfHealingAction,omitempty"`
	OSInstall            *bool  `json:"osInstall,omitempty"`
	SerialNumber         string `json:"serialNumber,omitempty"`
	ParentPackageID      string `json:"parentPackageId,omitempty"`
	BasePath             string `json:"basePath,omitempty"`
	SuppressUpdates      *bool  `json:"suppressUpdates,omitempty"`
	CloudTransferStatus  string `json:"cloudTransferStatus,omitempty"`
	IgnoreConflicts      *bool  `json:"ignoreConflicts,omitempty"`
	SuppressFromDock     *bool  `json:"suppressFromDock,omitempty"`
	SuppressEula         *bool  `json:"suppressEula,omitempty"`
	SuppressRegistration *bool  `json:"suppressRegistration,omitempty"`
	InstallLanguage      string `json:"installLanguage,omitempty"`
	MD5                  string `json:"md5,omitempty"`
	SHA256               string `json:"sha256,omitempty"`
	HashType             string `json:"hashType,omitempty"`
	HashValue            string `json:"hashValue,omitempty"`
	Size                 string `json:"size,omitempty"`
	OSInstallerVersion   string `json:"osInstallerVersion,omitempty"`
	Manifest             string `json:"manifest,omitempty"`
	ManifestFileName     string `json:"manifestFileName,omitempty"`
	Format               string `json:"format,omitempty"`
}

// ResponsePackageCreatedAndUpdated represents the response structure for creating and updating a package
type ResponsePackageCreatedAndUpdated struct {
	ID   string `json:"id,omitempty"`
	Href string `json:"href"`
}

// ResponsePackageHistoryList struct to capture the JSON response for package history list
type ResponsePackageHistoryList struct {
	TotalCount int                      `json:"totalCount"` // The total count attribute
	Results    []ResourcePackageHistory `json:"results"`    // The package history list
}

// ResourcePackageHistory struct to capture individual package history items in the list
type ResourcePackageHistory struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// CRUD

// GetPackages retrieves a list of packages with pagination, sorting, and filtering.
func (c *Client) GetPackages(sort, filter string) (*ResponsePackagesList, error) {
	const maxPageSize = 200 // Define a constant for the maximum page size

	var allResults []ResourcePackage
	var totalCount int
	page := 0

	for {
		u, err := url.Parse(uriPackages)
		if err != nil {
			return nil, fmt.Errorf("failed to parse URL: %v", err)
		}

		query := u.Query()
		query.Set("page", fmt.Sprintf("%d", page))
		query.Set("page-size", fmt.Sprintf("%d", maxPageSize))
		if sort != "" {
			query.Set("sort", sort)
		}
		if filter != "" {
			query.Set("filter", filter)
		}
		u.RawQuery = query.Encode()

		var paginatedResponse ResponsePackagesList
		resp, err := c.HTTP.DoRequest("GET", u.String(), nil, &paginatedResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch packages: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		totalCount = paginatedResponse.TotalCount
		allResults = append(allResults, paginatedResponse.Results...)

		if len(paginatedResponse.Results) < maxPageSize {
			break
		}
		page++
	}

	return &ResponsePackagesList{
		TotalCount: totalCount,
		Results:    allResults,
	}, nil
}

// GetPackageByID retrieves details of a specific package by its ID.
func (c *Client) GetPackageByID(id int) (*ResourcePackage, error) {
	endpoint := fmt.Sprintf("%s/%d", uriPackages, id)

	var response ResourcePackage
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "package", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetPackageHistoryByPackageID retrieves the history of a specific package by its ID with pagination, sorting, and filtering.
func (c *Client) GetPackageHistoryByPackageID(id int, sort, filter string) (*ResponsePackageHistoryList, error) {
	const maxPageSize = 200 // Define a constant for the maximum page size

	var allResults []ResourcePackageHistory
	var totalCount int
	page := 0

	for {
		u, err := url.Parse(fmt.Sprintf("%s/%d/history", uriPackages, id))
		if err != nil {
			return nil, fmt.Errorf("failed to parse URL: %v", err)
		}

		query := u.Query()
		query.Set("page", fmt.Sprintf("%d", page))
		query.Set("page-size", fmt.Sprintf("%d", maxPageSize))
		if sort != "" {
			query.Set("sort", sort)
		}
		if filter != "" {
			query.Set("filter", filter)
		}
		u.RawQuery = query.Encode()

		var paginatedResponse ResponsePackageHistoryList
		resp, err := c.HTTP.DoRequest("GET", u.String(), nil, &paginatedResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch package history: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		totalCount = paginatedResponse.TotalCount
		allResults = append(allResults, paginatedResponse.Results...)

		if len(paginatedResponse.Results) < maxPageSize {
			break
		}
		page++
	}

	return &ResponsePackageHistoryList{
		TotalCount: totalCount,
		Results:    allResults,
	}, nil
}

/*
Function: CreatePackage
Method: POST
Path: /api/v1/packages
Description: Creates a new package manifest in Jamf Pro. This is step one in the process of creating a package
and creates the package metadata to Jamf Pro. The package file must be uploaded separately using the
UploadPackage function.
Parameters: pkgManifest - A ResourcePackage struct containing the details of the package to be created.
Returns: ResponsePackageCreatedAndUpdated - The response containing the details of the created package.
Errors: Returns an error if the request fails.
Example:

	// Helper function to create a pointer to a bool
	func BoolPtr(b bool) *bool {
		return &b
	}

	pkg := jamfpro.ResourcePackage{
		PackageName:          "Firefox.dmg",
		FileName:             "Firefox.dmg",
		CategoryID:           "-1",
		Priority:             3,
		FillUserTemplate:     BoolPtr(false),
		SWU:                  BoolPtr(false),
		RebootRequired:       BoolPtr(false),
		OSInstall:            BoolPtr(false),
		SuppressUpdates:      BoolPtr(false),
		SuppressFromDock:     BoolPtr(false),
		SuppressEula:         BoolPtr(false),
		SuppressRegistration: BoolPtr(false),
	}

	response, err := client.CreatePackage(pkg)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(response)
*/
func (c *Client) CreatePackage(pkgManifest ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
	endpoint := uriPackages

	var response ResponsePackageCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, &pkgManifest, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "package", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UploadPackage uploads a package to the Jamf Pro server
func (c *Client) UploadPackage(id int, filePath string) (*ResponsePackageCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/%d/upload", uriPackages, id)

	// Create a buffer to hold the zipped file
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	// Open the file to be uploaded
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a writer for the file inside the zip
	w, err := zipWriter.Create(filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("failed to create zip writer: %v", err)
	}

	// Copy the file contents to the zip writer
	if _, err := io.Copy(w, file); err != nil {
		return nil, fmt.Errorf("failed to write file to zip: %v", err)
	}

	// Close the zip writer to finalize the zip file
	if err := zipWriter.Close(); err != nil {
		return nil, fmt.Errorf("failed to close zip writer: %v", err)
	}

	// Convert the zipped content to bytes
	zippedContent := buf.Bytes()

	// Create a map for the file to be uploaded
	files := map[string][]byte{
		"file": zippedContent,
	}

	var response ResponsePackageCreatedAndUpdated
	resp, err := c.HTTP.DoMultipartRequest("POST", endpoint, nil, files, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to upload package: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

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
