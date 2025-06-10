package main

import (
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

	// Fetch all user groups
	userGroupsList, err := client.GetUserGroups()
	if err != nil {
		log.Fatalf("Error fetching user groups: %v", err)
	}

	fmt.Println("User groups fetched. Starting deletion process:")

	// Iterate over each user group and delete
	for _, userGroup := range userGroupsList.UserGroup {
		fmt.Printf("Deleting user group ID: %d, Name: %s\n", userGroup.ID, userGroup.Name)

		err = client.DeleteUserGroupByID(strconv.Itoa(userGroup.ID))
		if err != nil {
			log.Printf("Error deleting user group ID %d: %v\n", userGroup.ID, err)
			continue // Move to the next user group if there's an error
		}

		fmt.Printf("User group ID %d deleted successfully.\n", userGroup.ID)
	}

	fmt.Println("User group deletion process completed.")
}
