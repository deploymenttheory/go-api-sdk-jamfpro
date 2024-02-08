// jcds2.go
// Jamf Pro Api
// Work in progress. waiting for jcds enabled jamf instance to
// TODO validate structs and logic flow.
// TODO create distinct create and update jcds package functions
// TODO move helper funcs to helpers.go
// TODO create package mains for create and update package funcs
// TODO remove repeat funcs and use packages.go where appropriate
// TODO create download package func with aws file manager
// TODO refactor to use v2 aws sdk for s3
package jamfpro

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const uriAPIJCDSFiles = "/api/v1/jcds/files"

type JCDSFilesResponse struct {
	Files []JCDSFile `json:"files" xml:"files"`
}

type JCDSFile struct {
	FileName string `json:"fileName" xml:"fileName"`
	MD5      string `json:"md5" xml:"md5"`
}

type JCDSUploadResponse struct {
	Credentials JCDSUploadCredentials `json:"Credentials"`
}

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

func NewUploadProgressPercentage(filename string) *UploadProgressPercentage {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}
	return &UploadProgressPercentage{
		Filename:  filename,
		TotalSize: fileInfo.Size(),
	}
}

func (p *UploadProgressPercentage) TrackProgress() {
	for {
		percentage := (float64(p.SeenSoFar) / float64(p.TotalSize)) * 100
		fmt.Printf("\r%s  %d / %d  (%.2f%%)", p.Filename, p.SeenSoFar, p.TotalSize, percentage)

		if p.SeenSoFar >= p.TotalSize {
			break
		}

		time.Sleep(1 * time.Second)
	}
	fmt.Println() // Move to the next line after completion
}

func (p *UploadProgressPercentage) AddBytes(bytes int64) {
	atomic.AddInt64(&p.SeenSoFar, bytes)
}

func (c *Client) CheckExistingPackageInJamfPro(idOrName string) (*ResponsePackage, error) {
	// Fetch the list of all packages
	packagesList, err := c.GetPackages()
	if err != nil {
		return nil, err
	}

	// Search through the packages list for the package with the given ID or Name
	for _, pkgList := range packagesList {
		for _, pkg := range pkgList.Packages {
			// Check if idOrName matches the package's ID or Name
			if strconv.Itoa(pkg.ID) == idOrName || pkg.Name == idOrName {
				// Return the details of the matched package
				return c.GetPackageByID(pkg.ID)
			}
		}
	}

	return nil, fmt.Errorf("package with identifier %s not found", idOrName)
}

func (c *Client) CheckExistingPackageInJCDS(pkgName string, pkgPath string) (bool, error) {
	// Step 2 - Get the list of existing packages in JCDS
	var jcdsFilesResponse JCDSFilesResponse
	if err := c.DoRequest("GET", uriAPIJCDSFiles, nil, nil, &jcdsFilesResponse); err != nil {
		return false, fmt.Errorf("failed to get existing packages in JCDS: %v", err)
	}

	// Step 3 - Iterate through the list to find the package and its MD5 hash
	var existingPkgFound bool
	var existingPkgMD5 string
	for _, jcdsFile := range jcdsFilesResponse.Files {
		if jcdsFile.FileName == pkgName {
			existingPkgFound = true
			existingPkgMD5 = jcdsFile.MD5
			break
		}
	}

	if !existingPkgFound {
		fmt.Println("Package not found in JCDS.")
		return true, nil
	}

	// Step 4 - Calculate MD5 hash of the local package
	pkgFile, err := os.Open(pkgPath)
	if err != nil {
		return false, fmt.Errorf("failed to open local package file: %v", err)
	}
	defer pkgFile.Close()

	pkgMD5, err := calculateMD5(pkgFile)
	if err != nil {
		return false, fmt.Errorf("failed to calculate MD5 hash of local package: %v", err)
	}

	// Step 5 - Compare MD5 hashes
	if existingPkgMD5 == pkgMD5 {
		fmt.Println("MD5 matches, not replacing existing package in JCDS")
		return false, nil
	}

	// Step 6 - Delete existing package and upload the new one
	deleteURL := fmt.Sprintf("%s/%s", uriAPIJCDSFiles, pkgName)
	if err := c.DoRequest("DELETE", deleteURL, nil, nil, nil, c.HTTP.Logger); err != nil {
		return false, fmt.Errorf("failed to delete existing package in JCDS: %v", err)
	}
	fmt.Println("Existing package deleted from JCDS")

	return true, nil
}

