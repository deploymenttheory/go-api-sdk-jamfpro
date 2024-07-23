package main

import (
	"encoding/xml"
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

	// New distribution point to create
	updateDistributionPoint := jamfpro.ResourceFileShareDistributionPoint{
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
	distributionPointID := "1" // Replace with the actual ID

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
