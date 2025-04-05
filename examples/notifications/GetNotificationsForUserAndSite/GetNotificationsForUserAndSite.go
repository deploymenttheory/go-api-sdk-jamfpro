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

	// Get notifications from Jamf Pro
	notifications, err := client.GetNotificationsForUserAndSite()
	if err != nil {
		log.Fatalf("Error retrieving notifications: %v", err)
	}

	// Pretty print the notifications in JSON format
	notificationsJSON, err := json.MarshalIndent(notifications, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling notifications: %v", err)
	}
	fmt.Println("Jamf Pro Notifications:\n", string(notificationsJSON))

	// Display count and summary of notifications
	fmt.Printf("\nTotal notifications: %d\n", len(*notifications))

	// Count notifications by type
	typeCount := make(map[string]int)
	for _, notification := range *notifications {
		typeCount[notification.Type]++
	}

	// Display notification counts by type
	fmt.Println("\nNotifications by type:")
	for notificationType, count := range typeCount {
		fmt.Printf("- %s: %d\n", notificationType, count)
	}
}
