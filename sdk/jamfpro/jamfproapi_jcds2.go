// jamfproapi_jcds2.go
// Jamf Pro Api - Jamf Cloud Distribution Service (JCDS)
// api reference: N/A
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriJCDS2 = "/api/v1/jcds"

// List

type ResponseJCDS2List struct {
	Files []JCDSFileListItem `json:"files" xml:"files"`
}

type JCDSFileListItem struct {
	FileName string `json:"fileName" xml:"fileName"`
	MD5      string `json:"md5" xml:"md5"`
}

// Response

type ResponseJCDSUploadCredentials struct {
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
	SessionToken    string `json:"sessionToken"`
	Region          string `json:"region"`
	BucketName      string `json:"bucketName"`
	Path            string `json:"path"`
	UUID            string `json:"uuid"`
}

type UploadProgressPercentage struct {
	Filename  string
	TotalSize int64
	SeenSoFar int64
}

// CRUD

// GetJCDS2Packages fetches a file list from Jamf Cloud Distribution Service
func (c *Client) GetJCDS2Packages() ([]JCDSFileListItem, error) {
	endpoint := uriJCDS2 + "/files"
	var out []JCDSFileListItem // Changed to a slice of JCDSFileListItem
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "JCDS 2.0", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return out, nil
}

// GetJCDS2PackageByName retrieves a file by name
func (c *Client) GetJCDS2PackageByName(id string) (*JCDSFileListItem, error) {
	endpoint := fmt.Sprintf("%s/%v", uriJCDS2+"/files", id)
	var out JCDSFileListItem
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "JCDS 2.0", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateJCDS2Package creates a new file in JCDS 2.0
func (c *Client) CreateJCDS2Package(JCDSpackage *JCDSFileListItem) (*JCDSFileListItem, error) {
	endpoint := uriJCDS2 + "/files"
	var ResponseJCDSPackageCreate JCDSFileListItem

	resp, err := c.HTTP.DoRequest("POST", endpoint, JCDSpackage, &ResponseJCDSPackageCreate)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "JCDS 2.0", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &ResponseJCDSPackageCreate, nil
}

// RenewJCDS2Credentials renews credentials for JCDS 2.0
func (c *Client) RenewJCDS2Credentials() (*ResponseJCDSUploadCredentials, error) {
	endpoint := uriJCDS2 + "/renew-credentials"
	var out ResponseJCDSUploadCredentials
	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "JCDS 2.0", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
