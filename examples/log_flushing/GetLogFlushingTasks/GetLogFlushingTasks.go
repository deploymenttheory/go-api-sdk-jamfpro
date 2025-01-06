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

	// Call function to get log flushing tasks
	logFlushingTasks, err := client.GetLogFlushingTasks()
	if err != nil {
		log.Fatalf("Error fetching log flushing tasks: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(logFlushingTasks, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling log flushing tasks data: %v", err)
	}
	fmt.Println("Fetched log flushing tasks:\n", string(response))
}
