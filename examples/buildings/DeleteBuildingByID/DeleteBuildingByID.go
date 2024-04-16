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

	// ID of the building to be deleted
	buildingID := "1347" // Replace with the actual ID of the building you want to delete

	err = client.DeleteBuildingByID(buildingID)
	if err != nil {
		log.Fatalf("Error deleting building: %v", err)
	}

	fmt.Println("Building deleted successfully")
}
