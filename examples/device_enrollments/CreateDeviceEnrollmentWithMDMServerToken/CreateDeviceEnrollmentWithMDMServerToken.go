package main

import (
	"encoding/json"
	"fmt"
	"log"

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

	// Create token upload request
	tokenUpload := &jamfpro.ResourceDeviceEnrollmentTokenUpload{
		TokenFileName: "Acme MDM Token",                                                                                                                               // Optional: provide a filename
		EncodedToken:  "VTI5dFpTQnlZVzVrYjIwZ1ltbDBJRzltSUhSbGVIUWdkRzhnZFhObElHRnVaQ0J6WldVZ2FXWWdZVzU1YjI1bElHRmpkSFZoYkd4NUlIUnlhV1Z6SUhSdklHUmxZMjlrWlNCcGRBPT0=", // Replace with actual base64 encoded token
	}

	// Upload token and create device enrollment
	response, err := client.CreateDeviceEnrollmentWithMDMServerToken(tokenUpload)
	if err != nil {
		log.Fatalf("Error creating device enrollment with MDM Server token: %v", err)
	}

	// Pretty print the created device enrollment response using JSON marshaling
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling device enrollment creation data: %v", err)
	}
	fmt.Println("Created Device Enrollment:", string(responseJSON))
}
