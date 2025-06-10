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

	// New directory binding data
	newBinding := &jamfpro.ResponseDirectoryBinding{
		Name:       "New Binding",
		Priority:   1,
		Domain:     "example.com",
		Username:   "user@example.com",
		Password:   "password",
		ComputerOU: "CN=Computers,DC=example,DC=com",
		Type:       "Active Directory",
	}

	// Create new directory binding
	createdBinding, err := client.CreateDirectoryBinding(newBinding)
	if err != nil {
		fmt.Println("Error creating directory binding:", err)
		return
	}

	// Pretty print the created directory binding in xml
	createdBindingXML, err := xml.MarshalIndent(createdBinding, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created binding data: %v", err)
	}
	fmt.Printf("Created Directory Binding:\n%s\n", string(createdBindingXML))
}
