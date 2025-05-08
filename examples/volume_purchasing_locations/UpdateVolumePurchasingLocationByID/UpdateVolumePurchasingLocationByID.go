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

	id := "14" // Replace with the actual ID of the volume purchasing location you want to update

	// Create the payload
	updatedVPL := &jamfpro.VolumePurchasingLocationCreateUpdateRequest{
		Name:                                  "Updated Test VPP Location",
		AutomaticallyPopulatePurchasedContent: true,
		SendNotificationWhenNoLongerAssigned:  false,
		AutoRegisterManagedUsers:              true,
		SiteID:                                "-1",                               // Replace with your actual site ID
		ServiceToken:                          "eyJleHBEYXRlIjoiMjAyNi0wNS0w....", // Replace with your actual service token
	}

	// Print the payload for debugging
	fmt.Printf("Sending payload: %+v\n", updatedVPL)

	// Call the CreateVolumePurchasingLocation function
	response, err := client.UpdateVolumePurchasingLocationByID(id, updatedVPL)
	if err != nil {
		log.Fatalf("Error updating volume purchasing location: %v", err)
	}

	// Print the response
	fmt.Printf("Updated Volume Purchasing Location: %+v\n", response)
}
