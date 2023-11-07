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
	attribute := &jamfpro.ResponseComputerExtensionAttribute{
		Name:      "Test Resource - Computer Extension Attribute - Pop Up Menu Test - Basic",
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
	attribute := &jamfpro.ResponseComputerExtensionAttribute{
		Name:             "Test Resource - Computer Extension Attribute - Pop Up Menu Test - Full",
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
	attribute := &jamfpro.ResponseComputerExtensionAttribute{
		Name:     "Test Resource - Computer Extension Attribute - Script Test - Basic",
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
	attribute := &jamfpro.ResponseComputerExtensionAttribute{
		Name:        "Test Resource - Computer Extension Attribute - Script Test - Full",
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
	attribute := &jamfpro.ResponseComputerExtensionAttribute{
		Name:     "Test Resource - Computer Extension Attribute - Text Field - Basic",
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
	attribute := &jamfpro.ResponseComputerExtensionAttribute{
		Name:        "Test Resource - Computer Extension Attribute - Text Field - Full",
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

func TestUpdateComputerExtensionAttributeByID_SwapAllAttributes(t *testing.T) {
	client := getClient(t)

	// Retrieve the two specific attributes using their names
	basicAttribute, err := client.GetComputerExtensionAttributeByName("Test Resource - Computer Extension Attribute - Pop Up Menu Test - Basic")
	if err != nil {
		t.Fatalf("Error retrieving Basic Attribute: %v", err)
	}

	fullAttribute, err := client.GetComputerExtensionAttributeByName("Test Resource - Computer Extension Attribute - Pop Up Menu Test - Full")
	if err != nil {
		t.Fatalf("Error retrieving Full Attribute: %v", err)
	}

	if basicAttribute == nil || fullAttribute == nil {
		t.Fatal("Could not find both attributes to swap")
	}

	// Swap all attributes between basicAttribute and fullAttribute
	tempAttribute := *basicAttribute
	basicAttribute = fullAttribute
	fullAttribute = &tempAttribute

	// Update the basic attribute (which now holds the full attribute's data)
	_, err = client.UpdateComputerExtensionAttributeByID(basicAttribute.ID, basicAttribute)
	if err != nil {
		t.Errorf("Error updating Basic Attribute by ID: %v", err)
	}

	// Update the full attribute (which now holds the basic attribute's data)
	_, err = client.UpdateComputerExtensionAttributeByID(fullAttribute.ID, fullAttribute)
	if err != nil {
		t.Errorf("Error updating Full Attribute by ID: %v", err)
	}
}

// TestUpdateComputerExtensionAttributeByID_Basic tests the basic update of a computer extension attribute by its ID.
func TestUpdateComputerExtensionAttributeByID_Basic(t *testing.T) {
	// getClient creates and returns a new Jamf Pro API client
	client := getClient(t)

	// GetComputerExtensionAttributes retrieves all computer extension attributes
	attributes, err := client.GetComputerExtensionAttributes()
	if err != nil {
		t.Fatalf("Error retrieving attributes: %v", err)
	}

	var attributeToUpdate *jamfpro.ResponseComputerExtensionAttribute
	for _, attr := range attributes.Results {
		if attr.Name == "Basic Pop Up Menu Test" {
			// Create a new variable of the correct type and assign relevant fields
			attributeToUpdate = &jamfpro.ResponseComputerExtensionAttribute{
				ID:   attr.ID,
				Name: attr.Name,
				// Assign other fields that are present in ComputerExtensionAttributeItem
			}
			break
		}
	}

	if attributeToUpdate == nil {
		t.Fatalf("Could not find attribute to update")
	}

	// Define updates for the attribute
	updatedAttribute := &jamfpro.ResponseComputerExtensionAttribute{
		Name: "Renamed Basic Pop Up Menu Test",
		// Assign the DataType or other fields if necessary and available
	}

	// UpdateComputerExtensionAttributeByID updates the computer extension attribute with the specified ID
	_, err = client.UpdateComputerExtensionAttributeByID(attributeToUpdate.ID, updatedAttribute)
	if err != nil {
		t.Errorf("Error updating Computer Extension Attribute by ID: %v", err)
	}
}

// TestUpdateComputerExtensionAttributeByID_Full tests a comprehensive update of a computer extension attribute by its ID.
func TestUpdateComputerExtensionAttributeByID_Full(t *testing.T) {
	client := getClient(t)

	attributes, err := client.GetComputerExtensionAttributes()
	if err != nil {
		t.Fatalf("Error retrieving attributes: %v", err)
	}

	var attributeToUpdate *jamfpro.ResponseComputerExtensionAttribute
	for _, attr := range attributes.Results {
		if attr.Name == "Renamed Basic Pop Up Menu Test" {
			attributeToUpdate = &jamfpro.ResponseComputerExtensionAttribute{
				ID:   attr.ID,
				Name: attr.Name,
			}
			break
		}
	}

	if attributeToUpdate == nil {
		t.Fatalf("Could not find attribute to update")
	}

	updatedAttribute := &jamfpro.ResponseComputerExtensionAttribute{
		Name:             "Updated Battery Cycle Count",
		Description:      "Updated number of charge cycles logged on the current battery",
		DataType:         "String",
		InputType:        jamfpro.ComputerExtensionAttributeInputType{Type: "Text Field"},
		InventoryDisplay: "General",
		ReconDisplay:     "Extension Attributes",
	}

	_, err = client.UpdateComputerExtensionAttributeByID(attributeToUpdate.ID, updatedAttribute)
	if err != nil {
		t.Errorf("Error updating Computer Extension Attribute by ID: %v", err)
	}
}

// TestUpdateComputerExtensionAttributeByName_Basic tests the basic update of a computer extension attribute by its name.
func TestUpdateComputerExtensionAttributeByName_Basic(t *testing.T) {
	client := getClient(t)

	// Retrieve the attribute to ensure it exists
	attributeToUpdate, err := client.GetComputerExtensionAttributeByName("Basic Pop Up Menu Test")
	if err != nil {
		t.Fatalf("Error retrieving attribute by name: %v", err)
	}

	if attributeToUpdate == nil {
		t.Fatalf("Could not find attribute to update")
	}

	updatedAttribute := &jamfpro.ResponseComputerExtensionAttribute{
		Name: "Renamed Basic Pop Up Menu Test",
	}

	_, err = client.UpdateComputerExtensionAttributeByName(attributeToUpdate.Name, updatedAttribute)
	if err != nil {
		t.Errorf("Error updating Computer Extension Attribute by Name: %v", err)
	}
}

// TestUpdateComputerExtensionAttributeByName_Full tests a comprehensive update of a computer extension attribute by its name.
func TestUpdateComputerExtensionAttributeByName_Full(t *testing.T) {
	client := getClient(t)

	attributeToUpdate, err := client.GetComputerExtensionAttributeByName("Renamed Basic Pop Up Menu Test")
	if err != nil {
		t.Fatalf("Error retrieving attribute by name: %v", err)
	}

	if attributeToUpdate == nil {
		t.Fatalf("Could not find attribute to update")
	}

	updatedAttribute := &jamfpro.ResponseComputerExtensionAttribute{
		Name:             "Updated Battery Cycle Count",
		Description:      "Updated number of charge cycles logged on the current battery",
		DataType:         "String",
		InputType:        jamfpro.ComputerExtensionAttributeInputType{Type: "Text Field"},
		InventoryDisplay: "General",
		ReconDisplay:     "Extension Attributes",
	}

	_, err = client.UpdateComputerExtensionAttributeByName(attributeToUpdate.Name, updatedAttribute)
	if err != nil {
		t.Errorf("Error updating Computer Extension Attribute by Name: %v", err)
	}
}

// TestDeleteComputerExtensionAttributesByName removes the resources created as part of the this test strategy
func TestDeleteComputerExtensionAttributesByName(t *testing.T) {
	client := getClient(t)

	// Names of resources created earlier and updated
	namesToDelete := []string{"TestAttribute1", "TestAttribute2", "TestAttribute3", "UpdatedTestAttribute1", "UpdatedTestAttribute2", "UpdatedTestAttribute3"}

	// Test Deletion by Name
	for _, name := range namesToDelete {
		err := client.DeleteComputerExtensionAttributeByNameByID(name)
		if err != nil {
			t.Errorf("Failed to delete resource with name %s: %v", name, err)
		}

		// Optional: Verify deletion
		// attrs, _ := client.GetComputerExtensionAttributes()
		// for _, attr := range attrs.Results {
		//     if attr.Name == name {
		//         t.Errorf("Resource with name %s still exists after deletion", name)
		//     }
		// }
	}
}

func getClient(t *testing.T) *jamfpro.Client {
	// Read environment variables
	clientID := os.Getenv("JAMFPRO_CLIENT_ID")
	clientSecret := os.Getenv("JAMFPRO_CLIENT_SECRET")
	instanceName := os.Getenv("JAMFPRO_INSTANCE_NAME")
	//debugMode := os.Getenv("JAMFPRO_DEBUG_MODE") == "true"

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
		InstanceName: instanceName,
		//DebugMode:                debugMode,
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
