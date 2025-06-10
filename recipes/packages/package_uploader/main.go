package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	uploader "github.com/deploymenttheory/go-api-sdk-jamfpro/recipes/packages/package_uploader/internal"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {

	// Load the client OAuth configuration
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Print the ASCII art
	uploader.PrintASCIILogo()

	// Set up a reader from standard input
	reader := bufio.NewReader(os.Stdin)

	// Define the directory containing the .pkg files
	fmt.Print("Enter the directory containing the .pkg files: ")
	directory, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read directory input: %v", err)
	}

	// Trim space to remove any new line character
	directory = filepath.Clean(strings.TrimSpace(directory))

	// Validate the directory path
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", directory)
	}

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

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Fetch existing packages from Jamf Pro
	jamfProPackages, err := client.GetPackages("", "")
	if err != nil {
		log.Fatalf("Failed to get Jamf Pro packages: %v", err)
	}

	// Create a map of existing JCDS entries
	jcdsMap := make(map[string]bool)
	for _, pkg := range jamfProPackages.Results {
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
		if uploader.PackageMetadataExists(jamfProPackages.Results, fileName) {
			fmt.Printf("Package metadata for %s already exists in Jamf Pro. Skipping metadata creation.\n", fileName)
			fmt.Println("-------------------------------------------------")
			continue
		} else {
			fmt.Printf("Package metadata for %s does not exist in Jamf Pro. Proceeding with package metadata creation.\n", fileName)
		}

		// Upload the package
		fmt.Printf("Uploading package: %s\n", fileName)

		// Create package metadata
		pkg := jamfpro.ResourcePackage{
			PackageName:          fileName,
			FileName:             fileName,
			CategoryID:           "-1",
			Priority:             3,
			FillUserTemplate:     uploader.BoolPtr(false),
			SWU:                  uploader.BoolPtr(false),
			RebootRequired:       uploader.BoolPtr(false),
			OSInstall:            uploader.BoolPtr(false),
			SuppressUpdates:      uploader.BoolPtr(false),
			SuppressFromDock:     uploader.BoolPtr(false),
			SuppressEula:         uploader.BoolPtr(false),
			SuppressRegistration: uploader.BoolPtr(false),
		}

		_, err = client.DoPackageUpload(filePath, &pkg)
		if err != nil {
			log.Fatalf("Error uploading package to Jamf Pro: %v", err)
		}

		fmt.Printf("Package %s uploaded and verified successfully.\n", fileName)
		fmt.Println("-------------------------------------------------")
	}
}
