package main

import (
	"encoding/xml"
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

	// Define the package update
	updatedPackage := jamfpro.ResourcePackage{
		PackageName:          "Updated Package Name",
		FileName:             "updated_package.zip",
		CategoryID:           "Productivity",
		Info:                 "This is an updated package for office productivity tools.",
		Notes:                "Please ensure compatibility before deploying.",
		Priority:             10,
		OSRequirements:       "10.15, 11.0",
		FillUserTemplate:     boolPtr(false),
		Indexed:              boolPtr(false),
		FillExistingUsers:    boolPtr(false),
		SWU:                  boolPtr(false),
		RebootRequired:       boolPtr(false),
		SelfHealNotify:       boolPtr(false),
		SelfHealingAction:    "nothing",
		OSInstall:            boolPtr(false),
		SerialNumber:         "123456",
		ParentPackageID:      "3",
		BasePath:             "/path/to/package",
		SuppressUpdates:      boolPtr(false),
		CloudTransferStatus:  "READY",
		IgnoreConflicts:      boolPtr(false),
		SuppressFromDock:     boolPtr(false),
		SuppressEula:         boolPtr(false),
		SuppressRegistration: boolPtr(false),
		InstallLanguage:      "en_US",
		MD5:                  "0cc175b9c0f1b6a831c399e269772661",
		SHA256:               "61be55a8e2f6b4e172338bddf184d6dbee29c98853e0a0485ecee7f27b9af0b4",
		HashType:             "MD5",
		HashValue:            "0cc175b9c0f1b6a831c399e269772661",
		Size:                 "234KB",
		OSInstallerVersion:   "10.3.x",
		Manifest:             "manifest",
		ManifestFileName:     "manifest.plist",
		Format:               "format",
	}

	packageID := "1" // Replace with actual ID

	// Update the package by ID
	updated, err := client.UpdatePackageManifestByID(packageID, updatedPackage)
	if err != nil {
		log.Fatalf("Error updating package by ID: %v", err)
	}

	// Print the updated package details
	packageXML, _ := xml.MarshalIndent(updated, "", "    ")
	fmt.Println("Updated Package:", string(packageXML))
}

// boolPtr is a helper function to create a pointer to a bool value
func boolPtr(b bool) *bool {
	return &b
}
