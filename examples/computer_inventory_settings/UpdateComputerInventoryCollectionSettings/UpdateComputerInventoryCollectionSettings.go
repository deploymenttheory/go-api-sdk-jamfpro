package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the new settings (match current SDK: preferences type and applicationPaths only)
	// Build the settings payload: only include preferences in the PATCH.
	// Do NOT include applicationPaths with arbitrary IDs — the API will return INVALID_ID if
	// the ID(s) do not already exist. To add custom paths use the dedicated custom-path endpoint.
	newSettings := &jamfpro.ResourceComputerInventoryCollectionSettings{
		ComputerInventoryCollectionPreferences: jamfpro.ComputerInventoryCollectionSettingsSubsetPreferences{
			MonitorApplicationUsage:                      false,
			IncludePackages:                              true,
			IncludeSoftwareUpdates:                       false,
			IncludeSoftwareId:                            true,
			IncludeAccounts:                              true,
			CalculateSizes:                               false,
			IncludeHiddenAccounts:                        false,
			IncludePrinters:                              true,
			IncludeServices:                              true,
			CollectSyncedMobileDeviceInfo:                false,
			UpdateLdapInfoOnComputerInventorySubmissions: false,
			MonitorBeacons:                               false,
			AllowChangingUserAndLocation:                 true,
			UseUnixUserPaths:                             true,
			CollectUnmanagedCertificates:                 true,
		},
	}

	// Update computer inventory collection settings (API returns 204 No Content on success)
	_, err = client.UpdateComputerInventoryCollectionSettings(newSettings)
	if err != nil {
		log.Fatalf("Error updating Computer Inventory Collection Settings: %s", err)
	}

	// Example: add a custom application path using the custom-path endpoint
	// (Do this after the PATCH; creating custom paths is a separate API call.)
	newPath := &jamfpro.ResourceComputerInventoryCollectionSettingsCustomPath{
		Scope: "APP",
		Path:  "/Example/Path/To/App4/",
	}
	created, err := client.CreateComputerInventoryCollectionSettingsCustomPath(newPath)
	if err != nil {
		log.Fatalf("Error creating custom application path: %s", err)
	}
	createdJSON, _ := json.MarshalIndent(created, "", "    ")
	fmt.Println("Created custom application path:")
	fmt.Println(string(createdJSON))
}
