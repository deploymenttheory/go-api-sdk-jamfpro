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

	// Create the payload
	newVPL := jamfpro.ResourceVolumePurchasingLocation{
		// TODO I've messed something up here
	}

	// Call the CreateVolumePurchasingLocation function
	response, err := client.CreateVolumePurchasingLocation(&newVPL)
	if err != nil {
		log.Fatalf("Error creating volume purchasing location: %v", err)
	}

	// Print the response
	fmt.Printf("Created Volume Purchasing Location: %+v\n", response)
}
