package jamfpro_test

import (
	"os"
	"testing"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const helloWorldScript = `#!/bin/bash\n` +
	`echo "Hello World"\n`

// TestCreateBasicComputerExtensionAttribute tests the creation of a basic computer extension attribute
func TestCreatePopUpMenuComputerExtensionAttribute_Basic(t *testing.T) {
	client := getClient(t)

	// Define a basic computer extension attribute
	attribute := &jamfpro.ComputerExtensionAttributeResponse{
		Name:      "Basic Pop Up Menu Test",
		InputType: jamfpro.ComputerExtensionAttributeInputType{Type: "Pop Up Menu", Choices: []string{"Option 1", "Option 2"}},
	}

	// Call CreateComputerExtensionAttribute function
	_, err := client.CreateComputerExtensionAttribute(attribute)
	if err != nil {
		t.Errorf("Error creating basic Computer Extension Attribute: %v", err)
	}
}

// TestCreateDetailedComputerExtensionAttribute tests the creation of a detailed computer extension attribute
func TestCreatePopUpMenuComputerExtensionAttribute_Full(t *testing.T) {
	client := getClient(t)

	// Define a detailed computer extension attribute
	attribute := &jamfpro.ComputerExtensionAttributeResponse{
		Name:             "Detailed Pop Up Menu Test",
		Description:      "A detailed pop up menu for testing",
		DataType:         "String",
		InputType:        jamfpro.ComputerExtensionAttributeInputType{Type: "Pop Up Menu", Choices: []string{"Choice 1", "Choice 2", "Choice 3"}},
		InventoryDisplay: "General",
		ReconDisplay:     "Extension Attributes",
	}

	// Call CreateComputerExtensionAttribute function
	_, err := client.CreateComputerExtensionAttribute(attribute)
	if err != nil {
		t.Errorf("Error creating detailed Computer Extension Attribute: %v", err)
	}
}

// TestCreateBasicScriptComputerExtensionAttribute tests the creation of a basic script-based computer extension attribute
func TestCreateScriptComputerExtensionAttribute_Basic(t *testing.T) {
	client := getClient(t)

	// Define a basic computer extension attribute using the embedded script
	attribute := &jamfpro.ComputerExtensionAttributeResponse{
		Name:     "Basic Script Test",
		DataType: "String",
		InputType: jamfpro.ComputerExtensionAttributeInputType{
			Type:   "Script",
			Script: helloWorldScript,
		},
	}

	// Call CreateComputerExtensionAttribute function
	_, err := client.CreateComputerExtensionAttribute(attribute)
	if err != nil {
		t.Errorf("Error creating basic script-based Computer Extension Attribute: %v", err)
	}
}

// TestCreateDetailedScriptComputerExtensionAttribute tests the creation of a detailed script-based computer extension attribute
func TestCreateScriptComputerExtensionAttribute_Full(t *testing.T) {
	client := getClient(t)

	// Define a detailed computer extension attribute using the embedded script
	attribute := &jamfpro.ComputerExtensionAttributeResponse{
		Name:        "Detailed Script Test",
		Description: "A detailed script for testing",
		DataType:    "String",
		InputType: jamfpro.ComputerExtensionAttributeInputType{
			Type:     "Script",
			Script:   helloWorldScript,
			Platform: "Mac",
		},
		InventoryDisplay: "General",
		ReconDisplay:     "Extension Attributes",
	}

	// Call CreateComputerExtensionAttribute function
	_, err := client.CreateComputerExtensionAttribute(attribute)
	if err != nil {
		t.Errorf("Error creating detailed script-based Computer Extension Attribute: %v", err)
	}
}

// TestCreateBasicTextComputerExtensionAttribute tests the creation of a basic text-based computer extension attribute
func TestCreateTextComputerExtensionAttribute_Basic(t *testing.T) {
	client := getClient(t)

	// Define a basic computer extension attribute with a text field
	attribute := &jamfpro.ComputerExtensionAttributeResponse{
		Name:     "Basic Text Field Test",
		DataType: "String",
		InputType: jamfpro.ComputerExtensionAttributeInputType{
			Type: "Text Field",
		},
	}

	// Call CreateComputerExtensionAttribute function
	_, err := client.CreateComputerExtensionAttribute(attribute)
	if err != nil {
		t.Errorf("Error creating basic text-based Computer Extension Attribute: %v", err)
	}
}

// TestCreateDetailedTextComputerExtensionAttribute tests the creation of a detailed text-based computer extension attribute
func TestCreateTextComputerExtensionAttribute_Full(t *testing.T) {
	client := getClient(t)

	// Define a detailed computer extension attribute with a text field
	attribute := &jamfpro.ComputerExtensionAttributeResponse{
		Name:        "Detailed Text Field Test",
		Description: "A detailed text field for testing",
		DataType:    "String",
		InputType: jamfpro.ComputerExtensionAttributeInputType{
			Type: "Text Field",
		},
		InventoryDisplay: "General",
		ReconDisplay:     "Extension Attributes",
	}

	// Call CreateComputerExtensionAttribute function
	_, err := client.CreateComputerExtensionAttribute(attribute)
	if err != nil {
		t.Errorf("Error creating detailed text-based Computer Extension Attribute: %v", err)
	}
}

// TestGetComputerExtensionAttributes tests the retrieval of all computer extension attributes.
func TestGetComputerExtensionAttributes(t *testing.T) {
	client := getClient(t)

	attributes, err := client.GetComputerExtensionAttributes()
	if err != nil {
		t.Errorf("Error retrieving Computer Extension Attributes: %v", err)
	}

	if len(attributes.Results) == 0 {
		t.Error("No Computer Extension Attributes found")
	}
}

// TestGetComputerExtensionAttributeByID tests the retrieval of a computer extension attribute by its ID.
func TestGetComputerExtensionAttributeByID(t *testing.T) {
	client := getClient(t)

	// First, retrieve all attributes to get a valid ID
	allAttributes, err := client.GetComputerExtensionAttributes()
	if err != nil {
		t.Fatalf("Error retrieving all Computer Extension Attributes: %v", err)
	}
	if len(allAttributes.Results) == 0 {
		t.Fatal("No Computer Extension Attributes found for testing")
	}

	// Use the ID of the first attribute for testing
	testID := allAttributes.Results[0].ID

	attribute, err := client.GetComputerExtensionAttributeByID(testID)
	if err != nil {
		t.Errorf("Error retrieving Computer Extension Attribute by ID: %v", err)
	}

	if attribute == nil || attribute.ID != testID {
		t.Errorf("Attribute not found or ID mismatch. Expected ID: %d", testID)
	}
}

// TestGetComputerExtensionAttributeByName tests the retrieval of a computer extension attribute by its name.
func TestGetComputerExtensionAttributeByName(t *testing.T) {
	client := getClient(t)

	// First, retrieve all attributes to get a valid name
	allAttributes, err := client.GetComputerExtensionAttributes()
	if err != nil {
		t.Fatalf("Error retrieving all Computer Extension Attributes: %v", err)
	}
	if len(allAttributes.Results) == 0 {
		t.Fatal("No Computer Extension Attributes found for testing")
	}

	// Use the name of the first attribute for testing
	testName := allAttributes.Results[0].Name

	attribute, err := client.GetComputerExtensionAttributeByName(testName)
	if err != nil {
		t.Errorf("Error retrieving Computer Extension Attribute by name: %v", err)
	}

	if attribute == nil || attribute.Name != testName {
		t.Errorf("Attribute not found or name mismatch. Expected Name: %s", testName)
	}
}

func getClient(t *testing.T) *jamfpro.Client {
	// Read environment variables
	clientID := os.Getenv("JAMFPRO_CLIENT_ID")
	clientSecret := os.Getenv("JAMFPRO_CLIENT_SECRET")
	instanceName := os.Getenv("JAMFPRO_INSTANCE_NAME")
	debugMode := os.Getenv("JAMFPRO_DEBUG_MODE") == "true"

	// Check if environment variables are set
	if clientID == "" {
		t.Fatalf("Environment variable JAMFPRO_CLIENT_ID is not set")
	}
	if clientSecret == "" {
		t.Fatalf("Environment variable JAMFPRO_CLIENT_SECRET is not set")
	}
	if instanceName == "" {
		t.Fatalf("Environment variable JAMFPRO_INSTANCE_NAME is not set")
	}

	// Create a config object
	config := jamfpro.Config{
		InstanceName:             instanceName,
		DebugMode:                debugMode,
		ClientID:                 clientID,
		ClientSecret:             clientSecret,
		MaxConcurrentRequests:    5,
		TokenLifespan:            30 * time.Minute,
		TokenRefreshBufferPeriod: 5 * time.Minute,
	}

	// Initialize the client
	client, err := jamfpro.NewClient(config)
	if err != nil {
		t.Fatalf("Failed to initialize client: %v", err)
	}

	return client
}
