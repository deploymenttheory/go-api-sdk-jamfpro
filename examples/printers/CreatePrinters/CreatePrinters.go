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

	// Printer details to create
	newPrinter := &jamfpro.ResourcePrinter{
		Name:        "HP 9th Floor 2",
		Category:    "",
		URI:         "lpd://10.1.20.204/",
		CUPSName:    "HP_9th_Floor",
		Location:    "string",
		Model:       "HP LaserJet 500 color MFP M575",
		Info:        "string",
		Notes:       "string",
		MakeDefault: true,
		UseGeneric:  true,
		PPD:         "9th_Floor_HP.ppd",
		PPDPath:     "/System/Library/Frameworks/ApplicationServices.framework/Versions/A/Frameworks/PrintCore.framework/Resources/Generic.ppd",
		PPDContents: "string",
	}

	createdPrinter, err := client.CreatePrinter(newPrinter)
	if err != nil {
		fmt.Println("Error creating printer:", err)
		return
	}

	configXML, err := xml.MarshalIndent(createdPrinter, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created configuration to XML: %v", err)
	}

	fmt.Printf("Created Individual Disk Encryption Configuration:\n%s\n", configXML)
}
