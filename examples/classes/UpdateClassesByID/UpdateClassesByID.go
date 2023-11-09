package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Define the class details you want to update.
	classToUpdate := &jamfpro.ResponseClasses{
		ID:          1, // The ID of the class you want to update.
		Source:      "N/A",
		Name:        "Math 101",
		Description: "Introduction to advanced algebra",
		Site: jamfpro.ClassSite{
			ID:   -1,
			Name: "None",
		},
		Teachers: []jamfpro.ClassTeacher{
			{Teacher: "John Doe"},
			{Teacher: "Jane Smith"},
		},
		// ... include other fields as necessary ...
	}

	// Call the update function with the class ID and the new details.
	err = client.UpdateClassesByID(classToUpdate.ID, classToUpdate)
	if err != nil {
		log.Fatalf("Error updating class: %v", err)
	} else {
		fmt.Println("Class updated successfully.")
	}

	// If you need to check the updated details, perform a GetClassByID call here and print or log the results.
	// For example:
	updatedClass, err := client.GetClassesByID(classToUpdate.ID)
	if err != nil {
		log.Fatalf("Error retrieving updated class: %v", err)
	}
	classXML, _ := xml.MarshalIndent(updatedClass, "", "  ")
	fmt.Printf("Updated Class Details:\n%s\n", classXML)
}
