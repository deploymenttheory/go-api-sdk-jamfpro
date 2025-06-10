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

	// Example task ID to delete
	taskID := "df705c19-c6ea-44c2-aea1-83bbb20cd85f"

	// Call function to delete log flushing task
	err = client.DeleteLogFlushingTaskByID(taskID)
	if err != nil {
		log.Fatalf("Error deleting log flushing task: %v", err)
	}

	fmt.Printf("Successfully deleted log flushing task with ID: %s\n", taskID)
}
