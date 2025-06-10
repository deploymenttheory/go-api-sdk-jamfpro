package main

import (
	"fmt"
	"log"
	"net/url"
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
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Ensure that the export directory exists
	if _, err := os.Stat(exportPath); os.IsNotExist(err) {
		err = os.MkdirAll(exportPath, 0750)
		if err != nil {
			log.Fatalf("Failed to create export directory: %v", err)
		}
	}

	// Get a list of all computer extension attributes
	// For more information on how to add parameters to this request, see docs/url_queries.md
	attributesList, err := client.GetComputerExtensionAttributes(url.Values{})
	if err != nil {
		log.Fatalf("Failed to fetch Computer Extension Attributes: %v", err)
	}

	for _, attribute := range attributesList.Results {
		// Log to indicate which extension attribute is being processed
		log.Printf("Processing extension attribute '%s' (ID: %s)\n", attribute.Name, attribute.ID)

		// Log the type of the attribute for debugging
		log.Printf("Attribute '%s' is of type '%s'\n", attribute.Name, attribute.InputType)

		// Check if the type is "SCRIPT"
		if attribute.InputType == "SCRIPT" {
			// Sanitize the attribute name to be used as a filename
			sanitizedFileName := sanitizeFileName(attribute.Name)

			// Export the script content
			scriptFileName := filepath.Join(exportPath, sanitizedFileName+".sh")

			// Log to indicate where the script intends to export the extension attribute
			log.Printf("Exporting computer extension attribute '%s' to: %s", attribute.Name, scriptFileName)

			// Check if file already exists
			if _, err := os.Stat(scriptFileName); err == nil {
				log.Printf("File %s already exists!", scriptFileName)
			} else if !os.IsNotExist(err) {
				log.Printf("Error checking if file exists: %v", err)
			}

			// Use os.Create to create or truncate the file for writing
			file, err := os.Create(scriptFileName)
			if err != nil {
				log.Printf("Failed to create file for attribute '%s': %v", attribute.Name, err)
				continue
			}
			defer file.Close()

			_, err = file.WriteString(attribute.ScriptContents)
			if err != nil {
				log.Printf("Failed to write script for attribute '%s' to file '%s': %v", attribute.Name, scriptFileName, err)
				continue
			}
			fmt.Printf("Exported script for attribute '%s' to file '%s'\n", attribute.Name, scriptFileName)
		} else {
			log.Printf("Attribute '%s' is not of type 'SCRIPT', skipping export.\n", attribute.Name)
		}

		if err != nil {
			log.Printf("Error encountered: %v", err)
		}
	}
}
