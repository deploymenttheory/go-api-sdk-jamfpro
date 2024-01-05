// classicapi_departments_test.go
// Jamf Pro Classic Api - Departments Testing
// api reference: https://developer.jamf.com/jamf-pro/reference/departments

package jamfpro_integration_test

import (
	"embed"
	"encoding/xml"
	"log"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

//go:embed classicapi_departments_test_data.xml
var testXMLData embed.FS

// DepartmentConfig defines the test configuration for a department, including only the name.
// IntegrationTestData represents the test data structure for department operations.
type IntegrationTestData struct {
	Departments DepartmentsTestData `xml:"Departments"`
}

// DepartmentsTestData holds configurations for creating and updating departments.
type DepartmentsTestData struct {
	Create struct {
		MinimumConfiguration DepartmentConfig `xml:"MinimumConfiguration"`
		MaximumConfiguration DepartmentConfig `xml:"MaximumConfiguration"`
	} `xml:"Create"`
	Update struct {
		MinimumConfiguration DepartmentConfig `xml:"MinimumConfiguration"`
		MaximumConfiguration DepartmentConfig `xml:"MaximumConfiguration"`
	} `xml:"Update"`
}
type DepartmentConfig struct {
	Name string `xml:"Name"`
}

// loadDepartmentTestData reads and unmarshals the XML file containing test data
// for department operations in integration tests.
func loadDepartmentTestData(t *testing.T) (*IntegrationTestData, error) {
	var testData IntegrationTestData

	// Read the XML file
	data, err := testXMLData.ReadFile("classicapi_departments_test_data.xml")
	if err != nil {
		t.Fatalf("Error reading XML file: %v\n", err)
	}
	t.Log("XML file read successfully")
	//t.Logf("Raw XML data: %s\n", string(data))

	// Unmarshal the XML data into the testData struct
	err = xml.Unmarshal(data, &testData)
	if err != nil {
		t.Fatalf("Error unmarshaling XML data: %v\n", err)
	}
	t.Logf("XML data unmarshaled successfully: %+v\n", testData)

	return &testData, nil
}

// TestJamfProIntegration_CreateDepartments
// Purpose: Tests the functionality to create new departments using the Jamf Pro Classic API. It verifies that the departments are successfully created with correct names and non-zero IDs.
// Process: Loads department configurations from XML test data, creates departments based on the configurations, and then verifies their creation by retrieving them by ID.
func TestJamfProIntegration_CreateDepartments(t *testing.T) {
	// Load department test data from XML
	testData, err := loadDepartmentTestData(t)
	if err != nil {
		t.Fatalf("Failed to load department test data: %v", err)
	}

	// Use the loaded department test data for create operation
	createTestData := testData.Departments.Create

	// Debug log
	log.Printf("Loaded Create Department Test Data: %+v\n", createTestData)

	// Create and assert departments using the loaded create test data
	for _, departmentConfig := range []DepartmentConfig{createTestData.MinimumConfiguration, createTestData.MaximumConfiguration} {
		department := jamfpro.ResourceDepartment{Name: departmentConfig.Name}
		createdDepartment, err := intTestClient.CreateDepartment(department.Name)
		if err != nil {
			t.Fatalf("Error creating department '%s': %v", department.Name, err)
		}

		// Assert non-zero ID
		if createdDepartment.ID == 0 {
			t.Errorf("Expected a non-zero ID for department '%s', got 0", department.Name)
		}

		// Retrieve and assert department
		retrievedDepartment, err := intTestClient.GetDepartmentByID(createdDepartment.ID)
		if err != nil {
			t.Fatalf("Error retrieving department '%s' by ID: %v", department.Name, err)
		}

		if retrievedDepartment.Name != department.Name {
			t.Errorf("Expected department name '%s', got '%s'", department.Name, retrievedDepartment.Name)
		}

		t.Logf("Retrieved Department: ID=%d, Name=%s", retrievedDepartment.ID, retrievedDepartment.Name)
	}
}

// TestJamfProIntegration_GetDepartments
// Purpose: Validates the ability to retrieve a list of all departments. It checks whether the departments created in the test data are present in the retrieved list.
// Process: Loads department configurations from XML test data, fetches the list of all departments, and then verifies the presence of the test departments in the list.
func TestJamfProIntegration_GetDepartments(t *testing.T) {
	// Initial log statement to confirm test execution
	t.Log("Starting TestJamfProIntegration_GetDepartments")

	// Load department test data from XML
	testData, err := loadDepartmentTestData(t)
	if err != nil {
		t.Fatalf("Failed to load department test data: %v", err)
	} else {
		t.Logf("Loaded department test data: %+v", testData)
	}

	// Call GetDepartments function to retrieve all departments
	departmentsList, err := intTestClient.GetDepartments()
	if err != nil {
		t.Fatalf("Error fetching departments: %v", err)
	} else {
		// Debug log: print the departmentsList for verification
		t.Logf("Retrieved Departments: %+v", departmentsList)
	}

	// Check for the presence of the departments defined in the test data
	var foundMinConfig, foundMaxConfig bool
	for _, department := range departmentsList.Results {
		if department.Name == testData.Departments.Create.MinimumConfiguration.Name {
			foundMinConfig = true
		}
		if department.Name == testData.Departments.Create.MaximumConfiguration.Name {
			foundMaxConfig = true
		}
	}

	// Assert that both departments are found
	if !foundMinConfig {
		t.Errorf("Department '%s' not found", testData.Departments.Create.MinimumConfiguration.Name)
	}
	if !foundMaxConfig {
		t.Errorf("Department '%s' not found", testData.Departments.Create.MaximumConfiguration.Name)
	}

	// Log the result for verification
	if foundMinConfig && foundMaxConfig {
		t.Logf("Both departments '%s' and '%s' are found",
			testData.Departments.Create.MinimumConfiguration.Name,
			testData.Departments.Create.MaximumConfiguration.Name)
	} else {
		if foundMinConfig {
			t.Logf("Department '%s' is found", testData.Departments.Create.MinimumConfiguration.Name)
		}
		if foundMaxConfig {
			t.Logf("Department '%s' is found", testData.Departments.Create.MaximumConfiguration.Name)
		}
	}
}

// TestJamfProIntegration_GetDepartmentByID
// Purpose: Tests retrieving department details by department ID. It ensures that departments can be correctly identified and retrieved using their unique IDs.
// Process: Iterates through each department defined in the test data, finds their IDs, retrieves them by these IDs, and verifies that the retrieved information matches the test data.
func TestJamfProIntegration_GetDepartmentByID(t *testing.T) {
	// Load department test data from XML
	testData, err := loadDepartmentTestData(t)
	if err != nil {
		t.Fatalf("Failed to load department test data: %v", err)
	}

	// Retrieve the list of all departments
	departmentsList, err := intTestClient.GetDepartments()
	if err != nil {
		t.Fatalf("Error fetching departments: %v", err)
	}

	// Test for each department in test data
	for _, config := range []DepartmentConfig{
		testData.Departments.Create.MinimumConfiguration,
		testData.Departments.Create.MaximumConfiguration,
	} {
		// Find the department by name and get its ID
		var departmentID int
		for _, dept := range departmentsList.Results {
			if dept.Name == config.Name {
				departmentID = dept.Id
				break
			}
		}

		// Assert that the department ID was found
		if departmentID == 0 {
			t.Fatalf("Department '%s' not found", config.Name)
		}

		// Retrieve the department details by its ID
		retrievedDepartment, err := intTestClient.GetDepartmentByID(departmentID)
		if err != nil {
			t.Fatalf("Error retrieving department by ID %d: %v", departmentID, err)
		}

		// Assert that the retrieved department has the expected name
		if retrievedDepartment.Name != config.Name {
			t.Errorf("Expected department name '%s', got '%s'", config.Name, retrievedDepartment.Name)
		}

		// Log the retrieved department for verification
		t.Logf("Retrieved Department: ID=%d, Name=%s", retrievedDepartment.ID, retrievedDepartment.Name)
	}
}

// TestJamfProIntegration_GetDepartmentByName
// Purpose: Verifies the functionality to retrieve department details by department name. It ensures that departments can be correctly identified and retrieved using their names.
// Process: Iterates through each department defined in the test data, retrieves them by their names, and verifies that the retrieved information matches the test dat
func TestJamfProIntegration_GetDepartmentByName(t *testing.T) {
	// Load department test data from XML
	testData, err := loadDepartmentTestData(t)
	if err != nil {
		t.Fatalf("Failed to load department test data: %v", err)
	}

	// Test for each department in test data
	for _, config := range []DepartmentConfig{
		testData.Departments.Create.MinimumConfiguration,
		testData.Departments.Create.MaximumConfiguration,
	} {
		// Retrieve the department by name
		retrievedDepartment, err := intTestClient.GetDepartmentByName(config.Name)
		if err != nil {
			t.Fatalf("Error retrieving department by name '%s': %v", config.Name, err)
		}

		// Assert that the retrieved department's name matches the expected name
		if retrievedDepartment.Name != config.Name {
			t.Errorf("Expected department name '%s', got '%s'", config.Name, retrievedDepartment.Name)
		}

		// Log the retrieved department for verification
		t.Logf("Retrieved Department: ID=%d, Name=%s", retrievedDepartment.ID, retrievedDepartment.Name)
	}
}

// TestJamfProIntegration_UpdateDepartmentByName
// Purpose: Tests the ability to update a department's name using the department's current name and tests that a minimum configuration can be replaced with a maximum configuration. It checks if the department name is correctly updated in the system.
// Process: Loads department configurations from XML test data, updates the name of a specified department, retrieves the department by its new ID to verify the update, and ensures the department's name has been updated.
func TestJamfProIntegration_UpdateDepartmentByName(t *testing.T) {
	// Load department test data from XML
	testData, err := loadDepartmentTestData(t)
	if err != nil {
		t.Fatalf("Failed to load department test data: %v", err)
	}

	// Define the original and new department names
	originalName := testData.Departments.Create.MinimumConfiguration.Name
	newName := testData.Departments.Update.MaximumConfiguration.Name

	// Update the department name
	updatedDepartment, err := intTestClient.UpdateDepartmentByName(originalName, newName)
	if err != nil {
		t.Fatalf("Error updating department name from '%s' to '%s': %v", originalName, newName, err)
	}

	// Use the ID from the update response to retrieve the updated department
	retrievedUpdatedDepartment, err := intTestClient.GetDepartmentByID(updatedDepartment.ID)
	if err != nil {
		t.Fatalf("Error retrieving updated department by ID %d: %v", updatedDepartment.ID, err)
	}

	// Assert that the retrieved updated department's name matches the new name
	if retrievedUpdatedDepartment.Name != newName {
		t.Errorf("Expected updated department name '%s', got '%s'", newName, retrievedUpdatedDepartment.Name)
	}

	// Log the updated department for verification
	t.Logf("Updated Department: ID=%d, Original Name=%s, New Name=%s", retrievedUpdatedDepartment.ID, originalName, retrievedUpdatedDepartment.Name)
}

// TestJamfProIntegration_UpdateDepartmentByID
// Purpose: Validates the functionality to update a department's name using the department's ID and tests that a minimum configuration can be replaced with a maximum configuration.. It ensures that department names can be accurately modified using their unique IDs.
// Process: Loads department configurations from XML test data, finds the ID of a specified department, updates the department using its ID, and verifies that the department's name has been updated as expected.
func TestJamfProIntegration_UpdateDepartmentByID(t *testing.T) {
	// Load department test data from XML
	testData, err := loadDepartmentTestData(t)
	if err != nil {
		t.Fatalf("Failed to load department test data: %v", err)
	}

	// Define the department to update and the new name
	originalName := testData.Departments.Create.MaximumConfiguration.Name
	newName := testData.Departments.Update.MinimumConfiguration.Name

	// Retrieve the list of all departments to find the target department's ID
	departmentsList, err := intTestClient.GetDepartments()
	if err != nil {
		t.Fatalf("Error fetching departments: %v", err)
	}

	var departmentID int
	for _, dept := range departmentsList.Results {
		if dept.Name == originalName {
			departmentID = dept.Id
			break
		}
	}

	if departmentID == 0 {
		t.Fatalf("Department '%s' not found", originalName)
	}

	// Update the department by ID
	_, err = intTestClient.UpdateDepartmentByID(departmentID, newName)
	if err != nil {
		t.Fatalf("Error updating department ID %d to name '%s': %v", departmentID, newName, err)
	}

	// Retrieve the updated department to confirm the update
	updatedDepartment, err := intTestClient.GetDepartmentByID(departmentID)
	if err != nil {
		t.Fatalf("Error retrieving updated department by ID %d: %v", departmentID, err)
	}

	// Assert that the updated department's name matches the new name
	if updatedDepartment.Name != newName {
		t.Errorf("Expected updated department name '%s', got '%s'", newName, updatedDepartment.Name)
	}

	// Log the updated department for verification
	t.Logf("Updated Department: ID=%d, Original Name=%s, New Name=%s", departmentID, originalName, updatedDepartment.Name)
}

// TestJamfProIntegration_DeleteDepartmentByID
// Purpose: Tests the deletion of a department using its ID. It confirms that departments can be successfully removed from the system by their unique IDs.
// Process: Determines the ID of a specified department, deletes the department by its ID, and logs the deletion action for verification.
func TestJamfProIntegration_DeleteDepartmentByID(t *testing.T) {
	// Define the updated department name to delete
	departmentName := "UpdateDepartmentsMinConfigIntTest1"

	// Retrieve the department ID by the updated name
	departmentID, err := intTestClient.GetDepartmentIdByName(departmentName)
	if err != nil {
		t.Fatalf("Error retrieving department ID for name '%s': %v", departmentName, err)
	}

	// Assert that the department ID was found
	if departmentID == 0 {
		t.Fatalf("Department '%s' not found", departmentName)
	}

	// Delete the department by its ID
	err = intTestClient.DeleteDepartmentByID(departmentID)
	if err != nil {
		t.Fatalf("Error deleting department by ID %d: %v", departmentID, err)
	}

	// Log the deletion for verification
	t.Logf("Deleted Department: ID=%d, Name=%s", departmentID, departmentName)
}

// TestJamfProIntegration_DeleteDepartmentByName
// Purpose: Validates the ability to delete a department using its name. It ensures that departments can be correctly identified and removed using their names.
// Process: Deletes a department using its name and logs the deletion action for verification.
func TestJamfProIntegration_DeleteDepartmentByName(t *testing.T) {
	// Define the updated department name to delete
	departmentName := "UpdateDepartmentsMaxConfigIntTest2"

	// Delete the department by its name
	err := intTestClient.DeleteDepartmentByName(departmentName)
	if err != nil {
		t.Fatalf("Error deleting department by name '%s': %v", departmentName, err)
	}

	// Log the deletion for verification
	t.Logf("Deleted Department: Name=%s", departmentName)
}
