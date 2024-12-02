// util_package_uploader.go
// This utility function uploads a package file to Jamf Pro by creating a new package and uploading the package file.
// Requires jamf pro v11.5 or later
package jamfpro

import (
	"fmt"
	"path/filepath"
	"time"
)

// DoPackageUpload
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

	time.Sleep(3 * time.Second)

	uploadedPackage, err := c.GetPackageByID(packageID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify uploaded package: %v", err)
	}

	if uploadedPackage.HashType != "SHA3_512" || uploadedPackage.HashValue != initialHash {
		return nil, fmt.Errorf("hash verification failed: initial=%s, uploaded=%s (type: %s)",
			initialHash, uploadedPackage.HashValue, uploadedPackage.HashType)
	}

	fmt.Printf("Package %s SHA3_512 verification was successful with validated hash of %s\n", pkgName, uploadedPackage.HashValue)
	return uploadResponse, nil
}
