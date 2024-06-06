// util_package_uploader.go
// This utility function uploads a package file to Jamf Pro by creating a new package and uploading the package file.
// Requires jamf pro v11.5 or later
package jamfpro

import (
	"fmt"
	"path/filepath"
	"strconv"
)

// DoPackageUpload creates a new package and uploads the package file to Jamf Pro
func (c *Client) DoPackageUpload(filePath string, packageData *ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
	// Step 1. Create package metadata in Jamf Pro
	pkgName := filepath.Base(filePath)
	packageData.FileName = pkgName

	metadataResponse, err := c.CreatePackage(*packageData)
	if err != nil {
		return nil, fmt.Errorf("failed to create package metadata in Jamf Pro: %v", err)
	}

	// Log the package creation response from Jamf Pro
	fmt.Printf("Jamf Pro package metadata created successfully with package ID: %s\n", metadataResponse.ID)

	// Step 2. Upload the package file using the newly created package ID
	packageID, err := strconv.Atoi(metadataResponse.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid package ID: %v", err)
	}

	filePaths := []string{filePath}
	uploadResponse, err := c.UploadPackage(packageID, filePaths)
	if err != nil {
		return nil, fmt.Errorf("failed to upload package file: %v", err)
	}

	fmt.Println("Package file uploaded successfully")

	// Return the package creation response and nil for no error
	return uploadResponse, nil
}
