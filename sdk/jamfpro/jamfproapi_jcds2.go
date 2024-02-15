// jamfproapi_jcds2.go
// Jamf Pro Api - Jamf Cloud Distribution Service (JCDS)
// api reference: N/A
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const uriJCDS2 = "/api/v1/jcds"

// List

type ResponseJCDS2List struct {
	FileName string `json:"fileName"` // The name of the file
	Length   int64  `json:"length"`   // The size of the file in bytes
	MD5      string `json:"md5"`      // The MD5 hash of the file
	Region   string `json:"region"`   // The AWS region where the file is stored
	SHA3     string `json:"sha3"`     // The SHA3 hash of the file
}

// Response

type ResponseJCDS2UploadCredentials struct {
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
	SessionToken    string `json:"sessionToken"`
	Region          string `json:"region"`
	BucketName      string `json:"bucketName"`
	Path            string `json:"path"`
	UUID            string `json:"uuid"`
}

type ResponseJCDS2File struct {
	URI string `json:"uri"`
}

type UploadProgressPercentage struct {
	Filename  string
	TotalSize int64
	SeenSoFar int64
}

// CRUD

// GetJCDS2Packages fetches a file list from Jamf Cloud Distribution Service
func (c *Client) GetJCDS2Packages() ([]ResponseJCDS2List, error) {
	endpoint := uriJCDS2 + "/files"
	var out []ResponseJCDS2List
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "JCDS 2.0", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return out, nil
}

// GetJCDS2PackageURIByName fetches a file URI from Jamf Cloud Distribution Service
func (c *Client) GetJCDS2PackageURIByName(id string) (*ResponseJCDS2File, error) {
	endpoint := fmt.Sprintf("%s/%v", uriJCDS2+"/files", id)
	var out ResponseJCDS2File
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out, c.HTTP.Logger)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "JCDS 2.0", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateJCDS2Package creates a new file in JCDS 2.0
func (c *Client) CreateJCDS2Package(filePath string) (*ResponseJCDS2File, error) {
	// Step 1: Obtain AWS credentials for the package upload endpoint
	var uploadCredentials ResponseJCDS2UploadCredentials
	resp, err := c.HTTP.DoRequest("POST", uriJCDS2+"/files", nil, &uploadCredentials, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain upload credentials: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Validate if we received necessary details
	if uploadCredentials.Region == "" || uploadCredentials.BucketName == "" || uploadCredentials.Path == "" {
		return nil, fmt.Errorf("incomplete upload credentials received")
	}

	// Step 2: Use the obtained credentials to configure the AWS SDK for Go
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(uploadCredentials.Region),
		Credentials: credentials.NewStaticCredentials(
			uploadCredentials.AccessKeyID,
			uploadCredentials.SecretAccessKey,
			uploadCredentials.SessionToken,
		),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %v", err)
	}

	svc := s3.New(sess)

	// Step 3: Use the AWS SDK for Go to upload the file to the specified S3 bucket
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(uploadCredentials.BucketName),
		Key:    aws.String(uploadCredentials.Path + file.Name()),
		Body:   file,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload file to S3: %v", err)
	}

	// Create a response object to return
	response := &ResponseJCDS2File{
		URI: fmt.Sprintf("s3://%s/%s%s", uploadCredentials.BucketName, uploadCredentials.Path, file.Name()),
	}

	return response, nil
}

// RenewJCDS2Credentials renews credentials for JCDS 2.0
func (c *Client) RenewJCDS2Credentials() (*ResponseJCDS2UploadCredentials, error) {
	endpoint := uriJCDS2 + "/renew-credentials"
	var out ResponseJCDS2UploadCredentials
	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "JCDS 2.0", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
