// jamfproapi_jcds2.go
// Jamf Pro Api - Jamf Cloud Distribution Service (JCDS)
// api reference: N/A
// Jamf Pro API requires the structs to support an JSON data structure.
// Ref: https://grahamrpugh.com/2023/08/21/introducing-jcds2.html

package jamfpro

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

// progressReader is a wrapper around an io.Reader that reports progress

type progressReader struct {
	reader     io.Reader
	totalBytes int64
	readBytes  int64
	progressFn func(int64, int64) // function to call to report progress
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

/*
// CreateJCDS2PackageV1 creates a new file in JCDS 2.0
func (c *Client) CreateJCDS2PackageV1(filePath string) (*ResponseJCDS2File, error) {
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

	// Step 3: Use s3manager.Uploader for uploading the file with progress tracking
	uploader := s3manager.NewUploader(sess)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}

	progressFn := func(readBytes, totalBytes int64) {
		fmt.Printf("\rUploaded %d / %d bytes (%.2f%%)", readBytes, totalBytes, float64(readBytes)/float64(totalBytes)*100)
	}

	reader := &progressReader{
		reader:     file,
		totalBytes: fileInfo.Size(),
		progressFn: progressFn,
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(uploadCredentials.BucketName),
		Key:    aws.String(uploadCredentials.Path + filepath.Base(filePath)),
		Body:   reader,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %v", err)
	}

	fmt.Println("\nUpload completed")

	// Create a response object to return
	response := &ResponseJCDS2File{
		URI: fmt.Sprintf("s3://%s/%s%s", uploadCredentials.BucketName, uploadCredentials.Path, filepath.Base(filePath)),
	}

	return response, nil
}
*/

// CreateJCDS2PackageV2 creates a new file in JCDS 2.0 using AWS SDK v2
func (c *Client) CreateJCDS2PackageV2(filePath string) (*ResponseJCDS2File, error) {
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

	// Open the file and use a progressReader to track the upload progress
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %v", err)
	}

	progressFn := func(bytesTransferred, totalBytes int64) {
		fmt.Printf("\rUploaded %d / %d bytes (%.2f%%)", bytesTransferred, totalBytes, float64(bytesTransferred)/float64(totalBytes)*100)
	}

	reader := &progressReader{
		reader:     file,
		totalBytes: fileInfo.Size(),
		progressFn: progressFn,
	}

	// Create the upload input
	uploadInput := &s3.PutObjectInput{
		Bucket: aws.String(uploadCredentials.BucketName),
		Key:    aws.String(uploadCredentials.Path + filepath.Base(filePath)),
		Body:   reader,
	}

	// Perform the upload
	_, err = uploader.Upload(context.TODO(), uploadInput)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %v", err)
	}

	fmt.Println("\nUpload completed")

	// Create a response object to return
	response := &ResponseJCDS2File{
		URI: fmt.Sprintf("s3://%s/%s%s", uploadCredentials.BucketName, uploadCredentials.Path, filepath.Base(filePath)),
	}

	return response, nil
}

// progressReader is a wrapper around an io.Reader that reports progress
func (r *progressReader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.readBytes += int64(n)
	r.progressFn(r.readBytes, r.totalBytes)
	return n, err
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
