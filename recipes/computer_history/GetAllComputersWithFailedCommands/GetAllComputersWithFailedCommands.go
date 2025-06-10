package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

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

	// Call the GetComputers method
	computers, err := client.GetComputers()
	if err != nil {
		log.Fatalf("Error fetching computers: %v", err)
	}

	foundFailedCommands := false // Flag to check if any computer has failed commands

	// Iterate through each computer to fetch its history
	for _, computer := range computers.Results {
		// Fetch computer history by ID
		computerHistory, err := client.GetComputerHistoryByComputerID(strconv.Itoa(computer.ID))
		if err != nil {
			log.Printf("Error fetching computer history for ID %d: %v", computer.ID, err)
			continue
		}

		// Check if the computer has any failed commands and no pending or completed commands
		if len(computerHistory.Commands.Failed) > 0 && len(computerHistory.Commands.Completed) == 0 && len(computerHistory.Commands.Pending) == 0 {
			foundFailedCommands = true // Set the flag to true as we found a computer with failed commands
			prettyXML, err := xml.MarshalIndent(computerHistory, "", "    ")
			if err != nil {
				log.Printf("Failed to generate pretty XML for computer ID %d: %v", computer.ID, err)
				continue
			}
			fmt.Printf("Computer ID: %d - With Failed MDM Commands \n", computer.ID)
			fmt.Printf("%s\n", prettyXML)
		}
	}

	// After the loop, check if no failed commands were found
	if !foundFailedCommands {
		fmt.Println("No failed failed MDM commands were found for any computer.")
	}
}
