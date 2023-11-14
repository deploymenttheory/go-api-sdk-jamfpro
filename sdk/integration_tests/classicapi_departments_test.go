// classicapi_departments_test.go
// Jamf Pro Classic Api - Departments Testing
// api reference: https://developer.jamf.com/jamf-pro/reference/departments

/*
	Test Strategy:

Global Setup: The testing process begins with the initialization of the Jamf Pro HTTP client.
This global setup phase involves creating a temporary API test role and setting up a corresponding
API client for integration testing. This ensures that all tests run in a consistent
and controlled environment.

Individual Test Execution: Each integration test, managed by testing.T, is executed according
to a predefined test plan. These tests utilize the temporary API client and test role established
in the setup phase. The use of testing.T facilitates granular error reporting and isolated
testing of specific functionalities within the Jamf Pro integration.

Global Teardown: Upon completion of all tests, the suite enters the teardown phase. This
involves a systematic cleanup of all test-generated resources, including the removal of
the temporary API client and the test role. This step is crucial for ensuring that the
testing environment is reset and no residual data impacts subsequent test runs.
*/
package jamfpro_integration_test

import (
	"log"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// setupJamfProDepartmentsClient initializes a Jamf Pro client for testing purposes.
// It reads the OAuth credentials from a specified configuration file and uses these credentials
// to create and return a new Jamf Pro client instance. This function is used to setup a client
// for various integration tests in this suite.
func setupJamfProDepartmentsClient() (*jamfpro.Client, error) {

	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelInfo // Set http_client logging level to debug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	return client, nil
}

// setupJamfProDepartmentsTemporaryTestRole sets up a temporary test role in Jamf Pro for testing purposes.
// It creates a new Jamf API role with specific privileges necessary for the test scenarios.
// This function returns the name of the created role for use in subsequent tests.
func setupJamfProDepartmentsTemporaryTestRole(client *jamfpro.Client) (string, error) {
	roleName := "go-api-sdk-jamfpro-apir-departments"

	newRole := &jamfpro.APIRole{
		DisplayName: roleName,
		Privileges:  []string{"Create Departments", "Read Departments", "Update Departments", "Delete Departments"},
	}

	newRole, err := client.CreateJamfApiRole(newRole)
	if err != nil {
		log.Fatalf("Error creating Jamf Departments API role: %v", err)
	}

	// Log the creation of the new API Integration
	log.Printf("Created API Role with Display Name: %s\n", newRole.DisplayName)
	log.Printf("Created API Role has the following priviledges: %s\n", newRole.Privileges)
	return roleName, nil
}

// setupJamfProDepartmentsTemporaryTestAPIClient creates a temporary API integration in Jamf Pro
// using specific API roles. This integration is used for testing and has privileges as defined in the roles.
func setupJamfProDepartmentsTemporaryTestAPIClient(client *jamfpro.Client, roleNames []string) (*jamfpro.ApiIntegration, error) {
	// Define the new API Integration using the provided role names
	newApiIntegration := &jamfpro.ApiIntegration{
		AuthorizationScopes:        roleNames, // Use the role names provided
		DisplayName:                "DepartmentsTemporaryTestAPIClient",
		Enabled:                    true,
		AccessTokenLifetimeSeconds: 1200,
	}

	// Call the function to create the new API Integration
	createdApiIntegration, err := client.CreateApiIntegration(newApiIntegration)
	if err != nil {
		log.Fatalf("Error creating API Integration: %v", err)
	}

	// Log the creation of the new API Integration
	log.Printf("Created API Integration with Display Name: %s\n", createdApiIntegration.DisplayName)

	return createdApiIntegration, nil
}

func TestJamfProIntegration_CreateDepartments(t *testing.T) {
	// Define names for the departments to be created
	departmentName1 := "NewDepartmentTest1"
	departmentName2 := "NewDepartmentTest2"

	// Function to create and assert a department
	createAndAssertDepartment := func(name string) {
		// Call CreateDepartment function using the department name
		createdDepartment, err := client.CreateDepartment(name)
		if err != nil {
			t.Fatalf("Error creating department '%s': %v", name, err)
		}

		// Assert that the created department has a non-zero ID
		if createdDepartment.ID == 0 {
			t.Errorf("Expected a non-zero ID for department '%s', got 0", name)
		}

		// Retrieve the created department by ID
		retrievedDepartment, err := client.GetDepartmentByID(createdDepartment.ID)
		if err != nil {
			t.Fatalf("Error retrieving department '%s' by ID: %v", name, err)
		}

		// Assert that the retrieved department has the expected name
		if retrievedDepartment.Name != name {
			t.Errorf("Expected department name '%s', got '%s'", name, retrievedDepartment.Name)
		}

		// Log the retrieved department for verification
		t.Logf("Retrieved Department: ID=%d, Name=%s", retrievedDepartment.ID, retrievedDepartment.Name)
	}

	// Create and assert first department
	createAndAssertDepartment(departmentName1)

	// Create and assert second department
	createAndAssertDepartment(departmentName2)
}

func TestJamfProIntegration_GetDepartments(t *testing.T) {
	// Call GetDepartments function to retrieve all departments
	departmentsList, err := client.GetDepartments()
	if err != nil {
		t.Fatalf("Error fetching departments: %v", err)
	}

	// Check for the presence of the created departments
	var found1, found2 bool
	for _, department := range departmentsList.Results {
		if department.Name == "NewDepartmentTest1" {
			found1 = true
		}
		if department.Name == "NewDepartmentTest2" {
			found2 = true
		}
	}

	// Assert that both departments are found
	if !found1 {
		t.Errorf("Department 'NewDepartmentTest1' not found")
	}
	if !found2 {
		t.Errorf("Department 'NewDepartmentTest2' not found")
	}

	// Log the result for verification
	if found1 && found2 {
		t.Logf("Both departments 'NewDepartmentTest1' and 'NewDepartmentTest2' are found")
	}
}

func TestJamfProIntegration_GetDepartmentByID(t *testing.T) {
	// Define the department name for which you want to get the ID
	departmentName := "NewDepartmentTest1"

	// Retrieve the list of all departments
	departmentsList, err := client.GetDepartments()
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
	retrievedDepartment, err := client.GetDepartmentByID(departmentID)
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

func TestJamfProIntegration_GetDepartmentByName(t *testing.T) {
	// Define the department name to retrieve
	departmentName := "NewDepartmentTest1"

	// Retrieve the department by name
	retrievedDepartment, err := client.GetDepartmentByName(departmentName)
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
	originalDepartment, err := client.GetDepartmentByName(originalName)
	if err != nil {
		t.Fatalf("Error retrieving department by name '%s': %v", originalName, err)
	}

	// Update the department name
	updatedDepartment, err := client.UpdateDepartmentByName(originalName, newName)
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
	departmentsList, err := client.GetDepartments()
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
	updatedDepartment, err := client.UpdateDepartmentByID(departmentID, newName)
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
	departmentID, err := client.GetDepartmentIdByName(departmentName)
	if err != nil {
		t.Fatalf("Error retrieving department ID for name '%s': %v", departmentName, err)
	}

	// Assert that the department ID was found
	if departmentID == 0 {
		t.Fatalf("Department '%s' not found", departmentName)
	}

	// Delete the department by its ID
	err = client.DeleteDepartmentByID(departmentID)
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
	err := client.DeleteDepartmentByName(departmentName)
	if err != nil {
		t.Fatalf("Error deleting department by name '%s': %v", departmentName, err)
	}

	// Log the deletion for verification
	t.Logf("Deleted Department: Name=%s", departmentName)
}

// Helper function to delete API integration
func teardownDepartmentsApiIntegration(client *jamfpro.Client, integrationName string) {
	if err := client.DeleteApiIntegrationByName(integrationName); err != nil {
		log.Fatalf("Failed to delete API integration: %v", err) // Exits the program if there's an error
	} else {
		log.Printf("API integration '%s' deleted successfully", integrationName)
	}
}

// Helper function to delete API role
func teardownDepartmentsApiRole(client *jamfpro.Client, roleName string) {
	if err := client.DeleteJamfApiRoleByName(roleName); err != nil {
		log.Fatalf("Failed to delete API role: %v", err) // Exits the program if there's an error
	} else {
		log.Printf("API role '%s' deleted successfully", roleName)
	}
}
