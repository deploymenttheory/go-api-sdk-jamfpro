package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/modules"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	// Specify the device name
	deviceName = "APIASDFGHJ"

	// Operation type: "lock" or "unlock"
	operation = "unlock"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Step 1: Get the management ID by device name
	managementID, err := modules.GetManagementIDByDeviceName(client, deviceName)
	if err != nil {
		log.Fatalf("Error retrieving management ID: %v", err)
	}
	log.Printf("Management ID for device %s: %s", deviceName, managementID)

	// Step 2: Process the operation based on type
	switch operation {
	case "lock":
		// Generate a PIN for locking
		pin := modules.GenerateRandomRecoveryLockPassword()
		log.Printf("Generated PIN: %s", pin)

		err = LockDevice(client, managementID, pin)
		if err != nil {
			log.Fatalf("Error locking device: %v", err)
		}
		log.Printf("Device %s locked successfully.", deviceName)

	case "unlock":
		err = UnlockDevice(client, managementID)
		if err != nil {
			log.Fatalf("Error unlocking device: %v", err)
		}
		log.Printf("Device %s unlock command sent successfully.", deviceName)

	default:
		log.Fatalf("Invalid operation: %s. Use 'lock' or 'unlock'", operation)
	}
}

// LockDevice sends an MDM command to lock the device with the given management ID.
func LockDevice(client *jamfpro.Client, managementID, pin string) error {
	mdmCommandRequest := &jamfpro.ResourceMDMCommandRequest{
		CommandData: jamfpro.CommandData{
			CommandType: "DeviceLock",
		},
		ClientData: []jamfpro.ClientData{
			{ManagementID: managementID},
		},
	}

	// Add a PIN code to the lock command if provided
	if pin != "" {
		mdmCommandRequest.CommandData.PIN = pin
	}

	response, err := client.SendMDMCommandForCreationAndQueuing(mdmCommandRequest)
	if err != nil {
		return fmt.Errorf("failed to send MDM lock command: %w", err)
	}

	log.Printf("MDM command sent successfully. Response: %+v", response)
	return nil
}

// UnlockDevice sends an MDM command to unlock the device with the given management ID.
func UnlockDevice(client *jamfpro.Client, managementID string) error {
	// For many MDM platforms, the "DeviceUnlock" command removes the pin/passcode
	mdmCommandRequest := &jamfpro.ResourceMDMCommandRequest{
		CommandData: jamfpro.CommandData{
			CommandType: "DeviceUnlock",
		},
		ClientData: []jamfpro.ClientData{
			{ManagementID: managementID},
		},
	}

	response, err := client.SendMDMCommandForCreationAndQueuing(mdmCommandRequest)
	if err != nil {
		return fmt.Errorf("failed to send MDM unlock command: %w", err)
	}

	log.Printf("MDM command sent successfully. Response: %+v", response)
	return nil
}
