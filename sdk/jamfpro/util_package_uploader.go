// util_package_uploader.go
// This utility function uploads a package file to Jamf Pro by creating a new package and uploading the package file.
// Requires jamf pro v11.5 or later
package jamfpro

import (
	"fmt"
	"path/filepath"
)

func (c *Client) DoPackageUpload(filePath string, packageData *ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
	initialHash, err := calculateSHA3_512(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate initial SHA3_512: %v", err)
	}

	pkgName := filepath.Base(filePath)
	packageData.FileName = pkgName
	packageData.HashType = "SHA3_512"
	packageData.HashValue = initialHash

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

	fmt.Println("Package file uploaded successfully")

	uploadedPackage, err := c.GetPackageByID(packageID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify uploaded package: %v", err)
	}

	if uploadedPackage.HashValue != initialHash {
		return nil, fmt.Errorf("SHA3_512 verification failed: initial=%s, uploaded=%s", initialHash, uploadedPackage.HashValue)
	}

	fmt.Printf("Package SHA3_512 verification was successful with validated hash of %s\n", uploadedPackage.HashValue)
	return uploadResponse, nil
}
