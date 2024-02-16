// jamfproapi_jcds2.go
// Jamf Pro Api - Jamf Cloud Distribution Service (JCDS)
// api reference: N/A
// Jamf Pro API requires the structs to support an JSON data structure.
// Ref: https://grahamrpugh.com/2023/08/21/introducing-jcds2.html
// Ref: https://aws.github.io/aws-sdk-go-v2/docs/sdk-utilities/s3/

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

// progressReader is a wrapper around an io.Reader that reports progress in kilobytes and megabytes.
type progressReader struct {
	reader     io.Reader
	totalBytes int64
	readBytes  int64
	progressFn func(int64, int64, string) // function to call to report progress
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

	progressFn := func(read, total int64, unit string) {
		fmt.Printf("\rUploaded %d / %d %s (%.2f%%)", read, total, unit, float64(read)/float64(total)*100)
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

	// Step 4. Perform the upload
	_, err = uploader.Upload(context.TODO(), uploadInput)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %v", err)
	}

	fmt.Println("\nUpload completed Successfully")

	// Step 5. Upload package metadata to Jamf Pro
	pkgName := filepath.Base(filePath)
	pkg := ResourcePackage{
		Name:     pkgName,
		Filename: pkgName,
		// Add other package metadata fields as necessary
	}

	metadataResponse, err := c.CreatePackage(pkg)
	if err != nil {
		return nil, fmt.Errorf("failed to create package metadata in Jamf Pro: %v", err)
	}

	// Log the package creation response from Jamf Pro
	fmt.Printf("Jamf Pro package metadata created successfully with package ID: %d\n", metadataResponse.ID)

	// Combine JCDS file URI and Jamf Pro package creation response for the final response
	finalResponse := &ResponseJCDS2File{
		URI: fmt.Sprintf("s3://%s/%s%s", uploadCredentials.BucketName, uploadCredentials.Path, filepath.Base(filePath)),
		// Include relevant fields from the Jamf Pro package creation response if necessary
	}

	return finalResponse, nil
}

// Read implements the io.Reader interface for progressReader, reporting upload progress in kilobytes and megabytes.
func (r *progressReader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	r.readBytes += int64(n)

	// Report progress in more human-readable units (KB or MB)
	const kb = 1024
	const mb = 1024 * kb
	readKB := r.readBytes / kb
	totalKB := r.totalBytes / kb
	if totalKB > kb { // If the total size is larger than 1 MB, report in MB
		readMB := r.readBytes / mb
		totalMB := r.totalBytes / mb
		r.progressFn(readMB, totalMB, "MB")
	} else { // For smaller files, report in KB
		r.progressFn(readKB, totalKB, "KB")
	}

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
