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

	// Specify the device name
	deviceName := "APIASDFGHJ"

	// Step 1: Get the management ID by device name
	managementID, err := GetManagementIDByDeviceName(client, deviceName)
	if err != nil {
		log.Fatalf("Error retrieving management ID: %v", err)
	}
	log.Printf("Management ID for device %s: %s", deviceName, managementID)

	// Step 2: Lock the device using the management ID
	err = LockDevice(client, managementID)
	if err != nil {
		log.Fatalf("Error locking device: %v", err)
	}
	log.Printf("Device %s locked successfully.", deviceName)
}

// GetManagementIDByDeviceName retrieves the management ID for a device by its name.
func GetManagementIDByDeviceName(client *jamfpro.Client, deviceName string) (string, error) {
	inventories, err := client.GetComputersInventory("")
	if err != nil {
		return "", fmt.Errorf("failed to get computer inventory: %w", err)
	}

	for _, inventory := range inventories.Results {
		if inventory.General.Name == deviceName {
			return inventory.General.ManagementId, nil
		}
	}

	return "", fmt.Errorf("device with name %s not found", deviceName)
}

// LockDevice sends an MDM command to lock the device with the given management ID.
func LockDevice(client *jamfpro.Client, managementID string) error {
	mdmCommandRequest := &jamfpro.ResourceMDMCommandRequest{
		CommandData: jamfpro.CommandData{
			CommandType: "DeviceLock",
		},
		ClientData: []jamfpro.ClientData{
			{ManagementID: managementID},
		},
	}

	response, err := client.SendMDMCommandForCreationAndQueuing(mdmCommandRequest)
	if err != nil {
		return fmt.Errorf("failed to send MDM lock command: %w", err)
	}

	log.Printf("MDM command sent successfully. Response: %+v", response)
	return nil
}
