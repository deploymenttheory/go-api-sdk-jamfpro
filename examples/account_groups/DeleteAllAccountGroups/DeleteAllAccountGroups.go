package main

import (
	"fmt"
	"log"
	"strconv"

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

	// Fetch all accounts to retrieve account groups
	accounts, err := client.GetAccounts()
	if err != nil {
		log.Fatalf("Error fetching accounts: %v", err)
	}

	fmt.Println("Account groups fetched. Starting deletion process:")

	// Iterate over each group in the accounts list and delete
	for _, group := range accounts.Groups {
		fmt.Printf("Deleting account group ID: %d, Name: %s\n", group.ID, group.Name)

		// Using the function to delete account groups by ID
		err = client.DeleteAccountGroupByID(strconv.Itoa(group.ID))
		if err != nil {
			log.Printf("Error deleting account group ID %d: %v\n", group.ID, err)
			continue // Move to the next group if there's an error
		}

		fmt.Printf("Account group ID %d deleted successfully.\n", group.ID)
	}

	fmt.Println("Account group deletion process completed.")
}
