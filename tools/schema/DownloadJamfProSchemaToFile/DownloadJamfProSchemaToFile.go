// main.go
package main

import (
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

	// Define the path to save the schema JSON file
	schemaFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/schema/11.5/schema.json"

	// Call DownloadJamfProSchemaToFile function
	err = client.DownloadJamfProSchemaToFile(schemaFilePath)
	if err != nil {
		log.Fatalf("Error downloading schema: %v", err)
	}

	fmt.Printf("Schema successfully downloaded to %s\n", schemaFilePath)
}
