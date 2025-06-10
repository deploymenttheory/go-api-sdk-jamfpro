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

	// Specify the name of the user group to delete
	userGroupName := "Teachers" // Replace with the actual name

	// Call DeleteUserGroupByName
	err = client.DeleteUserGroupByName(userGroupName)
	if err != nil {
		fmt.Println("Error deleting user group:", err)
		return
	}

	fmt.Println("User group deleted successfully")
}
