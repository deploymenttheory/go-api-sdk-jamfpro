package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	maxConcurrentRequestsAllowed = 5 // Maximum allowed concurrent requests.
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
	scriptFilePath               = "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/scriptfile.sh" // Path to the script file
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Load script contents from a file
	file, err := os.Open(scriptFilePath)
	if err != nil {
		log.Fatalf("Failed to open script file: %v", err)
	}
	defer file.Close()

	scriptContents, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read script file: %v", err)
	}

	// Define a sample script for testing
	updatedScript := &jamfpro.ResourceScript{
		ID:             "3", // Assuming ID 194 for this example
		Name:           "Updated Sample Script",
		CategoryId:     "None",
		Info:           "Updated Script information",
		Notes:          "Updated Sample Script",
		Priority:       "After",
		OSRequirements: "string",
		ScriptContents: string(scriptContents),
	}

	// Call UpdateScriptByID function
	resultScript, err := client.UpdateScriptByID(updatedScript.ID, updatedScript)
	if err != nil {
		log.Fatalf("Error updating script: %v", err)
	}

	// Pretty print the updated script details in XML
	resultScriptXML, err := xml.MarshalIndent(resultScript, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated script data: %v", err)
	}
	fmt.Println("Updated Script Details:\n", string(resultScriptXML))
}
