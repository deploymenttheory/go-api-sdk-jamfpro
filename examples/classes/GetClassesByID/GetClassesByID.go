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

	classID := "1" // Replace with the actual ID of the class you want to fetch

	// Get class by ID
	class, err := client.GetClassByID(classID)
	if err != nil {
		log.Fatalf("Error fetching class by ID: %s\n", err)
	}

	fmt.Printf("Class ID: %v, Name: %s, Description: %s\n", strconv.Itoa(class.ID), class.Name, class.Description)
	// Output additional class details as needed
}
