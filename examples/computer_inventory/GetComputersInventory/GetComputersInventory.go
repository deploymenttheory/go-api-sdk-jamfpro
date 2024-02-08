package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
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
