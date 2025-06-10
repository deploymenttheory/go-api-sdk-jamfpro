package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	// Create the Plist content as a string
	plistContent := `"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE plist PUBLIC \"-//Apple//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n<plist version=\"1.0\">\n  <dict/>\n</plist>`

	// Get the current directory where the program is running
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	// Read JSON schema from the file in the same directory
	schemaPath := filepath.Join(currentDir, "cis_lvl1.json")
	jsonSchemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("Failed to read JSON schema file: %v", err)
	}
	jsonSchema := string(jsonSchemaBytes)

	// Construct the ForcedSettings
	forcedSettings := jamfpro.ForcedSettings{
		Plist:        plistContent,
		JsonSchema:   jsonSchema,
		SchemaSource: "Custom",
		SchemaDomain: "org.cis_lvl1.audit",
	}

	// Construct the PayloadContentItem
	payloadItem := jamfpro.PayloadContentItem{
		PayloadType:      "com.apple.ManagedClient.preferences",
		Forced:           &forcedSettings,
		PreferenceDomain: "org.cis_lvl1.audit",
	}

	// Create the ResourceConfigProfile
	profile := jamfpro.ResourceConfigProfile{
		Level:          "SYSTEM",
		PayloadContent: []jamfpro.PayloadContentItem{payloadItem},
	}

	result, err := client.CreateConfigProfileWithCustomSettingsSchema(&profile)
	if err != nil {
		log.Fatalf("Failed to create configuration profile: %v", err)
	}

	// Print the UUID of the created profile
	fmt.Printf("Successfully created CIS Level 1 Audit Configuration Profile with UUID: %s\n", result.UUID)

	// Optional: Get the newly created profile to verify its content
	createdProfile, err := client.GetConfigProfileByPayloadUUID(result.UUID)
	if err != nil {
		log.Fatalf("Failed to get newly created profile: %v", err)
	}

	// Print some information about the created profile
	fmt.Printf("Profile UUID: %s\n", createdProfile.PayloadUUID)
	fmt.Printf("Profile Level: %s\n", createdProfile.Level)
	fmt.Printf("Number of payload items: %d\n", len(createdProfile.PayloadContent))
	fmt.Printf("First payload type: %s\n", createdProfile.PayloadContent[0].PayloadType)
	fmt.Printf("Preference domain: %s\n", createdProfile.PayloadContent[0].PreferenceDomain)
}
