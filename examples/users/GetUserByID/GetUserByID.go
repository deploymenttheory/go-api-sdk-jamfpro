package main

import (
	"encoding/xml"
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

	// Specify the ID of the user you want to retrieve
	userID := 1 // Replace with the desired user ID

	// Call the GetUserByID function
	user, err := client.GetUserByID(userID)
	if err != nil {
		log.Fatalf("Error fetching user by ID: %v", err)
	}

	// Pretty print the user details in XML
	userXML, err := xml.MarshalIndent(user, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling user data: %v", err)
	}
	fmt.Println("User Details:\n", string(userXML))
}
