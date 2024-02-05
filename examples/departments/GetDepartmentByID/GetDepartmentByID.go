package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/joseph/github/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// New
	cid := "c018aae1-2b8c-4dcb-97b1-e4091ccda7c6"
	cs := "KToFqXe4KUynj4pKku3G_aN2MQeTp90RJYsK4IMrStR7Q7Ib-gM9NqiojlWqqJc7"

	// Joe
	// cid := "c5593021-b062-49df-a5fc-a32b54358480"
	// cs := "ZXoh0ooXQafidcfS2I_ZIsvfKbopjxmM3oUslxIuGi0jGXVxf-0IUy29-5sI2kxd"

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:             authConfig.InstanceName,
		OverrideBaseDomain:       authConfig.OverrideBaseDomain,
		LogLevel:                 logLevel,
		Logger:                   logger,
		ClientID:                 cid,
		ClientSecret:             cs,
		TokenRefreshBufferPeriod: 60 * time.Second,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the department ID you want to retrieve
	departmentID := "24979" // Replace with the desired department ID

	// Call GetDepartmentByID function
	department, err := client.GetDepartmentByID(departmentID)
	if err != nil {
		log.Fatalf("Error fetching department by ID: %v", err)
	}

	// Pretty print the department in JSON
	departmentJSON, err := json.MarshalIndent(department, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling department data: %v", err)
	}
	fmt.Println("Fetched Department:\n", string(departmentJSON))
}