func calculateMD5(file *os.File) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (c *Client) GetPackageUploadCredentials() (*JCDSUploadCredentials, error) {
	var response JCDSUploadResponse
	if err := c.DoRequest("POST", uriAPIJCDSFiles, nil, nil, &response); err != nil {
		return nil, fmt.Errorf("failed to get upload credentials: %v", err)
	}
	return &response.Credentials, nil
}

func (c *Client) UploadPackageToJCDS(pkgPath string, creds *JCDSUploadCredentials) error {
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // 5MB per part
		u.LeavePartsOnError = true   // If the upload fails, don't delete the parts
	})

	// Open the file at pkgPath
	file, err := os.Open(pkgPath)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", pkgPath, err)
	}
	defer file.Close()

	tracker := NewUploadProgressPercentage(pkgPath)
	progressReader := &ProgressReader{
		r:       file,
		tracker: tracker,
	}
	go tracker.TrackProgress()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &creds.BucketName,
		Key:    &creds.Path,
		Body:   progressReader,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	return nil
}

type ProgressReader struct {
	r       io.Reader
	tracker *UploadProgressPercentage
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.r.Read(p)
	pr.tracker.AddBytes(int64(n))
	return n, err
}

func (c *Client) UploadPackageMetadataToJamfPro(pkgName string, pkgID int) error {
	pkgData := fmt.Sprintf("<package><name>%s</name><filename>%s</filename></package>", pkgName, pkgName)
	reqType := "POST"
	if pkgID > 0 {
		reqType = "PUT"
	}
	if err := c.DoRequest(reqType, fmt.Sprintf("/JSSResource/packages/id/%d", pkgID), pkgData, nil, nil, c.HTTP.Logger); err != nil {
		return fmt.Errorf("failed to upload package metadata: %v", err)
	}
	return nil
}

/*
// BucketBasics encapsulates the Amazon Simple Storage Service (Amazon S3) actions
// used in the examples.
// It contains S3Client, an Amazon S3 service client that is used to perform bucket
// and object actions.
type BucketBasics struct {
	S3Client *s3.Client
}

// BucketExists checks whether a bucket exists in the current account.
func (basics BucketBasics) BucketExists(bucketName string) (bool, error) {
	_, err := basics.S3Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	exists := true
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.(type) {
			case *types.NotFound:
				log.Printf("Bucket %v is available.\n", bucketName)
				exists = false
				err = nil
			default:
				log.Printf("Either you don't have access to bucket %v or another error occurred. "+
					"Here's what happened: %v\n", bucketName, err)
			}
		}
	} else {
		log.Printf("Bucket %v exists and you already own it.", bucketName)
	}

	return exists, err
}

// BucketBasics encapsulates the Amazon Simple Storage Service (Amazon S3) actions
// used in the examples.
// It contains S3Client, an Amazon S3 service client that is used to perform bucket
// and object actions.

// DownloadFile gets an object from a bucket and stores it in a local file.
func (basics BucketBasics) DownloadFile(bucketName string, objectKey string, fileName string) error {
	result, err := basics.S3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, objectKey, err)
		return err
	}
	defer result.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Couldn't create file %v. Here's why: %v\n", fileName, err)
		return err
	}
	defer file.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err)
	}
	_, err = file.Write(body)
	return err
}

// BucketBasics encapsulates the Amazon Simple Storage Service (Amazon S3) actions
// used in the examples.
// It contains S3Client, an Amazon S3 service client that is used to perform bucket
// and object actions.

// UploadFile reads from a file and puts the data into an object in a bucket.
func (basics BucketBasics) UploadFile(bucketName string, objectKey string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
	} else {
		defer file.Close()
		_, err = basics.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   file,
		})
		if err != nil {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileName, bucketName, objectKey, err)
		}
	}
	return err
}
*/
