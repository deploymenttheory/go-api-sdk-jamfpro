package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/tools/schema/helpers/file"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/tools/schema/helpers/schemaprocessing"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Download the schema
	schema, err := client.DownloadJamfProSchema()
	if err != nil {
		fmt.Printf("Error downloading schema: %v\n", err)
		return
	}

	// Parse the JSON schema to generate Go structs
	structs, err := schemaprocessing.ParseJSONSchema(schema)
	if err != nil {
		fmt.Printf("Error parsing JSON schema: %v\n", err)
		return
	}

	// Save the generated structs to a file
	filePath := "generated_structs.go"
	err = file.SaveStructsToFile(structs, filePath)
	if err != nil {
		fmt.Printf("Error saving structs to file: %v\n", err)
		return
	}

	fmt.Printf("Generated structs saved to %s\n", filePath)
}
