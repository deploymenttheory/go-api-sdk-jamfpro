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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Define the package details
	pkg := jamfpro.ResourcePackage{
		Name:                       "Firefox.dmg",
		Category:                   "Unknown",
		Filename:                   "Firefox.dmg",
		Info:                       "string",
		Notes:                      "string",
		Priority:                   5,
		RebootRequired:             true,
		FillUserTemplate:           true,
		FillExistingUsers:          true,
		BootVolumeRequired:         true,
		AllowUninstalled:           true,
		OSRequirements:             "string",
		RequiredProcessor:          "None",
		SwitchWithPackage:          "Do Not Install",
		InstallIfReportedAvailable: true,
		ReinstallOption:            "Do Not Reinstall",
		TriggeringFiles:            "string",
		SendNotification:           true,
	}

	// Use the CreatePackage function with the package payload string
	response, err := client.CreatePackage(pkg)
	if err != nil {
		fmt.Println("Error creating package:", err)
		return
	}

	// Pretty print the created script details in XML
	packageXML, err := xml.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created script data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(packageXML))
}
