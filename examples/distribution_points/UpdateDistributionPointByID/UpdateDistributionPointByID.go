package main

import (
	"encoding/xml"
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
	logLevel := http_client.LogLevelDebug

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

	// New distribution point to create
	updateDistributionPoint := jamfpro.ResourceDistributionPoint{
		Name:                     "Tokyo Share",
		IPAddress:                "tokyo.company.com",
		IsMaster:                 true,
		EnableLoadBalancing:      false,
		SSHUsername:              "casperadmin",
		Password:                 "password",
		ConnectionType:           "SMB",
		ShareName:                "Caspershare",
		WorkgroupOrDomain:        "COMPANY",
		SharePort:                139,
		ReadOnlyUsername:         "casperinstall",
		ReadOnlyPassword:         "password",
		ReadWriteUsername:        "casperwrite",
		ReadWritePassword:        "password",
		HTTPDownloadsEnabled:     true,
		HTTPURL:                  "http://ny.company.com/CasperShare",
		Context:                  "CasperShare",
		Protocol:                 "http",
		Port:                     80,
		NoAuthenticationRequired: false,
		UsernamePasswordRequired: true,
		HTTPUsername:             "casperinstall",
		HTTPPassword:             "password",
	}

	// ID of the distribution point to update
	distributionPointID := 1 // Replace with the actual ID

	// Call UpdateDistributionPointByID function
	updatedDistributionPoint, err := client.UpdateDistributionPointByID(distributionPointID, &updateDistributionPoint)
	if err != nil {
		log.Fatalf("Error updating distribution point: %v", err)
	}

	// Pretty print the updated distribution point in XML
	updatedDistributionPointXML, err := xml.MarshalIndent(updatedDistributionPoint, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated distribution point data: %v", err)
	}
	fmt.Println("Updated Distribution Point:\n", string(updatedDistributionPointXML))
}
