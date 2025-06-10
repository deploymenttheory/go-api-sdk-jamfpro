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

	// Specify the enrollment customization ID and panel ID
	customizationID := "22" // Replace with your actual customization ID
	panelID := "17"         // Replace with your actual panel ID

	// Prepare the updated text pane settings
	updatedTextPane := jamfpro.ResourceEnrollmentCustomizationTextPane{
		DisplayName:        "Updated Text Pane",
		Rank:               0,
		Title:              "Welcome to Enrollment",
		Body:               "This is the updated text content for the enrollment process.",
		Subtext:            "Please follow the instructions to complete enrollment.",
		BackButtonText:     "Back",
		ContinueButtonText: "Continue",
	}

	// Update the text prestage pane
	result, err := client.UpdateTextPrestagePaneByID(customizationID, panelID, updatedTextPane)
	if err != nil {
		log.Fatalf("Failed to update text prestage pane: %v", err)
	}

	// Pretty print the result in JSON
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling result: %v", err)
	}
	fmt.Println("Updated Text Prestage Pane:\n", string(prettyJSON))

	// Display individual fields as well
	fmt.Printf("\nText Pane Details:\n")
	fmt.Printf("ID: %d\n", result.ID)
	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Display Name: %s\n", result.DisplayName)
	fmt.Printf("Rank: %d\n", result.Rank)
	fmt.Printf("Title: %s\n", result.Title)
	fmt.Printf("Body: %s\n", result.Body)
	fmt.Printf("Subtext: %s\n", result.Subtext)
	fmt.Printf("Back Button Text: %s\n", result.BackButtonText)
	fmt.Printf("Continue Button Text: %s\n", result.ContinueButtonText)
}
