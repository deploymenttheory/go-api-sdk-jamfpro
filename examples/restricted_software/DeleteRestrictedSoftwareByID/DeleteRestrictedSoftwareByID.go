package main

import (
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

	restrictedSoftwareID := "1" // Replace with the actual ID

	err = client.DeleteRestrictedSoftwareByID(restrictedSoftwareID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Restricted software deleted successfully.")
}
