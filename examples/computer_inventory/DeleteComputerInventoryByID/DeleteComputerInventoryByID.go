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

	// Example: Delete a computer inventory with a specific ID
	computerID := "9" // Replace with the actual computer ID

	err = client.DeleteComputerInventoryByID(computerID)
	if err != nil {
		fmt.Println("Error deleting computer inventory:", err)
	} else {
		fmt.Println("Computer inventory deleted successfully")
	}
}
