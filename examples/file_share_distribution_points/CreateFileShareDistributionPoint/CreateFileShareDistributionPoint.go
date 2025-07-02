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

	// // New distribution point to create
	// newDistributionPoint := jamfpro.ResourceFileShareDistributionPoint{
	// 	ID:                        "323",
	// 	Name:                      "distribution-point-example",
	// 	ServerName:                "servername",
	// 	Principal:                 false,
	// 	BackupDistributionPointID: "",
	// 	FileSharingConnectionType: "",
	// 	HTTPSEnabled:              true,
	// 	HTTPSPort:                 443,
	// 	HTTPSSecurityType:         "",
	// 	HTTPSContext:              "mebo",
	// 	HTTPSUsername:             "meep",
	// 	EnableLoadBalancing:       false,
	// 	ShareName:                 "sharename",
	// 	Workgroup:                 "workgroup",
	// 	Port:                      443,
	// 	ReadWriteUsername:         "username",
	// 	ReadWritePassword:         "password",
	// 	ReadOnlyUsername:          "username",
	// 	ReadOnlyPassword:          "password",
	// 	SSHUsername:               "username",
	// 	SSHPassword:               "password",
	// 	LocalPathToShare:          "parf",
	// }

	// New distribution point to create
	newDistributionPoint := jamfpro.ResourceFileShareDistributionPoint{
		// ID:                        "323",
		Name:       "distribution-point-examplesswwsssssw",
		ServerName: "servername",
		// Principal:                 false,
		// BackupDistributionPointID: "",
		FileSharingConnectionType: "NONE",
		HTTPSEnabled:              true,
		HTTPSPort:                 443,
		HTTPSSecurityType:         "NONE",
		// HTTPSContext:              "mebo",
		// HTTPSUsername:             "meep",
		// EnableLoadBalancing:       false,
		// ShareName:                 "sharename",
		// Workgroup:                 "workgroup",
		// Port:                      443,
		// ReadWriteUsername:         "username",
		// ReadWritePassword:         "password",
		// ReadOnlyUsername:          "username",
		// ReadOnlyPassword:          "password",
		// SSHUsername:               "username",
		// SSHPassword:               "password",
		// LocalPathToShare:          "parf",
	}

	// Call CreateDistributionPoint function
	createdDistributionPoint, err := client.CreateDistributionPoint(&newDistributionPoint)
	if err != nil {
		log.Fatalf("Error creating distribution point: %v, %v", err, newDistributionPoint)
	}

	// Pretty print the newly created distribution point in XML
	createdDistributionPointXML, err := json.MarshalIndent(createdDistributionPoint, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created distribution point data: %v", err)
	}
	fmt.Println("Created Distribution Point:\n", string(createdDistributionPointXML))
}
