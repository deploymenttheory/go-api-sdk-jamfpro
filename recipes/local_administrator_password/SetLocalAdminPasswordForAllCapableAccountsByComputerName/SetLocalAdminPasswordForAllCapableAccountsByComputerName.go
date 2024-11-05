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

	// Example usage
	computerName := "MyMacBook"            // Replace with actual computer name
	newPassword := "NewSecurePassword123!" // Replace with desired password

	err = SetLocalAdminPasswordForAllCapableAccountsByComputerName(client, computerName, newPassword)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// SetLocalAdminPasswordForAllCapableAccountsByComputerName sets the LAPS password for all capable accounts on a device using its computer name
func SetLocalAdminPasswordForAllCapableAccountsByComputerName(client *jamfpro.Client, computerName string, newPassword string) error {
	// First get the computer inventory to get the management ID
	inventory, err := client.GetComputerInventoryByName(computerName)
	if err != nil {
		return fmt.Errorf("failed to get computer inventory for %s: %v", computerName, err)
	}

	managementID := inventory.General.ManagementId
	if managementID == "" {
		return fmt.Errorf("no management ID found for computer %s", computerName)
	}

	// Get all LAPS capable accounts for this device
	capableAccounts, err := client.GetLocalAdminPasswordCapableAccountsByClientManagementID(managementID)
	if err != nil {
		return fmt.Errorf("failed to get LAPS capable accounts for device %s: %v", computerName, err)
	}

	// Create password list for all capable accounts
	var passwordList jamfpro.ResourceLapsPasswordList
	for _, account := range capableAccounts.Results {
		passwordList.LapsUserPasswordList = append(passwordList.LapsUserPasswordList,
			jamfpro.LapsUserPassword{
				Username: account.Username,
				Password: newPassword,
			},
		)
	}

	// Set the new passwords
	response, err := client.SetLocalAdminPasswordByClientManagementID(managementID, &passwordList)
	if err != nil {
		return fmt.Errorf("failed to set LAPS passwords: %v", err)
	}

	// Pretty print the response for verification
	prettyResponse, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling response: %v", err)
	}

	fmt.Printf("Successfully set LAPS passwords for computer '%s'\n", computerName)
	fmt.Printf("Management ID: %s\n", managementID)
	fmt.Printf("Response:\n%s\n", string(prettyResponse))

	return nil
}
