package jamfpropackageuploader

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// findPkgFiles searches the given directory for files ending with .pkg and returns their paths
func FindPkgFiles(directory string) ([]string, error) {
	var pkgFiles []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".pkg") {
			pkgFiles = append(pkgFiles, path)
		}
		return nil
	})
	return pkgFiles, err
}

// PackageExistsInJCDS checks if a package exists in JCDS by comparing both the file name and MD5 hash.
func PackageExistsInJCDS(packages []jamfpro.ResponseJCDS2List, fileName, fileMD5 string) bool {
	for _, pkg := range packages {
		fmt.Printf("JCDS Entry: FileName: %s, MD5: %s\n", pkg.FileName, pkg.MD5)
		if pkg.FileName == fileName && strings.EqualFold(pkg.MD5, fileMD5) {
			fmt.Printf("Match Found: Package %s with MD5 %s already exists in JCDS.\n", fileName, fileMD5)
			return true
		}
	}
	fmt.Printf("No Match: Package %s with MD5 %s does not exist in JCDS.\n", fileName, fileMD5)
	return false
}

// PackageMetadataExists checks if package metadata exists in Jamf Pro by comparing the package name
func PackageMetadataExists(packages []jamfpro.PackageListItem, fileName string) bool {
	for _, pkg := range packages {
		if pkg.Name == fileName {
			return true
		}
	}
	return false
}

// CalculateFileSHA3 calculates the SHA3 hash of a file
func CalculateFileSHA3(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	hash := sha512.New512_256()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatalf("Failed to calculate SHA3 for file %s: %v", filePath, err)
	}

	return hex.EncodeToString(hash.Sum(nil))
}

// CalculateFileMD5 calculates the MD5 hash of a file and returns it as a hexadecimal string.
func CalculateFileMD5(filePath string) string {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file %s for MD5 calculation: %v", filePath, err)
	}
	defer file.Close()

	// Create a new MD5 hash instance
	hash := md5.New()

	// Copy the file content into the hash instance
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatalf("Failed to calculate MD5 for file %s: %v", filePath, err)
	}

	// Compute the MD5 checksum
	md5sum := hash.Sum(nil)

	// Return the hexadecimal representation of the MD5 checksum
	return hex.EncodeToString(md5sum)
}
