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

	// Define the DELETE_USER command
	deleteUserCommand := &jamfpro.ResourceMDMCommandRequest{
		CommandData: jamfpro.CommandData{
			CommandType:    "DELETE_USER",
			UserName:       "Barry White",
			ForceDeletion:  true,
			DeleteAllUsers: false,
		},
		ClientData: []jamfpro.ClientData{
			{ManagementID: "aaaaaaaa-3f1e-4b3a-a5b3-ca0cd7430937"},
			{ManagementID: "aaaaaaaa-3f1e-4b3a-a5b3-ca0cd7430937"},
		},
	}

	// Send the DELETE_USER command
	response, err := client.SendMDMCommandForCreationAndQueuing(deleteUserCommand)
	if err != nil {
		log.Fatalf("Error sending MDM command: %v", err)
	}

	// Pretty print the response details in JSON
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling response data: %v", err)
	}
	fmt.Println("MDM command response Details:\n", string(responseJSON))
}
