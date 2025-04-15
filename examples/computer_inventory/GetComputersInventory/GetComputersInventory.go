package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

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
	inventoryList, err := client.GetComputersInventory(url.Values{})
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
