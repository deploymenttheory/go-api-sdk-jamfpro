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

	// Name of the building to be deleted
	buildingName := "Apple Park" // Replace with the actual name of the building you want to delete

	err = client.DeleteBuildingByName(buildingName)
	if err != nil {
		log.Fatalf("Error deleting building: %v", err)
	}

	fmt.Println("Building deleted successfully")
}
