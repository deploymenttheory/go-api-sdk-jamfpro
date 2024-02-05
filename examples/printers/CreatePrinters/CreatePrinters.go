package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
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
