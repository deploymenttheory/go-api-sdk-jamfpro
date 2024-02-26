// Notes: https://learn.jamf.com/bundle/jamf-pro-documentation-current/page/Creating_and_Exporting_an_Institutional_Recovery_Key.html

package main

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"os"

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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Path to the certificate
	filePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/filevaultcertificate/FilevaultMaster-sdk.p12"

	// Read the contents of the file
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Encode the file contents to base64
	base64Encoded := base64.StdEncoding.EncodeToString(fileContents)

	institutionalConfig := &jamfpro.ResourceDiskEncryptionConfiguration{
		Name:                  "jamfpro-sdk-example-InstitutionalRecoveryKey-config",
		KeyType:               "Institutional",        // Institutional / Individual and Institutional
		FileVaultEnabledUsers: "Current or Next User", // Management Account / Current or Next User
		InstitutionalRecoveryKey: &jamfpro.DiskEncryptionConfigurationInstitutionalRecoveryKey{
			CertificateType: "PKCS12",
			Password:        "secretThing",
			Data:            base64Encoded,
		},
	}

	createdConfig, err := client.CreateDiskEncryptionConfiguration(institutionalConfig)
	if err != nil {
		log.Fatalf("Error creating Institutional Key Configuration: %v", err)
	}

	configXML, err := xml.MarshalIndent(createdConfig, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created configuration to XML: %v", err)
	}

	fmt.Printf("Created Institutional Disk Encryption Configuration:\n%s\n", configXML)
}
