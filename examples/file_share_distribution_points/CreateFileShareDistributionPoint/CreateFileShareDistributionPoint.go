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
	distributionPoints := []jamfpro.ResourceFileShareDistributionPoint{
		{
			Name:       "example-distribution-point-min",
			ServerName: "servername",
			FileSharingConnectionType: "NONE",
			HTTPSEnabled:              true,
			HTTPSPort:                 443,
			HTTPSSecurityType:         "NONE",
		},
		{
			Name:       "example-distribution-point-max-smb",
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
		},
	}

	// Loop through the list and create each distribution point
	for _, distributionPoint := range distributionPoints {
		createDistributionPointandLog(client, distributionPoint)
	}

}

func createDistributionPointandLog(client *jamfpro.Client, newDistributionPoint jamfpro.ResourceFileShareDistributionPoint) {
	// Call CreateDistributionPoint function
	createdDistributionPoint, err := client.CreateDistributionPoint(&newDistributionPoint)
	if err != nil {
		log.Fatalf("Error creating distribution point: %v, %v", err, newDistributionPoint)
	}

	createdDistributionPointJSON, err := json.MarshalIndent(createdDistributionPoint, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created distribution point data: %v", err)
	}
	fmt.Println("Created Distribution Point:\n", string(createdDistributionPointJSON))
}
