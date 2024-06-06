package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the package deploy request
	deployPackageRequest := &jamfpro.ResourceDeployPackage{
		Manifest: jamfpro.PackageManifest{
			HashType:         "MD5",
			URL:              "https://example.jamf.com/this/package",
			Hash:             "dcb02a41cd6d842943459a88c96a5f72",
			DisplayImageURL:  "https://example.jamf.com/img/display/this/package.jpg",
			FullSizeImageURL: "https://example.jamf.com/img/full/this/package.jpg",
			BundleID:         "com.jamf.example",
			BundleVersion:    "0.1.0",
			Subtitle:         "Subtitle",
			Title:            "Title",
			SizeInBytes:      12345,
		},
		InstallAsManaged: true,
		Devices:          []int{1, 2, 3},
		GroupID:          "1",
	}

	// Call SendMDMCommandForPackageDeployment function
	response, err := client.SendMDMCommandForPackageDeployment(deployPackageRequest)
	if err != nil {
		log.Fatalf("Error deploying package: %v", err)
	}

	// Pretty print the response details in JSON
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling response data: %v", err)
	}
	fmt.Println("Deploy package response details:\n", string(responseJSON))
}
