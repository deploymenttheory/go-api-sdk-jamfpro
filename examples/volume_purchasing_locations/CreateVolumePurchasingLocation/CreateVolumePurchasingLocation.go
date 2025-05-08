package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create the payload
	newVPL := &jamfpro.VolumePurchasingLocationCreateUpdateRequest{
		Name:                                  "Test VPP Location",
		AutomaticallyPopulatePurchasedContent: true,
		SendNotificationWhenNoLongerAssigned:  false,
		AutoRegisterManagedUsers:              true,
		SiteID:                                "-1",                     // Replace with your actual site ID
		ServiceToken:                          "eyJleHBEYXRlIjoiMjA...", // Replace with your actual service token
	}

	// Print the payload for debugging
	fmt.Printf("Sending payload: %+v\n", newVPL)

	// Call the CreateVolumePurchasingLocation function
	response, err := client.CreateVolumePurchasingLocation(newVPL)
	if err != nil {
		log.Fatalf("Error creating volume purchasing location: %v", err)
	}

	// Print the response
	fmt.Printf("Created Volume Purchasing Location: %+v\n", response)
}
