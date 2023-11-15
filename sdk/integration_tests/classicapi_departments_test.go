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
		department := jamfpro.DepartmentItem{Name: departmentConfig.Name}
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

func TestJamfProIntegration_GetDepartmentByID(t *testing.T) {
	// Define the department name for which you want to get the ID
	departmentName := "NewDepartmentTest1"

	// Retrieve the list of all departments
	departmentsList, err := intTestClient.GetDepartments()
	if err != nil {
		t.Fatalf("Error fetching departments: %v", err)
	}

	// Find the department by name and get its ID
	var departmentID int
	for _, dept := range departmentsList.Results {
		if dept.Name == departmentName {
			departmentID = dept.Id
			break
		}
	}

	// Assert that the department ID was found
	if departmentID == 0 {
		t.Fatalf("Department '%s' not found", departmentName)
	}

	// Retrieve the department details by its ID
	retrievedDepartment, err := intTestClient.GetDepartmentByID(departmentID)
	if err != nil {
		t.Fatalf("Error retrieving department by ID %d: %v", departmentID, err)
	}

	// Assert that the retrieved department has the expected name
	if retrievedDepartment.Name != departmentName {
		t.Errorf("Expected department name '%s', got '%s'", departmentName, retrievedDepartment.Name)
	}

	// Log the retrieved department for verification
	t.Logf("Retrieved Department: ID=%d, Name=%s", retrievedDepartment.ID, retrievedDepartment.Name)
}

/*
func TestJamfProIntegration_GetDepartmentByName(t *testing.T) {
	// Define the department name to retrieve
	departmentName := "NewDepartmentTest1"

	// Retrieve the department by name
	retrievedDepartment, err := intTestClient.GetDepartmentByName(departmentName)
	if err != nil {
		t.Fatalf("Error retrieving department by name '%s': %v", departmentName, err)
	}

	// Assert that the retrieved department's name matches the expected name
	if retrievedDepartment.Name != departmentName {
		t.Errorf("Expected department name '%s', got '%s'", departmentName, retrievedDepartment.Name)
	}

	// Log the retrieved department for verification
	t.Logf("Retrieved Department: ID=%d, Name=%s", retrievedDepartment.ID, retrievedDepartment.Name)
}

func TestJamfProIntegration_UpdateDepartmentByName(t *testing.T) {
	// Define the original and new department names
	originalName := "NewDepartmentTest1"
	newName := "UpdatedDepartmentTest1"

	// Retrieve the department by its original name
	originalDepartment, err := intTestClient.GetDepartmentByName(originalName)
	if err != nil {
		t.Fatalf("Error retrieving department by name '%s': %v", originalName, err)
	}

	// Update the department name
	updatedDepartment, err := intTestClient.UpdateDepartmentByName(originalName, newName)
	if err != nil {
		t.Fatalf("Error updating department name from '%s' to '%s': %v", originalName, newName, err)
	}

	// Assert that the updated department's ID matches the original department's ID
	if updatedDepartment.ID != originalDepartment.ID {
		t.Errorf("Expected updated department ID to match original, got: %d, want: %d", updatedDepartment.ID, originalDepartment.ID)
	}

	// Assert that the updated department's name matches the new name
	if updatedDepartment.Name != newName {
		t.Errorf("Expected updated department name '%s', got '%s'", newName, updatedDepartment.Name)
	}

	// Log the updated department for verification
	t.Logf("Updated Department: ID=%d, Original Name=%s, New Name=%s", updatedDepartment.ID, originalName, updatedDepartment.Name)
}

func TestJamfProIntegration_UpdateDepartmentByID(t *testing.T) {
	// Define the original department name and the new name
	originalName := "NewDepartmentTest2"
	newName := "UpdatedDepartmentTest2"

	// Retrieve the list of all departments
	departmentsList, err := intTestClient.GetDepartments()
	if err != nil {
		t.Fatalf("Error fetching departments: %v", err)
	}

	// Find the department by name and get its ID
	var departmentID int
	for _, dept := range departmentsList.Results {
		if dept.Name == originalName {
			departmentID = dept.Id
			break
		}
	}

	// Assert that the department ID was found
	if departmentID == 0 {
		t.Fatalf("Department '%s' not found", originalName)
	}

	// Update the department by its ID
	updatedDepartment, err := intTestClient.UpdateDepartmentByID(departmentID, newName)
	if err != nil {
		t.Fatalf("Error updating department ID %d: %v", departmentID, err)
	}

	// Assert that the updated department's name matches the new name
	if updatedDepartment.Name != newName {
		t.Errorf("Expected updated department name '%s', got '%s'", newName, updatedDepartment.Name)
	}

	// Log the updated department for verification
	t.Logf("Updated Department: ID=%d, Original Name=%s, New Name=%s", updatedDepartment.ID, originalName, updatedDepartment.Name)
}

func TestJamfProIntegration_DeleteDepartmentByID(t *testing.T) {
	// Define the updated department name to delete
	departmentName := "UpdatedDepartmentTest1"

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

func TestJamfProIntegration_DeleteDepartmentByName(t *testing.T) {
	// Define the updated department name to delete
	departmentName := "UpdatedDepartmentTest2"

	// Delete the department by its name
	err := intTestClient.DeleteDepartmentByName(departmentName)
	if err != nil {
		t.Fatalf("Error deleting department by name '%s': %v", departmentName, err)
	}

	// Log the deletion for verification
	t.Logf("Deleted Department: Name=%s", departmentName)
}

*/
