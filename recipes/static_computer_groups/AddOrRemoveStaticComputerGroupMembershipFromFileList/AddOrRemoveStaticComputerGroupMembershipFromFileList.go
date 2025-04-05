package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	// Config path
	configFilePath = "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Static group ID to update
	staticGroupID = "581"

	// Action required: "ADD" or "REMOVE"
	actionRequired = "ADD"

	// Path to file containing list of serial numbers
	serialNumberListPath = "~/Desktop/sourceList.txt"
)

// Read serial numbers from file
func readSerialNumbers(path string) ([]string, error) {
	// Expand ~ to home directory if needed
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("could not get home directory: %v", err)
		}
		path = filepath.Join(home, path[2:])
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open serial number file: %v", err)
	}
	defer file.Close()

	var serials []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		serial := strings.TrimSpace(scanner.Text())
		if serial != "" {
			serials = append(serials, serial)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading serial number file: %v", err)
	}

	if len(serials) == 0 {
		return nil, fmt.Errorf("no serial numbers found in file")
	}

	return serials, nil
}

// Add computers to static group
func addComputersToGroup(client *jamfpro.Client, groupID string, serials []string) ([]string, []string, error) {
	var successSerials []string
	var failureSerials []string

	// First, get the current group
	existingGroup, err := client.GetComputerGroupByID(groupID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get computer group: %v", err)
	}

	// For each serial, create a computer entry and add it to the group
	for _, serial := range serials {
		// Create the computer entry
		computer := jamfpro.ComputerGroupSubsetComputer{
			SerialNumber: serial,
		}

		// Create a modified copy of the existing group with this computer added
		// We need to make a copy to avoid modifying the original slice while iterating
		var updatedComputers []jamfpro.ComputerGroupSubsetComputer
		if existingGroup.Computers != nil {
			updatedComputers = append(updatedComputers, *existingGroup.Computers...)
		}
		updatedComputers = append(updatedComputers, computer)

		// Update the group with new computers
		existingGroup.Computers = &updatedComputers

		// Update the group in Jamf
		groupIDInt, _ := strconv.Atoi(groupID)
		existingGroup.ID = groupIDInt

		_, err := client.UpdateComputerGroupByID(groupID, existingGroup)
		if err != nil {
			failureSerials = append(failureSerials, serial)
			fmt.Printf("Failed to add serial %s: %v\n", serial, err)
			continue
		}

		successSerials = append(successSerials, serial)
		fmt.Printf("Successfully added serial: %s\n", serial)
	}

	return successSerials, failureSerials, nil
}

// Remove computers from static group
func removeComputersFromGroup(client *jamfpro.Client, groupID string, serials []string) ([]string, []string, error) {
	var successSerials []string
	var failureSerials []string

	// First, get the current group
	existingGroup, err := client.GetComputerGroupByID(groupID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get computer group: %v", err)
	}

	// If no computers in group, nothing to remove
	if existingGroup.Computers == nil || len(*existingGroup.Computers) == 0 {
		return nil, serials, nil
	}

	// For each serial, find and remove it from the group
	for _, serial := range serials {
		// Create new slice without the specified serial
		var updatedComputers []jamfpro.ComputerGroupSubsetComputer
		found := false

		for _, comp := range *existingGroup.Computers {
			if comp.SerialNumber != serial {
				updatedComputers = append(updatedComputers, comp)
			} else {
				found = true
			}
		}

		// If serial wasn't found, mark as failure
		if !found {
			failureSerials = append(failureSerials, serial)
			fmt.Printf("Serial %s not found in group\n", serial)
			continue
		}

		// Update the group with new computer list
		existingGroup.Computers = &updatedComputers

		// Update the group in Jamf
		groupIDInt, _ := strconv.Atoi(groupID)
		existingGroup.ID = groupIDInt

		_, err := client.UpdateComputerGroupByID(groupID, existingGroup)
		if err != nil {
			failureSerials = append(failureSerials, serial)
			fmt.Printf("Failed to remove serial %s: %v\n", serial, err)
			continue
		}

		successSerials = append(successSerials, serial)
		fmt.Printf("Successfully removed serial: %s\n", serial)
	}

	return successSerials, failureSerials, nil
}

func main() {
	// Initialize Jamf Pro client
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Read serial numbers
	serials, err := readSerialNumbers(serialNumberListPath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Validate action
	if actionRequired != "ADD" && actionRequired != "REMOVE" {
		log.Fatalf("Error: actionRequired must be 'ADD' or 'REMOVE', got '%s'", actionRequired)
	}

	fmt.Printf("Processing %d serial numbers...\n", len(serials))

	var successSerials, failureSerials []string

	// Process the serial numbers based on action
	if actionRequired == "ADD" {
		successSerials, failureSerials, err = addComputersToGroup(client, staticGroupID, serials)
	} else { // REMOVE
		successSerials, failureSerials, err = removeComputersFromGroup(client, staticGroupID, serials)
	}

	if err != nil {
		log.Fatalf("Error processing serial numbers: %v", err)
	}

	// Display summary
	fmt.Printf("\nSummary:\n")
	fmt.Printf("Success: %d Serial Numbers\n", len(successSerials))
	fmt.Printf("---------------------------------------\n")
	for _, serial := range successSerials {
		fmt.Println(serial)
	}

	fmt.Printf("\nFailed: %d Serial Numbers\n", len(failureSerials))
	fmt.Printf("---------------------------------------\n")
	for _, serial := range failureSerials {
		fmt.Println(serial)
	}

	// Get updated group information to show population
	computerGroup, err := client.GetComputerGroupByID(staticGroupID)
	if err != nil {
		log.Fatalf("Failed to get computer group details: %v", err)
	}

	// Pretty print the computer group
	fmt.Printf("\nUpdated Computer Group Details:\n")
	fmt.Printf("Group Name: %s\n", computerGroup.Name)
	fmt.Printf("Smart Group: %t\n", computerGroup.IsSmart)

	// Print computers in the group
	if computerGroup.Computers != nil && len(*computerGroup.Computers) > 0 {
		fmt.Printf("Computers in group: %d\n", len(*computerGroup.Computers))

		// Convert to JSON for easier reading
		jsonData, err := json.MarshalIndent(*computerGroup.Computers, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling computer group data: %v", err)
		}
		fmt.Printf("Computers: %s\n", string(jsonData))
	} else {
		fmt.Println("No computers in group")
	}

	fmt.Println("Process completed.")
}
