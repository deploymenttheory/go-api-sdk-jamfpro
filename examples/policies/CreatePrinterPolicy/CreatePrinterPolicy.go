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
		// General
		General: jamfpro.PolicySubsetGeneral{
			Name:                       "jamfpro-sdk-example-printer-policy-config",
			Enabled:                    jamfpro.FalsePtr(),
			TriggerCheckin:             jamfpro.FalsePtr(),
			TriggerEnrollmentComplete:  jamfpro.FalsePtr(),
			TriggerLogin:               jamfpro.FalsePtr(),
			TriggerLogout:              jamfpro.FalsePtr(),
			TriggerNetworkStateChanged: jamfpro.FalsePtr(),
			TriggerStartup:             jamfpro.FalsePtr(),
			TriggerOther:               "EVENT",
			Frequency:                  "Once per computer",
			RetryEvent:                 "none",
			RetryAttempts:              -1,
			NotifyOnEachFailedRetry:    jamfpro.FalsePtr(),
			LocationUserOnly:           jamfpro.FalsePtr(),
			TargetDrive:                "/",
			Offline:                    jamfpro.FalsePtr(),
		},
		// Printers
		Printers: jamfpro.PolicySubsetPrinters{
			LeaveExistingDefault: jamfpro.FalsePtr(),
			Printer: []jamfpro.PolicySubsetPrinter{
				{
					ID:          1740,
					Name:        "HP 9th Floor 2",
					Action:      "install", // install / uninstall
					MakeDefault: jamfpro.TruePtr(),
				},
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
