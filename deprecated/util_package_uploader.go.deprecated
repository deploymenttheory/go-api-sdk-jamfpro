// util_package_uploader.go
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

// DoPackageUpload creates a new file in JCDS 2.0 using AWS SDK v2
func (c *Client) DoPackageUpload(filePath string, packageData *ResourcePackage) (*ResponseJCDS2File, *ResponsePackageCreatedAndUpdated, error) {
	// Step 1: Obtain AWS credentials for the package upload endpoint
	var uploadCredentials ResponseJCDS2UploadCredentials
	resp, err := c.HTTP.DoRequest("POST", uriJCDS2+"/files", nil, &uploadCredentials)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to obtain upload credentials: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Validate if we received necessary details
	if uploadCredentials.Region == "" || uploadCredentials.BucketName == "" || uploadCredentials.Path == "" {
		return nil, nil, fmt.Errorf("incomplete upload credentials received")
	}

	// Step 2: Use the obtained credentials to configure AWS SDK v2
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(uploadCredentials.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(uploadCredentials.AccessKeyID, uploadCredentials.SecretAccessKey, uploadCredentials.SessionToken)),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create AWS config: %v", err)
	}

	// Create S3 service client
	s3Client := s3.NewFromConfig(cfg)

	// Step 3: Create an Uploader with the configuration and default options
	uploader := manager.NewUploader(s3Client)

	// Step 3: Use the secure file reading helper
	fileReader, fileSize, err := helpers.ReadJCDSPackageTypes(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read package file securely: %v", err)
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
		return nil, nil, fmt.Errorf("failed to upload file: %v", err)
	}

	fmt.Println("\nUpload completed Successfully")

	// Step 5. Upload package metadata to Jamf Pro
	pkgName := filepath.Base(filePath)
	pkg := ResourcePackage{
		Name:                       packageData.Name,
		Filename:                   pkgName,
		Category:                   packageData.Category,
		Info:                       packageData.Info,
		Notes:                      packageData.Notes,
		Priority:                   packageData.Priority,
		RebootRequired:             packageData.RebootRequired,
		FillUserTemplate:           packageData.FillUserTemplate,
		FillExistingUsers:          packageData.FillExistingUsers,
		BootVolumeRequired:         packageData.BootVolumeRequired,
		AllowUninstalled:           packageData.AllowUninstalled,
		OSRequirements:             packageData.OSRequirements,
		RequiredProcessor:          packageData.RequiredProcessor,
		SwitchWithPackage:          packageData.SwitchWithPackage,
		InstallIfReportedAvailable: packageData.InstallIfReportedAvailable,
		ReinstallOption:            packageData.ReinstallOption,
		TriggeringFiles:            packageData.TriggeringFiles,
		SendNotification:           packageData.SendNotification,
	}

	// Step 5. Upload package metadata to Jamf Pro
	metadataResponse, err := c.CreatePackage(pkg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create package metadata in Jamf Pro: %v", err)
	}

	// Log the package creation response from Jamf Pro
	fmt.Printf("Jamf Pro package metadata created successfully with package ID: %d\n", metadataResponse.ID)

	// Construct the final file upload response
	packageUploadresponse := &ResponseJCDS2File{
		URI: fmt.Sprintf("s3://%s/%s%s", uploadCredentials.BucketName, uploadCredentials.Path, filepath.Base(filePath)),
	}

	// Construct the jamf pro package creation response
	jamfPackageMetaData := &ResponsePackageCreatedAndUpdated{
		ID: metadataResponse.ID,
	}

	// Return the file upload response, the package creation response, and nil for no error
	return packageUploadresponse, jamfPackageMetaData, nil
}
