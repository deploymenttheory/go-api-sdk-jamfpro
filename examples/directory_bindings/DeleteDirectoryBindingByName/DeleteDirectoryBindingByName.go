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

	// Example name of the directory binding to delete

	bindingName := "New Binding" // Assuming an existing binding name

	// Delete directory binding by Name
	err = client.DeleteDirectoryBindingByName(bindingName)
	if err != nil {
		fmt.Println("Error deleting directory binding by Name:", err)
		return
	}
	fmt.Println("Successfully deleted Directory Binding by Name")
}
