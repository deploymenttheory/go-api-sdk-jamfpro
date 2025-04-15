package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

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

	// Define a sort filter (you can modify this as needed)
	params := url.Values{}
	params.Add("sort", "date:desc")

	// Call GetJamfProtectHistory function
	history, err := client.GetJamfProtectHistory(params)
	if err != nil {
		log.Fatalf("Error fetching Jamf Protect history: %v", err)
	}

	// Print the total count of history entries
	fmt.Printf("Total Jamf Protect history entries: %d\n\n", history.TotalCount)

	// Print details of each history entry
	for _, entry := range history.Results {
		fmt.Printf("ID: %d\n", entry.ID)
		fmt.Printf("Username: %s\n", entry.Username)
		fmt.Printf("Date: %s\n", entry.Date)
		fmt.Printf("Note: %s\n", entry.Note)
		fmt.Printf("Details: %s\n", entry.Details)
		fmt.Println("--------------------")
	}

	// Optionally, you can also print the entire response as JSON
	historyJSON, err := json.MarshalIndent(history, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling Jamf Protect history to JSON: %v", err)
	}
	fmt.Println("Jamf Protect History (JSON):")
	fmt.Println(string(historyJSON))
}
