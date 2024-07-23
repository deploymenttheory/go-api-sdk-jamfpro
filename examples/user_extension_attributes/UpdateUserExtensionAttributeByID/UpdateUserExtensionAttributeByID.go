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

	// Prepare the attribute to update
	attributeToUpdate := &jamfpro.ResourceUserExtensionAttribute{
		Name:        "User Attributes",                    // Updated name
		Description: "Text field for logging custom data", // Updated description
		DataType:    "String",                             // Updated data type
		InputType: jamfpro.ResourceUserExtensionAttributeSubsetInputType{
			Type: "Text Field", // Updated input type
		},
	}

	// Example ID of the user extension attribute to rupdate
	attributeID := "1"

	// Update the attribute
	updatedAttribute, err := client.UpdateUserExtensionAttributeByID(attributeID, attributeToUpdate) // Use the correct ID
	if err != nil {
		log.Fatalf("Error updating user extension attribute: %v", err)
	}

	// Print the updated attribute
	fmt.Printf("Updated Attribute: %+v\n", updatedAttribute)
}
