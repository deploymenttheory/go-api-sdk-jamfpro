// common_test_harness_test.go
// Testing Harness
package jamfpro_integration_test

/*
	Test Strategy:

Global Setup: The testing process begins with the initialization of the Jamf Pro HTTP intTestClient.
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

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// GlobalConfig is the top-level structure for the entire configuration.
type IntegrationTestGlobalConfig struct {
	JamfPro JamfProConfig `json:"JamfPro"`
}

// JamfProConfig contains the configurations specific to Jamf Pro.
type JamfProConfig struct {
	ApiClient ApiClientConfig          `json:"ApiClient"`
	ApiRoles  map[string]ApiRoleConfig `json:"ApiRoles"`
}

// ApiClientConfig defines the configuration for the API client.
type ApiClientConfig struct {
	DisplayName                string `json:"displayName"`
	Enabled                    bool   `json:"enabled"`
	AccessTokenLifetimeSeconds int    `json:"accessTokenLifetimeSeconds"`
}

// ApiRoleConfig defines the configuration for a specific API role.
type ApiRoleConfig struct {
	Name       string   `json:"name"`
	Privileges []string `json:"privileges"`
}

//go:embed common_test_harness_data.json
var common_test_harness_data embed.FS

// setupJamfProClient initializes the Jamf Pro client using environment variables.
// It Requires a pre existing bootstrap api client with api roles and api integration crud permissions
func setupJamfProClientWithBootstrapAccount() (*jamfpro.Client, error) {
	// Retrieve client configuration from environment variables
	instanceName := os.Getenv("JAMFPRO_INSTANCE_NAME")
	clientID := os.Getenv("JAMFPRO_CLIENT_ID")
	clientSecret := os.Getenv("JAMFPRO_CLIENT_SECRET")

	// Check if all required environment variables are set
	if instanceName == "" || clientID == "" || clientSecret == "" {
		log.Fatalf("Environment variables for Jamf Pro client are not set")
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // Set http_client logging level to debug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: instanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	// Create a new jamfpro client instance
	Client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro bootstrap client: %v", err)
	}

	// Log the status of the client
	if Client == nil {
		log.Println("Jamf Pro bootstrap client is nil after setup")
	} else {
		log.Println("Jamf Pro bootstrap client successfully initialized")
	}

	return Client, nil
}

// setupJamfProClientWithTestIntegrationAccount initializes the Jamf Pro client using
// the integration test account
func setupJamfProClientWithTestIntegrationAccount(bootstrapClient *jamfpro.Client) (*jamfpro.Client, error) {
	// Load configuration from the embedded file
	config, err := loadIntegrationTestDataConfig("common_test_harness_data.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Extract displayName for the integration test API client
	testApiClientDisplayName := config.JamfPro.ApiClient.DisplayName

	// Before calling GetApiIntegrationNameByID, log the status of the client
	if bootstrapClient == nil {
		log.Fatalf("Bootstrap client is nil before calling GetApiIntegrationNameByID")
	}

	// Use bootstrap client to get API integration details
	apiIntegration, err := bootstrapClient.GetApiIntegrationNameByID(testApiClientDisplayName)
	if err != nil {
		log.Fatalf("Failed to get API Integration by name: %v", err)
	}

	// Extract clientId from the API Integration response
	clientID := apiIntegration.ClientID

	// Extract the resourceID from the API Integration response
	resourceID := apiIntegration.ID

	// Retrieve client configurations
	instanceName := os.Getenv("JAMFPRO_INSTANCE_NAME")

	// Use bootstrap client to update client credentials by API Integration ID
	credentials, err := bootstrapClient.UpdateClientCredentialsByApiIntegrationID(fmt.Sprintf("%d", resourceID))
	if err != nil {
		log.Fatalf("Failed to update client credentials for API Integration ID %d: %v", resourceID, err)
	}

	// Extract the client secret from the credentials
	clientSecret := credentials.ClientSecret

	// Check if instance name and client ID are set
	if instanceName == "" || clientID == "" || clientSecret == "" {
		log.Fatalf("Missing required configuration for Jamf Pro client")
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // Set http_client logging level to debug

	// Configuration for the jamfpro
	jamfConfig := jamfpro.Config{
		InstanceName: instanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	// Create a new jamfpro client instance
	Client, err := jamfpro.NewClient(jamfConfig)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	return Client, nil
}

// setupAllIntegrationTestRoles is a helper function that sets up all temporary test roles defined in the json configuration.
func setupAllIntegrationTestRoles(client *jamfpro.Client, rolesConfig map[string]ApiRoleConfig) (map[string]string, error) {
	createdRoles := make(map[string]string) // Map to store the names of created roles

	for key, roleConfig := range rolesConfig {
		roleName, err := setupIntegrationTestRole(client, roleConfig)
		if err != nil {
			// Log the error but continue attempting to create other roles
			log.Printf("Error setting up role '%s': %v", key, err)
			continue
		}
		createdRoles[key] = roleName
	}

	return createdRoles, nil
}

// setupIntegrationTestRoles sets up a temporary test role in Jamf Pro for testing purposes.
// It creates a new Jamf API role based on the provided configuration.
func setupIntegrationTestRole(client *jamfpro.Client, roleConfig ApiRoleConfig) (string, error) {
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

// loadIntegrationTestDataConfig loads the configuration from the embedded JSON file.
func loadIntegrationTestDataConfig(fileName string) (IntegrationTestGlobalConfig, error) {
	fileData, err := common_test_harness_data.ReadFile(fileName)
	if err != nil {
		return IntegrationTestGlobalConfig{}, err
	}

	var config IntegrationTestGlobalConfig
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return IntegrationTestGlobalConfig{}, err
	}

	return config, nil
}

// teardownApiIntegration is a helper function that deletes API client integration at the end of the
// test suite.
func teardownApiIntegration(client *jamfpro.Client, integrationName string) {
	if err := client.DeleteApiIntegrationByName(integrationName); err != nil {
		log.Fatalf("Failed to delete API integration: %v", err) // Exits the program if there's an error
	} else {
		log.Printf("API integration '%s' deleted successfully", integrationName)
	}
}

// teardownApiRole is a helper function that deletes API roles at the end of the test suite.
func teardownApiRole(client *jamfpro.Client, roleName string) {
	if err := client.DeleteJamfApiRoleByName(roleName); err != nil {
		log.Fatalf("Failed to delete API role: %v", err) // Exits the program if there's an error
	} else {
		log.Printf("API role '%s' deleted successfully", roleName)
	}
}
