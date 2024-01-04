// jamfproapi_jcds2.go
// Jamf Pro Api - Jamf Cloud Distribution Service (JCDS)
// api reference: N/A
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriJCDS2 = "/api/v1/jcds/files"

// List

type ResponseJCDS2List struct {
	Files []JCDSFileListItem `json:"files" xml:"files"`
}

type JCDSFileListItem struct {
	FileName string `json:"fileName" xml:"fileName"`
	MD5      string `json:"md5" xml:"md5"`
}

// Response

type JCDSUploadResponse struct {
	Credentials JCDSUploadCredentials `json:"Credentials"`
}

// Other

type JCDSUploadCredentials struct {
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

// GetJCDS2Files fetches a file list from Jamf Cloud Distribution Service
func (c *Client) GetJCDS2Files() (*ResponseJCDS2List, error) {
	var out ResponseJCDS2List

	resp, err := c.HTTP.DoRequest("GET", uriJCDS2, nil, &out)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf(errMsgFailedGet, "sso failover settings", err)
		return nil, err
	}

	return &out, nil
}
