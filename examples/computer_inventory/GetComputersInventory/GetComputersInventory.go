package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	/*
		Fields allowed in the sort: general.name, udid, id, general.assetTag,
		general.jamfBinaryVersion, general.lastContactTime, general.lastEnrolledDate, general.lastCloudBackupDate,
		general.reportDate, general.remoteManagement.managementUsername,
		general.mdmCertificateExpiration, general.platform,
		hardware.make, hardware.model,
		operatingSystem.build, operatingSystem.supplementalBuildVersion, operatingSystem.rapidSecurityResponse, operatingSystem.name, operatingSystem.version,
		userAndLocation.realname,
		purchasing.lifeExpectancy, purchasing.warrantyDate

		Example: sort=udid:desc,general.name:asc.
	*/

	// Define your sorting criteria and section filters if needed
	// sortCriteria := []string{"udid:desc,general.name:asc"}        // Example sort criteria
	// sections := []string{"GENERAL", "DISK_ENCRYPTION", "STORAGE"} // Example sections

	// Call the GetComputersInventory function
	inventoryList, err := client.GetComputersInventory("")
	if err != nil {
		log.Fatalf("Error fetching computer inventory: %v", err)
	}

	// Pretty print the response
	prettyJSON, err := json.MarshalIndent(inventoryList, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %v", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}
