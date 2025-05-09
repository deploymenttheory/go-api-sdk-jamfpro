package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// prettyPrintJSON converts a struct to a pretty-printed JSON string
func prettyPrintJSON(prefix string, data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}
	fmt.Printf("%s:\n%s\n", prefix, string(jsonData))
}

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the SSO settings to update
	updateSettings := &jamfpro.ResourceSsoSettings{
		ConfigurationType: "SAML",
		OidcSettings: &jamfpro.OidcSettings{
			UserMapping: "USERNAME",
		},
		SamlSettings: &jamfpro.SamlSettings{
			TokenExpirationDisabled: false,
			UserAttributeEnabled:    false,
			UserAttributeName:       " ",
			UserMapping:             "USERNAME",
			GroupAttributeName:      "http://schemas.xmlsoap.org/claims/Group",
			GroupRdnKey:             " ",
			IdpProviderType:         "ADFS",
			IdpUrl:                  "https://example.idp.com/app/id/sso/saml/metadata",
			EntityId:                "saml/metadata",
			MetadataFileName:        "if MetadataSource is set to URL, remove this field",
			OtherProviderTypeName:   " ",
			FederationMetadataFile:  "WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09",
			MetadataSource:          "URL",
			SessionTimeout:          480,
		},
		SsoForEnrollmentEnabled:       false,
		SsoBypassAllowed:              false,
		SsoEnabled:                    false,
		SsoForMacOsSelfServiceEnabled: false,
		EnrollmentSsoForAccountDrivenEnrollmentEnabled: false,
		EnrollmentSsoConfig: &jamfpro.EnrollmentSsoConfig{
			Hosts: []string{
				"dev-12324233.okta.com",
				"example.okta.com",
			},
			ManagementHint: "",
		},
		GroupEnrollmentAccessEnabled: false,
		GroupEnrollmentAccessName:    " ",
	}

	// Update SSO settings
	fmt.Println("Updating SSO settings...")
	updatedSettings, err := client.UpdateSsoSettings(*updateSettings)
	if err != nil {
		fmt.Printf("Error updating SSO settings: %v\n", err)
		prettyPrintJSON("Request Body:", updateSettings)
		return
	}

	// Print the JSON response for the updated SSO settings
	jsonData, err := json.MarshalIndent(updatedSettings, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Printf("Updated SSO Settings: %s\n", jsonData)
}
