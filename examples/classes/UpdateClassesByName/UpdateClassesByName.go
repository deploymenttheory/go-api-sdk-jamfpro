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

	// Configuration for the Jamf Pro API client
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new Jamf Pro API client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the class details you want to update by its name.
	classToUpdate := &jamfpro.ResponseClasses{
		Name:        "Math 101", // The name of the class you want to update.
		Description: "Introduction to advanced algebra",
		Site: jamfpro.ClassSite{
			ID:   -1,
			Name: "None",
		},
		Teachers: []jamfpro.ClassTeacher{
			{Teacher: "John Doe"},
			{Teacher: "Jane Smith"},
		},
		TeacherIDs: []jamfpro.ClassTeacherID{
			{ID: 123}, // Assume 123 is the ID of John Doe
			{ID: 456}, // Assume 456 is the ID of Jane Smith
		},
		// ... include other fields as necessary ...
	}

	// Call the update function with the class name and the new details.
	err = client.UpdateClassesByName(classToUpdate.Name, classToUpdate)
	if err != nil {
		log.Fatalf("Error updating class by name: %v", err)
	} else {
		fmt.Println("Class updated successfully by name.")
	}

	// If you need to check the updated details, perform a GetClassesByName call here and print or log the results.
	// For example:
	updatedClass, err := client.GetClassesByName(classToUpdate.Name)
	if err != nil {
		log.Fatalf("Error retrieving updated class by name: %v", err)
	}
	classXML, err := xml.MarshalIndent(updatedClass, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling updated class to XML: %v", err)
	}
	fmt.Printf("Updated Class Details by Name:\n%s\n", classXML)
}
