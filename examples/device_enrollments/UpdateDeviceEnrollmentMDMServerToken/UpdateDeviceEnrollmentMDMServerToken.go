package main

import (
	"encoding/json"
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

	// Create token update request
	tokenUpload := &jamfpro.ResourceDeviceEnrollmentTokenUpload{
		TokenFileName: "updated_token.p7m",        // Optional: provide a filename
		EncodedToken:  "base64EncodedTokenString", // Replace with actual base64 encoded token
	}

	// Update device enrollment token
	deviceEnrollmentID := "1" // Using the known device enrollment ID
	response, err := client.UpdateDeviceEnrollmentMDMServerToken(deviceEnrollmentID, tokenUpload)
	if err != nil {
		log.Fatalf("Error updating device enrollment token: %v", err)
	}

	// Pretty print the update response using JSON marshaling
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling device enrollment token update data: %v", err)
	}
	fmt.Println("Updated Device Enrollment Token:", string(responseJSON))
}
