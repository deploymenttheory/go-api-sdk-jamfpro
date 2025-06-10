package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Printer details to create
	updatedPrinter := &jamfpro.ResourcePrinter{
		Name:        "HP 9th Floor",
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

	id := "1" // Replace with the actual printer ID

	response, err := client.UpdatePrinterByID(id, updatedPrinter)
	if err != nil {
		fmt.Println("Error updating printer:", err)
		return
	}

	fmt.Printf("Printer updated successfully, ID: %d\n", response.ID)

	// Fetch the full details of the updated printer
	updatedPrinterDetails, err := client.GetPrinterByID(strconv.Itoa(response.ID))
	if err != nil {
		fmt.Println("Error fetching updated printer details:", err)
		return
	}

	// Marshal the updated printer details to XML for display
	printerXML, err := xml.MarshalIndent(updatedPrinterDetails, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated printer to XML: %v", err)
	}

	fmt.Printf("Updated Printer Details:\n%s\n", printerXML)
}
