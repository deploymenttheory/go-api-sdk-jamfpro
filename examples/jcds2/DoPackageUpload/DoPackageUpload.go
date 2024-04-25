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

	// Specify the path to the file you want to upload
	filePath := "/Users/dafyddwatkins/localtesting/terraform/support_files/packages/microsoft-edge-121-0-2277-106.pkg"

	// Define package metadata
	packageData := jamfpro.ResourcePackage{
		Name:     "Microsoft Edge",
		Category: "Web Browsers",
		// Filename is a required field to assign the uploaded JCDS2 package to the Jamf Pro package reference.
		// Without this field the package stays in a pending state. Value must match the filename of the file being uploaded.
		Filename:                   "microsoft-edge-121-0-2277-106.pkg",
		Info:                       "Microsoft Edge Browser Package",
		Notes:                      "Version 121.0.2277.106. This package installs Microsoft Edge on macOS devices.",
		Priority:                   10,                                      // Set priority (lower number means higher priority)
		RebootRequired:             false,                                   // Set to true if a reboot is required after installation
		FillUserTemplate:           false,                                   // Set to true if the package should fill the user template
		FillExistingUsers:          false,                                   // Set to true if the package should fill existing user directories
		BootVolumeRequired:         true,                                    // Set to true if the package must be installed on the boot volume
		AllowUninstalled:           false,                                   // Set to true if the package can be uninstalled
		OSRequirements:             "macOS 10.15.x, macOS 11.x, macOS 12.x", // Specify OS requirements
		RequiredProcessor:          "",                                      // Specify if a particular processor is required, leave blank if no specific requirement
		SwitchWithPackage:          "",                                      // Specify package ID to switch with this package, leave blank if not applicable
		InstallIfReportedAvailable: false,                                   // Set to true to install the package even if it's reported as available
		ReinstallOption:            "Do Not Reinstall",                      // Specify reinstall option, possible values might include "Do Not Reinstall", "Reinstall on Same Version", or "Reinstall on Different Version"
		TriggeringFiles:            "",                                      // Specify triggering files, leave blank if not applicable
		SendNotification:           false,                                   // Set to true to send a notification when the package is deployed
	}

	// Call DoPackageUpload with the file path and package metadata
	fileResponse, packageResponse, err := client.DoPackageUpload(filePath, &packageData)
	if err != nil {
		log.Fatalf("Failed to upload package: %v", err)
	}

	// Combine responses into a single map and marshal
	combinedResponses := map[string]interface{}{
		"fileUploadResponse":      fileResponse,
		"packageCreationResponse": packageResponse,
	}

	combinedResponseBytes, err := json.Marshal(combinedResponses)
	if err != nil {
		log.Fatalf("Failed to marshal DoPackageUpload responses: %v", err)
	}

	// Print the response
	fmt.Println("Response:", string(combinedResponseBytes))

}
