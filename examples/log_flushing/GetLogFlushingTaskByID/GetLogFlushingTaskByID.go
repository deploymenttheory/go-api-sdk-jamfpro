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

	// Example task ID
	taskID := "42c54d16-1492-45a7-bd21-704115bd0a1c"

	// Call function to get log flushing task by ID
	task, err := client.GetLogFlushingTaskByID(taskID)
	if err != nil {
		log.Fatalf("Error fetching log flushing task: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(task, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling log flushing task data: %v", err)
	}
	fmt.Println("Fetched log flushing task:\n", string(response))
}
