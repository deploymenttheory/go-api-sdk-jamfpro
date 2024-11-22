// jamfproapi_jcds2.go
// Jamf Pro Api - Jamf Cloud Distribution Service (JCDS)
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-jcds-files
// Jamf Pro API requires the structs to support an JSON data structure.
// Ref: https://grahamrpugh.com/2023/08/21/introducing-jcds2.html
// Ref: https://aws.github.io/aws-sdk-go-v2/docs/sdk-utilities/s3/

package jamfpro

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/helpers"
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

type JCDS2Properties struct {
	JCDS2Enabled              bool `json:"jcds2Enabled"`
	FileStreamEndpointEnabled bool `json:"fileStreamEndpointEnabled"`
	MaxChunkSize              int  `json:"maxChunkSize"`
}

// CRUD

// GetJCDS2Packages fetches a file list from Jamf Cloud Distribution Service
func (c *Client) GetJCDS2Packages() ([]ResponseJCDS2List, error) {
	endpoint := uriJCDS2 + "/files"
	var out []ResponseJCDS2List
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "JCDS 2.0", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return out, nil
}

// GetJCDS2Properties fetches properties from Jamf Cloud Distribution Service
func (c *Client) GetJCDS2Properties() (*JCDS2Properties, error) {
	endpoint := uriJCDS2 + "/properties"
	var out JCDS2Properties
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "JCDS 2.0", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetJCDS2PackageURIByName fetches a file URI from Jamf Cloud Distribution Service
func (c *Client) GetJCDS2PackageURIByName(id string) (*ResponseJCDS2File, error) {
	endpoint := fmt.Sprintf("%s/%v", uriJCDS2+"/files", id)
	var out ResponseJCDS2File
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "JCDS 2.0", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// RenewJCDS2Credentials renews credentials for JCDS 2.0
func (c *Client) RenewJCDS2Credentials() (*ResponseJCDS2UploadCredentials, error) {
	endpoint := uriJCDS2 + "/renew-credentials"
	var out ResponseJCDS2UploadCredentials
	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "JCDS 2.0", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateJCDS2PackageV2 creates a new file in JCDS 2.0 using AWS SDK v2 without creating package metadata in Jamf Pro.
func (c *Client) CreateJCDS2PackageV2(filePath string) (*ResponseJCDS2File, error) {
	// Step 1: Obtain AWS credentials for the package upload endpoint
	var uploadCredentials ResponseJCDS2UploadCredentials
	resp, err := c.HTTP.DoRequest("POST", uriJCDS2+"/files", nil, &uploadCredentials)
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

	// Step 2: Use the obtained credentials to configure AWS SDK v2
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(uploadCredentials.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(uploadCredentials.AccessKeyID, uploadCredentials.SecretAccessKey, uploadCredentials.SessionToken)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS config: %v", err)
	}

	// Create S3 service client
	s3Client := s3.NewFromConfig(cfg)

	// Step 3: Create an Uploader with the configuration and default options
	uploader := manager.NewUploader(s3Client)

	// Step 3: Use the secure file reading helper
	fileReader, fileSize, err := helpers.ReadJCDSPackageTypes(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read package file securely: %v", err)
	}

	// Create a progress reader
	progressReader := &ProgressReader{
		reader:     fileReader,
		totalBytes: fileSize,
		progressFn: func(read, total int64, unit string) {
			fmt.Printf("\rUploaded %d / %d %s (%.2f%%)", read, total, unit, float64(read)/float64(total)*100)
		},
	}

	// Create the upload input
	uploadInput := &s3.PutObjectInput{
		Bucket: aws.String(uploadCredentials.BucketName),
		Key:    aws.String(uploadCredentials.Path + filepath.Base(filePath)),
		Body:   progressReader,
	}

	// Step 4. Perform the upload
	_, err = uploader.Upload(context.TODO(), uploadInput)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %v", err)
	}

	fmt.Println("\nUpload completed Successfully")

	// Construct the final file upload response
	finalResponse := &ResponseJCDS2File{
		URI: fmt.Sprintf("s3://%s/%s%s", uploadCredentials.BucketName, uploadCredentials.Path, filepath.Base(filePath)),
	}

	return finalResponse, nil
}

// DeleteJCDS2PackageV2 deletes an existing file from JCDS 2.0 using AWS SDK v2.
func (c *Client) DeleteJCDS2PackageV2(filePath string) error {
	// Step 1: Obtain AWS credentials for the package deletion endpoint
	var uploadCredentials ResponseJCDS2UploadCredentials
	resp, err := c.HTTP.DoRequest("POST", uriJCDS2+"/files", nil, &uploadCredentials)
	if err != nil {
		return fmt.Errorf("failed to obtain deletion credentials: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Validate if we received necessary details
	if uploadCredentials.Region == "" || uploadCredentials.BucketName == "" || uploadCredentials.Path == "" {
		return fmt.Errorf("incomplete deletion credentials received")
	}

	// Step 2: Use the obtained credentials to configure AWS SDK v2
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(uploadCredentials.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(uploadCredentials.AccessKeyID, uploadCredentials.SecretAccessKey, uploadCredentials.SessionToken)),
	)
	if err != nil {
		return fmt.Errorf("failed to create AWS config: %v", err)
	}

	// Create S3 service client
	s3Client := s3.NewFromConfig(cfg)

	// Step 3: Define the object to delete
	objectToDelete := &s3.DeleteObjectInput{
		Bucket: aws.String(uploadCredentials.BucketName),
		Key:    aws.String(uploadCredentials.Path + filepath.Base(filePath)),
	}

	// Step 4: Perform the deletion
	_, err = s3Client.DeleteObject(context.TODO(), objectToDelete)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	fmt.Printf("File '%s' successfully deleted from JCDS 2.0.\n", filepath.Base(filePath))
	return nil
}

// RefreshJCDS2Inventory refreshes the inventory and status of uploads in Jamf Pro.
func (c *Client) RefreshJCDS2Inventory() error {
	endpoint := fmt.Sprintf("%s/refresh-inventory", uriJCDS2)

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedAction, "refresh JCDS 2.0 inventory", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
