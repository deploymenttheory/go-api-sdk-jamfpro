package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/lloyds/Documents/clientconfig.json"
	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define variables for pointer fields

	// List of distribution points to create
	updatedDistributionPoint := jamfpro.ResourceFileShareDistributionPoint{

		
		Name:       "example-updated-distribution-pointss",
		ServerName: "servername",
		Principal:  false,
		// If EnableLoadBalancing is true, BackupDistributionPointID needs to have the
		// valid ID if another distribution point
		EnableLoadBalancing:       true,
		BackupDistributionPointID: "108",
		FileSharingConnectionType: "SMB",
		HTTPSEnabled:              true,
		HTTPSPort:                 443,
		HTTPSSecurityType:         "USERNAME_PASSWORD",
		HTTPSContext:              "context",
		HTTPSUsername:             "username",
		HTTPSPassword:             "password",
		ShareName:                 "sharename",
		Workgroup:                 "workgroup",
		Port:                      443,
		ReadWriteUsername:         "username",
		ReadWritePassword:         "password",
		ReadOnlyUsername:          "username",
		ReadOnlyPassword:          "password",
		SSHUsername:               "username",
		SSHPassword:               "password",
		LocalPathToShare:          "parf",
		
	}

	targetDistributionPoint := "135"

	// Loop through the list and create each distribution point
	updateDistributionPointandLog(client, updatedDistributionPoint, targetDistributionPoint)


}

func updateDistributionPointandLog(client *jamfpro.Client, distributionPointUpdateData jamfpro.ResourceFileShareDistributionPoint, id string) {
	// Call CreateDistributionPoint function
	updatedDistributionPoint, err := client.UpdateDistributionPointByID(id, &distributionPointUpdateData)
	if err != nil {
		log.Fatalf("Error creating distribution point: %v, %v", err, distributionPointUpdateData)
	}

	// Pretty print the newly created distribution point in XML
	updatedDistributionPointJSON, err := json.MarshalIndent(updatedDistributionPoint, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created distribution point data: %v", err)
	}
	fmt.Println("Updated Distribution Point:\n", string(updatedDistributionPointJSON))
}
