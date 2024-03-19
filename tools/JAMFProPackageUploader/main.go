package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	uploader "github.com/deploymenttheory/go-api-sdk-jamfpro/tools/JAMFProPackageUploader/internal"
)

func main() {
	// Print the ASCII art
	uploader.PrintASCIILogo()

	// Define the directory containing the .pkg files
	fmt.Print("Enter the directory containing the .pkg files: ")
	var directory string
	fmt.Scanln(&directory)

	// Find all .pkg files in the directory
	pkgFiles, err := uploader.FindPkgFiles(directory)
	if err != nil {
		log.Fatalf("Failed to find .pkg files: %v", err)
	}

	if len(pkgFiles) == 0 {
		fmt.Println("No .pkg files found in the specified directory.")
		return
	}

	// List out the .pkg files to be uploaded
	fmt.Println("The following .pkg files will be uploaded:")
	for _, file := range pkgFiles {
		fmt.Println(filepath.Base(file))
	}

	// Load the client OAuth configuration
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	jcdsPackages, err := client.GetJCDS2Packages()
	if err != nil {
		log.Fatalf("Failed to get JCDS 2.0 packages: %v", err)
	}

	jamfProPackages, err := client.GetPackages()
	if err != nil {
		log.Fatalf("Failed to get Jamf Pro packages: %v", err)
	}

	// Create a map of existing JCDS entries
	jcdsMap := make(map[string]bool)
	for _, pkg := range jcdsPackages {
		key := fmt.Sprintf("%s-%s", pkg.FileName, pkg.MD5)
		jcdsMap[key] = true
	}

	for _, filePath := range pkgFiles {
		fileName := filepath.Base(filePath)
		fileMD5 := uploader.CalculateFileMD5(filePath) // Use the MD5 calculation function
		jcdsKey := fmt.Sprintf("%s-%s", fileName, fileMD5)

		// Log the start of processing for this file
		fmt.Println("-------------------------------------------------")
		fmt.Printf("Processing package: %s\n", fileName)
		fmt.Printf("Calculated MD5: %s\n", fileMD5)

		// Check JCDS for the package with MD5 using the map
		fmt.Printf("Checking JCDS for existing package...\n")
		if _, exists := jcdsMap[jcdsKey]; exists {
			fmt.Printf("Package %s with MD5 %s already exists in JCDS. Skipping package upload.\n", fileName, fileMD5)
			fmt.Println("-------------------------------------------------")
			continue
		} else {
			fmt.Printf("Package %s with MD5 %s does not exist in JCDS. Proceeding with package upload.\n", fileName, fileMD5)
		}

		// Check Jamf Pro for existing package metadata
		fmt.Printf("Checking Jamf Pro for existing package metadata...\n")
		if uploader.PackageMetadataExists(jamfProPackages.Package, fileName) {
			fmt.Printf("Package metadata for %s already exists in Jamf Pro. Skipping metadata creation.\n", fileName)
			fmt.Println("-------------------------------------------------")
			continue
		} else {
			fmt.Printf("Package metadata for %s does not exist in Jamf Pro. Proceeding with package metadata creation.\n", fileName)
		}

		// Upload the package
		fmt.Printf("Uploading package: %s\n", fileName)

		// Assuming packageMetadata is correctly defined somewhere in your code
		var packageMetadata *jamfpro.ResourcePackage

		responseJCDS2File, responsePackageCreatedAndUpdated, err := client.DoPackageUpload(filePath, packageMetadata)
		if err != nil {
			log.Fatalf("Failed to upload %s: %v", filePath, err)
		}

		// Marshal and log the JCDS2File response
		responseJCDS2Bytes, err := json.Marshal(responseJCDS2File)
		if err != nil {
			log.Fatalf("Failed to marshal JCDS2File response for %s: %v", filePath, err)
		}
		fmt.Printf("JCDS2File Upload response for %s: %s\n", fileName, string(responseJCDS2Bytes))

		// Marshal and log the PackageCreatedAndUpdated response
		responsePackageCreatedAndUpdatedBytes, err := json.Marshal(responsePackageCreatedAndUpdated)
		if err != nil {
			log.Fatalf("Failed to marshal PackageCreatedAndUpdated response for %s: %v", filePath, err)
		}
		fmt.Printf("PackageCreatedAndUpdated response for %s: %s\n", fileName, string(responsePackageCreatedAndUpdatedBytes))

		fmt.Println("-------------------------------------------------")
	}
}
