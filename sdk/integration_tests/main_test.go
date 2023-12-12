// main_test.go
package jamfpro_integration_test

import (
	"log"
	"os"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// Global variables
var (
	bootstrapClient    *jamfpro.Client                 // initialzed client for bootstrap account
	intTestClient      *jamfpro.Client                 // initialzed client for integration testing account
	jamfApiIntegration *jamfpro.ResourceApiIntegration // jamf pro api integration / client
)

// TestMain function performs global setup of the test suite, executes
// the test suite and then performs the teardown of any resources created.
func TestMain(m *testing.M) {
	var err error
	// Load configuration from the embedded file
	config, err := loadIntegrationTestDataConfig("common_test_harness_data.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the Jamf Pro client with bootstrap credentials
	bootstrapClient, err := setupJamfProClientWithBootstrapAccount()
	if err != nil {
		log.Fatalf("Failed to setup Jamf Pro bootstrap HTTP client: %v", err)
	}

	// Setup all integration test roles using bootstrap client
	createdRoles, err := setupAllIntegrationTestRoles(bootstrapClient, config.JamfPro.ApiRoles)
	if err != nil {
		log.Fatalf("Failed to setup roles from configuration: %v", err)
	}

	// Get role names for API client setup
	var roleNames []string
	for _, roleName := range createdRoles {
		roleNames = append(roleNames, roleName)
	}

	// Setup the API client for testing using bootstrap client and configuration
	_, err = setupIntegrationTestAPIClient(bootstrapClient, config.JamfPro.ApiClient, roleNames)
	if err != nil {
		log.Fatalf("Failed to setup API client from configuration: %v", err)
	}

	// Initialize the Jamf Pro Integration Test client we just created
	intTestClient, err = setupJamfProClientWithTestIntegrationAccount(bootstrapClient)
	if err != nil {
		log.Fatalf("Failed to setup Jamf Pro Test Integration HTTP client: %v", err)
	}

	// Run the tests using the Integration Test client
	exitVal := m.Run()

	// Global Teardown: remove the integration test API client, then the api roles
	// using the bootstrap http client.
	teardownApiIntegration(bootstrapClient, jamfApiIntegration.DisplayName)
	for _, roleName := range createdRoles {
		teardownApiRole(bootstrapClient, roleName)
	}

	os.Exit(exitVal)
}
