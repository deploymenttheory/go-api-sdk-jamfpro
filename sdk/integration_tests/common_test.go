// common_test.go
// Testing Harness
package jamfpro_integration_test

import (
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

var (
	client    *jamfpro.Client
	roleName  string
	apiClient *jamfpro.ApiIntegration
)

// Global client to be shared across tests
var globalClient *jamfpro.Client

// GlobalConfig is the top-level structure for the entire configuration.
type IntegrationTestGlobalConfig struct {
    JamfPro JamfProConfig `json:"JamfPro"`
}

// JamfProConfig contains the configurations specific to Jamf Pro.
type JamfProConfig struct {
    ApiClient           ApiClientConfig          `json:"ApiClient"`
    ApiRoles            map[string]ApiRoleConfig `json:"ApiRoles"`
    IntegrationTestData IntegrationTestData      `json:"IntegrationTestData"`
}

// ApiClientConfig defines the configuration for the API client.
type ApiClientConfig struct {
    DisplayName               string `json:"displayName"`
    Enabled                   bool   `json:"enabled"`
    AccessTokenLifetimeSeconds int    `json:"accessTokenLifetimeSeconds"`
}

// ApiRoleConfig defines the configuration for a specific API role.
type ApiRoleConfig struct {
    Name       string   `json:"name"`
    Privileges []string `json:"privileges"`
}

// IntegrationTestData contains test data for different integration tests.
type IntegrationTestData struct {
	Departments DepartmentsTestData `json:"departments"`
	SsoFailover map[string]interface{} `json:"ssoFailover"`
}

type DepartmentsTestData struct {
	Create DepartmentsCreateUpdateTestData `json:"create"`
	Update DepartmentsCreateUpdateTestData `json:"update"`
}

type DepartmentsCreateUpdateTestData struct {
	MinimumConfiguration DepartmentResource `json:"MinimumConfiguration"`
	MaximumConfiguration DepartmentResource `json:"MaximumConfiguration"`
}

type DepartmentResource struct {
	Names []string `json:"names"`
}

// Global variables shared across tests
var (
	globalClient    *jamfpro.Client
	globalRoleName  string
	globalApiClient *jamfpro.ApiIntegration
)

// TestMain function performs global setup of the test suite, executes
// the test suite and then performs the teardown of any resources created.
func TestMain(m *testing.M) {
	var err error
	// Initialize the Jamf Pro client
	client, err = setupJamfProHTTPClient()
	if err != nil {
		log.Fatalf("Failed to setup Jamf Pro client: %v", err)
	}

	// Create an API role and get its name
	config, err := LoadConfig("path/to/config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup all roles defined in the configuration
	createdRoles, err := setupAllIntegrationTestRoles(globalClient, config.JamfPro.ApiRoles)
	if err != nil {
		log.Fatalf("Failed to setup roles from configuration: %v", err)
	}

	// Setup the API client using the configuration
	createdApiClient, err := setupIntegrationTestAPIClient(globalClient, config.JamfPro.ApiClient, roleNames)
	if err != nil {
		log.Fatalf("Failed to setup API client from configuration: %v", err)
	}

	// Run the tests
	exitVal := m.Run()

	// Global Teardown: first remove the API client, then the role
	teardownApiIntegration(client, apiClient.DisplayName)
	teardownApiRole(client, roleName)

	os.Exit(exitVal)
}

// globalSetup performs common setup tasks for all test suites.
func globalSetup() error {
	var err error
	globalClient, err = setupJamfProClient() // Initializes the Jamf Pro client.
	if err != nil {
		return err
	}

	globalRoleName, err = setupTemporaryTestRole(globalClient) // Creates a temporary test role.
	if err != nil {
		return err
	}

	globalApiClient, err = setupTemporaryTestAPIClient(globalClient, []string{globalRoleName}) // Sets up a temporary API client.
	return err
}

// globalTeardown performs common teardown tasks for all test suites.
func globalTeardown() {
	teardownApiIntegration(globalClient, globalApiClient.DisplayName)
	teardownApiRole(globalClient, globalRoleName)
}

// setupJamfProClient initializes the Jamf Pro client.
func setupJamfProClient() (*jamfpro.Client, error) {

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

// setupIntegrationTestRoles sets up a temporary test role in Jamf Pro for testing purposes.
// It creates a new Jamf API role based on the provided configuration.
func setupIntegrationTestRoles(client *jamfpro.Client, roleConfig ApiRoleConfig) (string, error) {
	newRole := &jamfpro.APIRole{
			DisplayName: roleConfig.Name,
			Privileges:  roleConfig.Privileges,
	}

	newRole, err := client.CreateJamfApiRole(newRole)
	if err != nil {
			log.Fatalf("Error creating Jamf API role '%s': %v", roleConfig.Name, err)
	}

	// Log the creation of the new API role
	log.Printf("Created API Role with Display Name: %s\n", newRole.DisplayName)
	log.Printf("Created API Role has the following privileges: %s\n", newRole.Privileges)

	return newRole.DisplayName, nil
}

// setupJamfProTemporaryTestAPIClient creates a temporary API integration in Jamf Pro.
// It uses the provided API client configuration and role names.
func setupIntegrationTestAPIClient(client *jamfpro.Client, apiClientConfig ApiClientConfig, roleNames []string) (*jamfpro.ApiIntegration, error) {
	newApiIntegration := &jamfpro.ApiIntegration{
		AuthorizationScopes:        roleNames,
		DisplayName:                apiClientConfig.DisplayName,
		Enabled:                    apiClientConfig.Enabled,
		AccessTokenLifetimeSeconds: apiClientConfig.AccessTokenLifetimeSeconds,
	}

	createdApiIntegration, err := client.CreateApiIntegration(newApiIntegration)
	if err != nil {
		log.Fatalf("Error creating API Integration: %v", err)
	}

	log.Printf("Created API Integration with Display Name: %s\n", createdApiIntegration.DisplayName)
	return createdApiIntegration, nil
}



// teardownApiIntegration is a  Helper function to delete API integration
func teardownApiIntegration(client *jamfpro.Client, integrationName string) {
	if err := client.DeleteApiIntegrationByName(integrationName); err != nil {
		log.Fatalf("Failed to delete API integration: %v", err) // Exits the program if there's an error
	} else {
		log.Printf("API integration '%s' deleted successfully", integrationName)
	}
}

// teardownApiRole is a Helper function to delete API role
func teardownApiRole(client *jamfpro.Client, roleName string) {
	if err := client.DeleteJamfApiRoleByName(roleName); err != nil {
		log.Fatalf("Failed to delete API role: %v", err) // Exits the program if there's an error
	} else {
		log.Printf("API role '%s' deleted successfully", roleName)
	}
}

func setupFromConfig(config GlobalConfig) error {
	// Use config.JamfPro.Roles and config.JamfPro.TestData
	// to setup roles and test data
}

// In your test main or setup functions:
config, err := LoadConfig("path/to/config.json")
if err != nil {
	log.Fatalf("Failed to load configuration: %v", err)
}
err = setupFromConfig(config)
if err != nil {
	log.Fatalf("Failed to setup from configuration: %v", err)
}

// setupAllIntegrationTestRoles sets up all temporary test roles defined in the configuration.
func setupAllIntegrationTestRoles(client *jamfpro.Client, rolesConfig map[string]ApiRoleConfig) (map[string]string, error) {
	createdRoles := make(map[string]string) // Map to store the names of created roles

	for key, roleConfig := range rolesConfig {
		roleName, err := setupIntegrationTestRoles(client, roleConfig)
		if err != nil {
			// Log the error but continue attempting to create other roles
			log.Printf("Error setting up role '%s': %v", key, err)
			continue
		}
		createdRoles[key] = roleName
	}

	return createdRoles, nil
}

// loadConfig loads the configuration from a JSON file.
func loadConfig(filePath string) (IntegrationTestGlobalConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return IntegrationTestGlobalConfig{}, err
	}
	defer file.Close()

	var config IntegrationTestGlobalConfig
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return IntegrationTestGlobalConfig{}, err
	}

	return config, nil
}