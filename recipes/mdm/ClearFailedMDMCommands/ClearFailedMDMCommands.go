package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/modules"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	// Config path - can be passed as an argument or set here
	configFilePath = "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
)

func main() {
	// Check JSS connection
	if err := modules.CheckJamfProConnection(); err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Get JSS URL from preferences
	jssURL, err := modules.GetJamfProURL()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Using JSS URL: %s\n", jssURL)

	// Initialize the Jamf Pro client
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Get hardware UUID
	macUUID, err := modules.GetHardwareUUIDFromSystemProfiler()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Hardware UUID: %s\n", macUUID)

	// Get computers inventory to find computer by UUID
	inventory, err := client.GetComputersInventory("")
	if err != nil {
		log.Fatalf("Error fetching computer inventory: %v", err)
	}

	// Find computer with matching UUID
	var computerID string
	var computerName string

	for _, comp := range inventory.Results {
		if comp.UDID == macUUID {
			computerID = comp.ID
			computerName = comp.General.Name
			break
		}
	}

	if computerID == "" {
		log.Fatalf("No computer found with UUID: %s", macUUID)
	}

	fmt.Printf("Found computer: %s (ID: %s)\n", computerName, computerID)

	// Check failed MDM commands by checking the computer management
	// Create MDM command to check status
	commandData := jamfpro.CommandData{
		CommandType: "DeviceInformation",
	}

	clientData := jamfpro.ClientData{
		ManagementID: computerID, // Use the management ID from the inventory
	}

	mdmRequest := &jamfpro.ResourceMDMCommandRequest{
		CommandData: commandData,
		ClientData:  []jamfpro.ClientData{clientData},
	}

	// Send a test command to see if it gets queued properly
	response, err := client.SendMDMCommandForCreationAndQueuing(mdmRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Previous MDM command failed") {
			fmt.Println("Failed MDM commands detected. Clearing...")
			if err := client.ClearFailedComputerMDMCommandsByComputerID(computerID); err != nil {
				log.Fatalf("Error clearing failed commands: %v", err)
			}
		} else {
			log.Fatalf("Error checking MDM status: %v", err)
		}
	} else {
		fmt.Printf("Test command sent successfully. No failed commands detected. Command ID: %s\n", response.ID)
	}

	fmt.Println("Process completed successfully")
}
