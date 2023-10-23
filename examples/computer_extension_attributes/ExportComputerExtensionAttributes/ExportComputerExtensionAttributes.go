package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// path to export the scripts
const exportPath = "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/export"

func sanitizeFileName(fileName string) string {
	// Remove any characters that are not letters, numbers, hyphens, or underscores
	reg := regexp.MustCompile("[^a-zA-Z0-9-_]+")
	sanitized := reg.ReplaceAllString(fileName, "_")

	return strings.Trim(sanitized, "_")
}

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Get a list of all computer extension attributes
	attributesList, err := client.GetComputerExtensionAttributes()
	if err != nil {
		log.Fatalf("Failed to fetch Computer Extension Attributes: %v", err)
	}

	for _, attributeItem := range attributesList.Results {
		// Get details for each attribute
		attribute, err := client.GetComputerExtensionAttributeByID(attributeItem.ID)
		if err != nil {
			log.Printf("Failed to fetch details for attribute %s (ID: %d): %v", attributeItem.Name, attributeItem.ID, err)
			continue
		}

		// Check if the type is "Script"
		if attribute.InputType.Type == "Script" {
			// Sanitize the attribute name to be used as a filename
			sanitizedFileName := sanitizeFileName(attribute.Name)

			// Export the script content
			scriptFileName := filepath.Join(exportPath, sanitizedFileName+".sh")

			err = os.WriteFile(scriptFileName, []byte(attribute.InputType.Script), 0644)
			if err != nil {
				log.Printf("Failed to write script for attribute '%s' to file '%s': %v", attribute.Name, scriptFileName, err)
				continue
			}
			fmt.Printf("Exported script for attribute '%s' to file '%s'\n", attribute.Name, scriptFileName)
		}
	}
}
