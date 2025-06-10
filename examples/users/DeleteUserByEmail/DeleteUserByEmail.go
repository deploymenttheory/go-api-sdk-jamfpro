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

	userEmail := "aharrison@company.com" // Example user email to delete

	err = client.DeleteUserByEmail(userEmail)
	if err != nil {
		log.Fatalf("Error deleting user by email: %v", err)
	}

	fmt.Println("User deleted successfully by email")
}
