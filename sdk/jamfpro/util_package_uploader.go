// util_package_uploader.go
// This utility function uploads a package file to Jamf Pro by creating a new package and uploading the package file.
// Requires jamf pro v11.5 or later
package jamfpro

import (
	"fmt"
	"path/filepath"
	"time"
)

/*
DoPackageUpload uploads a local package file to Jamf Pro, creates the package record,
and verifies the uploaded package’s SHA3_512 hash.

Steps:
1. Calculate the SHA3_512 of the local file before uploading (initialHash).
2. Create package metadata in Jamf Pro (Package name, etc.).
3. Upload the actual file to Jamf Pro.
4. Poll Jamf Pro until the uploaded package’s SHA3_512 hash is present or until max retries are reached.
5. Compare Jamf Pro’s hash with the initially calculated hash to ensure data integrity.

Arguments:
- filePath: The path to the local package file to be uploaded.
- packageData: The metadata (ResourcePackage) associated with the package to be uploaded.

Returns:
- ResponsePackageCreatedAndUpdated: Contains package details after creation and upload.
- error: Non-nil if any step fails.

Usage:
- Instantiate a Client that can communicate with Jamf Pro.
- Prepare a ResourcePackage struct with needed fields (FileName is set automatically here).
- Call DoPackageUpload with the path of your package and that resource data.
*/
func (c *Client) DoPackageUpload(filePath string, packageData *ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
	initialHash, err := CalculateSHA3_512(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate initial SHA3_512: %v", err)
	}

	pkgName := filepath.Base(filePath)
	packageData.FileName = pkgName
	metadataResponse, err := c.CreatePackage(*packageData)
	if err != nil {
		return nil, fmt.Errorf("failed to create package metadata in Jamf Pro: %v", err)
	}
	fmt.Printf("Jamf Pro package metadata created successfully with package ID: %s\n", metadataResponse.ID)

	packageID := metadataResponse.ID
	filePaths := []string{filePath}
	uploadResponse, err := c.UploadPackage(packageID, filePaths)
	if err != nil {
		return nil, fmt.Errorf("failed to upload package file: %v", err)
	}
	fmt.Printf("Package %s file uploaded successfully\n", pkgName)

	maxAttempts := 10
	sleepBetween := 2 * time.Second

	var uploadedPackage *ResourcePackage
	for i := 1; i <= maxAttempts; i++ {
		uploadedPackage, err = c.GetPackageByID(packageID)
		if err != nil {
			return nil, fmt.Errorf("failed to get package by ID (attempt %d/%d): %v", i, maxAttempts, err)
		}

		if uploadedPackage.HashType == "SHA3_512" && uploadedPackage.HashValue != "" {
			break
		}

		if i < maxAttempts {
			time.Sleep(sleepBetween)
		}
	}

	if uploadedPackage.HashType != "SHA3_512" || uploadedPackage.HashValue == "" {
		return nil, fmt.Errorf("timed out waiting for SHA3_512 to be populated in Jamf Pro")
	}

	if uploadedPackage.HashValue != initialHash {
		return nil, fmt.Errorf("hash verification failed: initial=%s, uploaded=%s (type: %s)",
			initialHash, uploadedPackage.HashValue, uploadedPackage.HashType)
	}

	fmt.Printf("Package %s SHA3_512 verification was successful with validated hash of %s\n", pkgName, uploadedPackage.HashValue)
	return uploadResponse, nil
}
