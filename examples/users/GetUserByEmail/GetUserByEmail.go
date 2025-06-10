package main

import (
	"encoding/xml"
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

	// Specify the email of the user you want to retrieve
	email := "aharrison@company.com" // Replace with the desired email address

	// Call the GetUserByEmail function
	user, err := client.GetUserByEmail(email)
	if err != nil {
		log.Fatalf("Error fetching user by email: %v", err)
	}

	// Pretty print the user details in XML
	userXML, err := xml.MarshalIndent(user, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling user data: %v", err)
	}
	fmt.Println("User Details:\n", string(userXML))
}
