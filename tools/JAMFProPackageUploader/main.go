package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
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
	configFilePath := "/Users/dafyddwatkins/localtesting/clientconfig.json"
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	jcdsPackages, err := client.GetJCDS2Packages()
	if err != nil {
		log.Fatalf("Failed to get JCDS 2.0 packages: %v", err)
	}

	jamfProPackages, err := client.GetPackages()
	if err != nil {
		log.Fatalf("Failed to get Jamf Pro packages: %v", err)
	}

	for _, filePath := range pkgFiles {
		fileName := filepath.Base(filePath)
		fileMD5 := uploader.CalculateFileMD5(filePath) // Use the MD5 calculation function

		// Log the start of processing for this file
		fmt.Println("-------------------------------------------------")
		fmt.Printf("Processing package: %s\n", fileName)
		fmt.Printf("Calculated MD5: %s\n", fileMD5)

		// Check JCDS for the package with MD5
		fmt.Printf("Checking JCDS for existing package...\n")
		if exists := uploader.PackageExistsInJCDS(jcdsPackages, fileName, fileMD5); exists {
			fmt.Printf("Package %s with MD5 %s already exists in JCDS. Skipping upload.\n", fileName, fileMD5)
			fmt.Println("-------------------------------------------------")
			continue
		} else {
			fmt.Printf("Package %s with MD5 %s does not exist in JCDS. Proceeding with upload.\n", fileName, fileMD5)
		}

		// Check Jamf Pro for existing package metadata
		fmt.Printf("Checking Jamf Pro for existing package metadata...\n")
		if uploader.PackageMetadataExists(jamfProPackages.Package, fileName) {
			fmt.Printf("Package metadata for %s already exists in Jamf Pro. Skipping metadata upload.\n", fileName)
			fmt.Println("-------------------------------------------------")
			continue
		} else {
			fmt.Printf("Package metadata for %s does not exist in Jamf Pro. Proceeding with metadata upload.\n", fileName)
		}

		// Upload the package
		fmt.Printf("Uploading package: %s\n", fileName)
		response, err := client.CreateJCDS2PackageV2(filePath)
		if err != nil {
			log.Fatalf("Failed to upload %s: %v", filePath, err)
		}

		// Marshal and log the response
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Failed to marshal response for %s: %v", filePath, err)
		}
		fmt.Printf("Upload response for %s: %s\n", fileName, string(responseBytes))
		fmt.Println("-------------------------------------------------")
	}
}
