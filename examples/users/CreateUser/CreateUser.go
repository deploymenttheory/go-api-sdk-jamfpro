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

	// Create a sample user to be created
	newUser := &jamfpro.ResourceUser{
		ID:           1,
		Name:         "AHarrison",
		FullName:     "Ashley Harrison",
		Email:        "aharrison@company.com",
		EmailAddress: "aharrison@company.com",
		PhoneNumber:  "123-555-6789",
		Position:     "Teacher",
		Sites: []jamfpro.SharedResourceSite{
			{
				ID:   -1,
				Name: "None",
			},
		},
	}

	// Call the CreateUser function
	createdUser, err := client.CreateUser(newUser)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	// Print the details of the created user
	fmt.Printf("Created User Details:\nID: %d\nName: %s\nEmail: %s\n", createdUser.ID, createdUser.Name, createdUser.Email)
}
