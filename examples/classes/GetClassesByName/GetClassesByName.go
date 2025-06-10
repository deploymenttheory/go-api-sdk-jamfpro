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

	className := "Math 101" // Replace with the actual name of the class you want to fetch

	// Get class by name
	class, err := client.GetClassByName(className)
	if err != nil {
		log.Fatalf("Error fetching class by name: %s\n", err)
	}

	fmt.Printf("Class Name: %s, ID: %d, Description: %s\n", class.Name, class.ID, class.Description)
	// Output additional class details as needed
}
