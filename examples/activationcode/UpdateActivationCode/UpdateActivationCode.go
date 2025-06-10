package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create the ResourceActivationCode struct
	activationCode := &jamfpro.ResourceActivationCode{
		OrganizationName: "Organization Name",
		Code:             "UW5M-xxxx-CNAP-TCDT-xxxx-DNTV-ZAGT-xxxx",
	}

	// Update the activation code
	err = client.UpdateActivationCode(activationCode)
	if err != nil {
		fmt.Printf("Error updating activation code: %v\n", err)
		return
	}

	fmt.Println("Activation code updated successfully")
}
