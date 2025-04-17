package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	// Path to the Jamf Pro client configuration file
	configFilePath = "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Path for the output CSV file
	outputCSVPath = "/Users/dafyddwatkins/Documents/multi_computer_users.csv"
)

func main() {
	// Initialize the Jamf Pro client with the configuration file
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Get computer inventory with all data
	// For more information on how to add parameters to this request, see docs/url_queries.md
	params := url.Values{}
	inventoryList, err := client.GetComputersInventory(params)
	if err != nil {
		log.Fatalf("Failed to get computer inventory data: %v", err)
	}

	computerInventories := inventoryList.Results

	// Process the computer data to find users with multiple computers
	multiComputerUsers := findUsersWithMultipleComputers(computerInventories)

	// Output the number of users with more than one computer
	fmt.Printf("Number of users with more than one computer: %d\n", len(multiComputerUsers))

	// Output the multi-computer users as JSON
	jsonOutput, err := json.MarshalIndent(multiComputerUsers, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal multi-computer users to JSON: %v", err)
	}
	fmt.Println(string(jsonOutput))

	// Write the data to a CSV file
	err = writeMultiComputerUsersToCSV(multiComputerUsers, outputCSVPath)
	if err != nil {
		log.Fatalf("Failed to write CSV file: %v", err)
	}

	fmt.Printf("CSV output has been written to %s\n", outputCSVPath)
}

// findUsersWithMultipleComputers processes computer inventory data to identify users with multiple computers
func findUsersWithMultipleComputers(computers []jamfpro.ResourceComputerInventory) map[string][]string {
	// Map to store username -> list of computer names
	users := make(map[string][]string)

	// Process each computer
	for _, computer := range computers {
		// Get username from the UserAndLocation section
		username := computer.UserAndLocation.Username

		// Skip empty usernames
		if strings.TrimSpace(username) == "" {
			continue
		}

		// Add computer name to the user's list
		users[username] = append(users[username], computer.General.Name)
	}

	// Filter out users with only one computer
	multiComputerUsers := make(map[string][]string)
	for username, computers := range users {
		if len(computers) > 1 {
			multiComputerUsers[username] = computers
		}
	}

	return multiComputerUsers
}

// writeMultiComputerUsersToCSV writes the multi-computer users data to a CSV file
func writeMultiComputerUsersToCSV(multiComputerUsers map[string][]string, outputCSV string) error {
	// Create the CSV file
	file, err := os.Create(outputCSV)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	// Create the CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Find the maximum number of computers any user has
	maxComputers := 0
	for _, computers := range multiComputerUsers {
		if len(computers) > maxComputers {
			maxComputers = len(computers)
		}
	}

	// Create the header row
	header := []string{"Username"}
	for i := 0; i < maxComputers; i++ {
		header = append(header, fmt.Sprintf("Computer %d", i+1))
		header = append(header, fmt.Sprintf("Status %d", i+1))
	}
	header = append(header, "Notes")

	// Write the header row
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header to CSV: %w", err)
	}

	// Write data rows for each user
	for username, computers := range multiComputerUsers {
		row := []string{username}

		// Add computer names and empty status columns
		for i := 0; i < maxComputers; i++ {
			if i < len(computers) {
				row = append(row, computers[i]) // Computer name
				row = append(row, "")           // Empty status column
			} else {
				row = append(row, "") // Empty computer name
				row = append(row, "") // Empty status column
			}
		}

		// Add empty notes column
		row = append(row, "")

		// Write the user row
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write data row to CSV: %w", err)
		}
	}

	return nil
}
