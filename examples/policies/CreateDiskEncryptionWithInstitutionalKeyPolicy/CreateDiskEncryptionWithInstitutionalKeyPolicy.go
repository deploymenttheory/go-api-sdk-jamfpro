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

	// Define a new policy with all required fields
	newPolicy := &jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			Name:                       "jamfpro-sdk-example-diskencryptionWithInstitutionalKey-policy-config-10",
			Enabled:                    false,
			TriggerCheckin:             false,
			TriggerEnrollmentComplete:  false,
			TriggerLogin:               false,
			TriggerLogout:              false,
			TriggerNetworkStateChanged: false,
			TriggerStartup:             false,
			TriggerOther:               "EVENT",
			Frequency:                  "Once per computer",
			RetryEvent:                 "none",
			RetryAttempts:              -1,
			NotifyOnEachFailedRetry:    false,
			LocationUserOnly:           false,
			TargetDrive:                "/",
			Offline:                    false,
		},
		DiskEncryption: jamfpro.PolicySubsetDiskEncryption{
			Action:                        "apply",
			DiskEncryptionConfigurationID: 3,
			AuthRestart:                   false,
		},
		PackageConfiguration: jamfpro.PolicySubsetPackageConfiguration{
			Packages:          []jamfpro.PolicySubsetPackageConfigurationPackage{}, // Empty packages list
			DistributionPoint: "default",
		},
		AccountMaintenance: jamfpro.PolicySubsetAccountMaintenance{
			ManagementAccount: &jamfpro.PolicySubsetAccountMaintenanceManagementAccount{
				Action:                "doNotChange",
				ManagedPassword:       "",
				ManagedPasswordLength: 0,
			},
			OpenFirmwareEfiPassword: &jamfpro.PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword{
				OfMode:           "none",
				OfPassword:       "",
				OfPasswordSHA256: "",
			},
		},
	}

	policyXML, err := xml.MarshalIndent(newPolicy, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling policy data: %v", err)
	}
	fmt.Println("Policy Details to be Sent:\n", string(policyXML))

	// Call CreatePolicy function
	createdPolicy, err := client.CreatePolicy(newPolicy)
	if err != nil {
		log.Fatalf("Error creating policy: %v", err)
	}

	// Pretty print the created policy details in XML
	policyXML, err = xml.MarshalIndent(createdPolicy, "", "    ") // Indent with 4 spaces and use '='
	if err != nil {
		log.Fatalf("Error marshaling policy details data: %v", err)
	}
	fmt.Println("Created Policy Details:\n", string(policyXML))
}
