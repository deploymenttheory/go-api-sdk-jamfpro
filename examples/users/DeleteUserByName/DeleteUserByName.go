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

	userName := "AHarrison" // Example user name to delete

	err = client.DeleteUserByName(userName)
	if err != nil {
		log.Fatalf("Error deleting user by name: %v", err)
	}

	fmt.Println("User deleted successfully by name")
}
