package main

import (
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	packageID := "1"

	if err := client.FailCloudDistributionPointUploadV1(packageID); err != nil {
		log.Fatalf("Error marking cloud distribution point upload as failed: %v", err)
	}

	log.Printf("Cloud distribution point upload for package %s marked as failed.", packageID)
}
