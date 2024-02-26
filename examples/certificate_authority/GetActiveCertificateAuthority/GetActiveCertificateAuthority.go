// get_active_certificate_authority_main.go
package main

import (
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

	// Retrieve the active certificate authority details
	activeCertAuth, err := client.GetActiveCertificateAuthority()
	if err != nil {
		log.Fatalf("Error retrieving active certificate authority: %s", err)
	}

	fmt.Printf("Active Certificate Authority: %+v\n", activeCertAuth)
}
