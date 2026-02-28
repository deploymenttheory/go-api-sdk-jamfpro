package main

import (
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/Shared/GitHub/deploymenttheory/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	fileName := "acs-pkg-policy-autoupdate-jamfconnectlogin-branding-1.1.pkg"

	if err := client.RefreshCloudDistributionPointInventoryV1(fileName); err != nil {
		log.Fatalf("Error refreshing Cloud Distribution Point inventory: %v", err)
	}

	log.Printf("Cloud Distribution Point inventory refresh triggered successfully for file: %s", fileName)
}
