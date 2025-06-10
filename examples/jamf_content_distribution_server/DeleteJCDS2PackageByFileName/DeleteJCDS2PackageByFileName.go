package main

import (
	"fmt"
	"log"

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

	// Define the name of the macOS Configuration Profile you want to delete
	packageFileName := "microsoft-edge-121-0-2277-106.pkg"

	// Call the DeleteMacApplicationByName function
	err = client.DeleteJCDS2PackageV2(packageFileName)
	if err != nil {
		log.Fatalf("Failed to delete JCDS 2.0 file with name '%s': %v", packageFileName, err)
	}

	fmt.Printf("JCDS 2.0 file with name '%s' deleted successfully\n", packageFileName)
}
