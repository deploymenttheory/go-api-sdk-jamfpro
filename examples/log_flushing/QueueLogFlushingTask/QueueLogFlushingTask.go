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

	// Create new log flushing task request
	newTask := &jamfpro.ResourceLogFlushingTask{
		Qualifier:           "policy", // The qualifier of the retention policy
		RetentionPeriod:     3,        // The period beyond which data will be flushed
		RetentionPeriodUnit: "MONTH",  // DAY, WEEK, MONTH, YEAR
	}

	// Call function to create log flushing task
	createdTask, err := client.QueueLogFlushingTask(newTask)
	if err != nil {
		log.Fatalf("Error creating log flushing task: %v", err)
	}

	// Pretty print the JSON response
	response, err := json.MarshalIndent(createdTask, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created task data: %v", err)
	}
	fmt.Println("Created log flushing task:\n", string(response))
}
