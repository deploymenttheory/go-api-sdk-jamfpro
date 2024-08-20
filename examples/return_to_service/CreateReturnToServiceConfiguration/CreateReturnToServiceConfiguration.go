package main

import (
	"encoding/json"
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

	// Create the Return to Service configuration
	config := jamfpro.ResourceReturnToServiceConfiguration{
		DisplayName:                   "thing",
		SsoForEnrollmentEnabled:       false,
		SsoBypassAllowed:              false,
		SsoEnabled:                    false,
		SsoForMacOsSelfServiceEnabled: false,
		TokenExpirationDisabled:       false,
		UserAttributeEnabled:          false,
		UserAttributeName:             " ",
		UserMapping:                   "USERNAME",
		EnrollmentSsoForAccountDrivenEnrollmentEnabled: false,
		GroupEnrollmentAccessEnabled:                   false,
		GroupAttributeName:                             "http://schemas.xmlsoap.org/claims/Group",
		GroupRdnKey:                                    " ",
		GroupEnrollmentAccessName:                      " ",
		IdpProviderType:                                "ADFS",
		OtherProviderTypeName:                          " ",
		MetadataSource:                                 "URL",
		SessionTimeout:                                 480,
		Title:                                          "Quantity",
		Description:                                    "How many of these would you like?",
		Priority:                                       1,
		WifiProfileID:                                  "3800",
	}

	// Call CreateReturnToServiceConfiguration function
	createdConfig, err := client.CreateReturnToServiceConfiguration(config)
	if err != nil {
		log.Fatalf("Error creating Return to Service configuration: %v", err)
	}

	// Pretty print the created configuration in JSON
	response, err := json.MarshalIndent(createdConfig, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created configuration data: %v", err)
	}
	fmt.Println("Created Return to Service configuration:\n", string(response))
}
