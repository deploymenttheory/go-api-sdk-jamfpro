// update_device_communication_settings.go
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

	// Create updated settings
	updatedSettings := jamfpro.ResourceDeviceCommunicationSettings{
		AutoRenewMobileDeviceMdmProfileWhenCaRenewed:                  true,
		AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring: true,
		AutoRenewComputerMdmProfileWhenCaRenewed:                      true,
		AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring:     true,
		MdmProfileMobileDeviceExpirationLimitInDays:                   180,
		MdmProfileComputerExpirationLimitInDays:                       180,
	}

	// Update device communication settings
	response, err := client.UpdateDeviceCommunicationSettings(updatedSettings)
	if err != nil {
		log.Fatalf("Error updating device communication settings: %v", err)
	}

	// Pretty print the updated device communication settings using JSON marshaling
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated device communication settings data: %v", err)
	}
	fmt.Println("Updated Device Communication Settings:", string(responseJSON))
}
