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

	// IDs of the buildings to be deleted
	buildingIDs := []string{"5", "6"} // Replace with the actual IDs of the buildings you want to delete

	err = client.DeleteMultipleBuildingsByID(buildingIDs)
	if err != nil {
		log.Fatalf("Error deleting multiple buildings: %v", err)
	}

	fmt.Println("Buildings deleted successfully")
}
