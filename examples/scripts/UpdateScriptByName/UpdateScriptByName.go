package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	maxConcurrentRequestsAllowed = 5
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute

	// Embedded script
	embeddedScript = `echo "Sample script update by jamf sdk"`

	scriptNameToUpdate = "Embedded Sample Script"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	authConfig, err := http_client.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	sampleScript := &jamfpro.ResponseScript{
		Name:     scriptNameToUpdate,
		Category: "None",
		Filename: "string",
		Info:     "Script information",
		Notes:    "Sample Script",
		Priority: "Before",
		Parameters: jamfpro.Parameters{
			Parameter4:  "string",
			Parameter5:  "string",
			Parameter6:  "string",
			Parameter7:  "string",
			Parameter8:  "string",
			Parameter9:  "string",
			Parameter10: "string",
			Parameter11: "string",
		},
		OSRequirements: "string",
		ScriptContents: embeddedScript,
	}

	updatedScript, err := client.UpdateScriptByName(sampleScript)
	if err != nil {
		log.Fatalf("Error updating script by name: %v", err)
	}

	updatedScriptXML, err := xml.MarshalIndent(updatedScript, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated script data: %v", err)
	}
	fmt.Println("Updated Script Details:\n", string(updatedScriptXML))
}
